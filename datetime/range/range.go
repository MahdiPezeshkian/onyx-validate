package daterange

import (
	"time"

	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

type Options struct {
	Inclusive bool
}

func Validate(start, end time.Time, opt Options) error {
	if start.IsZero() || end.IsZero() {
		return validate.Required()
	}
	if opt.Inclusive {
		if end.Before(start) {
			return validate.OutOfRange("end must be on or after start")
		}
		return nil
	}
	if !end.After(start) {
		return validate.OutOfRange("end must be after start")
	}
	return nil
}

func IsValid(start, end time.Time, opt Options) bool {
	return Validate(start, end, opt) == nil
}
