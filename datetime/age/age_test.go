package age_test

import (
	"testing"
	"time"

	"github.com/MahdiPezeshkian/onyx-validate/datetime/age"
)

func TestAge(t *testing.T) {
	birth := time.Date(2000, 6, 15, 0, 0, 0, 0, time.UTC)
	ref := time.Date(2024, 6, 14, 0, 0, 0, 0, time.UTC)
	if age.Years(birth, ref) != 23 {
		t.Fatalf("years=%d", age.Years(birth, ref))
	}
	ref2 := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
	if age.Years(birth, ref2) != 24 {
		t.Fatalf("years=%d", age.Years(birth, ref2))
	}
	if err := age.ValidateMin(birth, ref2, 18); err != nil {
		t.Fatal(err)
	}
}
