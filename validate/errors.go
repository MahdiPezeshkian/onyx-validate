package validate

import "fmt"

// ValidationError is the shared, machine-readable validation failure type.
// Message is for humans; consumers should prefer Code.
// Values must never include sensitive input such as passwords or full card numbers.
type ValidationError struct {
	Code    string
	Message string
	Field   string
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	if e == nil {
		return ""
	}
	if e.Field != "" {
		return fmt.Sprintf("%s: %s (%s)", e.Field, e.Message, e.Code)
	}
	return fmt.Sprintf("%s (%s)", e.Message, e.Code)
}

// New creates a ValidationError with the given code and message.
func New(code, message string) *ValidationError {
	return &ValidationError{Code: code, Message: message}
}

// Newf creates a ValidationError with a formatted message.
func Newf(code, format string, args ...any) *ValidationError {
	return &ValidationError{Code: code, Message: fmt.Sprintf(format, args...)}
}

// WithField returns a copy of the error with Field set.
func (e *ValidationError) WithField(field string) *ValidationError {
	if e == nil {
		return nil
	}
	cp := *e
	cp.Field = field
	return &cp
}

// Is reports whether target is a ValidationError with the same Code.
func (e *ValidationError) Is(target error) bool {
	t, ok := target.(*ValidationError)
	if !ok || e == nil || t == nil {
		return false
	}
	return e.Code == t.Code
}

// Required returns a required-field validation error.
func Required() *ValidationError {
	return New(CodeRequired, "value is required")
}

// InvalidLength returns an invalid-length validation error.
func InvalidLength(message string) *ValidationError {
	if message == "" {
		message = "value has invalid length"
	}
	return New(CodeInvalidLength, message)
}

// InvalidFormat returns an invalid-format validation error.
func InvalidFormat(message string) *ValidationError {
	if message == "" {
		message = "value has invalid format"
	}
	return New(CodeInvalidFormat, message)
}

// InvalidChecksum returns an invalid-checksum validation error.
func InvalidChecksum(message string) *ValidationError {
	if message == "" {
		message = "value failed checksum validation"
	}
	return New(CodeInvalidChecksum, message)
}

// InvalidCharacter returns an invalid-character validation error.
func InvalidCharacter(message string) *ValidationError {
	if message == "" {
		message = "value contains invalid characters"
	}
	return New(CodeInvalidCharacter, message)
}

// OutOfRange returns an out-of-range validation error.
func OutOfRange(message string) *ValidationError {
	if message == "" {
		message = "value is out of range"
	}
	return New(CodeOutOfRange, message)
}
