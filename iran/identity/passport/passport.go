// Package passport validates Iranian passport numbers structurally.
//
// Supported format (ordinary booklet, commonly documented):
//   - One Latin letter (A-Z) followed by exactly 8 digits, e.g. A12345678
//
// Input policy: trim; upper-case letter; Persian/Arabic digits converted.
// Other categories (service/diplomatic) are not claimed here.
// Structural validity does not prove issuance or status.
package passport

import (
	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value string) bool { return Validate(value) == nil }

func Validate(value string) error {
	s := normalize.Digits(normalize.TrimSpace(value))
	s = normalize.UpperASCII(s)
	if s == "" {
		return validate.Required()
	}
	if len(s) != 9 {
		return validate.InvalidLength("ordinary Iranian passport must be 1 letter + 8 digits")
	}
	if s[0] < 'A' || s[0] > 'Z' {
		return validate.InvalidFormat("passport must start with a Latin letter")
	}
	for i := 1; i < 9; i++ {
		if s[i] < '0' || s[i] > '9' {
			return validate.InvalidCharacter("passport number digits are invalid")
		}
	}
	return nil
}
