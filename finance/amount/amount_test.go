package amount_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/finance/amount"
)

func TestAmount(t *testing.T) {
	opt := amount.DefaultOptions()
	if err := amount.Validate("12.34", opt); err != nil {
		t.Fatal(err)
	}
	if amount.IsValid("12.345", opt) {
		t.Fatal("precision")
	}
	if amount.IsValid("-1", opt) {
		t.Fatal("negative")
	}
	if err := amount.Validate("-1.00", amount.Options{AllowNegative: true, MaxDecimals: 2}); err != nil {
		t.Fatal(err)
	}
}
