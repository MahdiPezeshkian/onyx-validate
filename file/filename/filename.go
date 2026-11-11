package filename

import (
	"strings"
	"unicode"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value string) bool { return Validate(value) == nil }

func Validate(value string) error {
	s := normalize.TrimSpace(value)
	if s == "" {
		return validate.Required()
	}
	if len([]rune(s)) > 255 {
		return validate.InvalidLength("filename is too long")
	}
	if s == "." || s == ".." {
		return validate.InvalidFormat("filename must not be '.' or '..'")
	}
	if strings.Contains(s, "/") || strings.Contains(s, "\\") {
		return validate.InvalidCharacter("filename must not contain path separators")
	}
	if strings.Contains(s, "\x00") {
		return validate.InvalidCharacter("filename must not contain NUL")
	}
	for _, r := range s {
		if unicode.IsControl(r) {
			return validate.InvalidCharacter("filename must not contain control characters")
		}
	}
	return nil
}
