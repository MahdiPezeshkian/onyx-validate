package percentage_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/finance/percentage"
)

func TestPercentage(t *testing.T) {
	opt := percentage.DefaultOptions()
	if err := percentage.Validate("12.5", opt); err != nil {
		t.Fatal(err)
	}
	if percentage.IsValid("101", opt) {
		t.Fatal("range")
	}
}
