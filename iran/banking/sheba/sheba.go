// Package sheba validates Iranian Sheba numbers (IR IBAN).
//
// A Sheba is exactly IR + 24 digits (26 characters). Validation reuses the
// general IBAN MOD-97 algorithm. Bank code is BBAN digits[0:3].
//
// Input policy: trim; convert Persian/Arabic digits; remove spaces/hyphens
// via normalize.FoldIRANBankSeparators; upper-case. Bare 24-digit bodies are
// accepted by prepending IR.
package sheba

import (
	"github.com/MahdiPezeshkian/onyx-validate/banking/iban"
	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

// IsValid reports whether value is a structurally valid Iranian Sheba.
func IsValid(value string) bool {
	return Validate(value) == nil
}

// Validate checks Iranian Sheba structure and checksum.
func Validate(value string) error {
	s, err := normalizeSheba(value)
	if err != nil {
		return err
	}
	return iban.Validate(s)
}

// Normalize returns compact IR+24digits form, or empty if blank/invalid shape.
func Normalize(value string) string {
	s, err := normalizeSheba(value)
	if err != nil {
		return ""
	}
	return s
}

// Format returns the spaced written form of a valid Sheba.
func Format(value string) (string, error) {
	if err := Validate(value); err != nil {
		return "", err
	}
	return iban.Format(Normalize(value))
}

// BankCode returns the 3-digit bank identifier from a Sheba BBAN.
func BankCode(value string) (string, error) {
	s, err := normalizeSheba(value)
	if err != nil {
		return "", err
	}
	if len(s) != 26 {
		return "", validate.InvalidLength("sheba must be 26 characters")
	}
	return s[4:7], nil
}

func normalizeSheba(value string) (string, error) {
	s := normalize.Digits(normalize.TrimSpace(value))
	s = normalize.FoldIRANBankSeparators(s)
	s = normalize.UpperASCII(s)
	if s == "" {
		return "", validate.Required()
	}
	if len(s) == 24 {
		// Bare BBAN+check? Actually bare 24 digits means body without IR.
		allDigits := true
		for i := 0; i < 24; i++ {
			if s[i] < '0' || s[i] > '9' {
				allDigits = false
				break
			}
		}
		if allDigits {
			s = "IR" + s
		}
	}
	if len(s) < 2 || s[:2] != "IR" {
		return "", validate.InvalidFormat("sheba must start with IR")
	}
	return s, nil
}
