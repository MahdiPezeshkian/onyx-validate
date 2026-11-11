package postalcode_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/address/postalcode"
)

func TestPostalCode(t *testing.T) {
	if err := postalcode.Validate("1234567890"); err != nil {
		t.Fatal(err)
	}
	if err := postalcode.Validate("۱۲۳۴۵۶۷۸۹۰"); err != nil {
		t.Fatal(err)
	}
	if postalcode.IsValid("12345") || postalcode.IsValid("") {
		t.Fatal("expected invalid")
	}
}
