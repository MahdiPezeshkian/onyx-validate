package iban

// countryLengths maps ISO country codes to IBAN lengths from the SWIFT IBAN Registry.
// Dataset: SWIFT IBAN Registry excerpts commonly published for validation libraries.
// Coverage: major IBAN countries; not every territory is guaranteed.
// Updated: 2026-11-11
// License of this embedding: MIT (this repository). Source data is factual registry lengths.
var countryLengths = map[string]int{
	"AL": 28, "AD": 24, "AT": 20, "AZ": 28, "BH": 22, "BY": 28, "BE": 16, "BA": 20,
	"BR": 29, "BG": 22, "CR": 22, "HR": 21, "CY": 28, "CZ": 24, "DK": 18, "DO": 28,
	"TL": 23, "EE": 20, "FO": 18, "FI": 18, "FR": 27, "GE": 22, "DE": 22, "GI": 23,
	"GR": 27, "GL": 18, "GT": 28, "HU": 28, "IS": 26, "IQ": 23, "IE": 22, "IL": 23,
	"IT": 27, "JO": 30, "KZ": 20, "XK": 20, "KW": 30, "LV": 21, "LB": 28, "LI": 21,
	"LT": 20, "LU": 20, "MK": 19, "MT": 31, "MR": 27, "MU": 30, "MC": 27, "MD": 24,
	"ME": 22, "NL": 18, "NO": 15, "PK": 24, "PS": 29, "PL": 28, "PT": 25, "QA": 29,
	"RO": 24, "LC": 32, "SM": 27, "ST": 25, "SA": 24, "RS": 22, "SC": 31, "SK": 24,
	"SI": 19, "ES": 24, "SE": 24, "CH": 21, "TN": 24, "TR": 26, "UA": 29, "AE": 23,
	"GB": 22, "VA": 22, "VG": 24, "IR": 26,
}

// SupportedCountry reports whether the country code has a registered IBAN length.
func SupportedCountry(code string) bool {
	_, ok := countryLengths[normalizeCountry(code)]
	return ok
}

func normalizeCountry(code string) string {
	if len(code) != 2 {
		return code
	}
	b0, b1 := code[0], code[1]
	if b0 >= 'a' && b0 <= 'z' {
		b0 -= 32
	}
	if b1 >= 'a' && b1 <= 'z' {
		b1 -= 32
	}
	return string([]byte{b0, b1})
}
