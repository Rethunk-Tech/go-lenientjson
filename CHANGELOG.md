# Changelog

Notable changes to `go-lenientjson`. The format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and versions follow
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- `Clean` rewrites the bytes Json.NET reads (comments, trailing commas and elisions, single quotes, unquoted names, `NaN`/`Infinity`/`undefined`, bare decimal points, raw control characters, UTF-8/16/32 BOMs) as strict JSON for `encoding/json`, keeping source line numbers.

[Unreleased]: https://github.com/Rethunk-Tech/go-lenientjson/commits/main
