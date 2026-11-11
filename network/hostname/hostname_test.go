package hostname_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/network/hostname"
)

func TestHostname(t *testing.T) {
	if err := hostname.Validate("example.com"); err != nil {
		t.Fatal(err)
	}
	if hostname.IsValid("-bad.com") || hostname.IsValid("") {
		t.Fatal("invalid")
	}
}
