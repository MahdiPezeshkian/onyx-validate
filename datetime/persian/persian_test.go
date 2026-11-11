package persian_test

import (
	"testing"
	"time"

	"github.com/MahdiPezeshkian/onyx-validate/datetime/persian"
)

func TestPersianNowruz2024(t *testing.T) {
	d := persian.Date{Year: 1403, Month: 1, Day: 1}
	if err := persian.Validate(d.Year, d.Month, d.Day); err != nil {
		t.Fatal(err)
	}
	g, err := persian.ToGregorian(d)
	if err != nil {
		t.Fatal(err)
	}
	if g.Format("2006-01-02") != "2024-03-20" {
		t.Fatalf("got %s", g.Format("2006-01-02"))
	}
	back, err := persian.FromGregorian(time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	if back != d {
		t.Fatalf("roundtrip %#v", back)
	}
}

func TestPersianInvalid(t *testing.T) {
	if persian.IsValid(1403, 12, 30) && !persian.IsLeap(1403) {
		// 1403 leap? check day 30 of month 12 only if leap
	}
	if persian.IsValid(1403, 13, 1) {
		t.Fatal("month")
	}
}
