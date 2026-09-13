# 11 – Scope

## What You Will Learn
- The three levels of scope in Go: package, function, and block
- What "shadowing" means and how a variable name can hide an outer one
- Why `:=` can accidentally create a new variable instead of reusing an existing one

## Why This Matters
As programs grow past a single function, you need to know exactly where a variable is "alive" and visible — and where it silently stops existing. Getting scope wrong is one of the most common sources of confusing bugs for beginners: you change a variable in one place and expect it to affect another, but because of scope rules, it doesn't. Understanding scope now also sets up closures (the next lesson), which are built entirely on how Go tracks which variables a piece of code can "see."

## Concept Explanation
Go has three scope levels:
1. **Package scope** – declared outside any function with `var`. Visible to every function in the same package (and file).
2. **Function scope** – declared inside a function. Visible only within that function.
3. **Block scope** – declared inside `{ }` (an `if`, `for`, or any bare block). Visible only within that block; it disappears once the closing brace is reached.

**Shadowing** happens when an inner scope declares a variable with the *same name* as one in an outer scope. Inside that inner scope, every reference to the name resolves to the new, inner variable — the outer one still exists, but it's temporarily hidden. This matters especially with `:=` (short variable declaration), which *always* creates a new variable in the current scope, even if a variable with that name already exists one level up. That's easy to do by accident inside an `if` or `for` block.

## Simple Example
From [main.go](main.go), a package-level variable being shadowed inside `main`:

```go
var x int = 25 // package-level

func main() {
    var x int = 100 // shadows package-level x
    fmt.Printf("x = %d\n", x) // prints 100, not 25
}
```

And a block-scoped variable with `:=`:

```go
z := "outer z"
{
    z := "inner z" // new variable, only inside this block
    fmt.Println(z) // "inner z"
}
fmt.Println(z) // "outer z" — unchanged
```

## How It Works
The package-level `x` and `y` are created once, before `main` even runs, and are reachable from any function in the file. When `main` declares its own `var x int = 100`, that's a completely separate variable that happens to share the name `x`; every subsequent use of `x` inside `main` refers to this local one. The package-level `x` is untouched — it still holds `25`, just inaccessible by that name from within `main`. The same idea applies to the `if true { blockVar := "..." }` example: `blockVar` only exists between the braces and would be a compile error if referenced afterward. The nested-block `z` example shows that even `:=` inside plain `{ }` (not just `if`/`for`) creates a fresh variable that vanishes when the block ends.

## Common Mistakes
- **Accidental shadowing with `:=` inside `if`/`for`.** Writing `if val, err := doSomething(); err == nil { ... }` creates a `val` that only exists inside the `if` block — trying to use it afterward fails to compile, and reusing the pattern elsewhere in the same function can silently shadow an outer `val` instead of updating it.
- **Assuming a shadowed outer variable was modified.** Declaring `x := 100` inside a function when an outer or package-level `x` already exists does *not* change the outer `x` — it just hides it locally.
- **Relying on block-scoped variables outside their block**, which is always a compile-time error in Go, not a warning.

## Best Practices
- Give inner-scope variables distinct names when the code also needs the outer value — don't shadow unless you deliberately want to hide the outer variable for that block.
- Prefer the smallest scope that works: declare variables as close as possible to where they're used, inside the block that needs them.
- Minimize package-level (global) variables — they make it harder to reason about who can change a value from where.

## Real-World Example
A common real bug: `if err := step1(); err != nil { return err }` followed later by `if err := step2(); err != nil { return err }` — each `err` is a fresh, block-scoped variable. This is intentional and fine here, but if you meant to check one accumulated `err` variable across multiple steps, using `:=` repeatedly instead of `=` would shadow it each time instead of reusing it, hiding logic errors.

## Exercise
Write a package-level variable `counter int` initialized to `0`. In `main`, write an `if` block that declares a local `counter := 100` and prints it, then print the package-level `counter` again right after the block to show it's still `0`.

## Mini Project
Build a small "settings" program: declare package-level `defaultTimeout int = 30`. In `main`, simulate three different "requests" using three separate blocks (`{ }`), each shadowing `defaultTimeout` with a different local value (e.g., 10, 60, 5) and printing it. After all three blocks, print `defaultTimeout` once more to confirm the package-level value never changed.

## Summary
Go variables live in package, function, or block scope, and a variable's visibility ends the moment its enclosing scope closes. Shadowing lets an inner scope reuse an outer name without touching the outer variable — powerful, but a frequent source of subtle bugs, especially with `:=` inside `if`/`for` blocks. Always know which variable a name is currently pointing to.

## What to Learn Next
Previous: [10 – Functions](../10Functions/README.md)
Next: [12 – Closures](../12Closures/README.md)
