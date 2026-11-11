package mac

import (
	"strings"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value string) bool { return Validate(value) == nil }

func Validate(value string) error {
	s := normalize.TrimSpace(value)
	if s == "" {
		return validate.Required()
	}
	var hex []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == ':' || c == '-' || c == '.':
			continue
		case (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F'):
			hex = append(hex, c)
		default:
			return validate.InvalidCharacter("MAC contains invalid characters")
		}
	}
	if len(hex) != 12 {
		return validate.InvalidLength("MAC must contain 12 hex digits")
	}
	return nil
}

func Normalize(value string) string {
	s := strings.ToLower(normalize.TrimSpace(value))
	var hex []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') {
			hex = append(hex, c)
		}
	}
	if len(hex) != 12 {
		return ""
	}
	var b strings.Builder
	for i := 0; i < 12; i++ {
		if i > 0 && i%2 == 0 {
			b.WriteByte(':')
		}
		b.WriteByte(hex[i])
	}
	return b.String()
}
