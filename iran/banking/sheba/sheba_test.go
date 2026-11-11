package sheba_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/banking/sheba"
)

func TestSheba(t *testing.T) {
	v := "IR820540102680020817909002"
	if err := sheba.Validate(v); err != nil {
		t.Fatal(err)
	}
	if err := sheba.Validate("IR82 0540 1026 8002 0817 9090 02"); err != nil {
		t.Fatal(err)
	}
	if err := sheba.Validate("۸۲۰۵۴۰۱۰۲۶۸۰۰۲۰۸۱۷۹۰۹۰۰۲"); err != nil {
		t.Fatal(err)
	}
	code, err := sheba.BankCode(v)
	if err != nil || code != "054" {
		t.Fatalf("BankCode=%q %v", code, err)
	}
	if sheba.IsValid("IR820540102680020817909001") {
		t.Fatal("bad checksum should fail")
	}
}
