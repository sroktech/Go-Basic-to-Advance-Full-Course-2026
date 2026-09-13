# 23 – Type Casting (Type Conversion & Type Assertion)

## What You Will Learn
- The difference between **type conversion** (`T(value)`) and **type assertion** (`value.(T)`)
- Why Go never converts types for you automatically
- How numeric conversions can silently truncate or overflow
- How to convert between `string`, `[]byte`, and `[]rune`
- How to convert between strings and numbers using `strconv`
- How to safely extract a concrete type out of an `interface{}`/`any` value, using the two-value assertion and type switches

## Why This Matters
JavaScript and Python will happily turn a string into a number, or a float into an int, whenever the context "seems to want it." That convenience is also a source of subtle bugs — `"5" + 3` behaves differently depending on the language and mood of the interpreter. Go refuses to guess. Every type change must be written explicitly, so a reviewer (or you, six months later) can see exactly where and why data changes shape. This is the same philosophy behind Go's explicit error returns: nothing important happens invisibly.

Type assertion is a related but different tool: it doesn't change a value's type, it reveals what concrete type is already hiding inside an interface value. You just learned interfaces in the previous lesson — an `interface{}` (or `any`) is the extreme case, an interface with zero methods, so it can hold *any* value. Type assertion is how you get a concrete, usable value back out of it.

## Concept Explanation
**Type conversion** (`targetType(value)`) works between compatible concrete types — e.g. `int` to `float64`, or `int32` to `int64` — and always produces a value, even if data is lost (truncation, overflow, precision loss). The compiler enforces that the types are convertible; it does not enforce that the conversion is *safe*.

**Type assertion** (`value.(ConcreteType)`) only applies to interface values. It asks, "does this interface currently hold a value of exactly this concrete type?" It comes in two forms:
- Single-value: `v := val.(string)` — **panics** if `val` doesn't actually hold a `string`.
- Two-value ("comma-ok"): `v, ok := val.(string)` — never panics; `ok` is `false` and `v` is the zero value on mismatch.

A **type switch** (`switch t := v.(type) { case int: ... }`) is the clean way to branch over several possible concrete types at once.

## Simple Example
Numeric conversion requiring an explicit cast, and overflow when narrowing:

```go
var i int = 42
var f float64 = float64(i) // explicit — Go won't do this for you

var big int = 300
small := int8(big) // int8 max is 127; 300 overflows and wraps
```

Safe type assertion versus a type switch, from [main.go](main.go):

```go
var val interface{} = "Hello, Go!"

str2, ok := val.(string) // ok=true, no panic
num2, ok := val.(int)    // ok=false, num2 is 0 (zero value)

switch t := v.(type) {
case int:
    fmt.Printf("int: %d\n", t)
case string:
    fmt.Printf("string: %q\n", t)
}
```

## How It Works
`float64(i)` doesn't "ask" the runtime anything — the compiler already knows both types at compile time, so it just re-encodes the bits into the new type's representation. That's why conversion never fails at runtime; it can only lose information.

A type assertion is different: `val.(string)` is a **runtime** check, because an interface value carries its concrete type as hidden metadata alongside the data. `val.(string)` asks that metadata "are you actually a string?" The two-value form checks and reports the answer; the single-value form checks and panics if the answer is no — which is why the single-value form is a landmine you should reserve for that rare case where you are 100% certain of the type.

## Common Mistakes
- **Silent truncation/overflow**: `int8(300)` doesn't error — it wraps around to some unrelated small number. Narrowing conversions need a range check first if the input isn't already known to fit.
- **Using the single-value assertion**: `v := val.(string)` crashes the program the instant `val` isn't a string. Prefer `v, ok := val.(string)` unless a panic is genuinely the correct behavior (e.g., a bug you want to surface loudly).
- **Confusing conversion with assertion**: `T(x)` converts a *concrete* value between compatible concrete types; `x.(T)` extracts a concrete type from an *interface* value. `x.(T)` on a non-interface `x` is a compile error, and `T(x)` on incompatible types is also a compile error — they solve different problems and aren't interchangeable.

## Best Practices
- Always use the comma-ok form of type assertion unless you can prove the type can't be wrong.
- When narrowing numeric types (e.g., `int` → `int8`), validate the range first, or document why it's known to be safe.
- Prefer `strconv` functions (`Atoi`, `ParseFloat`, etc.) over manual parsing, and always check their returned `error`.
- Use a type switch instead of a chain of `if _, ok := ...` assertions when handling more than two possible types.

## Real-World Example
Type assertions show up constantly when decoding loosely-typed data — JSON into `map[string]interface{}`, config values read as `any`, or plugin systems that pass data through a generic interface. A config loader might read a value as `interface{}` and then type-switch on it to decide whether it's a string, a number, or a nested map. Numeric conversions matter in domains like image processing or audio (narrowing `int32` sample values to `int16`) where the possibility of overflow needs to be handled deliberately rather than accidentally.

## Exercise
Write a function `toCelsius(f interface{}) (float64, bool)` that accepts an `interface{}` which might be an `int`, `float64`, or `string` (like `"98.6"`). Use a type switch to detect which case it is, convert/parse it to `float64` accordingly, then convert the Fahrenheit value to Celsius (`(f - 32) * 5 / 9`). Return `false` for any other type instead of panicking.

## Mini Project
Build a tiny "form field parser": given a `map[string]interface{}` representing form input (e.g., `{"age": "34", "score": 91, "active": true}`), write a function that extracts the `"age"` field as an `int` (using `strconv.Atoi` since it arrives as a string), the `"score"` field as an `int` via safe type assertion, and prints a validation error (without panicking) for any field of an unexpected type. This combines `strconv` conversions with comma-ok assertions and type switches from this lesson.

## Summary
Go requires explicit type conversion (`T(value)`) between concrete types, and explicit type assertion (`value.(T)`) to pull a concrete type out of an interface — two different mechanisms solving two different problems. Conversions can silently lose data (truncation, overflow, precision), so treat narrowing conversions with care. Assertions can panic unless you use the two-value form or a type switch — prefer those in almost all real code.

## What to Learn Next
Previous: [22Interfaces](../22Interfaces/README.md)
Next: [24Defer](../24Defer/README.md)
