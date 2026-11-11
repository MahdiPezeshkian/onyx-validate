// Package card validates Iranian bank cards by reusing generic Luhn validation
// and optional BIN lookup.
//
// Iranian cards are typically 16 digits. This package accepts 16-digit numbers
// that pass Luhn. Unknown BINs are allowed: an unknown BIN is not treated as
// invalid.
package card

import (
	generic "github.com/MahdiPezeshkian/onyx-validate/banking/card"
	"github.com/MahdiPezeshkian/onyx-validate/iran/banking/bank"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

// IsValid reports whether number is a valid Iranian-style card number.
func IsValid(number string) bool {
	return Validate(number) == nil
}

// Validate checks a 16-digit Iranian card number (Luhn).
func Validate(number string) error {
	digits := generic.Normalize(number)
	if digits == "" {
		return validate.Required()
	}
	if len(digits) != 16 {
		return validate.InvalidLength("Iranian card number must be exactly 16 digits")
	}
	return generic.Validate(number)
}

// Normalize / Format / Mask delegate to the generic card package after Validate.
func Normalize(number string) string { return generic.Normalize(number) }

func Format(number string) (string, error) {
	if err := Validate(number); err != nil {
		return "", err
	}
	return generic.Format(number)
}

func Mask(number string) (string, error) {
	if err := Validate(number); err != nil {
		return "", err
	}
	return generic.Mask(number)
}

// LookupBank returns bank metadata for the card BIN when known.
func LookupBank(number string) (bank.Bank, bool) {
	digits := Normalize(number)
	if len(digits) < 6 {
		return bank.Bank{}, false
	}
	return bank.ByBIN(digits[:6])
}
