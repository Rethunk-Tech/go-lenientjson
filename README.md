<h1 align="center">go-lenientjson</h1>

<div align="center">

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go](https://img.shields.io/badge/go-1.27.3+-00ADD8?logo=go&logoColor=white)](go.mod)

</div>

---

Turns the lenient JSON that Content Patcher and Json.NET accept (comments, trailing commas, loose numbers) into strict JSON for encoding/json. Package `lenientjson`, module `github.com/Rethunk-Tech/go-lenientjson`.

One function, `Clean`, takes the bytes a .NET program would hand to Json.NET (Newtonsoft.Json) and returns strict JSON that `encoding/json` reads to the same values. The target is Json.NET's reader, not JSONC or JSON5. A caller that needs the cleaned bytes (for `json.Decoder`, a line-numbered error) gets them; one that does not writes `json.Unmarshal(lenientjson.Clean(b), &v)`.

## Quick start

```bash
go get github.com/Rethunk-Tech/go-lenientjson
```

A worked example and the contracts: [HUMANS.md](HUMANS.md).

## Features

- Reads the encoding `File.ReadAllText` picks from the BOM: UTF-8, UTF-16 LE/BE, UTF-32 LE/BE.
- Strips `//` and `/* */` comments and trailing commas; an elided array element or empty property value (`[1,,2]`, `{"o":,}`) becomes `null`.
- Accepts single-quoted strings, unquoted property names, `.5` and `5.`, and raw control characters in strings.
- `NaN`, `Infinity`, `-Infinity` and `undefined` become `null`, because `encoding/json` has no NaN.
- Keeps every newline, so a `*json.SyntaxError`'s line over the output is the line in the source. Columns are not kept.
- Known gap: hex (`0x1F`) and octal (`012`) integers stay invalid; Json.NET reads them, but a rewrite depends on the target type.

## Documentation

| Doc | Audience |
| --- | --- |
| [HUMANS.md](HUMANS.md) | Running and using the package |
| [AGENTS.md](AGENTS.md) | Internals, invariants, and why |
| [CONTRIBUTING.md](CONTRIBUTING.md) | Process — commits, tests, docs |
| [SECURITY.md](SECURITY.md) | Vulnerability disclosure |
| [CHANGELOG.md](CHANGELOG.md) | Release notes |

## License

MIT — see [LICENSE](LICENSE). © 2026 Rethunk.Tech, LLC.
