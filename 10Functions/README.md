# 10 – Functions

## What You Will Learn
- How to declare and call a function in Go
- Functions with no return value, one return value, and **multiple** return values
- The difference between passing a value by copy versus sharing data through a slice

## Why This Matters
Every program beyond a few lines needs a way to name a piece of logic once and reuse it. Functions are that unit of reuse. Go also makes an unusual design choice here: instead of throwing exceptions or writing results into "out parameters" (common in other languages), Go lets a function return **more than one value** — typically a result and an error. This keeps error handling explicit and visible at every call site, rather than hidden in a try/catch somewhere far away. Understanding functions well is the foundation for almost everything that follows in this course.

## Concept Explanation
A function is declared with `func`, a name, a parameter list (each parameter needs an explicit type), and an optional return type:

```go
func functionName(paramName paramType) returnType {
    // body
}
```

If a function doesn't produce a value, the return type is simply omitted. If it needs to hand back more than one value, list the types in parentheses: `func swap(x, y string) (string, string)`. The caller must then capture every returned value (or explicitly discard one with `_`).

Go also distinguishes between two ways parameters behave once inside a function:
- **Call by value** – most types (like `int`, `string`, `bool`) are copied into the function. Changes inside the function never affect the caller's original variable.
- **Sharing data** – a few types, such as slices (an ordered collection we'll cover properly in a later lesson), internally point at shared data. Passing one into a function lets that function's changes be visible back in the caller. This isn't the same as passing a pointer explicitly — that's also a later topic — it's just how slices are built.

## Simple Example
From [main.go](main.go), a function with a single return value:

```go
func max(n1, n2 int) int {
    if n1 > n2 {
        return n1
    } else {
        return n2
    }
}
```

And one returning two values at once:

```go
func swap(x, y string) (string, string) {
    return y, x
}
```

Called as `firstName, LastName := swap("John", "Doe")`.

## How It Works
When `max(100, 200)` runs, Go copies `100` into `n1` and `200` into `n2`, compares them, and `return` immediately sends the winning value back and exits the function. For `swap`, both return values are packed together and the caller must receive both — `firstName, LastName := swap(...)` — because Go doesn't let you silently ignore return values with a plain assignment. The call-by-value demo (`increment`) shows a copy being changed with no effect on the original `x` in `main`, while the call-by-reference demo (`modify`) shows a slice's underlying data being changed and that change staying visible in `main` afterward.

## Common Mistakes
- **Assuming Go supports default parameter values or overloading.** Unlike some languages, Go has neither — every function has exactly one signature, and every call must supply all its parameters explicitly.
- **Forgetting to capture all return values.** If a function returns two values, you cannot assign the call to a single variable; you must use `_` for any value you don't need.
- **Expecting a plain `int`, `string`, or `struct` copy to reflect changes made inside a function** — only reference-like types (such as slices) behave that way.

## Best Practices
- Keep functions short and focused on one task — if the name needs "and" to describe it, consider splitting it.
- Group parameters of the same type together, e.g. `func max(n1, n2 int) int` instead of repeating the type.
- Use multiple return values for the common Go pattern `(result, error)` rather than panicking or using a global "last error" variable.

## Real-World Example
Nearly every function in Go's standard library that can fail returns `(value, error)`, e.g. `os.Open(name string) (*os.File, error)`. Callers check the error immediately after the call, right where the failure could happen, instead of jumping to a distant catch block.

## Exercise
Write a function `isEven(n int) bool` that returns whether `n` is even, and a function `divide(a, b int) (int, int)` that returns both the quotient and the remainder of `a / b`. Call both from `main` and print the results.

## Mini Project
Build a small "temperature converter" program: write `toFahrenheit(c float64) float64` and `toCelsius(f float64) float64`, plus a function `describe(tempC float64) string` that returns `"freezing"`, `"cold"`, `"warm"`, or `"hot"` based on the value using if/else. Call all three from `main` with a few sample temperatures.

## Summary
Functions let you name and reuse logic. Go's parameter types must be explicit, there's no overloading or default values, and functions can return multiple values — a design that makes errors an ordinary, visible part of a function's result rather than an exceptional control-flow jump. Most types are passed as copies; a few, like slices, share their underlying data with the caller.

## What to Learn Next
Previous: [09 Control Flow & Loops](../09Control_Flow_Loops/README.md)
Next: [11 – Scope](../11Scope/README.md)
