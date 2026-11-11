# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed
- `validate.Fail(nil)` no longer produces Valid=false with a nil Err.


### Added

- Modular package architecture for Iranian and international validators.
- Shared `validate` error model with stable codes.
- `normalize` helpers for Persian/Arabic digits and whitespace.
- Iranian national ID, legal ID, passport validators.
- Iranian banking: card (Luhn), Sheba (IBAN), partial bank/BIN lookup.
- Iranian contact (mobile/landline/phone), postal code, vehicle helpers.
- International banking (card/IBAN/BIC), contact (email/E.164), identity
  (UUID/ULID/username), network, text, finance, datetime (incl. Jalali),
  security, and file packages.
- Fuzz tests for national ID and IBAN; benchmarks for hot paths.
- ARCHITECTURE.md and ROADMAP.md.

### Changed

- Root package reflection helpers moved to `legacy` to free `validate/` for
  shared error contracts.

### Notes

- Structural validity does not imply real-world issuance or activation.
- Some datasets are partial by design; see package docs and ROADMAP.

## [0.1.0] - 2026-11-11

### Added

- Initial extraction of reflection-based validation helpers from
  `Onyx.Common/validation`.

[Unreleased]: https://github.com/MahdiPezeshkian/onyx-validate/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/MahdiPezeshkian/onyx-validate/releases/tag/v0.1.0
