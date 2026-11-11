package size_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/file/size"
)

func TestSize(t *testing.T) {
	opt := size.Options{MinBytes: 1, MaxBytes: 1024}
	if err := size.Validate(10, opt); err != nil {
		t.Fatal(err)
	}
	if size.IsValid(0, opt) || size.IsValid(2048, opt) {
		t.Fatal("invalid")
	}
}
