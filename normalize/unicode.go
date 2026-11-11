package normalize

import "strings"

// FoldIRANBankSeparators removes spaces, hyphens, and common invisible marks
// used when pasting Iranian banking identifiers. It does not remove other
// punctuation. Digits are not converted; call Digits separately when needed.
func FoldIRANBankSeparators(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch r {
		case ' ', '\t', '\n', '\r', '-', '–', '—',
			'\u00a0', // NBSP
			'\u200c', // ZWNJ
			'\u200e', // LRM
			'\u200f': // RLM
			continue
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}
