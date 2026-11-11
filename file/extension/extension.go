package extension

import (
	"path"
	"strings"

	"github.com/MahdiPezeshkian/onyx-validate/normalize"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func IsValid(filename string, allowed []string) bool {
	return Validate(filename, allowed) == nil
}

func Validate(filename string, allowed []string) error {
	filename = normalize.TrimSpace(filename)
	if filename == "" {
		return validate.Required()
	}
	ext := strings.ToLower(path.Ext(filename))
	if ext == "" {
		return validate.InvalidFormat("filename has no extension")
	}
	for _, a := range allowed {
		a = strings.ToLower(a)
		if !strings.HasPrefix(a, ".") {
			a = "." + a
		}
		if ext == a {
			return nil
		}
	}
	return validate.New(validate.CodeNotAllowed, "file extension is not allowed")
}
