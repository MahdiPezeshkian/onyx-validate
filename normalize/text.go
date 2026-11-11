// Package normalize provides explicit text normalization helpers.
// Normalization never mutates caller-owned data; functions return new strings.
package normalize

import "strings"

// UpperASCII returns s with ASCII letters upper-cased.
// Non-ASCII characters are left unchanged.
func UpperASCII(s string) string {
	return strings.ToUpper(s)
}

// LowerASCII returns s with ASCII letters lower-cased.
// Non-ASCII characters are left unchanged.
func LowerASCII(s string) string {
	return strings.ToLower(s)
}

// DigitsOnly keeps ASCII digits after Digits conversion and drops everything else.
// Use only when a validator's documented policy allows stripping non-digits.
func DigitsOnly(s string) string {
	s = Digits(s)
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			b.WriteByte(c)
		}
	}
	return b.String()
}
