package mime

import (
	"strings"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value string, allowed []string) bool { return Validate(value, allowed) == nil }

// Validate checks a MIME type string against an allow-list (type/subtype).
// This is not content sniffing and does not prove file safety.
func Validate(value string, allowed []string) error {
	s := strings.ToLower(normalize.TrimSpace(value))
	if s == "" {
		return validate.Required()
	}
	parts := strings.Split(s, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return validate.InvalidFormat("MIME type must be type/subtype")
	}
	for _, a := range allowed {
		if s == strings.ToLower(strings.TrimSpace(a)) {
			return nil
		}
	}
	return validate.New(validate.CodeNotAllowed, "MIME type is not allowed")
}
