package card_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/banking/card"
)

func TestCardLuhn(t *testing.T) {
	// Visa test PAN (public test number).
	valid := "4111111111111111"
	if err := card.Validate(valid); err != nil {
		t.Fatal(err)
	}
	if err := card.Validate("4111 1111 1111 1111"); err != nil {
		t.Fatal(err)
	}
	if card.IsValid("4111111111111112") {
		t.Fatal("bad Luhn")
	}
	masked, err := card.Mask(valid)
	if err != nil || masked != "411111******1111" {
		t.Fatalf("Mask=%q %v", masked, err)
	}
}
