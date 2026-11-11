package nationalid_test

import (
	"errors"
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/identity/nationalid"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func TestNationalID(t *testing.T) {
	valid := []string{
		"0013542419", // independently known valid vector
		"0079028748",
		"۰۰۱۳۵۴۲۴۱۹", // Persian digits
		" 0013542419 ",
	}
	for _, v := range valid {
		if err := nationalid.Validate(v); err != nil {
			t.Fatalf("Validate(%q)=%v", v, err)
		}
		if !nationalid.IsValid(v) {
			t.Fatalf("IsValid(%q)=false", v)
		}
	}

	invalid := []struct {
		in   string
		code string
	}{
		{"", validate.CodeRequired},
		{"123", validate.CodeInvalidLength},
		{"001354241a", validate.CodeInvalidCharacter},
		{"0000000000", validate.CodeInvalidFormat},
		{"1111111111", validate.CodeInvalidFormat},
		{"0013542418", validate.CodeInvalidChecksum}, // wrong check digit
	}
	for _, tc := range invalid {
		err := nationalid.Validate(tc.in)
		var ve *validate.ValidationError
		if !errors.As(err, &ve) || ve.Code != tc.code {
			t.Fatalf("Validate(%q): got %#v want code %s", tc.in, err, tc.code)
		}
	}
}
