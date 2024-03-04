package repository

import (
	"pomodoro-api/domain"
	"time"

	"gorm.io/gorm"
)

type ITimeRepository interface {
	GetTotalFocusTime(userId uint) (int, error)
	GetConsecutiveDays(userId uint) (int, error)
	GetDailyReport(userId uint) ([]domain.DailyReport, error)
	GetWeeklyReport(userId uint, startDate, endDate time.Time) ([]domain.WeeklyReport, error)
	StoreTime(time *domain.Time) error
}

type timeRepository struct {
	db *gorm.DB
}

func NewTimeRepository(db *gorm.DB) ITimeRepository {
	return &timeRepository{db}
}

func (tr *timeRepository) GetTotalFocusTime(userId uint) (int, error) {
	var totalFocusTime int
	err := tr.db.Model(&domain.Time{}).
		Select("sum(focus_time) as total_focus_time").
		Where("user_id = ?", userId).
		Scan(&totalFocusTime).Error

	if err != nil {
		return 0, err
	}

	return totalFocusTime, nil
}

func (tr *timeRepository) GetConsecutiveDays(userId uint) (int, error) {
	var consecutiveDays int
	query := `
	WITH ordered_dates AS (
	  SELECT
		user_id,
		created_at::date AS record_date,
		created_at::date - lag(created_at::date) OVER (PARTITION BY user_id ORDER BY created_at::date) AS diff
	  FROM
		times
	  WHERE
		user_id = ?
	),
	consecutive_counts AS (
	  SELECT
		user_id,
		record_date,
		SUM(CASE WHEN diff = 1 THEN 0 ELSE 1 END) OVER (PARTITION BY user_id ORDER BY record_date DESC) AS group_id
	  FROM
		ordered_dates
	)
	SELECT
	  MAX(group_id) AS consecutive_days
	FROM
	  consecutive_counts
	WHERE
	  user_id = ?
	  AND record_date <= current_date
	`

	err := tr.db.Raw(query, userId, userId).Row().Scan(&consecutiveDays)
	if err != nil {
		return 0, err
	}

	return consecutiveDays, nil
}

func (tr *timeRepository) GetDailyReport(userId uint) ([]domain.DailyReport, error) {
	var dailyReport []domain.DailyReport
	err := tr.db.Model(&domain.Time{}).
		Select("date_trunc('hour', created_at) as Time, sum(focus_time) as focus_time").
		Where("user_id = ?", userId).
		Where("created_at >= current_date AND created_at < current_date + interval '1 day'").
		Group("date_trunc('hour', created_at)").
		Scan(&dailyReport).Error

	if err != nil {
		return nil, err
	}

	return dailyReport, nil
}

func (tr *timeRepository) GetWeeklyReport(userId uint, startDate, endDate time.Time) ([]domain.WeeklyReport, error) {
	var weeklyReport []domain.WeeklyReport
	err := tr.db.Model(&domain.Time{}).
		Select(`
			DATE(created_at) as date,
			SUM(focus_time) as focus_time`).
		Where("user_id = ? AND created_at BETWEEN ? AND ?", userId, startDate, endDate).
		Group("DATE(created_at)").
		Order("DATE(created_at) ASC").
		Scan(&weeklyReport).Error

	if err != nil {
		return nil, err
	}

	return weeklyReport, nil
}

func (tr *timeRepository) StoreTime(time *domain.Time) error {
	if err := tr.db.Create(time).Error; err != nil {
		return err
	}

	return nil
}
