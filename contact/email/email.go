// Package email provides pragmatic email syntax validation.
//
// It accepts common addr-spec forms: local@domain with dotted domain labels.
// It does not implement the full RFC 5322 grammar, does not verify mailboxes,
// and does not perform network lookups. Unicode local-parts / IDNA domains are
// only partially supported (domain labels must be ASCII in this version).
package email

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
	if len(s) > 254 {
		return validate.InvalidLength("email address is too long")
	}
	at := strings.LastIndexByte(s, '@')
	if at <= 0 || at == len(s)-1 {
		return validate.InvalidFormat("email must contain local-part and domain")
	}
	local, domain := s[:at], s[at+1:]
	if err := validateLocal(local); err != nil {
		return err
	}
	return validateDomain(domain)
}

func validateLocal(local string) error {
	if local == "" || len(local) > 64 {
		return validate.InvalidLength("email local-part length is invalid")
	}
	if local[0] == '.' || local[len(local)-1] == '.' || strings.Contains(local, "..") {
		return validate.InvalidFormat("email local-part has invalid dots")
	}
	for _, r := range local {
		if unicode.IsControl(r) || unicode.IsSpace(r) {
			return validate.InvalidCharacter("email local-part contains whitespace/control")
		}
		if r == '@' {
			return validate.InvalidFormat("email local-part contains '@'")
		}
	}
	return nil
}

func validateDomain(domain string) error {
	if domain == "" || len(domain) > 253 {
		return validate.InvalidLength("email domain length is invalid")
	}
	if strings.Contains(domain, "..") || domain[0] == '.' || domain[len(domain)-1] == '.' {
		return validate.InvalidFormat("email domain is invalid")
	}
	labels := strings.Split(domain, ".")
	if len(labels) < 2 {
		return validate.InvalidFormat("email domain must have at least two labels")
	}
	for _, label := range labels {
		if label == "" || len(label) > 63 {
			return validate.InvalidLength("email domain label length is invalid")
		}
		if label[0] == '-' || label[len(label)-1] == '-' {
			return validate.InvalidFormat("email domain label must not start/end with '-'")
		}
		for i := 0; i < len(label); i++ {
			c := label[i]
			ok := (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-'
			if !ok {
				return validate.InvalidCharacter("email domain must be ASCII labels in this version")
			}
		}
	}
	return nil
}
