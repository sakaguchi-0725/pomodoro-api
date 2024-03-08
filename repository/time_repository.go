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
	WITH recursive dates AS (
		SELECT 0 AS day, current_date - INTERVAL '1 day' AS date
		UNION ALL
		SELECT d.day + 1, d.date - INTERVAL '1 day'
		FROM dates d
		WHERE EXISTS (
		  SELECT 1
		  FROM times
		  WHERE user_id = ?
			AND DATE(created_at) = d.date
		)
	  )
	  SELECT max(day) AS consecutive_login_days
	  FROM dates;
	`

	err := tr.db.Raw(query, userId).Row().Scan(&consecutiveDays)
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
