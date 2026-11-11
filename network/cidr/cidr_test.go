package cidr_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/network/cidr"
)

func TestCIDR(t *testing.T) {
	if err := cidr.Validate("192.168.0.0/24"); err != nil {
		t.Fatal(err)
	}
	if err := cidr.Validate("2001:db8::/32"); err != nil {
		t.Fatal(err)
	}
	if cidr.IsValid("192.168.0.0/99") {
		t.Fatal("bad prefix")
	}
}
