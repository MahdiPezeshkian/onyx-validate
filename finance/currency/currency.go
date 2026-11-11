// Package currency validates ISO 4217 alphabetic currency codes.
//
// Dataset: ISO 4217 commonly assigned alphabetic codes (partial active set).
// Updated: 2026-11-11. Completeness is not claimed.
package currency

import (
	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value string) bool { return Validate(value) == nil }

func Validate(value string) error {
	s := normalize.UpperASCII(normalize.TrimSpace(value))
	if s == "" {
		return validate.Required()
	}
	if len(s) != 3 {
		return validate.InvalidLength("currency code must be 3 letters")
	}
	for i := 0; i < 3; i++ {
		if s[i] < 'A' || s[i] > 'Z' {
			return validate.InvalidCharacter("currency code must be letters")
		}
	}
	if !codes[s] {
		return validate.New(validate.CodeUnsupportedCountry, "currency code is not in the local ISO 4217 subset")
	}
	return nil
}
