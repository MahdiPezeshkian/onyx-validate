# Roadmap

## Implemented

- Shared `validate` error model and codes
- `normalize` (digits, whitespace, banking separators, case helpers)
- Iranian national ID (checksum + repeated-digit rejection)
- Iranian legal-entity ID (documented checksum; known vector `10380284790`)
- Iranian ordinary passport structural form (`Letter` + 8 digits)
- Iranian bank card (16-digit Luhn) + optional BIN lookup
- Iranian Sheba (IR IBAN via shared MOD-97)
- Partial Iranian bank code / BIN metadata
- Iranian mobile / landline / phone dispatcher
- Iranian postal code (10-digit structural)
- Passenger plate structural validation (selected letter set)
- VIN structural (+ optional NA check digit)
- Engine number conservative structural checks
- Generic card (Luhn), IBAN (+ country lengths), BIC
- Email (pragmatic), international E.164 phone
- UUID, ULID, configurable username
- Network: IP, CIDR, hostname, domain, MAC, port
- Text: length (runes/bytes), charset, format, pattern
- Finance: amount (no float math), currency subset, percentage
- Date/range/age + Jalali calendar conversion/validation
- Password policy + PIN policy
- File extension/MIME/filename/size
- Legacy reflection helpers preserved under `legacy/`

## Partially implemented

- Iranian bank metadata / BIN tables (explicitly partial)
- IBAN country length registry (major countries; not every territory)
- ISO 4217 currency subset (common codes)
- Iranian plates (passenger only; other categories deferred)
- International phone (E.164 only; no national numbering plans)
- Email (ASCII domains; limited Unicode/IDNA)

## Deferred

| Feature | Reason | Needed to complete |
|---|---|---|
| Generic Iranian account numbers | No single shared checksum/format across banks | Per-bank authoritative rules |
| Free-zone / diplomatic / temporary plates | Specs not fully verified here | Official plate category specs + samples |
| Full passport categories | Only ordinary booklet form documented | Official category formats |
| Operator assignment for mobiles | Avoid stale hardcoding | Maintained operator prefix dataset + license |
| Postal code geographic restrictions | No reliable offline rules adopted | Authoritative postal dataset |
| libphonenumber-style phone metadata | Large dependency; not required for E.164 | Evaluated dependency + update process |
| Weak-password dictionary | Licensing/update concerns | Trusted list + license clearance |
| Content-based MIME sniffing beyond allow-lists | Safety claims would be misleading | Clear API boundaries + stdlib sniffing helpers |

## Planned

- Expand bank BIN coverage with cited sources
- Expand IBAN country registry to full SWIFT registry snapshot
- Optional compiled username/password policy presets
- Additional Jalali formatting locales
- More fuzz targets for parsers
