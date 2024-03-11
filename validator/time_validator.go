package validator

import (
	"pomodoro-api/domain"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITimeValidator interface {
	TimeValidate(time domain.Time) error
	ReportValidate(reportType, starteDate, endDate string) error
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

func (tv *timeValidator) ReportValidate(reportType, startDate, endDate string) error {
	err := validation.Validate(reportType,
		validation.Required,
		validation.In("all", "weekly"),
	)
	if err != nil {
		return err
	}

	// YYYY-MM-DDにマッチする正規表現
	datePattern := `^\d{4}-\d{2}-\d{2}$`

	err = validation.Validate(startDate,
		validation.Match(regexp.MustCompile(datePattern)),
	)
	if err != nil {
		return err
	}

	err = validation.Validate(endDate,
		validation.Match(regexp.MustCompile(datePattern)),
	)
	if err != nil {
		return err
	}

	return nil
}
