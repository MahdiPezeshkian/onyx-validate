package ip_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/network/ip"
)

func TestIP(t *testing.T) {
	if err := ip.Validate("127.0.0.1"); err != nil {
		t.Fatal(err)
	}
	if err := ip.Validate("::1"); err != nil {
		t.Fatal(err)
	}
	if err := ip.ValidateVersion("192.168.0.1", ip.V4); err != nil {
		t.Fatal(err)
	}
	if ip.IsValid("not-an-ip") {
		t.Fatal("invalid")
	}
}
