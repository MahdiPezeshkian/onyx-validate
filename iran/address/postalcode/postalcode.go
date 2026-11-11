// Package postalcode validates Iranian 10-digit postal codes structurally.
//
// Input policy: trim; Persian/Arabic digits accepted; no separators accepted
// inside the code. Leading zeros are preserved.
//
// Structural validity does not identify a real delivery point.
package postalcode

import (
	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value string) bool { return Validate(value) == nil }

func Validate(value string) error {
	s := normalize.Digits(normalize.TrimSpace(value))
	if s == "" {
		return validate.Required()
	}
	if len(s) != 10 {
		return validate.InvalidLength("postal code must be exactly 10 digits")
	}
	for i := 0; i < 10; i++ {
		if s[i] < '0' || s[i] > '9' {
			return validate.InvalidCharacter("postal code must contain only digits")
		}
	}
	return nil
}

func Normalize(value string) string {
	return normalize.Digits(normalize.TrimSpace(value))
}
