// Package length validates string lengths by runes or bytes.
package length

import "github.com/MahdiPezeshkian/onyx-validate/validate"

type Unit int

const (
	Runes Unit = iota
	Bytes
)

func IsValid(value string, min, max int) bool { return Validate(value, min, max) == nil }

// Validate checks rune length in [min, max]. Use max < 0 for unbounded max.
func Validate(value string, min, max int) error {
	return ValidateUnit(value, min, max, Runes)
}

func ValidateUnit(value string, min, max int, unit Unit) error {
	var n int
	switch unit {
	case Bytes:
		n = len(value)
	default:
		n = len([]rune(value))
	}
	if min < 0 {
		return validate.OutOfRange("minimum length must be non-negative")
	}
	if n < min {
		return validate.InvalidLength("value is shorter than minimum")
	}
	if max >= 0 && n > max {
		return validate.InvalidLength("value is longer than maximum")
	}
	return nil
}

func Exact(value string, n int) error { return Validate(value, n, n) }

func Required(value string) error {
	if value == "" {
		return validate.Required()
	}
	return nil
}
