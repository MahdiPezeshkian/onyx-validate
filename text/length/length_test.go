package length_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/text/length"
)

func TestLength(t *testing.T) {
	if err := length.Validate("سلام", 1, 10); err != nil {
		t.Fatal(err)
	}
	// "سلام" is 4 runes but 8 bytes in UTF-8.
	if err := length.ValidateUnit("سلام", 8, 8, length.Bytes); err != nil {
		t.Fatal(err)
	}
	if length.IsValid("سلام", 5, 10) {
		t.Fatal("too short in runes")
	}
}
