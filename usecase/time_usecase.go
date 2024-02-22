package usecase

import (
	"pomodoro-api/domain"
	"pomodoro-api/repository"
	"pomodoro-api/validator"
)

type ITimeUsecase interface {
	GetAllTimes(userId uint) ([]domain.TimeResponse, error)
	StoreTime(time domain.Time) (domain.TimeResponse, error)
}

type timeUsecase struct {
	tr repository.ITimeRepository
	tv validator.ITimeValidator
}

func NewTimeUsecase(tr repository.ITimeRepository, tv validator.ITimeValidator) ITimeUsecase {
	return &timeUsecase{tr, tv}
}

func (tu *timeUsecase) GetAllTimes(userId uint) ([]domain.TimeResponse, error) {
	times := []domain.Time{}
	if err := tu.tr.GetAllTimes(&times, userId); err != nil {
		return nil, err
	}

	resTimes := []domain.TimeResponse{}
	for _, v := range times {
		t := domain.TimeResponse{
			ID:        v.ID,
			FocusTime: v.FocusTime,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTimes = append(resTimes, t)
	}

	return resTimes, nil
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
