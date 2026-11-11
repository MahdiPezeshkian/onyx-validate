// Package persian validates and converts Jalali (Persian Solar Hijri) dates.
//
// Conversion uses the classic Jalali↔Gregorian algorithm widely published for
// Iranian software. Leap years follow that algorithm’s 33-year cycle method.
//
// Reference vectors are covered in tests (e.g. 1403-01-01 ↔ 2024-03-20).
package persian

import (
	"fmt"
	"time"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

// Date is a Jalali calendar date.
type Date struct {
	Year  int
	Month int
	Day   int
}

// IsValid reports whether year/month/day form a real Jalali date.
func IsValid(year, month, day int) bool { return Validate(year, month, day) == nil }

// Validate checks Jalali date components.
func Validate(year, month, day int) error {
	if year < 1 || month < 1 || month > 12 || day < 1 {
		return validate.New(validate.CodeInvalidDate, "Jalali date components out of range")
	}
	if day > daysInMonth(year, month) {
		return validate.New(validate.CodeInvalidDate, "Jalali day is invalid for month/year")
	}
	return nil
}

// IsLeap reports whether year is a Jalali leap year in the classic cycle.
func IsLeap(year int) bool {
	return isJalaliLeapClassic(year)
}

func daysInMonth(year, month int) int {
	if month <= 6 {
		return 31
	}
	if month <= 11 {
		return 30
	}
	if IsLeap(year) {
		return 30
	}
	return 29
}

// Parse parses YYYY-MM-DD Jalali dates (ASCII or Persian/Arabic digits).
func Parse(value string) (Date, error) {
	s := normalize.Digits(normalize.TrimSpace(value))
	if s == "" {
		return Date{}, validate.Required()
	}
	var y, m, d int
	if _, err := fmt.Sscanf(s, "%d-%d-%d", &y, &m, &d); err != nil {
		return Date{}, validate.New(validate.CodeInvalidDate, "Jalali date must use YYYY-MM-DD")
	}
	if err := Validate(y, m, d); err != nil {
		return Date{}, err
	}
	return Date{Year: y, Month: m, Day: d}, nil
}

// Format renders a Jalali date as YYYY-MM-DD.
func Format(d Date) (string, error) {
	if err := Validate(d.Year, d.Month, d.Day); err != nil {
		return "", err
	}
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day), nil
}

// ToGregorian converts a Jalali date to Gregorian UTC midnight.
func ToGregorian(d Date) (time.Time, error) {
	if err := Validate(d.Year, d.Month, d.Day); err != nil {
		return time.Time{}, err
	}
	gy, gm, gd := jalaliToGregorian(d.Year, d.Month, d.Day)
	return time.Date(gy, time.Month(gm), gd, 0, 0, 0, 0, time.UTC), nil
}

// FromGregorian converts a Gregorian date to Jalali (date part only).
func FromGregorian(t time.Time) (Date, error) {
	jy, jm, jd := gregorianToJalali(t.Year(), int(t.Month()), t.Day())
	d := Date{Year: jy, Month: jm, Day: jd}
	if err := Validate(d.Year, d.Month, d.Day); err != nil {
		return Date{}, err
	}
	return d, nil
}
