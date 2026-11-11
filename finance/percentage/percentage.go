package percentage

import (
	"strconv"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

type Options struct {
	Min         float64
	Max         float64
	MaxDecimals int
}

func DefaultOptions() Options {
	return Options{Min: 0, Max: 100, MaxDecimals: 4}
}

func IsValid(value string, opt Options) bool { return Validate(value, opt) == nil }

func Validate(value string, opt Options) error {
	s := normalize.Digits(normalize.TrimSpace(value))
	if s == "" {
		return validate.Required()
	}
	if stringsHasBadDecimals(s, opt.MaxDecimals) {
		return validate.New(validate.CodeInvalidPrecision, "percentage has too many decimals")
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return validate.InvalidFormat("percentage is not numeric")
	}
	if f < opt.Min || f > opt.Max {
		return validate.OutOfRange("percentage out of configured range")
	}
	return nil
}

func stringsHasBadDecimals(s string, max int) bool {
	if max < 0 {
		return false
	}
	i := -1
	for idx := 0; idx < len(s); idx++ {
		if s[idx] == '.' {
			i = idx
			break
		}
	}
	if i < 0 {
		return false
	}
	return len(s)-i-1 > max
}
