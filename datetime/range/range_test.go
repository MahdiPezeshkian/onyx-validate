package daterange_test

import (
	"testing"
	"time"

	daterange "github.com/MahdiPezeshkian/onyx-validate/datetime/range"
)

func TestRange(t *testing.T) {
	a := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	b := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	if err := daterange.Validate(a, b, daterange.Options{}); err != nil {
		t.Fatal(err)
	}
	if daterange.IsValid(b, a, daterange.Options{}) {
		t.Fatal("inverted")
	}
}
