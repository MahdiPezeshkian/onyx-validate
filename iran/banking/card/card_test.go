package card_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/banking/card"
)

func TestIranCard(t *testing.T) {
	// Public Luhn-valid 16-digit PAN used as a structural fixture.
	valid := "4111111111111111"
	if err := card.Validate(valid); err != nil {
		t.Fatal(err)
	}
	if !card.IsValid(valid) {
		t.Fatal("expected valid 16-digit Luhn card")
	}
	if card.IsValid("411111111111111") { // 15 digits
		t.Fatal("15 digits must fail Iranian length gate")
	}
	if card.IsValid("4111111111111112") {
		t.Fatal("bad Luhn must fail")
	}
}
