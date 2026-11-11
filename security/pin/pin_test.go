package pin_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/security/pin"
)

func TestPIN(t *testing.T) {
	opt := pin.DefaultOptions()
	if err := pin.Validate("1234", opt); err != nil {
		t.Fatal(err)
	}
	if pin.IsValid("12a4", opt) || pin.IsValid("123", opt) {
		t.Fatal("invalid")
	}
}
