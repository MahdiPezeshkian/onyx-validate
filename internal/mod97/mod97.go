// Package mod97 implements ISO 7064 / ISO 13616 MOD-97 for IBAN validation.
package mod97

// ValidIBAN reports whether an IBAN (already normalized: uppercase, no spaces)
// has a correct MOD-97 checksum (remainder == 1).
func ValidIBAN(iban string) bool {
	if len(iban) < 5 {
		return false
	}
	// Move first four characters to the end.
	rearranged := iban[4:] + iban[:4]
	return mod97(rearranged) == 1
}

// mod97 computes the remainder of the alphanumeric string modulo 97
// without allocating a big integer, processing digit-by-digit.
func mod97(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			total = (total*10 + int(c-'0')) % 97
		case c >= 'A' && c <= 'Z':
			v := int(c-'A') + 10
			// Letters expand to two digits (A=10 ... Z=35).
			total = (total*10 + v/10) % 97
			total = (total*10 + v%10) % 97
		default:
			return -1
		}
	}
	return total
}
