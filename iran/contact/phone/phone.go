// Package phone dispatches Iranian phone validation to mobile or landline rules.
package phone

import (
	"github.com/MahdiPezeshkian/onyx-validate/iran/contact/landline"
	"github.com/MahdiPezeshkian/onyx-validate/iran/contact/mobile"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

// Kind classifies an Iranian phone number.
type Kind string

const (
	KindMobile   Kind = "mobile"
	KindLandline Kind = "landline"
	KindUnknown  Kind = "unknown"
)

// IsValid reports whether value is a valid Iranian mobile or landline number.
func IsValid(value string) bool { return Validate(value) == nil }

// Validate accepts either a valid mobile or landline number.
func Validate(value string) error {
	if mobile.IsValid(value) || landline.IsValid(value) {
		return nil
	}
	return validate.InvalidFormat("value is neither a valid Iranian mobile nor landline number")
}

// Classify returns the phone kind when valid.
func Classify(value string) Kind {
	switch {
	case mobile.IsValid(value):
		return KindMobile
	case landline.IsValid(value):
		return KindLandline
	default:
		return KindUnknown
	}
}
