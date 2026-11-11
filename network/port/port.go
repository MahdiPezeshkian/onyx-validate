package port

import (
	"strconv"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

type Options struct {
	AllowZero bool
}

func IsValid(value string) bool { return Validate(value) == nil }

func Validate(value string) error { return ValidateWith(value, Options{}) }

func ValidateWith(value string, opt Options) error {
	s := normalize.TrimSpace(value)
	if s == "" {
		return validate.Required()
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return validate.InvalidFormat("port must be an integer")
	}
	min := 1
	if opt.AllowZero {
		min = 0
	}
	if n < min || n > 65535 {
		return validate.OutOfRange("port out of range")
	}
	return nil
}

func ValidateInt(n int, opt Options) error {
	min := 1
	if opt.AllowZero {
		min = 0
	}
	if n < min || n > 65535 {
		return validate.OutOfRange("port out of range")
	}
	return nil
}
