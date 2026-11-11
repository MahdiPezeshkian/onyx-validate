// Package phone validates international phone numbers in E.164 form.
//
// Accepted input:
//   - E.164: + followed by 8 to 15 digits (country code + subscriber)
//   - Optional spaces/hyphens/parentheses are stripped before validation
//
// This does not validate national numbering plans or assignment.
// Full libphonenumber-style metadata is deferred (dependency size/maintenance).
package phone

import (
	"strings"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value string) bool { return Validate(value) == nil }

func Validate(value string) error {
	s, err := normalizeE164(value)
	if err != nil {
		return err
	}
	if s[0] != '+' {
		return validate.InvalidFormat("international phone must be in E.164 form starting with +")
	}
	digits := s[1:]
	if len(digits) < 8 || len(digits) > 15 {
		return validate.InvalidLength("E.164 number must have 8 to 15 digits after +")
	}
	if digits[0] == '0' {
		return validate.InvalidFormat("E.164 country code must not start with 0")
	}
	return nil
}

func Normalize(value string) string {
	s, err := normalizeE164(value)
	if err != nil {
		return ""
	}
	return s
}

func normalizeE164(value string) (string, error) {
	s := normalize.Digits(normalize.TrimSpace(value))
	if s == "" {
		return "", validate.Required()
	}
	var b strings.Builder
	for i, r := range s {
		switch {
		case r == '+' && i == 0:
			b.WriteRune(r)
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == ' ' || r == '-' || r == '(' || r == ')':
			continue
		default:
			if r == '+' {
				return "", validate.InvalidFormat("'+' only allowed at start")
			}
			return "", validate.InvalidCharacter("phone contains invalid characters")
		}
	}
	out := b.String()
	if out == "" || out == "+" {
		return "", validate.Required()
	}
	if out[0] != '+' {
		// Interpret bare digits as incomplete without country code.
		return "", validate.InvalidFormat("international phone must include country code with +")
	}
	return out, nil
}
