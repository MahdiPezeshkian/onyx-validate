// Package bank provides Iranian bank metadata lookup.
//
// Dataset status: PARTIAL.
// Sources:
//   - Bank Keshavarzi Sheba documentation (bank codes): https://sheba.bki.ir/DynamicPage.aspx
//   - Commonly published BIN prefixes from public banking docs / open-source tables
//
// Update date: 2026-11-11
//
// This dataset is incomplete. Unknown codes return ok=false and must not be
// treated as proof of invalidity by card/sheba validators.
package bank

// Bank describes Iranian bank metadata.
type Bank struct {
	Code        string // 3-digit Sheba bank code, e.g. "017"
	Name        string
	PersianName string
}

// ByCode looks up a bank by 3-digit Sheba bank code.
func ByCode(code string) (Bank, bool) {
	if len(code) == 2 {
		code = "0" + code
	}
	b, ok := byCode[code]
	return b, ok
}

// ByShebaBankCode is an alias of ByCode.
func ByShebaBankCode(code string) (Bank, bool) { return ByCode(code) }

// ByBIN looks up a bank by a 6-digit card BIN. Partial coverage only.
func ByBIN(bin string) (Bank, bool) {
	if len(bin) < 6 {
		return Bank{}, false
	}
	code, ok := binToCode[bin[:6]]
	if !ok {
		return Bank{}, false
	}
	return ByCode(code)
}
