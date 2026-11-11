# onyx-validate

[![Go Version](https://img.shields.io/badge/go-1.27-00ADD8.svg)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Reusable reflection-based validation helpers. Extracted from
`Onyx.Common/validation`.

## Installation

```bash
go get github.com/MahdiPezeshkian/onyx-validate@v0.1.0
```

## Import path

```go
import "github.com/MahdiPezeshkian/onyx-validate"
```

Package name: `validate`

## Basic usage

```go
if err := validate.ValidateNotEmpty(email, password); err != nil {
	return err
}
if err := validate.ValidateEmail(email); err != nil {
	return err
}
return validate.ValidateRange(18, 120, age)
```

## Public API overview

| API | Purpose |
|---|---|
| `ValidateNotEmpty` | Reject zero/empty values |
| `ValidateNonNegative` / `ValidatePositive` | Numeric sign checks |
| `ValidateMin` / `ValidateMax` / `ValidateRange` | Numeric bounds |
| `ValidateLength` | Exact length for string/slice/array/map |
| `ValidateEmail` | Lightweight email shape check (`@` and `.`) |

## Design decisions

- Independent of HTTP frameworks and application models.
- Uses reflection for flexible variadic helpers, matching the original package.
- Email validation is intentionally simple (not RFC 5322).

## Dependencies

None beyond the Go standard library.

## Limitations

- Not a struct-tag validation framework.
- Email checks are heuristic.
- No localization of error messages.

## License

MIT — see [LICENSE](LICENSE).
