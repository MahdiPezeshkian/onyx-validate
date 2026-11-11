package format

import (
	"unicode"

	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsSlug(value string) bool { return ValidateSlug(value) == nil }
func ValidateSlug(value string) error {
	if value == "" {
		return validate.Required()
	}
	if value[0] == '-' || value[len(value)-1] == '-' {
		return validate.InvalidFormat("slug must not start/end with '-'")
	}
	prevHyphen := false
	for i := 0; i < len(value); i++ {
		c := value[i]
		switch {
		case c >= 'a' && c <= 'z', c >= '0' && c <= '9':
			prevHyphen = false
		case c == '-':
			if prevHyphen {
				return validate.InvalidFormat("slug must not contain consecutive hyphens")
			}
			prevHyphen = true
		default:
			return validate.InvalidCharacter("slug allows lowercase a-z, digits, and '-'")
		}
	}
	return nil
}

func IsSnake(value string) bool { return ValidateSnake(value) == nil }
func ValidateSnake(value string) error {
	return validateCase(value, '_', false)
}

func IsKebab(value string) bool { return ValidateKebab(value) == nil }
func ValidateKebab(value string) error {
	return validateCase(value, '-', false)
}

func IsCamel(value string) bool { return ValidateCamel(value) == nil }
func ValidateCamel(value string) error {
	if value == "" {
		return validate.Required()
	}
	if value[0] < 'a' || value[0] > 'z' {
		return validate.InvalidFormat("camelCase must start with lowercase letter")
	}
	for _, r := range value {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r)) || r > unicode.MaxASCII {
			return validate.InvalidCharacter("camelCase allows ASCII letters/digits only")
		}
	}
	return nil
}

func IsPascal(value string) bool { return ValidatePascal(value) == nil }
func ValidatePascal(value string) error {
	if value == "" {
		return validate.Required()
	}
	if value[0] < 'A' || value[0] > 'Z' {
		return validate.InvalidFormat("PascalCase must start with uppercase letter")
	}
	for _, r := range value {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r)) || r > unicode.MaxASCII {
			return validate.InvalidCharacter("PascalCase allows ASCII letters/digits only")
		}
	}
	return nil
}

func validateCase(value string, sep byte, _ bool) error {
	if value == "" {
		return validate.Required()
	}
	if value[0] == sep || value[len(value)-1] == sep {
		return validate.InvalidFormat("value must not start/end with separator")
	}
	prevSep := false
	for i := 0; i < len(value); i++ {
		c := value[i]
		switch {
		case c >= 'a' && c <= 'z', c >= '0' && c <= '9':
			prevSep = false
		case c == sep:
			if prevSep {
				return validate.InvalidFormat("consecutive separators are not allowed")
			}
			prevSep = true
		default:
			return validate.InvalidCharacter("invalid character for format")
		}
	}
	return nil
}
