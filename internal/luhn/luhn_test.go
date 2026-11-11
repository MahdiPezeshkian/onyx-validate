package luhn_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/internal/luhn"
)

func TestLuhn(t *testing.T) {
	if !luhn.Valid("4111111111111111") {
		t.Fatal("visa test pan")
	}
	if luhn.Valid("4111111111111112") {
		t.Fatal("bad")
	}
	if luhn.Valid("") || luhn.Valid("41a1") {
		t.Fatal("invalid input")
	}
}

func BenchmarkLuhn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = luhn.Valid("4111111111111111")
	}
}
