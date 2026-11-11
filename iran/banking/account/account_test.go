package account_test

import (
	"errors"
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/banking/account"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func TestAccountDeferred(t *testing.T) {
	err := account.Validate("017", "123")
	var ve *validate.ValidationError
	if !errors.As(err, &ve) || ve.Code != validate.CodeUnsupportedCountry {
		t.Fatalf("got %#v", err)
	}
	if account.IsValid("017", "123") {
		t.Fatal("deferred must be false")
	}
}
