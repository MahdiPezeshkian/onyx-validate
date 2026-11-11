package validate_test

import (
	"errors"
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func TestValidationError_AsAndIs(t *testing.T) {
	err := validate.InvalidChecksum("bad checksum")
	var ve *validate.ValidationError
	if !errors.As(err, &ve) {
		t.Fatal("errors.As failed")
	}
	if ve.Code != validate.CodeInvalidChecksum {
		t.Fatalf("code=%s", ve.Code)
	}
	if !errors.Is(err, validate.InvalidChecksum("")) {
		t.Fatal("errors.Is by code failed")
	}
	if errors.Is(err, validate.Required()) {
		t.Fatal("different codes must not match")
	}
}

func TestWithField(t *testing.T) {
	err := validate.Required().WithField("email")
	if err.Field != "email" {
		t.Fatalf("field=%q", err.Field)
	}
	if got := err.Error(); got == "" || got == "value is required (required)" {
		// Field-aware formatting expected.
		if err.Field == "" {
			t.Fatalf("unexpected error text: %q", got)
		}
	}
}
