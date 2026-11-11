package mobile_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/contact/mobile"
)

func TestMobile(t *testing.T) {
	valid := []string{"09121234567", "+989121234567", "00989121234567", "۰۹۱۲۱۲۳۴۵۶۷", "0912 123 4567"}
	for _, v := range valid {
		if err := mobile.Validate(v); err != nil {
			t.Fatalf("%q: %v", v, err)
		}
	}
	if mobile.IsValid("02112345678") {
		t.Fatal("landline must fail")
	}
	if mobile.IsValid("") {
		t.Fatal("empty")
	}
	f, err := mobile.Format("09121234567")
	if err != nil || f != "+98 912 123 4567" {
		t.Fatalf("Format=%q %v", f, err)
	}
}
