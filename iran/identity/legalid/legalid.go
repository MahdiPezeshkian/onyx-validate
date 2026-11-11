// Package legalid validates Iranian legal-entity national IDs (شناسه ملی حقوقی).
//
// Algorithm references (independently documented):
//   - http://www.aliarash.com/article/shenasameli/shenasa_meli.htm
//   - Widely reproduced technical write-ups (e.g. persian-tools / Abzar docs)
//
// Known valid vector used in tests: 10380284790
//
// Input policy: trim spaces; accept Persian/Arabic digits; exactly 11 ASCII digits.
// A valid checksum does not prove company registration or activity.
package legalid

import (
	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

var coefficients = [5]int{29, 27, 23, 19, 17}

// IsValid reports whether value is a structurally valid legal-entity ID.
func IsValid(value string) bool {
	return Validate(value) == nil
}

// Validate checks an Iranian legal-entity national ID.
func Validate(value string) error {
	s := normalize.Digits(normalize.TrimSpace(value))
	if s == "" {
		return validate.Required()
	}
	if len(s) != 11 {
		return validate.InvalidLength("legal ID must be exactly 11 digits")
	}
	for i := 0; i < 11; i++ {
		if s[i] < '0' || s[i] > '9' {
			return validate.InvalidCharacter("legal ID must contain only digits")
		}
	}
	allZero := true
	for i := 0; i < 11; i++ {
		if s[i] != '0' {
			allZero = false
			break
		}
	}
	if allZero {
		return validate.InvalidFormat("legal ID must not be all zeros")
	}
	// Digits at positions 3..8 (0-based) must not all be zero.
	midZero := true
	for i := 3; i <= 8; i++ {
		if s[i] != '0' {
			midZero = false
			break
		}
	}
	if midZero {
		return validate.InvalidFormat("legal ID middle segment must not be all zeros")
	}

	d := int(s[9]-'0') + 2
	sum := 0
	for i := 0; i < 10; i++ {
		sum += (d + int(s[i]-'0')) * coefficients[i%5]
	}
	rem := sum % 11
	if rem == 10 {
		rem = 0
	}
	if rem != int(s[10]-'0') {
		return validate.InvalidChecksum("legal ID checksum is invalid")
	}
	return nil
}

// Normalize returns the ASCII-digit form after trimming and digit conversion.
func Normalize(value string) string {
	return normalize.Digits(normalize.TrimSpace(value))
}
