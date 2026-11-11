// Package bic validates BIC/SWIFT codes structurally (ISO 9362).
//
// Formats:
//   - 8 characters: AAAACCXX (bank + country + location)
//   - 11 characters: AAAACCXXBBB (optional branch)
//
// Syntax validity does not prove the BIC is assigned or active.
package bic

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
	if len(s) != 8 && len(s) != 11 {
		return validate.InvalidLength("BIC must be 8 or 11 characters")
	}
	for i := 0; i < 4; i++ {
		if s[i] < 'A' || s[i] > 'Z' {
			return validate.InvalidFormat("BIC bank code must be letters")
		}
	}
	for i := 4; i < 6; i++ {
		if s[i] < 'A' || s[i] > 'Z' {
			return validate.InvalidFormat("BIC country code must be letters")
		}
	}
	for i := 6; i < 8; i++ {
		c := s[i]
		if (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
			return validate.InvalidCharacter("BIC location code invalid")
		}
	}
	if len(s) == 11 {
		for i := 8; i < 11; i++ {
			c := s[i]
			if (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
				return validate.InvalidCharacter("BIC branch code invalid")
			}
		}
	}
	return nil
}
