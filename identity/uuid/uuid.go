// Package uuid validates UUID text forms (RFC 4122).
package uuid

import (
	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value string) bool { return Validate(value) == nil }

// Validate accepts 8-4-4-4-12 hex form with hyphens (canonical).
func Validate(value string) error {
	s := normalize.TrimSpace(value)
	if s == "" {
		return validate.Required()
	}
	if len(s) != 36 {
		return validate.InvalidLength("UUID must be 36 characters in canonical form")
	}
	for i, c := range s {
		switch i {
		case 8, 13, 18, 23:
			if c != '-' {
				return validate.InvalidFormat("UUID hyphens are misplaced")
			}
		default:
			ok := (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
			if !ok {
				return validate.InvalidCharacter("UUID contains non-hex characters")
			}
		}
	}
	return nil
}

// Version returns the UUID version nibble when structurally valid.
func Version(value string) (int, error) {
	if err := Validate(value); err != nil {
		return 0, err
	}
	s := normalize.TrimSpace(value)
	v := s[14]
	switch {
	case v >= '0' && v <= '9':
		return int(v - '0'), nil
	case v >= 'a' && v <= 'f':
		return int(v-'a') + 10, nil
	case v >= 'A' && v <= 'F':
		return int(v-'A') + 10, nil
	default:
		return 0, validate.InvalidFormat("UUID version nibble invalid")
	}
}
