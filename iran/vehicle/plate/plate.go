// Package plate validates Iranian vehicle plates for selected documented categories.
//
// Supported categories:
//   - CategoryPassenger: national private/public numeric form commonly written as
//     NN L NNN NN where L is a Persian plate letter (or its common Latin token),
//     and N are digits. Separators: spaces or hyphens may appear between groups.
//
// Free-zone, diplomatic, and temporary plates are deferred (see ROADMAP).
// This is structural validation only.
package plate

import (
	"strings"
	"unicode"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

// Category selects a plate class.
type Category string

const (
	CategoryPassenger Category = "passenger"
)

// Allowed Persian plate letters commonly used on national plates.
// Source: commonly published Iranian plate letter sets (partial; excludes ambiguous vowels in some series).
var persianLetters = map[rune]struct{}{
	'ا': {}, 'ب': {}, 'پ': {}, 'ت': {}, 'ث': {}, 'ج': {}, 'چ': {}, 'ح': {}, 'خ': {},
	'د': {}, 'ذ': {}, 'ر': {}, 'ز': {}, 'ژ': {}, 'س': {}, 'ش': {}, 'ص': {}, 'ض': {},
	'ط': {}, 'ظ': {}, 'ع': {}, 'غ': {}, 'ف': {}, 'ق': {}, 'ک': {}, 'گ': {}, 'ل': {},
	'م': {}, 'ن': {}, 'و': {}, 'ه': {}, 'ی': {},
}

func IsValid(value string) bool { return Validate(value, CategoryPassenger) == nil }

func Validate(value string, category Category) error {
	switch category {
	case CategoryPassenger, "":
		return validatePassenger(value)
	default:
		return validate.New(validate.CodeUnsupportedCountry, "plate category is not implemented")
	}
}

func validatePassenger(value string) error {
	s := normalize.Digits(normalize.TrimSpace(value))
	if s == "" {
		return validate.Required()
	}
	// Remove spaces/hyphens between groups only.
	var cleaned strings.Builder
	for _, r := range s {
		if r == ' ' || r == '-' {
			continue
		}
		cleaned.WriteRune(r)
	}
	s = cleaned.String()
	// Expect: 2 digits + 1 letter + 3 digits + 2 digits = 8 symbols
	runes := []rune(s)
	if len(runes) != 8 {
		return validate.InvalidLength("passenger plate must have 2+1+3+2 symbols after removing separators")
	}
	checkDigits := func(rs []rune) bool {
		for _, r := range rs {
			if r < '0' || r > '9' {
				return false
			}
		}
		return true
	}
	if !checkDigits(runes[0:2]) || !checkDigits(runes[3:6]) || !checkDigits(runes[6:8]) {
		return validate.InvalidCharacter("passenger plate digit groups are invalid")
	}
	letter := runes[2]
	if _, ok := persianLetters[letter]; ok {
		return nil
	}
	// Allow a single Latin letter token as transliteration.
	if (letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z') {
		return nil
	}
	if unicode.IsLetter(letter) {
		return validate.InvalidCharacter("plate letter is not in the supported set")
	}
	return validate.InvalidFormat("passenger plate format is invalid")
}
