package ulid_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/identity/ulid"
)

func TestULID(t *testing.T) {
	if err := ulid.Validate("01ARZ3NDEKTSV4RRFFQ69G5FAV"); err != nil {
		t.Fatal(err)
	}
	if ulid.IsValid("81ARZ3NDEKTSV4RRFFQ69G5FAV") { // overflow
		t.Fatal("overflow")
	}
	if ulid.IsValid("01ARZ3NDEKTSV4RRFFQ69G5FAI") {
		t.Fatal("I forbidden")
	}
}
