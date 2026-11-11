package domain

import (
	"strings"

	"github.com/MahdiPezeshkian/onyx-validate/network/hostname"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(value string) bool { return Validate(value) == nil }

// Validate requires a hostname with at least two labels (name + TLD).
func Validate(value string) error {
	if err := hostname.Validate(value); err != nil {
		return err
	}
	if strings.Count(value, ".") < 1 {
		return validate.InvalidFormat("domain must include a top-level label")
	}
	return nil
}
