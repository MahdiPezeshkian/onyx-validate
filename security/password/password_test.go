package password_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/security/password"
)

func TestPassword(t *testing.T) {
	opt := password.DefaultOptions()
	if err := password.Validate("Abcdefg1", opt); err != nil {
		t.Fatal(err)
	}
	if password.IsValid("short1A", opt) || password.IsValid("alllowercase1", opt) {
		t.Fatal("weak")
	}
}
