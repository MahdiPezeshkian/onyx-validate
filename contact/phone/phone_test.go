package phone_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/contact/phone"
)

func TestPhone(t *testing.T) {
	if err := phone.Validate("+14155552671"); err != nil {
		t.Fatal(err)
	}
	if err := phone.Validate("+98 912 123 4567"); err != nil {
		t.Fatal(err)
	}
	if phone.IsValid("09121234567") || phone.IsValid("+0123") {
		t.Fatal("invalid")
	}
}
