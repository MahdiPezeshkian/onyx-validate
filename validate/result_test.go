package validate_test

import (
	"errors"
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func TestResultOK(t *testing.T) {
	r := validate.OK()
	if !r.Valid || r.Err != nil || r.Error() != nil {
		t.Fatalf("OK invariant broken: %#v", r)
	}
}

func TestResultFailNil(t *testing.T) {
	r := validate.Fail(nil)
	if r.Valid || r.Err == nil {
		t.Fatalf("Fail(nil) must be invalid with non-nil Err: %#v", r)
	}
	if r.Err.Code != validate.CodeInvalidFormat {
		t.Fatalf("code=%s", r.Err.Code)
	}
	if r.Error() == nil {
		t.Fatal("Error()")
	}
}

func TestResultFail(t *testing.T) {
	err := validate.Required()
	r := validate.Fail(err)
	if r.Valid || !errors.Is(r.Error(), err) {
		t.Fatalf("%#v", r)
	}
}
