// Package iban validates International Bank Account Numbers (ISO 13616).
//
// Input policy:
//   - Spaces are removed.
//   - Letters are upper-cased.
//   - Persian/Arabic digits are converted to ASCII.
//   - Other separators are not removed; they cause invalid_character/format errors.
//
// Country lengths are based on the SWIFT IBAN Registry. Coverage is documented
// in lengths.go (dataset date recorded there). An unknown country code yields
// unsupported_country rather than a guessed length.
package iban

import (
	"github.com/MahdiPezeshkian/onyx-validate/internal/ascii"
	"github.com/MahdiPezeshkian/onyx-validate/internal/mod97"
	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

// IsValid reports whether value is a structurally valid IBAN.
func IsValid(value string) bool {
	return Validate(value) == nil
}

// Validate checks IBAN structure, country length, and MOD-97 checksum.
func Validate(value string) error {
	s, err := normalizeIBAN(value)
	if err != nil {
		return err
	}
	cc := s[:2]
	want, ok := countryLengths[cc]
	if !ok {
		return validate.New(validate.CodeUnsupportedCountry, "IBAN country code is not in the length registry")
	}
	if len(s) != want {
		return validate.InvalidLength("IBAN length does not match country registry")
	}
	if !ascii.IsAlphaNumUpper(s) {
		return validate.InvalidCharacter("IBAN must contain only A-Z and 0-9")
	}
	if !mod97.ValidIBAN(s) {
		return validate.InvalidChecksum("IBAN MOD-97 checksum is invalid")
	}
	return nil
}

// Normalize returns the compact uppercase ASCII IBAN form, or empty on blank input.
func Normalize(value string) string {
	s, err := normalizeIBAN(value)
	if err != nil {
		return ""
	}
	return s
}

// Format groups the normalized IBAN into blocks of four characters.
func Format(value string) (string, error) {
	if err := Validate(value); err != nil {
		return "", err
	}
	s := Normalize(value)
	var out []byte
	for i := 0; i < len(s); i++ {
		if i > 0 && i%4 == 0 {
			out = append(out, ' ')
		}
		out = append(out, s[i])
	}
	return string(out), nil
}

// CountryCode returns the two-letter country code from a normalized IBAN.
func CountryCode(value string) (string, error) {
	s, err := normalizeIBAN(value)
	if err != nil {
		return "", err
	}
	if len(s) < 2 {
		return "", validate.InvalidFormat("IBAN is too short")
	}
	return s[:2], nil
}

func normalizeIBAN(value string) (string, error) {
	s := normalize.Digits(normalize.TrimSpace(value))
	s = normalize.RemoveSpaces(s)
	s = normalize.UpperASCII(s)
	if s == "" {
		return "", validate.Required()
	}
	if len(s) < 5 {
		return "", validate.InvalidLength("IBAN is too short")
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
			return "", validate.InvalidCharacter("IBAN contains invalid characters")
		}
	}
	if s[0] < 'A' || s[0] > 'Z' || s[1] < 'A' || s[1] > 'Z' {
		return "", validate.InvalidFormat("IBAN must start with a country code")
	}
	if s[2] < '0' || s[2] > '9' || s[3] < '0' || s[3] > '9' {
		return "", validate.InvalidFormat("IBAN check digits must be numeric")
	}
	return s, nil
}
