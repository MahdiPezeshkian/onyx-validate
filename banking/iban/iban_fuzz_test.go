package iban_test

import (
	"testing"
	"unicode/utf8"

	"github.com/MahdiPezeshkian/onyx-validate/banking/iban"
)

func FuzzValidate(f *testing.F) {
	f.Add("GB82WEST12345698765432")
	f.Add("IR820540102680020817909002")
	f.Add("")
	f.Fuzz(func(t *testing.T, in string) {
		if !utf8.ValidString(in) {
			return
		}
		_ = iban.Validate(in)
	})
}
