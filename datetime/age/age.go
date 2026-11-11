package age

import (
	"time"

	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func Years(birth, ref time.Time) int {
	years := ref.Year() - birth.Year()
	anniversary := birth.AddDate(years, 0, 0)
	if ref.Before(anniversary) {
		years--
	}
	return years
}

func ValidateMin(birth, ref time.Time, minYears int) error {
	if birth.IsZero() || ref.IsZero() {
		return validate.Required()
	}
	if birth.After(ref) {
		return validate.OutOfRange("birth date is after reference date")
	}
	if Years(birth, ref) < minYears {
		return validate.OutOfRange("age is below minimum")
	}
	return nil
}

func ValidateMax(birth, ref time.Time, maxYears int) error {
	if birth.IsZero() || ref.IsZero() {
		return validate.Required()
	}
	if Years(birth, ref) > maxYears {
		return validate.OutOfRange("age is above maximum")
	}
	return nil
}

func IsValidMin(birth, ref time.Time, minYears int) bool {
	return ValidateMin(birth, ref, minYears) == nil
}
