package engine_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/vehicle/engine"
)

func TestEngine(t *testing.T) {
	if err := engine.Validate("ENG12345"); err != nil {
		t.Fatal(err)
	}
	if engine.IsValid("") || engine.IsValid("ab") {
		t.Fatal("invalid")
	}
}
