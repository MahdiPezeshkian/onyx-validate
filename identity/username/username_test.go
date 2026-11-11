package username_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/identity/username"
)

func TestUsername(t *testing.T) {
	opt := username.DefaultOptions()
	if err := username.Validate("ali_42", opt); err != nil {
		t.Fatal(err)
	}
	if username.IsValid("1ali", opt) || username.IsValid("ab", opt) {
		t.Fatal("invalid")
	}
}
