// Package validate provides shared validation error types and stable error codes
// used across Onyx.Validate packages.
//
// Import path: github.com/MahdiPezeshkian/onyx-validate/validate
package validate

// Result is an optional structured outcome for callers that prefer a value
// object over error-only APIs. Most validators return error directly.
//
// Invariants:
//   - OK() always has Valid=true and Err=nil
//   - Fail(err) always has Valid=false and a non-nil Err
//   - Fail(nil) uses a generic InvalidFormat error instead of an inconsistent state
type Result struct {
	Valid bool
	Err   *ValidationError
}

// OK returns a successful Result.
func OK() Result {
	return Result{Valid: true, Err: nil}
}

// Fail returns an unsuccessful Result wrapping err.
// If err is nil, Fail records InvalidFormat("validation failed") so Valid=false
// never pairs with a nil Err.
func Fail(err *ValidationError) Result {
	if err == nil {
		err = InvalidFormat("validation failed")
	}
	return Result{Valid: false, Err: err}
}

// Error returns nil for a valid Result, otherwise the validation error.
func (r Result) Error() error {
	if r.Valid {
		return nil
	}
	if r.Err != nil {
		return r.Err
	}
	return InvalidFormat("validation failed")
}
