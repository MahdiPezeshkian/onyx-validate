package email_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/contact/email"
)

func TestEmail(t *testing.T) {
	valid := []string{"a@b.co", "john.doe+tag@example.com", "x_y@mail-server.org"}
	for _, v := range valid {
		if err := email.Validate(v); err != nil {
			t.Fatalf("%q: %v", v, err)
		}
	}
	invalid := []string{"", "no-at", "@nodomain", "a@b", "a@-b.com", "a@b..com"}
	for _, v := range invalid {
		if email.IsValid(v) {
			t.Fatalf("expected invalid: %q", v)
		}
	}
}
