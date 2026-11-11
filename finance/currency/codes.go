package currency

// codes is a partial set of commonly used ISO 4217 alphabetic codes.
// Source: ISO 4217 lists / public currency tables. Updated: 2026-11-11.
var codes = map[string]bool{
	"USD": true, "EUR": true, "GBP": true, "JPY": true, "CHF": true, "CAD": true,
	"AUD": true, "NZD": true, "CNY": true, "HKD": true, "SGD": true, "INR": true,
	"IRR": true, "IQD": true, "AED": true, "SAR": true, "TRY": true, "RUB": true,
	"BRL": true, "MXN": true, "ZAR": true, "SEK": true, "NOK": true, "DKK": true,
	"PLN": true, "CZK": true, "HUF": true, "ILS": true, "KRW": true, "THB": true,
	"MYR": true, "IDR": true, "PHP": true, "VND": true, "PKR": true, "EGP": true,
	"XAU": true, "XAG": true, "XDR": true,
}
