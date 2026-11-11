package format_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/text/format"
)

func TestFormat(t *testing.T) {
	if err := format.ValidateSlug("hello-world"); err != nil {
		t.Fatal(err)
	}
	if err := format.ValidateSnake("hello_world"); err != nil {
		t.Fatal(err)
	}
	if err := format.ValidateCamel("helloWorld"); err != nil {
		t.Fatal(err)
	}
	if err := format.ValidatePascal("HelloWorld"); err != nil {
		t.Fatal(err)
	}
	if format.IsSlug("Hello") || format.IsCamel("Hello") {
		t.Fatal("invalid")
	}
}
