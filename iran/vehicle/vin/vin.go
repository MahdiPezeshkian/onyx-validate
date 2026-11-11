// Package vin validates Vehicle Identification Numbers using ISO 3779 structure.
//
// Rules implemented:
//   - Exactly 17 characters
//   - Characters A-Z and 0-9 excluding I, O, Q
//   - Optional North-American check digit verification (transliteration + weights)
//
// Use Validate with CheckDigit=true for ISO-style check digit verification
// commonly applied in North America. Iranian usage may not always apply the
// check digit; default Validate performs character/length checks only.
package vin

import (
	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

var translit = map[byte]int{
	'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7, 'H': 8,
	'J': 1, 'K': 2, 'L': 3, 'M': 4, 'N': 5, 'P': 7, 'R': 9,
	'S': 2, 'T': 3, 'U': 4, 'V': 5, 'W': 6, 'X': 7, 'Y': 8, 'Z': 9,
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
}

var weights = [17]int{8, 7, 6, 5, 4, 3, 2, 10, 0, 9, 8, 7, 6, 5, 4, 3, 2}

type Options struct {
	CheckDigit bool
}

func IsValid(value string) bool { return Validate(value) == nil }

func Validate(value string) error { return ValidateWith(value, Options{}) }

func ValidateWith(value string, opt Options) error {
	s := normalize.UpperASCII(normalize.TrimSpace(value))
	if s == "" {
		return validate.Required()
	}
	if len(s) != 17 {
		return validate.InvalidLength("VIN must be exactly 17 characters")
	}
	for i := 0; i < 17; i++ {
		c := s[i]
		if c == 'I' || c == 'O' || c == 'Q' {
			return validate.InvalidCharacter("VIN must not contain I, O, or Q")
		}
		if _, ok := translit[c]; !ok {
			return validate.InvalidCharacter("VIN contains invalid characters")
		}
	}
	if opt.CheckDigit {
		sum := 0
		for i := 0; i < 17; i++ {
			sum += translit[s[i]] * weights[i]
		}
		rem := sum % 11
		var expect byte = '0' + byte(rem)
		if rem == 10 {
			expect = 'X'
		}
		if s[8] != expect {
			return validate.InvalidChecksum("VIN check digit is invalid")
		}
	}
	return nil
}
