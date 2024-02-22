package validator

import (
	"pomodoro-api/domain"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITimeValidator interface {
	TimeValidate(time domain.Time) error
}

type timeValidator struct{}

func NewTimeValidator() ITimeValidator {
	return &timeValidator{}
}

func (tv *timeValidator) TimeValidate(time domain.Time) error {
	return validation.ValidateStruct(&time,
		validation.Field(
			&time.FocusTime,
			validation.Required.Error("focus_time is required"),
		),
	)
}
