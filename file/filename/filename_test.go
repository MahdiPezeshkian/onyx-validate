package filename_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/file/filename"
)

func TestFilename(t *testing.T) {
	if err := filename.Validate("report.pdf"); err != nil {
		t.Fatal(err)
	}
	if filename.IsValid("../x") || filename.IsValid("") {
		t.Fatal("invalid")
	}
}
