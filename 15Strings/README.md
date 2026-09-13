# 15 – Strings

## What You Will Learn
- How Go represents strings and why they are immutable
- The difference between interpreted (`"..."`) and raw (`` `...` ``) string literals
- Why `len(s)` counts bytes, not characters, and how to count characters correctly
- How to iterate over a string's characters safely with `for range`
- Common `strings` package functions for searching, trimming, splitting, and joining
- Converting between `string` and `[]byte`

## Why This Matters
Almost every program touches text: reading user input, parsing config files, building API
responses, formatting log messages. Go's string type is small and simple, but its byte-oriented
nature (not character-oriented) trips up beginners constantly — especially once non-English
text or emoji shows up. Understanding *why* strings behave this way now will save you from
subtle bugs later.

## Concept Explanation
A Go string is an **immutable sequence of bytes**, usually holding UTF-8 encoded text.
Immutable means: once a string is created, its contents can never be changed in place. Any
operation that looks like it "modifies" a string (uppercasing, trimming, replacing) actually
builds and returns a brand-new string, leaving the original untouched.

Why immutable? Two reasons that matter in practice:
- **Safety** — a string can be shared freely between functions and goroutines without anyone
  worrying that one part of the program will unexpectedly change it under another's feet.
- **Efficiency** — because the content never changes, the Go runtime can let multiple strings
  share the same underlying byte data (for example, when you take a substring) instead of
  copying it.

The tricky part is that UTF-8 is a *variable-width* encoding: an ASCII character like `'A'`
takes 1 byte, but many other characters (accented letters, CJK characters, emoji) take 2, 3, or
4 bytes. `len(s)` always reports the byte count. To count actual characters (called **runes** in
Go — a rune is one Unicode code point), use `utf8.RuneCountInString(s)` or a `for range` loop,
which automatically decodes one rune at a time instead of one byte at a time.

## Simple Example
From [main.go](main.go):
```go
emoji := "Go🚀"
fmt.Println("Byte length of Go🚀:", len(emoji))                  // 6 (🚀 = 4 bytes)
fmt.Println("Rune count of Go🚀:", utf8.RuneCountInString(emoji)) // 3 (3 characters)

for i, ch := range "Golang" {
    fmt.Printf("  index %d: %c\n", i, ch) // ch is a rune (Unicode code point)
}
```

## How It Works
`len(emoji)` returns 6 because "Go" is 2 ASCII bytes and the rocket emoji is 4 bytes in UTF-8 —
6 bytes total, even though a human sees 3 characters. `utf8.RuneCountInString` walks the byte
sequence and decodes it correctly, reporting 3. The `for i, ch := range "Golang"` loop does the
same decoding under the hood: `i` is the *byte* index where each rune starts, and `ch` is the
decoded rune (a `rune`, which is really an `int32`), not a raw byte — this is why `range` is the
safe way to walk characters, while `s[i]` just grabs a single byte.

## Common Mistakes
- Assuming `len(s)` gives the number of characters — it gives bytes. For any text that might
  contain non-ASCII characters, use `utf8.RuneCountInString` or `range` instead.
- Indexing a string directly (`s[0]`) to get "the first character" — this only works reliably
  for pure ASCII text; for multi-byte text it can slice into the middle of a character.
- Concatenating strings with `+` inside a loop. Each `+` allocates a new string, so building a
  large string this way is inefficient (many wasted allocations and copies).

## Best Practices
- Use `strings.Builder` when assembling a string piece by piece in a loop — it appends to an
  internal buffer instead of allocating a new string every time, as shown in `main.go`:
  ```go
  var builder strings.Builder
  for i := 0; i < 3; i++ {
      builder.WriteString("Go! ")
  }
  fmt.Println(builder.String())
  ```
- Reach for the standard `strings` package (`Contains`, `TrimSpace`, `Split`, `Join`, etc.)
  before writing manual character-by-character logic — it's tested, readable, and fast.
- Use raw string literals (`` `...` ``) for regex patterns, file paths, and multi-line text so
  you don't have to escape backslashes and quotes.

## Real-World Example
Parsing is everywhere: reading a CSV line, validating an email format, trimming whitespace from
user-submitted form input, or building a JSON string for an HTTP response with `fmt.Sprintf`.
Getting byte-vs-character counting right matters for things like enforcing a "max 280
characters" rule on user-generated text — if you count bytes instead of runes, you'll reject
perfectly valid short messages written in other languages.

## Exercise
Write a function that takes a string and returns how many vowels (`a, e, i, o, u`, upper or
lower case) it contains. Use a `for range` loop so it works correctly on any input, then test it
on both an ASCII word and a word containing an accented or multi-byte character.

## Mini Project
Build a small "text stats" program: given a sentence (a hardcoded string is fine), print its
byte length, its rune (character) count, whether it contains a specific substring, an
uppercase version, and a version with all spaces replaced by underscores. Use only `fmt` and
`strings`/`unicode/utf8` functions already shown in `main.go`.

## Summary
Go strings are immutable, UTF-8 byte sequences. `len()` counts bytes, so use runes (`range` or
`utf8.RuneCountInString`) when you need character counts. Concatenate occasionally with `+`, but
switch to `strings.Builder` for loops. The `strings` package covers most everyday text
operations, and `[]byte(s)` / `string(b)` let you convert to a mutable byte slice when you need
to edit content in place.

## What to Learn Next
Previous: [14Recursion](../14Recursion/README.md)
Next: [16Arrays](../16Arrays/README.md)
