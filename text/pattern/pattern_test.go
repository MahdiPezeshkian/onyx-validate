package pattern_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/text/pattern"
)

func TestPattern(t *testing.T) {
	m, err := pattern.Compile(`^[a-z]+$`)
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Validate("abc"); err != nil {
		t.Fatal(err)
	}
	if m.IsValid("Abc") {
		t.Fatal("case")
	}
	if _, err := pattern.Compile(`(`); err == nil {
		t.Fatal("bad regex")
	}
}
