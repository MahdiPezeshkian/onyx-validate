package vin_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/vehicle/vin"
)

func TestVIN(t *testing.T) {
	// Public example VIN with valid check digit.
	v := "1HGCM82633A004352"
	if err := vin.Validate(v); err != nil {
		t.Fatal(err)
	}
	if err := vin.ValidateWith(v, vin.Options{CheckDigit: true}); err != nil {
		t.Fatal(err)
	}
	if vin.IsValid("1HGCM82633A00435I") {
		t.Fatal("I forbidden")
	}
}
