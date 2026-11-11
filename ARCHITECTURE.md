# Architecture

## Package boundaries

`onyx-validate` is a multi-package module. Consumers import leaf packages
directly. The root of the module no longer exports validators; shared
contracts live in `validate/`.

```text
validate/          shared ValidationError + codes
normalize/         explicit string transforms
internal/          luhn, mod97, ascii helpers (not importable by apps)
iran/...           Iran-specific rules
banking|contact|identity|network|text|finance|datetime|security|file
legacy/            preserved v0.1.0 reflection helpers
```

## Dependency graph

```text
validate
   ↑
normalize, identity, banking, contact, network, text,
finance, datetime, security, file, iran/*

iran/banking/sheba  → banking/iban → internal/mod97
iran/banking/card   → banking/card → internal/luhn
iran/banking/card   → iran/banking/bank   (lookup only)
iran/contact/phone  → iran/contact/mobile + landline
network/domain      → network/hostname
```

No cycles. Iranian packages may reuse international algorithms when rules
genuinely overlap (IBAN/Luhn). Lookup datasets stay separate from checksum
code.

## Validation vs normalization vs formatting vs lookup

| Concern | Responsibility | Example |
|---|---|---|
| Validation | Accept/reject | `nationalid.Validate` |
| Normalization | Explicit transform | `normalize.Digits` |
| Formatting | Human-readable output | `sheba.Format` |
| Lookup | Local metadata | `bank.ByBIN` |
| Online verification | **Out of scope** | — |

Validators may normalize a temporary copy according to a documented input
policy; they never mutate the caller's string.

## Shared errors

`validate.ValidationError` carries a stable `Code` and a non-sensitive
`Message`. Prefer codes in application logic. `errors.As` and code-based
`errors.Is` are supported.

## Algorithms and datasets

- Algorithms live next to validators (`nationalid`, `legalid`, Luhn, MOD-97,
  Jalali conversion).
- Datasets live in dedicated files (`banking/iban/lengths.go`,
  `iran/banking/bank/data.go`, `finance/currency/codes.go`) with source and
  update date comments.
- Incomplete datasets must not cause checksum-valid values to fail solely
  because metadata is missing (e.g. unknown BIN).

## Iranian vs international rules

Iran-specific formats and checksums stay under `iran/`. International
packages remain usable without Iranian rules. Overlap is reused, not copied.

## Extension strategy — adding a validator

1. Choose the correct leaf package (or create one with a single clear job).
2. Document input policy (digits, separators, trim).
3. Export `IsValid` + `Validate` (or documented options API).
4. Return `*validate.ValidationError` with stable codes.
5. Add table-driven tests with independent known vectors.
6. Update ROADMAP/CHANGELOG.
7. Do not invent checksums or datasets.

## Legacy package

`legacy` preserves the original reflection helpers from v0.1.0 for
compatibility. New code should use modular packages.
