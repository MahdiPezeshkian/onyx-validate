// Package validate provides shared validation error types and stable error codes
// used across Onyx.Validate packages.
//
// Import path: github.com/MahdiPezeshkian/onyx-validate/validate
package validate

// Result is an optional structured outcome for callers that prefer a value
// object over error-only APIs. Most validators return error directly.
type Result struct {
	Valid bool
	Err   *ValidationError
}

// OK returns a successful Result.
func OK() Result {
	return Result{Valid: true}
}

// Fail returns an unsuccessful Result wrapping err.
func Fail(err *ValidationError) Result {
	return Result{Valid: false, Err: err}
}
