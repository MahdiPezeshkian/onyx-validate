package charset_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/text/charset"
)

func TestCharset(t *testing.T) {
	if err := charset.ValidateASCII("abc"); err != nil {
		t.Fatal(err)
	}
	if charset.IsASCII("سلام") {
		t.Fatal("persian not ascii")
	}
	if err := charset.ValidatePersian("سلام ۱۲"); err != nil {
		t.Fatal(err)
	}
	if charset.IsPersian("سلام a") {
		t.Fatal("mixed")
	}
}
