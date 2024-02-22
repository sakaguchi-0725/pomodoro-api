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
