package nationalid_test

import (
	"testing"
	"unicode/utf8"

	"github.com/MahdiPezeshkian/onyx-validate/iran/identity/nationalid"
)

func FuzzValidate(f *testing.F) {
	f.Add("0013542419")
	f.Add("")
	f.Add("۰۰۱۳۵۴۲۴۱۹")
	f.Add("abcdefghij")
	f.Fuzz(func(t *testing.T, in string) {
		if !utf8.ValidString(in) {
			return
		}
		_ = nationalid.Validate(in)
		_ = nationalid.IsValid(in)
	})
}
