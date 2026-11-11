package pattern

import (
	"regexp"
	"sync"

	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

// Matcher is a reusable compiled regular expression validator.
type Matcher struct {
	re *regexp.Regexp
}

var cache sync.Map // map[string]*Matcher

func Compile(expr string) (*Matcher, error) {
	if v, ok := cache.Load(expr); ok {
		return v.(*Matcher), nil
	}
	re, err := regexp.Compile(expr)
	if err != nil {
		return nil, validate.New(validate.CodeInvalidPattern, "regular expression is invalid")
	}
	m := &Matcher{re: re}
	actual, _ := cache.LoadOrStore(expr, m)
	return actual.(*Matcher), nil
}

func (m *Matcher) IsValid(value string) bool { return m.Validate(value) == nil }

func (m *Matcher) Validate(value string) error {
	if m == nil || m.re == nil {
		return validate.New(validate.CodeInvalidPattern, "matcher is not initialized")
	}
	if !m.re.MatchString(value) {
		return validate.InvalidFormat("value does not match pattern")
	}
	return nil
}

func IsValid(value, expr string) bool { return Validate(value, expr) == nil }

func Validate(value, expr string) error {
	m, err := Compile(expr)
	if err != nil {
		return err
	}
	return m.Validate(value)
}
