package cidr

import (
	"net/netip"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value string) bool { return Validate(value) == nil }

func Validate(value string) error {
	s := normalize.TrimSpace(value)
	if s == "" {
		return validate.Required()
	}
	if _, err := netip.ParsePrefix(s); err != nil {
		return validate.InvalidFormat("CIDR prefix is invalid")
	}
	return nil
}
