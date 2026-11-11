package validate

// Stable machine-readable validation error codes.
const (
	CodeRequired           = "required"
	CodeInvalidLength      = "invalid_length"
	CodeInvalidFormat      = "invalid_format"
	CodeInvalidChecksum    = "invalid_checksum"
	CodeInvalidCharacter   = "invalid_character"
	CodeOutOfRange         = "out_of_range"
	CodeUnsupportedCountry = "unsupported_country"
	CodeInvalidDate        = "invalid_date"
	CodeInvalidPrecision   = "invalid_precision"
	CodeInvalidPattern     = "invalid_pattern"
	CodeTooWeak            = "too_weak"
	CodeNotAllowed         = "not_allowed"
)
