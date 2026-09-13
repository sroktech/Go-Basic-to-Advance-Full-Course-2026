# 33 – JSON

## What You Will Learn
- Encoding Go values to JSON with `json.Marshal`/`json.MarshalIndent`
- Decoding JSON into Go structs and maps with `json.Unmarshal`
- Struct tags: `json:"name"`, `json:"name,omitempty"`, `json:"-"`
- Decoding into `map[string]interface{}` for JSON of unknown shape
- Streaming JSON with `json.Encoder`/`json.Decoder`
- Writing custom `MarshalJSON`/`UnmarshalJSON` methods

## Why This Matters
JSON is the lingua franca of web APIs — nearly every HTTP service you'll build or consume speaks it. Go's `encoding/json` package uses reflection and struct tags rather than a separate schema/IDL, so the same struct you already use in your business logic can double as your wire format, with tags telling the encoder exactly how to shape the output. Understanding *why* the tags exist — and their subtle rules around casing, `omitempty`, and exported fields — is what separates code that "seems to work" from code that produces correct API payloads under every input, including empty and partial ones.

## Concept Explanation
Marshal turns a Go value into JSON bytes; Unmarshal does the reverse:

```
Go struct/map ──► json.Marshal() ──► []byte (JSON string)
[]byte (JSON string) ──► json.Unmarshal() ──► Go struct/map
```

By default, a struct field's JSON key is its Go field name verbatim — which is almost never what an API expects (JSON convention is `camelCase` or `snake_case`, Go convention is `PascalCase` for exported fields). Struct tags fix the mismatch:

- `` `json:"name"` `` — always use `"name"` as the key.
- `` `json:"name,omitempty"` `` — use `"name"`, but drop the field entirely from the output if it holds its zero value (empty string, 0, nil, empty slice, etc.).
- `` `json:"-"` `` — never include this field in JSON output, regardless of its value.

`omitempty` and `json:"-"` solve different problems: `omitempty` is about *optional* data (a field that's sometimes absent is fine to leave out of JSON when unset), while `json:"-"` is about data that must **never** leave the process — a password hash, an internal-only flag — no matter what value it holds.

Only **exported** (capitalized) fields participate in JSON at all. `encoding/json` uses reflection, and reflection cannot see unexported fields from another package — so a lowercase field is silently skipped, with no error, which is one of the most common early Go/JSON bugs.

## Simple Example
The `Person` struct in [main.go](main.go) shows all three tag forms together:

```go
type Person struct {
    Name     string   `json:"name"`
    Age      int      `json:"age"`
    Email    string   `json:"email,omitempty"` // omit if empty string
    Password string   `json:"-"`               // NEVER include in JSON
    Address  Address  `json:"address"`
    Tags     []string `json:"tags,omitempty"`   // omit if nil/empty
}
```

## How It Works
When `person` (with an `Email` and `Password` set) is marshaled, `Password` never appears in the output — `json:"-"` guarantees that regardless of value. When a second `Person` is built with no `Email` set, `Email` is dropped from its JSON entirely because of `omitempty`, giving a smaller, cleaner payload instead of `"email":""`.

Decoding is the mirror image: `json.Unmarshal([]byte(jsonInput), &decoded)` fills a `Person` from a JSON string. Extra keys in the input that don't match any struct field (like `unknownField` in the sample) are silently ignored — Go does not error on unknown fields by default, which is convenient for forward compatibility with evolving APIs but means typos in your tags fail silently too (the field just won't populate).

When the shape of incoming JSON isn't known ahead of time, decode into `map[string]interface{}` instead of a struct — every value comes back as `interface{}`, and JSON numbers decode as `float64` by default, so a type assertion (`dynamic["x"].(float64)`) is required before use.

For large payloads, `json.NewEncoder(w)`/`json.NewDecoder(r)` stream directly to/from an `io.Writer`/`io.Reader` (a file, an HTTP body, a socket) instead of building the whole `[]byte` in memory first — this is exactly what the HTTP handlers in the next lesson use.

## Common Mistakes
- Forgetting struct tags entirely — the JSON output then uses Go's exact field names (`Name`, `Age`), which rarely matches what a real API client expects.
- Confusing `omitempty` (drop when zero-valued) with `json:"-"` (drop always) — using the wrong one either leaks a sensitive field when it happens to be empty, or drops a legitimate zero value (like an `Age` of `0`) that the API actually needed to see.
- Using lowercase (unexported) struct fields and being confused why they never appear in the JSON output — `encoding/json` can't see them at all.
- Assuming numbers decoded into `interface{}` are `int` — they are always `float64` unless you use `json.Decoder` with `UseNumber()`.

## Best Practices
- Tag every struct field you intend to serialize, even when the name would happen to match, so the JSON shape is explicit and doesn't silently break if the Go field is renamed later.
- Use `omitempty` for genuinely optional API fields; use `json:"-"` for anything that must never be serialized, security-sensitive or otherwise.
- Prefer typed structs over `map[string]interface{}` whenever the shape is known — you get compile-time safety and no type assertions.
- For big data or network streams, use `Encoder`/`Decoder` instead of `Marshal`/`Unmarshal` to avoid holding the entire payload in memory.

## Real-World Example
A typical REST API layer defines request/response structs with JSON tags matching the public API contract, decodes incoming request bodies with `json.NewDecoder(r.Body).Decode(&req)`, and encodes responses the same way — exactly the pattern the HTTP lesson builds on. A config loader reads a JSON (or YAML) file and unmarshals it into a typed `Config` struct at startup, failing fast if required fields are missing.

## Exercise
Define a struct for a "blog post" (title, body, author, an internal `viewCount` that should never be exposed, and an optional `publishedAt` that should be omitted when unset). Marshal a sample value, unmarshal it back, and print both to confirm the sensitive field never appears and the optional field is omitted correctly.

## Mini Project
Build a small JSON config loader: define a `Config` struct (e.g., `Port int`, `Host string`, `Debug bool`, `Tags []string`), write a sample `config.json` file to disk (using what you learned in the File I/O lesson), then read and unmarshal it into `Config` and print the loaded values. Handle and report the error clearly if the file is missing or the JSON is malformed.

## What to Learn Next
Continue to [34 – HTTP](../34HTTP/README.md), where the JSON encoding/decoding you just learned becomes the payload format for a real HTTP server and client.
