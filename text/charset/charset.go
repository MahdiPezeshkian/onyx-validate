package charset

import (
	"unicode"

	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsASCII(value string) bool { return ValidateASCII(value) == nil }
func ValidateASCII(value string) error {
	for _, r := range value {
		if r > unicode.MaxASCII {
			return validate.InvalidCharacter("value contains non-ASCII characters")
		}
	}
	return nil
}

func IsNumeric(value string) bool { return ValidateNumeric(value) == nil }
func ValidateNumeric(value string) error {
	if value == "" {
		return validate.Required()
	}
	for _, r := range value {
		if r < '0' || r > '9' {
			return validate.InvalidCharacter("value must be ASCII digits")
		}
	}
	return nil
}

func IsAlpha(value string) bool { return ValidateAlpha(value) == nil }
func ValidateAlpha(value string) error {
	if value == "" {
		return validate.Required()
	}
	for _, r := range value {
		if !unicode.IsLetter(r) {
			return validate.InvalidCharacter("value must be letters")
		}
	}
	return nil
}

func IsPersian(value string) bool { return ValidatePersian(value) == nil }

// ValidatePersian allows Persian letters, Persian/Arabic digits, and spaces.
func ValidatePersian(value string) error {
	if value == "" {
		return validate.Required()
	}
	for _, r := range value {
		switch {
		case unicode.Is(unicode.Space, r):
		case r >= 0x0600 && r <= 0x06FF: // Arabic block (covers Persian letters/digits)
		case r >= 0xFB50 && r <= 0xFDFF: // Arabic presentation forms-A
		case r >= 0xFE70 && r <= 0xFEFF: // Arabic presentation forms-B
		default:
			return validate.InvalidCharacter("value contains non-Persian script characters")
		}
	}
	return nil
}
