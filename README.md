# onyx-validate

[![Go Version](https://img.shields.io/badge/go-1.27-00ADD8.svg)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A modular Go validation toolkit for Iranian and international data.

Import only the packages you need. Validation, normalization, formatting, and
lookup are separate concerns. Online verification is out of scope.

## Installation

```bash
go get github.com/MahdiPezeshkian/onyx-validate@latest
```

Module path: `github.com/MahdiPezeshkian/onyx-validate`

## Quick start

```go
import (
    "github.com/MahdiPezeshkian/onyx-validate/iran/identity/nationalid"
    "github.com/MahdiPezeshkian/onyx-validate/normalize"
    "github.com/MahdiPezeshkian/onyx-validate/validate"
)

func check(id string) error {
    // Explicit normalization when you want ASCII digits:
    _ = normalize.Digits(id)

    if err := nationalid.Validate(id); err != nil {
        var ve *validate.ValidationError
        if errors.As(err, &ve) {
            // Prefer ve.Code over ve.Message.
            return err
        }
        return err
    }
    return nil
}
```

## API conventions

Most validators expose:

- `IsValid(value string) bool`
- `Validate(value string) error`

Parameterized validators take options or extra arguments:

```go
username.Validate(name, username.DefaultOptions())
age.ValidateMin(birth, referenceDate, 18)
```

Shared errors live in package `validate` and support `errors.As` / `errors.Is`
by stable machine-readable codes.

## Package map

| Area | Import path prefix | Notes |
|---|---|---|
| Shared errors | `.../validate` | Codes + `ValidationError` |
| Normalization | `.../normalize` | Digits, whitespace, separators |
| Iranian identity | `.../iran/identity/...` | National ID, legal ID, passport |
| Iranian banking | `.../iran/banking/...` | Card, Sheba, bank lookup |
| Iranian contact | `.../iran/contact/...` | Mobile, landline, dispatcher |
| Iranian address | `.../iran/address/postalcode` | 10-digit structural |
| Iranian vehicle | `.../iran/vehicle/...` | Plate, VIN, engine |
| International banking | `.../banking/...` | Card (Luhn), IBAN, BIC |
| International contact | `.../contact/...` | Email, E.164 phone |
| Identifiers | `.../identity/...` | UUID, ULID, username |
| Network | `.../network/...` | IP, CIDR, hostname, domain, MAC, port |
| Text | `.../text/...` | Length, charset, format, pattern |
| Finance | `.../finance/...` | Amount, currency, percentage |
| Date/time | `.../datetime/...` | Date, range, age, Persian calendar |
| Security | `.../security/...` | Password policy, PIN |
| File | `.../file/...` | Extension, MIME, filename, size |
| Legacy | `.../legacy` | Original v0.1.0 reflection helpers |

## Normalization policy

- Validators never mutate caller-owned strings.
- Persian and Arabic-Indic digits may be accepted by Iranian validators after
  an **internal copy** conversion; each package documents its policy.
- Use `normalize.Digits`, `normalize.TrimSpace`, etc. when you need explicit
  transformation of your own values.

## Error handling

```go
err := nationalid.Validate(input)
var ve *validate.ValidationError
if errors.As(err, &ve) {
    switch ve.Code {
    case validate.CodeRequired:
    case validate.CodeInvalidChecksum:
    }
}
```

Error messages never include passwords, PINs, or full sensitive secrets.

## Result helper

Package `validate` exposes an optional `Result` value object. Prefer ordinary
`error` returns from leaf validators. `Fail(nil)` is normalized to an
`invalid_format` error so `Valid=false` never pairs with a nil `Err`.

## Dependency policy

- Prefer the Go standard library.
- No web frameworks, databases, or `onyx-errors`.
- No network I/O in ordinary validators.
- Current direct dependencies: **none** beyond the Go standard library.

## Limitations

- Structural validity ≠ real-world existence, ownership, or activation.
- Bank BIN / bank-code / currency / IBAN-country datasets are **partial** and
  documented with sources and update dates.
- Iranian account-number validation is deferred (no shared bank algorithm).
- International phone validation is E.164 structural only (no numbering-plan
  metadata library).
- Email validation is pragmatic, not full RFC 5322 / IDNA.
- Vehicle free-zone / diplomatic plates are deferred.

See [ARCHITECTURE.md](ARCHITECTURE.md) and [ROADMAP.md](ROADMAP.md).

## License

MIT — see [LICENSE](LICENSE).
