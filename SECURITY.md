# Security

## Supported versions

The latest tagged release. There are no long-lived release branches, so a fix ships in the next tag
rather than being backported.

## Reporting a vulnerability

Use GitHub's **private vulnerability reporting** on this repository (Security → Report a
vulnerability). It stays private until a fix ships.

Please do not open a public issue for a suspected vulnerability first.

## What is in scope

- `Clean(b)` takes arbitrary bytes, usually a file a third party wrote. For any input it must not panic, must not index out of range on a truncated string, comment or UTF-16/32 sequence, and must allocate in proportion to the input (the output is at most a constant factor larger than the input, plus the newlines it keeps).

## What is out of scope

- Whether the cleaned bytes are valid and what they mean: that is `encoding/json`'s, and the caller must check its error.
- Json.NET behaviour that is deliberately not matched (hex and octal integers, string-target readings of `.5` or `NaN`): see the README.
- Resource limits on very large inputs: the caller bounds the size it reads.
