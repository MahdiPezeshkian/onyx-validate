package mime_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/file/mime"
)

func TestMIME(t *testing.T) {
	if err := mime.Validate("image/png", []string{"image/png", "image/jpeg"}); err != nil {
		t.Fatal(err)
	}
	if mime.IsValid("application/x-msdownload", []string{"image/png"}) {
		t.Fatal("denied")
	}
}
