package nationalid_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/identity/nationalid"
)

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = nationalid.Validate("0013542419")
	}
}
