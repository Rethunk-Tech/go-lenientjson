# Changelog

Notable changes to `go-lenientjson`. The format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and versions follow
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed

- The lint script removes the downloaded `actionlint` before running it.

## [0.2.0] - 2026-11-12

### Fixed

- `Clean` treats Json.NET whitespace as JSON space, skips a NUL outside a string, and ends `//` comments on CR as well as LF.
- An unclosed block comment is emitted as a bare `/*`, and a comment before a property colon stays invalid, as in Json.NET.
- Bare words stop at a Unicode space, and a property colon is looked for past comments.
- Unquoted property names are quoted when they start with a digit, and only Json.NET identifier characters are quoted in a bare name.
- A UTF-16 or UTF-32 input ending mid code unit is an error, not silently truncated.

## [0.1.1] - 2026-11-07

### Fixed

- `Clean` leaves a signed `NaN` or `undefined` invalid, as Json.NET does.

## [0.1.0] - 2026-10-29

### Added

- `Clean` rewrites the bytes Json.NET reads (comments, trailing commas and elisions, single quotes, unquoted names, `NaN`/`Infinity`/`undefined`, bare decimal points, raw control characters, UTF-8/16/32 BOMs) as strict JSON for `encoding/json`, keeping source line numbers.

[Unreleased]: https://github.com/Rethunk-Tech/go-lenientjson/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/Rethunk-Tech/go-lenientjson/compare/v0.1.1...v0.2.0
[0.1.1]: https://github.com/Rethunk-Tech/go-lenientjson/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/Rethunk-Tech/go-lenientjson/releases/tag/v0.1.0
