package currency_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/finance/currency"
)

func TestCurrency(t *testing.T) {
	if err := currency.Validate("IRR"); err != nil {
		t.Fatal(err)
	}
	if currency.IsValid("ZZZ") || currency.IsValid("us") {
		t.Fatal("invalid")
	}
}
