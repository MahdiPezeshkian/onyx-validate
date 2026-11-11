// Package password validates password policy constraints.
// It does not hash passwords and never returns the password in errors.
package password

import (
	"unicode"

	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

type Options struct {
	MinLength        int
	MaxLength        int
	RequireUpper     bool
	RequireLower     bool
	RequireDigit     bool
	RequireSymbol    bool
	ForbidWhitespace bool
}

func DefaultOptions() Options {
	return Options{
		MinLength: 8, MaxLength: 128,
		RequireUpper: true, RequireLower: true, RequireDigit: true,
		ForbidWhitespace: true,
	}
}

func IsValid(value string, opt Options) bool { return Validate(value, opt) == nil }

func Validate(value string, opt Options) error {
	if opt.MinLength == 0 && opt.MaxLength == 0 {
		opt = DefaultOptions()
	}
	if value == "" {
		return validate.Required()
	}
	n := len([]rune(value))
	if n < opt.MinLength || (opt.MaxLength > 0 && n > opt.MaxLength) {
		return validate.InvalidLength("password length out of policy range")
	}
	var hasUpper, hasLower, hasDigit, hasSymbol bool
	for _, r := range value {
		switch {
		case unicode.IsSpace(r):
			if opt.ForbidWhitespace {
				return validate.InvalidCharacter("password must not contain whitespace")
			}
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		default:
			hasSymbol = true
		}
	}
	if opt.RequireUpper && !hasUpper {
		return validate.New(validate.CodeTooWeak, "password must include an uppercase letter")
	}
	if opt.RequireLower && !hasLower {
		return validate.New(validate.CodeTooWeak, "password must include a lowercase letter")
	}
	if opt.RequireDigit && !hasDigit {
		return validate.New(validate.CodeTooWeak, "password must include a digit")
	}
	if opt.RequireSymbol && !hasSymbol {
		return validate.New(validate.CodeTooWeak, "password must include a symbol")
	}
	return nil
}
