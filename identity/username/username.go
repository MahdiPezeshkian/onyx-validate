// Package username validates usernames with explicit configurable rules.
package username

import (
	"unicode"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

// Options configures username rules. Zero value is not usable; use DefaultOptions.
type Options struct {
	MinLength       int
	MaxLength       int
	AllowUnicode    bool
	AllowDigits     bool
	AllowUnderscore bool
	AllowDot        bool
	AllowHyphen     bool
	MustStartAlpha  bool
}

func DefaultOptions() Options {
	return Options{
		MinLength: 3, MaxLength: 32,
		AllowDigits: true, AllowUnderscore: true,
		MustStartAlpha: true,
	}
}

func IsValid(value string, opt Options) bool { return Validate(value, opt) == nil }

func Validate(value string, opt Options) error {
	if opt.MinLength == 0 && opt.MaxLength == 0 {
		opt = DefaultOptions()
	}
	s := normalize.TrimSpace(value)
	if s == "" {
		return validate.Required()
	}
	n := len([]rune(s))
	if n < opt.MinLength || n > opt.MaxLength {
		return validate.InvalidLength("username length out of range")
	}
	runes := []rune(s)
	if opt.MustStartAlpha {
		r := runes[0]
		if opt.AllowUnicode {
			if !unicode.IsLetter(r) {
				return validate.InvalidFormat("username must start with a letter")
			}
		} else if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') {
			return validate.InvalidFormat("username must start with an ASCII letter")
		}
	}
	for _, r := range runes {
		switch {
		case r >= 'A' && r <= 'Z', r >= 'a' && r <= 'z':
		case opt.AllowDigits && r >= '0' && r <= '9':
		case opt.AllowUnderscore && r == '_':
		case opt.AllowDot && r == '.':
		case opt.AllowHyphen && r == '-':
		case opt.AllowUnicode && unicode.IsLetter(r):
		default:
			return validate.InvalidCharacter("username contains disallowed characters")
		}
	}
	return nil
}
