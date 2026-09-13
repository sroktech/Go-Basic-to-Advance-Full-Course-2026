# 03 – Data Types

## What You Will Learn
Go's core built-in data types — signed and unsigned integers, floating-point numbers, booleans, complex numbers, and strings — and why Go gives you so many numeric variants instead of just one generic "number" type.

## Why This Matters
Go is **statically typed**: every variable's type is fixed and known at compile time (recall Lesson 00). That means picking the right type isn't just a formality — it directly affects how much memory your program uses, what range of values it can safely hold, and whether the compiler will catch a mistake for you before the program ever runs. Languages like Python or JavaScript hide most of this from you; Go asks you to be deliberate about it.

## Concept Explanation
An **integer** is a whole number with no decimal component. Go gives you both **signed** integers (`int`, `int8`, `int16`, `int32`, `int64`), which can represent negative and positive values, and **unsigned** integers (`uint`, `uint8`, `uint16`, `uint32`, `uint64`), which can only represent zero and positive values. The number in the type name is the **bit size** — how many binary digits are used to store the value — which determines its range. A smaller bit size uses less memory but can hold a smaller range of numbers; because unsigned types don't need to represent negatives, they get double the positive range for the same bit size.

A **floating-point** type (`float32`, `float64`) represents numbers with a decimal component, following the IEEE 754 standard used by virtually every modern programming language. A larger bit size gives more digits of precision at the cost of using more memory.

A **boolean** (`bool`) holds exactly one of two values, `true` or `false`, and is the foundation of all conditional logic you'll write later.

Go also has native support for **complex numbers** (`complex64`, `complex128`) — numbers with a real and an imaginary part — built with the `complex()` function. This is a niche type for scientific/engineering computation; you likely won't reach for it often, but it's worth knowing it exists.

A **string** is an immutable (unchangeable once created) sequence of bytes, almost always representing UTF-8 text. "Immutable" means that once a string is created, its contents can't be modified in place — any "change" actually creates a new string.

## Simple Example
From [main.go](main.go) in this folder:
```go
var i8 int8 = 127
var u8 uint8 = 255

var f32 float32 = 10.6
var f64 float64 = 10.6

var isActive bool = true
var name string = "Kermet the frog!"
```

## How It Works
- `int8` ranges from −128 to 127 (8 bits split between negative and positive values), while `uint8` ranges from 0 to 255 (all 8 bits used for positive values) — same memory footprint, different range, because one reserves room for negative numbers and the other doesn't.
- `uint8` is also aliased as `byte`, since raw binary data is usually represented as unsigned 8-bit values.
- `int32` is also aliased as `rune`, Go's name for a single Unicode code point (roughly, one "character") — you'll meet this properly when working with strings in more depth later.
- `float64` is the default, preferred float type in Go for most math — `main.go` demonstrates that assigning a long decimal like `10123456789012345` to a `float32` loses precision compared to a `float64`, because `float32` only has about 6–7 significant digits versus `float64`'s 15–16.
- `len(someString)` returns the number of **bytes**, not necessarily the number of visible characters — a distinction that matters once you use non-English text, since some Unicode characters take more than one byte.

## Common Mistakes
- Assuming `int` is always the same size — it's platform-dependent (32-bit or 64-bit depending on the system), so never hardcode assumptions about its exact range in portable code.
- Overflowing a smaller type, e.g., trying to store 300 in an `int8` (max 127) — this either fails to compile (for constants) or silently wraps around (for computed values), which is a classic source of hard-to-find bugs.
- Using `float32`/`float64` for money or anything requiring exact decimal precision — floating-point numbers can't represent every decimal value exactly, which is a well-known source of rounding bugs (use a dedicated decimal type for currency in real applications, a topic beyond this lesson).
- Forgetting that strings are immutable — there's no way to change a single character of a string in place; you always build a new one.

## Best Practices
- Default to `int` for whole numbers and `float64` for decimals unless you have a specific reason (like memory constraints or interfacing with a specific binary format) to pick a smaller or unsigned type.
- Use `uint` types only when a negative value truly makes no sense (like a count or an index), not just because a value happens to always be positive today.
- Reach for `complex64`/`complex128` only if you actually need complex-number math — most everyday Go code never touches them.

## Real-World Example
Choosing the right data type is a real engineering decision, not academic trivia. An API that returns a user's age as an `int8` instead of `int` would silently break the moment someone tries to store 130 (still a plausible edge case for data validation testing) or if the field is repurposed for something like "days since signup." Similarly, choosing `byte` (`uint8`) arrays instead of `string` is exactly how Go represents raw file contents or network data efficiently, because those are unsigned, memory-tight, and don't imply "readable text."

## Exercise
Declare a `float32` and a `float64`, assign each the value `1.0 / 3.0` cast appropriately, print both with `fmt.Println`, and observe how many digits of precision each one actually keeps.

## Mini Project
Build a tiny "unit converter" printer: declare a distance in kilometers as a `float64`, convert it to miles (multiply by `0.621371`), and print both values along with their types using `%T` in `fmt.Printf`.

## Summary
Go gives you distinct integer, float, boolean, complex, and string types, each with explicit sizes and trade-offs around range, precision, and memory. Because Go is statically typed, choosing the right one matters — the compiler will hold you to it, which prevents a whole class of bugs common in languages that don't check types this strictly.

## What to Learn Next
Now that you know the types values can have, the next lesson covers variables in depth — the different ways to declare them and how Go's type inference actually works.

Prev: [../02Basic_Syntax/README.md](../02Basic_Syntax/README.md)
Next: [../04Variables/README.md](../04Variables/README.md)
