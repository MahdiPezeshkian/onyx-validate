package bank_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/banking/bank"
)

func TestBankLookup(t *testing.T) {
	b, ok := bank.ByCode("017")
	if !ok || b.Name == "" || b.PersianName == "" {
		t.Fatalf("ByCode 017: %#v %v", b, ok)
	}
	b, ok = bank.ByBIN("603799")
	if !ok || b.Code != "017" {
		t.Fatalf("ByBIN: %#v %v", b, ok)
	}
	if _, ok := bank.ByBIN("000000"); ok {
		t.Fatal("unknown BIN must be false")
	}
}
