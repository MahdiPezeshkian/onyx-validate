package plate_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/vehicle/plate"
)

func TestPlate(t *testing.T) {
	if err := plate.Validate("12ب34567", plate.CategoryPassenger); err != nil {
		t.Fatal(err)
	}
	if err := plate.Validate("12 B 345 67", plate.CategoryPassenger); err != nil {
		t.Fatal(err)
	}
	if plate.IsValid("123") {
		t.Fatal("short")
	}
}
