package size

import "github.com/MahdiPezeshkian/onyx-validate/validate"

type Options struct {
	MinBytes int64
	MaxBytes int64 // 0 means no max
}

func IsValid(n int64, opt Options) bool { return Validate(n, opt) == nil }

func Validate(n int64, opt Options) error {
	if n < 0 {
		return validate.OutOfRange("size must not be negative")
	}
	if n < opt.MinBytes {
		return validate.OutOfRange("size is below minimum")
	}
	if opt.MaxBytes > 0 && n > opt.MaxBytes {
		return validate.OutOfRange("size is above maximum")
	}
	return nil
}
