// Package amount validates monetary amounts as decimal strings without float64.
package amount

import (
	"strings"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

type Options struct {
	AllowNegative bool
	AllowZero     bool
	MaxDecimals   int // <0 means unlimited
}

func DefaultOptions() Options {
	return Options{AllowZero: true, MaxDecimals: 2}
}

func IsValid(value string, opt Options) bool { return Validate(value, opt) == nil }

func Validate(value string, opt Options) error {
	s := normalize.Digits(normalize.TrimSpace(value))
	if s == "" {
		return validate.Required()
	}
	neg := false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		neg = true
		s = s[1:]
	}
	if neg && !opt.AllowNegative {
		return validate.OutOfRange("negative amounts are not allowed")
	}
	if s == "" {
		return validate.InvalidFormat("amount is empty")
	}
	parts := strings.Split(s, ".")
	if len(parts) > 2 {
		return validate.InvalidFormat("amount has too many decimal points")
	}
	intPart := parts[0]
	if intPart == "" {
		intPart = "0"
	}
	for i := 0; i < len(intPart); i++ {
		if intPart[i] < '0' || intPart[i] > '9' {
			return validate.InvalidCharacter("amount integer part must be digits")
		}
	}
	frac := ""
	if len(parts) == 2 {
		frac = parts[1]
		if frac == "" {
			return validate.InvalidFormat("amount fractional part is empty")
		}
		for i := 0; i < len(frac); i++ {
			if frac[i] < '0' || frac[i] > '9' {
				return validate.InvalidCharacter("amount fractional part must be digits")
			}
		}
		if opt.MaxDecimals >= 0 && len(frac) > opt.MaxDecimals {
			return validate.New(validate.CodeInvalidPrecision, "amount has too many decimal places")
		}
	}
	if !opt.AllowZero {
		allZero := true
		for i := 0; i < len(intPart); i++ {
			if intPart[i] != '0' {
				allZero = false
				break
			}
		}
		for i := 0; i < len(frac); i++ {
			if frac[i] != '0' {
				allZero = false
				break
			}
		}
		if allZero {
			return validate.OutOfRange("zero amounts are not allowed")
		}
	}
	return nil
}
