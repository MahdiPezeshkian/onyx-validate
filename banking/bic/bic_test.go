package bic_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/banking/bic"
)

func TestBIC(t *testing.T) {
	if err := bic.Validate("DEUTDEFF"); err != nil {
		t.Fatal(err)
	}
	if err := bic.Validate("DEUTDEFF500"); err != nil {
		t.Fatal(err)
	}
	if bic.IsValid("DEUTDE") || bic.IsValid("DEUT12FF") {
		t.Fatal("invalid")
	}
}
