package landline_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/contact/landline"
)

func TestLandline(t *testing.T) {
	if err := landline.Validate("02188776655"); err != nil {
		t.Fatal(err)
	}
	if err := landline.Validate("+982188776655"); err != nil {
		t.Fatal(err)
	}
	if landline.IsValid("09121234567") {
		t.Fatal("mobile must fail")
	}
}
