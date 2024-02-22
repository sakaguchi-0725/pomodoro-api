package repository

import (
	"pomodoro-api/domain"

	"gorm.io/gorm"
)

type ITimeRepository interface {
	GetAllTimes(times *[]domain.Time, userId uint) error
	StoreTime(time *domain.Time) error
}

type timeRepository struct {
	db *gorm.DB
}

func NewTimeRepository(db *gorm.DB) ITimeRepository {
	return &timeRepository{db}
}

func (tr *timeRepository) GetAllTimes(times *[]domain.Time, userId uint) error {
	err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(times).Error
	if err != nil {
		return err
	}

	return nil
}

func (tr *timeRepository) StoreTime(time *domain.Time) error {
	if err := tr.db.Create(time).Error; err != nil {
		return err
	}

	return nil
}
