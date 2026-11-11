package pin

import "github.com/MahdiPezeshkian/onyx-validate/validate"

type Options struct {
	Length          int
	AllowOnlyDigits bool
}

func DefaultOptions() Options {
	return Options{Length: 4, AllowOnlyDigits: true}
}

func IsValid(value string, opt Options) bool { return Validate(value, opt) == nil }

func Validate(value string, opt Options) error {
	if opt.Length == 0 {
		opt = DefaultOptions()
	}
	if value == "" {
		return validate.Required()
	}
	if len(value) != opt.Length {
		return validate.InvalidLength("PIN length does not match policy")
	}
	if opt.AllowOnlyDigits {
		for i := 0; i < len(value); i++ {
			if value[i] < '0' || value[i] > '9' {
				return validate.InvalidCharacter("PIN must be digits")
			}
		}
	}
	return nil
}
