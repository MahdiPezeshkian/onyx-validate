package mac_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/network/mac"
)

func TestMAC(t *testing.T) {
	if err := mac.Validate("01:23:45:67:89:ab"); err != nil {
		t.Fatal(err)
	}
	if err := mac.Validate("01-23-45-67-89-AB"); err != nil {
		t.Fatal(err)
	}
	if mac.Normalize("01-23-45-67-89-AB") != "01:23:45:67:89:ab" {
		t.Fatal("normalize")
	}
}
