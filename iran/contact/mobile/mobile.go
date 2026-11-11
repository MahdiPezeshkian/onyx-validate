// Package mobile validates Iranian mobile numbers.
//
// Accepted formats (after digit normalization and optional separator removal):
//   - 09XXXXXXXXX          (11 digits)
//   - +989XXXXXXXXX
//   - 00989XXXXXXXXX
//
// Separators: surrounding whitespace trimmed; internal spaces and hyphens removed.
// Operator assignment is not validated (no hardcoded operator tables).
// Validity does not prove the number is active or assigned.
package mobile

import (
	"strings"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

// IsValid reports whether value is a structurally valid Iranian mobile number.
func IsValid(value string) bool { return Validate(value) == nil }

// Validate checks Iranian mobile structure.
func Validate(value string) error {
	n, err := normalizeMobile(value)
	if err != nil {
		return err
	}
	if len(n) != 11 || n[:2] != "09" {
		return validate.InvalidFormat("Iranian mobile must be 09 followed by 9 digits")
	}
	for i := 0; i < 11; i++ {
		if n[i] < '0' || n[i] > '9' {
			return validate.InvalidCharacter("mobile number must be digits")
		}
	}
	return nil
}

// Normalize returns the national 09XXXXXXXXX form, or empty if unparseable.
func Normalize(value string) string {
	n, err := normalizeMobile(value)
	if err != nil {
		return ""
	}
	return n
}

// Format returns +98 9XX XXX XXXX for a valid mobile.
func Format(value string) (string, error) {
	if err := Validate(value); err != nil {
		return "", err
	}
	n := Normalize(value) // 09xxxxxxxxx
	body := n[1:]         // 9xxxxxxxxx
	return "+98 " + body[:3] + " " + body[3:6] + " " + body[6:], nil
}

func normalizeMobile(value string) (string, error) {
	s := normalize.Digits(normalize.TrimSpace(value))
	if s == "" {
		return "", validate.Required()
	}
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= '0' && r <= '9', r == '+':
			b.WriteRune(r)
		case r == ' ' || r == '-':
			continue
		default:
			return "", validate.InvalidCharacter("mobile number contains invalid characters")
		}
	}
	s = b.String()
	switch {
	case strings.HasPrefix(s, "+98"):
		s = s[3:]
	case strings.HasPrefix(s, "0098"):
		s = s[4:]
	case strings.HasPrefix(s, "98") && len(s) == 12:
		s = s[2:]
	}
	if strings.HasPrefix(s, "9") && len(s) == 10 {
		s = "0" + s
	}
	return s, nil
}
