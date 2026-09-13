# 21 – Maps

## What You Will Learn
- What a map is and when to reach for one instead of a slice
- Creating maps with literals and `make`, and why a `nil` map is dangerous
- Adding, updating, and reading entries
- The comma-ok idiom for checking whether a key exists
- Deleting entries with `delete`
- Why map iteration order is random, and what that implies
- Maps are reference types, like slices
- Storing structs as map values, and updating a struct field inside one

## Why This Matters
A slice is great for an ordered list you access by position. But often you actually need a *lookup*: "given this user's ID, get their record," or "given this word, get its count." Scanning a slice from the start every time is slow and awkward. A map gives you that lookup directly by key — a dictionary or hash table — in roughly constant time regardless of size.

## Concept Explanation
A map's type is `map[KeyType]ValueType`. Keys must be *comparable* — strings, numbers, booleans, and structs of comparable fields work; slices and maps don't, since Go can't check two of them for equality.

Two ways to create a usable map:
```go
capitals := map[string]string{"USA": "Washington D.C."} // literal
ages := make(map[string]int)                            // empty, ready to use
```
A map declared with just `var m map[string]int` (no literal, no `make`) is a **nil map**. Reading from it is safe (you get the zero value back), but *writing* to it panics at runtime. Always initialize with `make` or a literal before writing.

Reading a missing key doesn't panic either — it silently returns the value type's zero value. That's convenient, but it means a plain lookup can't tell "key missing" apart from "key exists and happens to be zero." The **comma-ok idiom** solves this: `value, ok := m[key]` — `ok` is true only if the key actually exists.

Like slices, maps are **reference types** — assigning one to another variable doesn't copy its contents; both point at the same data.

Map iteration order is **intentionally randomized** by Go on every run — a deliberate choice, not an oversight, so programs never accidentally depend on an order that was never guaranteed. If you need predictable output, collect the keys into a slice and sort that.

## Simple Example
From [main.go](main.go):

```go
ages := make(map[string]int)
ages["Alice"] = 30

age, ok := ages["Alice"]  // age=30, ok=true
age, ok = ages["Nobody"]  // age=0,  ok=false

delete(ages, "Charlie")
```

## How It Works
- `make(map[string]int)` allocates an empty, ready-to-use map — safe to write to, unlike a `nil` map.
- `ages["Alice"] = 30` both adds and updates — the syntax is identical either way.
- `age, ok := ages["Alice"]` tells you definitively whether `"Alice"` is present, which a plain lookup can't.
- `delete(ages, "Charlie")` removes the entry; deleting a missing key is a safe no-op.
- The `students := map[string]Student{...}` section holds structs as values. Since map values aren't addressable, `students["Bob"].Score = 90` won't compile — instead the file copies the struct out (`s := students["Bob"]`), edits the copy, and writes the whole struct back (`students["Bob"] = s`).

## Common Mistakes
- **Writing to a nil map.** `var m map[string]int; m["x"] = 1` panics at runtime. Always `make` or literal-initialize first (reading is fine).
- **Assuming any iteration order matches insertion.** Go randomizes it deliberately; sort keys yourself if order matters.
- **Using a plain lookup to check existence** — `if ages["Nobody"] != 0` can't distinguish missing from zero. Use comma-ok: `if _, ok := ages["Nobody"]; ok { ... }`.
- **Modifying a struct field through a map directly** — `students["Bob"].Score = 90` fails to compile since map values aren't addressable; read, modify, reassign.

## Best Practices
- Always initialize maps with `make` or a literal before writing.
- Use comma-ok (`v, ok := m[k]`) whenever "does this key exist" matters.
- Never design around map iteration order; sort explicitly if needed.
- Copy key-by-key if you need an independent copy — maps are reference types.

## Real-World Example
Maps fit lookup tables and caches naturally: user ID → profile, SKU → price, status code → message, or memoizing expensive results so repeat calls are instant.

## Exercise
Write `countLetters(s string) map[rune]int` counting how many times each rune appears in a string. Print the result, then look up one specific letter's count using comma-ok.

## Mini Project
Build a word-frequency counter: given a slice of words, use a `map[string]int` to tally occurrences, then print each word with its count. Include a lookup for a word not present to demonstrate comma-ok returning `false`.

## Summary
Maps store key-value pairs for fast lookup by key instead of position. A map must be initialized with `make` or a literal before writing — writing to `nil` panics, though reading is safe. The comma-ok idiom is the correct way to distinguish a missing key from a zero value. Maps are reference types, and iteration order is deliberately randomized.

## What to Learn Next
Previous: [20 Range](../20Range/README.md)

You've used interfaces implicitly already — next we make that concept explicit: how Go lets different types share behavior without any `implements` keyword.

Next: [22 Interfaces](../22Interfaces/README.md)
