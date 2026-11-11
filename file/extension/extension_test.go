package extension_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/file/extension"
)

func TestExtension(t *testing.T) {
	if err := extension.Validate("a.PNG", []string{".png", "jpg"}); err != nil {
		t.Fatal(err)
	}
	if extension.IsValid("a.exe", []string{".png"}) {
		t.Fatal("denied")
	}
}
