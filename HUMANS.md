# HUMANS.md

Run and use `go-lenientjson`. Internals: [AGENTS.md](AGENTS.md).

## Quick start

```bash
go get github.com/Rethunk-Tech/go-lenientjson
```

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/Rethunk-Tech/go-lenientjson"
)

func main() {
	raw := []byte("{\n  // comment\n  name: 'demo', tags: [1, 2,],\n}")
	var v struct {
		Name string `json:"name"`
		Tags []int  `json:"tags"`
	}
	if err := json.Unmarshal(lenientjson.Clean(raw), &v); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v.Name, v.Tags)
}
```

## Contracts

| Contract | Behaviour |
| --- | --- |
| `Clean(b) []byte` | Never fails and never panics; the result may still be invalid, and `encoding/json` reports that with an offset. |
| Rejected input | Text Json.NET rejects is left invalid: an unclosed `/*`, `{,}`, `{"o":1,,"s":"x"}`, `+1`, `0x1F`, `012`, a backslash before a raw control character. |
| Strict JSON | Unmarshals to the same value before and after `Clean`. |
| Newlines | Every newline of the source, including those inside comments and raw-newline strings, is in the output, so the line of a `*json.SyntaxError` counted over the output is the source line. |
| Strings | Contents are never rewritten beyond escaping raw control characters and `'`-quoting. |
| `NaN`, `Infinity`, `undefined` | Become `null`, also under a string target where Json.NET would keep the text. |
| Bare decimal point | `.5` becomes `0.5` and `5.` becomes `5.0`, right for a numeric target only. |
| Input slice | Not modified; the result is a new slice. |

## Requirements

Go 1.27.3 or newer, plus `lefthook` and `gitleaks` on PATH if you want the git hooks.

```bash
lefthook install
```

That is per clone; git does not commit hooks. Pre-commit scans staged changes for secrets, and
pre-push runs `gate`.

## Verify

```bash
gate
```
