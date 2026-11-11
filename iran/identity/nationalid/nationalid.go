// Package nationalid validates Iranian national identification numbers (کد ملی).
//
// Input policy:
//   - Exact length 10 after digit normalization.
//   - Persian and Arabic-Indic digits are accepted and converted to ASCII.
//   - Surrounding whitespace is trimmed.
//   - No other separators are accepted.
//
// A valid checksum does not prove the ID is issued or belongs to a person.
package nationalid

import (
	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

// IsValid reports whether value is a structurally valid Iranian national ID.
func IsValid(value string) bool {
	return Validate(value) == nil
}

// Validate checks an Iranian national ID.
func Validate(value string) error {
	s := normalize.Digits(normalize.TrimSpace(value))
	if s == "" {
		return validate.Required()
	}
	if len(s) != 10 {
		return validate.InvalidLength("national ID must be exactly 10 digits")
	}
	for i := 0; i < 10; i++ {
		if s[i] < '0' || s[i] > '9' {
			return validate.InvalidCharacter("national ID must contain only digits")
		}
	}
	// Reject trivial repeated sequences (0000000000 ... 9999999999).
	same := true
	for i := 1; i < 10; i++ {
		if s[i] != s[0] {
			same = false
			break
		}
	}
	if same {
		return validate.InvalidFormat("national ID must not be a repeated digit sequence")
	}
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(s[i]-'0') * (10 - i)
	}
	rem := sum % 11
	check := int(s[9] - '0')
	ok := (rem < 2 && check == rem) || (rem >= 2 && check == 11-rem)
	if !ok {
		return validate.InvalidChecksum("national ID checksum is invalid")
	}
	return nil
}

// Normalize returns the ASCII-digit form after trimming spaces and converting
// Persian/Arabic digits. It does not validate.
func Normalize(value string) string {
	return normalize.Digits(normalize.TrimSpace(value))
}
