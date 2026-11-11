// Package landline validates Iranian geographic (landline) numbers in national form.
//
// Accepted national form after normalization: 0 + 2–3 digit area code + subscriber
// digits, totaling 11 digits starting with 0 but not 09 (mobile).
//
// International forms +98 / 0098 are accepted and converted to national form.
// This does not validate area-code assignment tables.
package landline

import (
	"strings"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value string) bool { return Validate(value) == nil }

func Validate(value string) error {
	n, err := normalizeLandline(value)
	if err != nil {
		return err
	}
	if len(n) != 11 {
		return validate.InvalidLength("Iranian landline national form must be 11 digits")
	}
	if n[0] != '0' || n[1] == '9' {
		return validate.InvalidFormat("landline must start with 0 and not be a mobile prefix 09")
	}
	for i := 0; i < len(n); i++ {
		if n[i] < '0' || n[i] > '9' {
			return validate.InvalidCharacter("landline must be digits")
		}
	}
	return nil
}

func Normalize(value string) string {
	n, err := normalizeLandline(value)
	if err != nil {
		return ""
	}
	return n
}

func normalizeLandline(value string) (string, error) {
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
			return "", validate.InvalidCharacter("landline contains invalid characters")
		}
	}
	s = b.String()
	switch {
	case strings.HasPrefix(s, "+98"):
		s = "0" + s[3:]
	case strings.HasPrefix(s, "0098"):
		s = "0" + s[4:]
	case strings.HasPrefix(s, "98") && len(s) >= 10:
		s = "0" + s[2:]
	}
	return s, nil
}
