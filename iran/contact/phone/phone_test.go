package phone_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/contact/phone"
)

func TestPhone(t *testing.T) {
	if !phone.IsValid("09121234567") || phone.Classify("09121234567") != phone.KindMobile {
		t.Fatal("mobile")
	}
	if !phone.IsValid("02188776655") || phone.Classify("02188776655") != phone.KindLandline {
		t.Fatal("landline")
	}
	if phone.IsValid("12345") {
		t.Fatal("invalid")
	}
}
