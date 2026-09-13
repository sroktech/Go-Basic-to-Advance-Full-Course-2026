# 04 – Variables

## What You Will Learn
The different ways to declare variables in Go — `var` with an explicit type, `var` with inferred type, and the shorthand `:=` — plus how Go's type inference works and where each style is (and isn't) allowed.

## Why This Matters
A **variable** is a named storage location that holds a value. Go gives you more than one way to declare one, and each style exists for a reason: sometimes you want to be maximally explicit (useful in shared, public code), and sometimes you want brevity (useful for quick, local logic). Knowing which tool fits which situation is a habit that carries through the rest of the course.

## Concept Explanation
Every variable in Go has a type, and that type is fixed once it's set — Go is **not** dynamically typed the way Python or JavaScript are, even when you let Go infer the type for you. "Type inference" just means Go looks at the value you're assigning and figures out the type on your behalf, at compile time; it does not mean the variable can later hold a different type of value.

Also important: every variable you declare **must be used somewhere**, or Go refuses to compile — the same rule you saw with imports in Lesson 02.

## Simple Example
From [main.go](main.go) in this folder:
```go
var mango string = "This is a big mango!"   // explicit type
var height int = 23                          // explicit type, same line

age := 54                                    // shorthand, type inferred as int
city := "Washington"                         // shorthand, type inferred as string

var apples, oranges int = 23, 78             // multiple vars, same type
var a, b, c = 758.52, 8, "foobar"            // multiple vars, different inferred types
```

## How It Works
- `var name type = value` is the fully explicit form: you state the type yourself, so anyone reading the code (or the compiler) doesn't need to infer anything.
- `var name = value` (no type) still uses the `var` keyword but lets Go infer the type from the value — useful when you want a package-level variable but the type is obvious from context.
- `name := value` — the **short variable declaration** — is the most common style inside functions. It declares the variable and infers its type from `value` in one step. `:=` **only works inside functions**; at the package level (outside any function) you must use `var`.
- `var a, b, c = 758.52, 8, "foobar"` shows that a single `var` statement can assign different types to different variables at once, as long as each has its own value — Go infers each independently (`float64`, `int`, `string` here).
- `var apples, oranges int = 23, 78` shows the same shorthand but with one shared explicit type applied to every name in the list.
- `%T` in `fmt.Printf` prints the actual Go type of a variable — handy for confirming what type inference actually picked.

## Common Mistakes
- Trying to use `:=` at the package level (outside any function) — this does not compile. `:=` is a function-local construct only; use `var` at the package level.
- Declaring a variable and never using it — just like unused imports, this is a hard compile error in Go, not a warning, so half-finished experiments won't even build.
- Assuming inferred types can change later. `age := 54` fixes `age` as an `int` forever in that scope — you cannot later assign `age = "fifty-four"`.
- Redeclaring the same variable with `:=` in the same scope — Go requires that at least one variable on the left side of `:=` be new; reusing `:=` for variables that already all exist is a compile error.

## Best Practices
- Use `:=` for local, everyday variables inside functions — it's idiomatic Go and keeps code concise.
- Use explicit `var name type = value` when the zero value or the exact type matters for clarity, especially in code others will read, or when declaring a variable before you have a value ready for it yet.
- Group related variables of the same type on one line (e.g., `var apples, oranges int = 23, 78`) when it improves readability, but don't force unrelated variables together just to save a line.

## Real-World Example
Explicit, fixed types for variables are exactly what makes large Go codebases predictable. In a web server handling thousands of requests, a variable declared as `var userID int` can never accidentally become a string somewhere down the call chain the way it might in a loosely typed language — eliminating an entire class of "expected number, got string" bugs that are common in production JavaScript or Python services.

## Exercise
Declare three variables describing yourself (e.g., name as a string, birth year as an int, height in meters as a float64) using the `:=` shorthand, then print each one along with its type using `fmt.Printf` and `%T`.

## Mini Project
Build a small "profile card printer": declare variables for a person's name, age, and city using a mix of `var` (explicit type) and `:=` (inferred), then print a formatted multi-line profile using `fmt.Printf`, including each field's type.

## Summary
Go offers three main ways to declare a variable — explicit `var`, inferred `var`, and shorthand `:=` — but every one of them produces a fixed, statically known type. `:=` is function-scoped only, and unused variables never compile, both of which push you toward clean, deliberate code.

## What to Learn Next
With variables covered, the next lesson introduces constants — values that, unlike variables, can never change after they're declared.

Prev: [../03Data_Types/README.md](../03Data_Types/README.md)
Next: [../05Constants/README.md](../05Constants/README.md)
