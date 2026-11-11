// Package date validates Gregorian dates with explicit layouts.
package date

import (
	"time"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value, layout string) bool { return Validate(value, layout) == nil }

func Validate(value, layout string) error {
	s := normalize.TrimSpace(value)
	if s == "" {
		return validate.Required()
	}
	if layout == "" {
		return validate.InvalidFormat("date layout is required")
	}
	if _, err := time.Parse(layout, s); err != nil {
		return validate.New(validate.CodeInvalidDate, "date does not match layout or is not a real calendar date")
	}
	return nil
}

func Before(value, layout string, ref time.Time) error {
	t, err := parse(value, layout)
	if err != nil {
		return err
	}
	if !t.Before(ref) {
		return validate.OutOfRange("date must be before reference")
	}
	return nil
}

func After(value, layout string, ref time.Time) error {
	t, err := parse(value, layout)
	if err != nil {
		return err
	}
	if !t.After(ref) {
		return validate.OutOfRange("date must be after reference")
	}
	return nil
}

func parse(value, layout string) (time.Time, error) {
	s := normalize.TrimSpace(value)
	if s == "" {
		return time.Time{}, validate.Required()
	}
	t, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{}, validate.New(validate.CodeInvalidDate, "date does not match layout or is not a real calendar date")
	}
	return t, nil
}
