package passport_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/identity/passport"
)

func TestPassport(t *testing.T) {
	if err := passport.Validate("A12345678"); err != nil {
		t.Fatal(err)
	}
	if err := passport.Validate("a۱۲۳۴۵۶۷۸"); err != nil {
		t.Fatal(err)
	}
	if passport.IsValid("12345678") || passport.IsValid("") {
		t.Fatal("invalid")
	}
}
