# 05 – Constants

## What You Will Learn
- What constants are and how they differ from variables
- Typed vs untyped constants, and why the distinction matters
- Declaring single constants and grouped `const` blocks
- Integer literals in decimal, octal, and hexadecimal
- Floating-point and scientific-notation constants
- String constants, escape sequences, and compile-time concatenation
- Boolean constants and constant expressions evaluated at compile time

## Why This Matters
Some values in a program should never change once set: a tax rate, the number of days in a week, a fixed API version, a feature flag. If you used a regular variable for these, nothing would stop you (or a teammate) from accidentally reassigning them later. Constants make that mistake impossible — the compiler rejects any attempt to change them. They also let Go compute values ahead of time (at compile time) instead of redoing the work every time the program runs, which is both safer and faster.

## Concept Explanation
A constant is declared with the `const` keyword instead of `var` or `:=`. Its value must be known at compile time — a literal like `100` or `"hello"`, or an expression built only from other constants. You cannot assign the result of a function call or anything computed at runtime to a constant.

Go constants come in two flavors:
- **Typed constants** — declared with an explicit type, e.g. `const PRICE int = 100`. A typed constant behaves like a fixed value of that exact type and cannot be mixed with a different type without conversion.
- **Untyped constants** — declared without a type, e.g. `const AVOGADRO = 6.022e23`. Go doesn't lock these to one type; instead it figures out the right type from how the constant is used. This makes untyped constants more flexible — the same constant can be used as a `float32` in one place and a `float64` in another.

Integer literals can be written in different numeral bases: plain digits for decimal (`255`), a leading `0` for octal (`0377`), and a leading `0x` for hexadecimal (`0xff`) — all three represent the same number, 255.

## Simple Example
From [const.go](const.go):

```go
const NAME string = "John Doe" // typed constant
const PRICE int = 100

const (
	DECIMAL     = 255  // base 10
	OCTAL       = 0377 // base 8
	HEXADECIMAL = 0xff // base 16
)

const AREA = LENGTH * WIDTH // computed at compile time
```

## How It Works
- `const NAME string = "John Doe"` fixes `NAME` to the type `string` forever; you could never later do `NAME = 42`.
- The grouped `const ( ... )` block declares several related constants together — useful when they belong to the same idea (here, three ways of writing the same number).
- `AREA = LENGTH * WIDTH` is not calculated while the program runs — the Go compiler works out `50 * 5 = 250` before the program is even built, and bakes `250` directly into the binary.
- String constants like `GREETING = "Hello, Earth!\n"` support escape sequences (`\n` for newline, `\"` for a literal quote) and can be joined at compile time with `+`, as `MULTILINE` and `CONCATENATED` show.

## Common Mistakes
- Trying to use `:=` with a constant — this only works for variables. Constants must use `const`.
- Reassigning a constant after declaration — the compiler will refuse to compile.
- Assuming a typed constant can be freely mixed with another numeric type without conversion (e.g. passing a typed `int` constant where a `float64` is expected).
- Forgetting that a leading `0` on an integer literal means octal, not decimal — `0377` is 255, not three hundred seventy-seven.

## Best Practices
- Use `UPPER_SNAKE_CASE` for constant names so they're visually distinct from variables at a glance.
- Prefer untyped constants when the value should be reusable across multiple numeric types (like `AVOGADRO`); use typed constants when you want to lock the type deliberately (like `PRICE int`).
- Group related constants in a single `const (...)` block for readability.
- Use `fmt.Print`, not `fmt.Printf`, when printing a plain string constant that has no format verbs — passing arbitrary text to `Printf` risks it being misread as a format string if it ever contains a stray `%`.

## Real-World Example
Configuration values that must never drift during a program's life are a natural fit for constants: an HTTP status code table, a maximum retry count, a currency's number of decimal places, or a fixed conversion factor used throughout a physics or finance library (like `AVOGADRO` here). Because these are resolved at compile time, using them costs nothing at runtime — unlike reading a config value from a file or database on every call.

## Exercise
Declare a grouped `const` block with the number of days in each of `MONDAY` through `FRIDAY` you work per week, a `HOURLY_RATE` float64 constant, and compute a `WEEKLY_PAY` constant from them. Print all values with `fmt.Println`.

## Mini Project
Write a small "unit price calculator": declare constants for `ITEM_PRICE`, `TAX_RATE` (e.g. `0.08`), and `QUANTITY`, then compute and print `SUBTOTAL`, `TAX`, and `TOTAL` — all as constant expressions computed at compile time.

## Summary
Constants give you fixed, compiler-checked values that can never change after declaration. Typed constants lock to one type; untyped constants stay flexible until used. Go evaluates constant expressions at compile time, making them both safer than variables and free at runtime. Integer literals can be written in decimal, octal, or hexadecimal, and string constants support escape sequences and compile-time concatenation.

## What to Learn Next
Previous: [04 Variables](../04Variables/README.md)

Now that you know how to declare fixed and changeable values, the next step is learning how to combine them using Go's operators.

Next: [06 Operators](../06Operators/README.md)
