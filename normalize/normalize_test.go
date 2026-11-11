package normalize_test

import (
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
)

func TestDigits(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"۰۱۲۳۴۵۶۷۸۹", "0123456789"},
		{"٠١٢٣٤٥٦٧٨٩", "0123456789"},
		{"کد ۱۲۳-٤٥", "کد 123-45"},
		{"abc", "abc"},
		{"", ""},
	}
	for _, tc := range cases {
		if got := normalize.Digits(tc.in); got != tc.want {
			t.Fatalf("Digits(%q)=%q want %q", tc.in, got, tc.want)
		}
	}
}

func TestWhitespace(t *testing.T) {
	if got := normalize.TrimSpace("  a  "); got != "a" {
		t.Fatalf("TrimSpace: %q", got)
	}
	if got := normalize.CompactSpaces("a \t  b\nc"); got != "a b c" {
		t.Fatalf("CompactSpaces: %q", got)
	}
	if got := normalize.RemoveSpaces("a b\tc"); got != "abc" {
		t.Fatalf("RemoveSpaces: %q", got)
	}
}

func TestDigitsOnlyAndCase(t *testing.T) {
	if got := normalize.DigitsOnly("۱۲-۳a"); got != "123" {
		t.Fatalf("DigitsOnly: %q", got)
	}
	if normalize.UpperASCII("ir") != "IR" || normalize.LowerASCII("IR") != "ir" {
		t.Fatal("ASCII case helpers failed")
	}
}
