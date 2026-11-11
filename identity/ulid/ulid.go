// Package ulid validates ULID strings structurally (Crockford Base32, 26 chars).
package ulid

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
	if len(s) != 26 {
		return validate.InvalidLength("ULID must be 26 characters")
	}
	for i := 0; i < 26; i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
		case c >= 'A' && c <= 'Z':
			if c == 'I' || c == 'L' || c == 'O' || c == 'U' {
				return validate.InvalidCharacter("ULID uses Crockford Base32 (no I,L,O,U)")
			}
		default:
			return validate.InvalidCharacter("ULID contains invalid characters")
		}
	}
	// First character must be 0-7 to keep timestamp 48-bit.
	if s[0] > '7' {
		return validate.OutOfRange("ULID timestamp overflow in first character")
	}
	return nil
}
