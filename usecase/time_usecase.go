package usecase

import (
	"pomodoro-api/domain"
	"pomodoro-api/repository"
	"pomodoro-api/validator"
	"time"
)

type ITimeUsecase interface {
	GetReport(userId uint, reportType, startDate, endDate string) (domain.ReportResponse, error)
	StoreTime(time domain.Time) (domain.TimeResponse, error)
}

type timeUsecase struct {
	tr repository.ITimeRepository
	tv validator.ITimeValidator
}

func NewTimeUsecase(tr repository.ITimeRepository, tv validator.ITimeValidator) ITimeUsecase {
	return &timeUsecase{tr, tv}
}

func (tu *timeUsecase) GetReport(userId uint, reportType, startDate, endDate string) (domain.ReportResponse, error) {
	if err := tu.tv.ReportValidate(reportType, startDate, endDate); err != nil {
		return domain.ReportResponse{}, err
	}

	var reportRes domain.ReportResponse
	parsedStartDate, _ := time.Parse("2006-01-02", startDate)
	parsedEndDate, _ := time.Parse("2006-01-02", endDate)

	totalFocusTime, err := tu.tr.GetTotalFocusTime(userId)
	if err != nil {
		return domain.ReportResponse{}, err
	}

	consecutiveDays, err := tu.tr.GetConsecutiveDays(userId)
	if err != nil {
		return domain.ReportResponse{}, err
	}

	reportRes.TotalFocusTime = totalFocusTime
	reportRes.ConsecutiveDays = consecutiveDays

	switch reportType {
	case "all":
		dailyReport, err := tu.tr.GetDailyReport(userId)
		if err != nil {
			return domain.ReportResponse{}, nil
		}

		weeklyReport, err := tu.tr.GetWeeklyReport(userId, parsedStartDate, parsedEndDate)
		if err != nil {
			return domain.ReportResponse{}, err
		}

		reportRes.DailyReport = dailyReport
		reportRes.WeeklyReport = weeklyReport
	case "weekly":
		weeklyReport, err := tu.tr.GetWeeklyReport(userId, parsedStartDate, parsedEndDate)
		if err != nil {
			return domain.ReportResponse{}, err
		}

		reportRes.WeeklyReport = weeklyReport
	}

	return reportRes, nil
}

func (tu *timeUsecase) StoreTime(time domain.Time) (domain.TimeResponse, error) {
	if err := tu.tv.TimeValidate(time); err != nil {
		return domain.TimeResponse{}, err
	}

	if err := tu.tr.StoreTime(&time); err != nil {
		return domain.TimeResponse{}, nil
	}

	resTime := domain.TimeResponse{
		ID:        time.ID,
		FocusTime: time.FocusTime,
		CreatedAt: time.CreatedAt,
		UpdatedAt: time.UpdatedAt,
	}

	return resTime, nil
}
