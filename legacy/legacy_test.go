package legacy_test

import (
	"strings"
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/legacy"
)

func TestValidateNotEmpty(t *testing.T) {
	if err := legacy.ValidateNotEmpty("a", 1, []int{1}); err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if err := legacy.ValidateNotEmpty(""); err == nil {
		t.Fatal("expected empty string error")
	}
	if err := legacy.ValidateNotEmpty(0); err == nil {
		t.Fatal("expected zero int error")
	}
	if err := legacy.ValidateNotEmpty([]string{}); err == nil {
		t.Fatal("expected empty slice error")
	}
	var p *int
	if err := legacy.ValidateNotEmpty(p); err == nil {
		t.Fatal("expected nil pointer error")
	}
}

func TestValidateNonNegative(t *testing.T) {
	if err := legacy.ValidateNonNegative(0, 1, uint(0), 0.0); err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if err := legacy.ValidateNonNegative(-1); err == nil {
		t.Fatal("expected negative error")
	}
	if err := legacy.ValidateNonNegative("x"); err == nil {
		t.Fatal("expected non-numeric error")
	}
}

func TestValidatePositive(t *testing.T) {
	if err := legacy.ValidatePositive(1, uint(2), 0.5); err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if err := legacy.ValidatePositive(0); err == nil {
		t.Fatal("expected non-positive error")
	}
	if err := legacy.ValidatePositive(-3.14); err == nil {
		t.Fatal("expected non-positive float error")
	}
}

func TestValidateMinMaxRange(t *testing.T) {
	if err := legacy.ValidateMin(10, 10, 11); err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if err := legacy.ValidateMin(10, 9); err == nil {
		t.Fatal("expected min error")
	}
	if err := legacy.ValidateMax(10, 9, 10); err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if err := legacy.ValidateMax(10, 11); err == nil {
		t.Fatal("expected max error")
	}
	if err := legacy.ValidateRange(1, 5, 1, 3, 5); err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if err := legacy.ValidateRange(1, 5, 0); err == nil {
		t.Fatal("expected range error")
	}
}

func TestValidateLength(t *testing.T) {
	if err := legacy.ValidateLength(3, "abc", []int{1, 2, 3}); err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if err := legacy.ValidateLength(2, "abc"); err == nil {
		t.Fatal("expected length error")
	}
	if err := legacy.ValidateLength(1, 42); err == nil {
		t.Fatal("expected non-length-bearing type error")
	}
}

func TestValidateEmail(t *testing.T) {
	if err := legacy.ValidateEmail("a@b.com"); err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if err := legacy.ValidateEmail("not-an-email"); err == nil {
		t.Fatal("expected email error")
	}
	err := legacy.ValidateEmail("bad")
	if err == nil || !strings.Contains(err.Error(), "not a valid email") {
		t.Fatalf("error message: %v", err)
	}
}
