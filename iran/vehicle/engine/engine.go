// Package engine provides conservative structural checks for engine numbers.
//
// There is no universal checksum. This package only rejects empty values and
// clearly invalid characters (spaces-only / control characters). Prefer
// domain-specific rules in applications.
package engine

import (
	"unicode"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value string) bool { return Validate(value) == nil }

func Validate(value string) error {
	s := normalize.TrimSpace(value)
	if s == "" {
		return validate.Required()
	}
	if len([]rune(s)) < 3 || len([]rune(s)) > 32 {
		return validate.InvalidLength("engine number length out of accepted structural range 3-32")
	}
	for _, r := range s {
		if unicode.IsControl(r) {
			return validate.InvalidCharacter("engine number contains control characters")
		}
	}
	return nil
}
