package iban_test

import (
	"errors"
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/banking/iban"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func TestIBANKnownVectors(t *testing.T) {
	// Official/common published examples (MOD-97 == 1).
	valid := []string{
		"GB82 WEST 1234 5698 7654 32",
		"DE89370400440532013000",
		"IR820540102680020817909002",
		"ir82 0540 1026 8002 0817 9090 02",
	}
	for _, v := range valid {
		if err := iban.Validate(v); err != nil {
			t.Fatalf("Validate(%q)=%v", v, err)
		}
	}
	if err := iban.Validate("GB82WEST12345698765431"); err == nil {
		t.Fatal("expected checksum failure")
	}
	var ve *validate.ValidationError
	err := iban.Validate("ZZ820540102680020817909002")
	if !errors.As(err, &ve) || ve.Code != validate.CodeUnsupportedCountry {
		t.Fatalf("unsupported country: %#v", err)
	}
	formatted, err := iban.Format("DE89370400440532013000")
	if err != nil || formatted != "DE89 3704 0044 0532 0130 00" {
		t.Fatalf("Format: %q %v", formatted, err)
	}
}
