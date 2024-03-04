package domain

import "time"

type Time struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FocusTime int       `json:"focus_time" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type TimeResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FocusTime int       `json:"focus_time" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ReportResponse struct {
	TotalFocusTime  int            `json:"total_focus_time" gorm:"column:total_focus_time"`
	ConsecutiveDays int            `json:"consecutive_days" gorm:"column:consecutive_days"`
	DailyReport     []DailyReport  `json:"daily_report"`
	WeeklyReport    []WeeklyReport `json:"weekly_report"`
}

type DailyReport struct {
	Time      string `json:"time" gorm:"column:time"`
	FocusTime int    `json:"focus_time" gorm:"column:focus_time"`
}

type WeeklyReport struct {
	Date      string `json:"date" gorm:"column:date"`
	DayOfWeek string `json:"day_of_week" gorm:"column:day_of_week"`
	FocusTime int    `json:"focus_time" gorm:"column:focus_time"`
}
