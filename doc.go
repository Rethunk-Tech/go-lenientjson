// Package lenientjson turns the JSON that Json.NET (Newtonsoft.Json) reads into strict JSON for encoding/json: it
// decodes a UTF-8, UTF-16 or UTF-32 BOM and strips comments and trailing commas; accepts single-quoted strings, unquoted
// property names, bare decimal points and NaN/Infinity/undefined; and escapes control characters inside strings.
package lenientjson
