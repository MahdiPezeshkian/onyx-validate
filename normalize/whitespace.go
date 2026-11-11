package normalize

import (
	"strings"
	"unicode"
)

// TrimSpace removes leading and trailing Unicode spaces.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// CompactSpaces collapses runs of Unicode spaces into a single ASCII space
// and trims the ends. Internal non-space characters are preserved.
func CompactSpaces(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	prevSpace := false
	for _, r := range strings.TrimSpace(s) {
		if unicode.IsSpace(r) {
			if !prevSpace {
				b.WriteByte(' ')
				prevSpace = true
			}
			continue
		}
		prevSpace = false
		b.WriteRune(r)
	}
	return b.String()
}

// RemoveSpaces removes all Unicode space characters.
func RemoveSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}
