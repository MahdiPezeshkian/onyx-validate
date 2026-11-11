// Package account is reserved for bank-specific Iranian account-number rules.
//
// Deferred: Iranian banks do not share a single account-number format or
// checksum. No generic validator is provided to avoid inventing rules.
package account

import "github.com/MahdiPezeshkian/onyx-validate/validate"

// Validate is intentionally unimplemented for generic accounts.
// It always returns an unsupported error until bank-specific rules are added.
func Validate(bankCode, account string) error {
	_ = bankCode
	_ = account
	return validate.New(validate.CodeUnsupportedCountry, "Iranian account validation is deferred; bank-specific rules are required")
}

// IsValid always returns false until bank-specific validators are implemented.
func IsValid(bankCode, account string) bool {
	return Validate(bankCode, account) == nil
}
