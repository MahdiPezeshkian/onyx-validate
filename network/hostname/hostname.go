package hostname

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
	if len(s) > 253 {
		return validate.InvalidLength("hostname too long")
	}
	labels := strings.Split(s, ".")
	for _, label := range labels {
		if label == "" || len(label) > 63 {
			return validate.InvalidLength("hostname label length invalid")
		}
		if label[0] == '-' || label[len(label)-1] == '-' {
			return validate.InvalidFormat("hostname label must not start/end with '-'")
		}
		for i := 0; i < len(label); i++ {
			c := label[i]
			ok := (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-'
			if !ok {
				return validate.InvalidCharacter("hostname contains invalid characters")
			}
		}
	}
	return nil
}
