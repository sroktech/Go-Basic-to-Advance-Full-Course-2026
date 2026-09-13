# 20 – Range

## What You Will Learn
- What `range` is and why it's the idiomatic way to iterate in Go
- Ranging over slices, arrays, strings, and (briefly) maps
- Using `_` to discard the index or the value you don't need
- Why the value `range` gives you is a *copy*, and how to modify elements in place instead
- How ranging over a string differs from indexing it (byte index vs. rune)
- Building a new slice by filtering while ranging

## Why This Matters
Every language lets you loop with a manual counter (`for i := 0; i < len(s); i++`), but that's rarely what you want to write by hand — bounds are easy to get wrong, and it doesn't read as clearly as "for each item, do this." `range` expresses intent directly and works uniformly across very different data structures — slices, arrays, strings, maps — so once you know `range`, you know how to iterate over almost anything in Go.

## Concept Explanation
`range` pairs with `for` to walk a collection, producing two values each step: `for index, value := range collection { }`. What those mean depends on the collection: a **slice or array** gives index + element copy; a **string** gives byte index + `rune` (a Unicode code point, not a byte); a **map** gives key + value (order is *random* — more in the next lesson). Use `_` to discard whichever value you don't need: `for _, v := range s` skips the index, `for i := range s` skips the value.

A crucial detail: `value` is a **copy** of the element, not a reference. Assigning to it inside the loop changes only the local copy — the original is untouched. To modify the original, index back in: `collection[index] = ...`.

Ranging over a `string` decodes it as UTF-8, giving you each rune (a full character, 1–4 bytes) plus the *byte* index where it starts — not a sequential character count. That's why the byte index can jump by more than 1 when a multi-byte character (like an emoji) appears.

This lesson also ranges over a map to show the pattern looks identical to a slice. We'll cover maps properly (declaring, checking missing keys) in the next lesson — for now just notice `range` works the same way, except the order is random.

## Simple Example
From [main.go](main.go):

```go
fruits := []string{"apple", "banana", "cherry", "date"}
for i, fruit := range fruits {
    fmt.Printf("  [%d] = %s\n", i, fruit)
}

word := "Go🚀"
for i, ch := range word {
    fmt.Printf("  byte index %d: %c (Unicode: U+%04X)\n", i, ch, ch)
}
```

## How It Works
- `for i, fruit := range fruits` gives the index and a copy of each element, in storage order.
- `for i, ch := range word` decodes `"Go🚀"` rune by rune: `'G'` at byte 0, `'o'` at byte 1, `🚀` at byte 2 — the next rune after that lands at byte 6, since the emoji is 4 bytes in UTF-8. Direct indexing (`word[2]`) would instead give one raw byte, not the whole character.
- The "wrong way" block, `for _, v := range numbers { v *= 2 }`, only changes the loop's local copy, so `numbers` stays unchanged. The "correct way," `for i := range numbers { numbers[i] *= 2 }`, indexes back in to mutate the real element.

## Common Mistakes
- **Expecting `value` to be a reference** — it's a copy; use the index to write back if you need to change elements.
- **Modifying a slice's length while ranging over it** (e.g. appending to it inside the loop) — length is evaluated once at the start, so behavior gets confusing; build a *new* slice instead.
- **Modifying a map while ranging over it** — adding/deleting keys mid-iteration is unsafe in general (deleting the *current* key is fine); collect changes and apply them after.
- **Forgetting `_`** and getting an "unused variable" error when you only need one of the two values.
- **Closure-capture surprises in older Go code.** Before Go 1.22, every iteration reused the same loop variables, so a closure created inside the loop could capture the final value instead of its own iteration's. Go 1.22+ (this module's target) gives each iteration a fresh copy, fixing that class of bug — worth knowing if you read older Go code.

## Best Practices
- Prefer `range` over manual index-counting whenever you're simply visiting every element.
- Use `_` explicitly to document "intentionally ignoring this value."
- Index back into the collection (`s[i] = ...`) to mutate elements, rather than relying on the range value.
- Append into a fresh slice when filtering/transforming, rather than editing the source in place.

## Real-World Example
`range` shows up everywhere: iterating database result rows, safely processing user input byte-by-byte or rune-by-rune (handling non-English text and emoji correctly), walking HTTP header slices, or summing report values.

## Exercise
Write a function that takes a `[]int` and returns a new slice containing only the values at even indices. Use `range` with the index, discarding the value where unneeded.

## Mini Project
Build a word-length reporter: given a slice of words, range over them printing each word and its length, then compute the average length. Also range over one word as a string, printing each rune with `%c`.

## Summary
`range` is Go's idiomatic iteration tool, working uniformly across slices, arrays, strings, and maps while returning an index/key and a value each step. The value is always a copy, so in-place edits require indexing back into the original. Ranging over strings decodes proper Unicode runes rather than raw bytes.

## What to Learn Next
Previous: [19 Slice](../19Slice/README.md)

You've seen maps briefly through `range` — next, we cover maps properly: creating them, checking missing keys, and why iteration order is random.

Next: [21 Maps](../21Maps/README.md)
