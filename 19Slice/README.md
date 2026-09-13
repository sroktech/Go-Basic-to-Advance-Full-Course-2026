# 19 – Slices

## What You Will Learn
- What a slice is and how it differs from the fixed-size arrays you already know
- The three parts of a slice: pointer, length, and capacity
- Creating slices with literals, `make`, and the zero-value `nil` slice
- Growing a slice with `append`, including appending one slice into another
- Slicing syntax (`s[low:high]`) to create sub-slices
- Why slices can share an underlying array, and how `copy()` avoids that
- Removing an element (Go has no built-in `remove`) and 2D slices

## Why This Matters
Arrays are fixed in size the moment they're declared — `[5]int` is always exactly 5 ints. Real programs rarely know the exact item count ahead of time: users from a database, lines from a file, filtered search results. Slices wrap an array with a flexible, resizable view, which is why they — not arrays — are the data structure you'll reach for almost every time in Go.

## Concept Explanation
A slice looks like an array literal without the size (`[]string{...}`), but underneath it's a small struct with three fields: a **pointer** to where its data starts in an underlying array, a **length** (`len`, elements currently visible), and a **capacity** (`cap`, room available before a new array must be allocated). Copying a slice variable copies that pointer/len/cap, not the data — so two slices can point at the *same* backing array, which is powerful but also the source of Go's most common slice bug (see below).

`append` adds elements. If spare capacity exists, it writes into the existing array. If not, Go allocates a bigger array, copies everything over, and returns a slice pointing at the new array — which is exactly why you must always write `s = append(s, x)` rather than discard the result.

## Simple Example
From [main.go](main.go):

```go
nums := make([]int, 3, 5) // length=3, capacity=5

scores := []int{10, 20, 30}
scores = append(scores, 40)

letters := []string{"a", "b", "c", "d", "e"}
sub := letters[1:4] // [b c d] — a view into the SAME array as letters
sub[0] = "B"        // also changes letters[1]
```

## How It Works
- `make([]int, 3, 5)` allocates capacity 5 but exposes only 3 slots; `len` is 3, `cap` is 5.
- `append(scores, 40)` has room within capacity, so it writes into the next free slot and bumps `len`.
- `letters[1:4]` doesn't copy anything — `sub` and `letters` share the same backing array from different offsets, so `sub[0] = "B"` is visible through `letters` too.
- `copy(cloned, original)` copies element values into a separate array, breaking that link.
- Removing an element joins `items[:i]` and `items[i+1:]...` with `append`, since Go has no built-in `remove`.

## Common Mistakes
- **Assuming a sub-slice is independent.** `sub := letters[1:4]` shares memory with `letters` — mutating one can silently mutate the other.
- **The append-sometimes-mutates gotcha.** Whether `append` affects a shared array depends entirely on *capacity*: spare capacity means it writes in place (visible elsewhere); exceeded capacity means a new array is allocated and the original is untouched. Same-looking code, different behavior — never rely on aliasing; use `copy()` when you need real independence.
- **Confusing `len` and `cap`.** `len` is what's currently visible; `cap` is how much room exists before reallocation.
- **Forgetting `s = append(s, x)`** — discarding `append`'s return silently drops a possibly reallocated result.

## Best Practices
- Always reassign: `s = append(s, ...)`.
- Use `make([]T, len, cap)` when you know roughly how many elements you'll add.
- Use `copy()` for a slice that must be fully independent of its source.
- Prefer slices over arrays; use a fixed array only when the size is a true, permanent constant.

## Real-World Example
Slices represent nearly every dynamic list in Go: database query rows, a shopping cart's line items, tokens parsed from input, a network read buffer. Any time the item count isn't known until runtime, reach for a slice.

## Exercise
Write `removeDuplicates(nums []int) []int` that returns a new slice with duplicates removed, preserving first-occurrence order. Use only slices, `append`, and loops.

## Mini Project
Build a small inventory tracker: an `Item` struct (`Name string`, `Quantity int`), a `[]Item` slice, and functions to add an item, remove one by name (using the append-removal trick), and print the inventory.

## Summary
Slices are Go's flexible, resizable view over an array, described by a pointer, length, and capacity. `append` grows a slice, reallocating only when capacity runs out — which is also why shared sub-slices sometimes alias and sometimes don't. `copy()` breaks that aliasing when true independence is needed.

## What to Learn Next
Previous: [18 Structures](../18Structures/README.md)

Now that you can build and grow dynamic lists with slices, next is `range`, Go's idiomatic way to iterate over them cleanly.

Next: [20 Range](../20Range/README.md)
