package port_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/network/port"
)

func TestPort(t *testing.T) {
	if err := port.Validate("8080"); err != nil {
		t.Fatal(err)
	}
	if port.IsValid("0") {
		t.Fatal("zero")
	}
	if err := port.ValidateWith("0", port.Options{AllowZero: true}); err != nil {
		t.Fatal(err)
	}
}
