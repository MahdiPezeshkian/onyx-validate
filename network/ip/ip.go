// Package ip validates IP address strings using net/netip.
package ip

import (
	"net/netip"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

type Version int

const (
	Any Version = iota
	V4
	V6
)

func IsValid(value string) bool { return Validate(value) == nil }

func Validate(value string) error { return ValidateVersion(value, Any) }

func ValidateVersion(value string, ver Version) error {
	s := normalize.TrimSpace(value)
	if s == "" {
		return validate.Required()
	}
	addr, err := netip.ParseAddr(s)
	if err != nil {
		return validate.InvalidFormat("IP address is invalid")
	}
	switch ver {
	case V4:
		if !addr.Is4() {
			return validate.InvalidFormat("IPv4 address required")
		}
	case V6:
		if !addr.Is6() {
			return validate.InvalidFormat("IPv6 address required")
		}
	}
	return nil
}
