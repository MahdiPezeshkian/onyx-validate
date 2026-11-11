// Package card provides generic payment-card number validation (Luhn).
//
// Input policy:
//   - Trim surrounding whitespace.
//   - Convert Persian/Arabic digits.
//   - Remove spaces and hyphens only.
//   - Length must be 12–19 digits (ISO/IEC 7812 common range).
//
// Luhn validity does not prove issuance or account status.
package card

import (
	"strings"

	"github.com/MahdiPezeshkian/onyx-validate/internal/luhn"
	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

// IsValid reports whether number passes structural and Luhn checks.
func IsValid(number string) bool {
	return Validate(number) == nil
}

// Validate checks card number structure and Luhn checksum.
func Validate(number string) error {
	digits, err := normalizeCard(number)
	if err != nil {
		return err
	}
	if len(digits) < 12 || len(digits) > 19 {
		return validate.InvalidLength("card number must be 12 to 19 digits")
	}
	if !luhn.Valid(digits) {
		return validate.InvalidChecksum("card number failed Luhn checksum")
	}
	return nil
}

// Normalize returns compact ASCII digits, or empty if blank after cleanup.
func Normalize(number string) string {
	digits, err := normalizeCard(number)
	if err != nil {
		return ""
	}
	return digits
}

// Format groups digits in blocks of 4.
func Format(number string) (string, error) {
	if err := Validate(number); err != nil {
		return "", err
	}
	d := Normalize(number)
	var b strings.Builder
	for i := 0; i < len(d); i++ {
		if i > 0 && i%4 == 0 {
			b.WriteByte(' ')
		}
		b.WriteByte(d[i])
	}
	return b.String(), nil
}

// Mask keeps the first 6 and last 4 digits visible when length allows.
func Mask(number string) (string, error) {
	if err := Validate(number); err != nil {
		return "", err
	}
	d := Normalize(number)
	if len(d) < 10 {
		return strings.Repeat("*", len(d)), nil
	}
	masked := d[:6] + strings.Repeat("*", len(d)-10) + d[len(d)-4:]
	return masked, nil
}

func normalizeCard(number string) (string, error) {
	s := normalize.Digits(normalize.TrimSpace(number))
	if s == "" {
		return "", validate.Required()
	}
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch {
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == ' ' || r == '-':
			continue
		default:
			return "", validate.InvalidCharacter("card number contains invalid characters")
		}
	}
	return b.String(), nil
}
