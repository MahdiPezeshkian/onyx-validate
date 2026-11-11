package normalize

import "strings"

var digitReplacer = strings.NewReplacer(
	// Persian digits
	"۰", "0", "۱", "1", "۲", "2", "۳", "3", "۴", "4",
	"۵", "5", "۶", "6", "۷", "7", "۸", "8", "۹", "9",
	// Arabic-Indic digits
	"٠", "0", "١", "1", "٢", "2", "٣", "3", "٤", "4",
	"٥", "5", "٦", "6", "٧", "7", "٨", "8", "٩", "9",
)

// Digits converts Persian and Arabic-Indic digits to ASCII digits.
// Non-digit characters are preserved unchanged.
func Digits(s string) string {
	return digitReplacer.Replace(s)
}
