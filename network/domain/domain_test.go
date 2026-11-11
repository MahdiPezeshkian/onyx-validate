package domain_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/network/domain"
)

func TestDomain(t *testing.T) {
	if err := domain.Validate("foo.bar"); err != nil {
		t.Fatal(err)
	}
	if domain.IsValid("localhost") {
		t.Fatal("single label")
	}
}
