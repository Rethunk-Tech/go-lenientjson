# AGENTS.md

Package `lenientjson` rewrites the JSON that Json.NET reads into strict JSON for `encoding/json`. It exports `Clean` only and imports the standard library only; it never imports a consumer package.

## Gate

`gate` is the command. Each override in `.gate.toml` is there because a runner
cannot resolve that check from PATH.

## Shape

- Json.NET's reader is the behaviour target, not a spec. Where its result depends on the target type (hex, octal, a bare point under a string field) `Clean` does not guess: a rewrite is written back into config files by some callers and must keep the value exactly.
- `Clean` is `decode`, `normalise`, `dropTrailingCommas` in that order. Comment and string handling live in `normalise`; container-aware comma handling needs the cleaned text, so it runs second.
- Newlines are never dropped (comments keep theirs; a raw newline in a string is paid back after the closing quote), because callers derive a source line from a `*json.SyntaxError` over the output.
- A rejected input stays invalid rather than being repaired: an unclosed `/*` is emitted as a bare `/*`, and a backslash before a raw control character is passed through, so `encoding/json` refuses both as Json.NET does.
- Everything stays unexported; the single entry point keeps callers independent of the pass order.

## Files

| File | Holds |
| --- | --- |
| `lenientjson.go` | `Clean` and its passes: BOM decoding, token rewriting, comma and elision handling |
| `doc.go` | Package documentation |
