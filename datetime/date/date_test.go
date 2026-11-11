package date_test

import (
	"testing"
	"time"

	"github.com/MahdiPezeshkian/onyx-validate/datetime/date"
)

func TestDate(t *testing.T) {
	if err := date.Validate("2024-02-29", "2006-01-02"); err != nil {
		t.Fatal(err)
	}
	if date.IsValid("2023-02-29", "2006-01-02") {
		t.Fatal("non-leap")
	}
	ref := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	if err := date.Before("2023-12-31", "2006-01-02", ref); err != nil {
		t.Fatal(err)
	}
}
