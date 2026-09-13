# 12 – Closures

## What You Will Learn
- What an anonymous function is and how to define one inline
- What a closure is, and how it "remembers" variables from where it was created
- How to pass functions as arguments (higher-order functions) and why that's useful

## Why This Matters
So far, every function has been a fixed, named block of code. Closures let a function carry its own private state around with it — this is how you build things like counters, generators, or "wrap this behavior around that function" tools (middleware) without needing a class or object. Because Go treats functions as ordinary values (you can store them in a variable, pass them around, return them), closures fall out naturally from the language rather than being a special add-on feature.

## Concept Explanation
An **anonymous function** is a function literal with no name, defined right where it's needed:

```go
greet := func(name string) string {
    return "Hello, " + name + "!"
}
```

A **closure** is a function that references a variable declared *outside* its own body, in the scope where it was created. Go keeps that variable alive for as long as the closure exists — even after the outer function that declared it has already returned. Each time the outer function runs again, it creates a brand-new, independent variable for the new closure to capture.

This matters because it lets a function have private, persistent state without any global variable: the state lives only inside the closure, invisible to anything else.

## Simple Example
From [main.go](main.go):

```go
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

```go
counterA := makeCounter()
counterB := makeCounter()
counterA() // 1
counterA() // 2
counterB() // 1 — independent from counterA
```

## How It Works
`makeCounter` declares `count := 0` and then returns an inner function that increments and returns it. Normally a local variable disappears when its function returns — but here, the returned closure still references `count`, so Go keeps it alive. Every call to `makeCounter()` creates a *new* `count`, so `counterA` and `counterB` never interfere with each other. The same pattern powers `makeAdder(x int)`, which returns a function that always adds the same captured `x` to whatever number it's later called with — `add5` and `add10` each remember their own `x`.

Closures also enable **higher-order functions** — functions that take other functions as arguments, like `apply(nums, operation)` and `filter(nums, predicate)` in the example file, which run a passed-in function against every element of a slice (an ordered collection we'll cover properly in a later lesson).

## Common Mistakes
- **The classic loop-variable capture bug.** In older Go versions, a closure created inside a `for` loop that referenced the loop variable directly (e.g. `for i := 0; i < 3; i++ { funcs[i] = func() { fmt.Print(i) } }`) would capture the *same* variable across every iteration, so every closure printed the loop's final value instead of the value at the time it was created. **As of Go 1.22, the language changed this**: each loop iteration gets its own fresh copy of the loop variable, so this pattern now works as most people expect. Even so, the explicit `i := i` pattern (shadowing the loop variable with a fresh copy inside the loop body) still works and makes the intent obvious to any reader, regardless of Go version.
- **Forgetting a closure keeps its captured variables alive.** This can hold onto memory longer than expected if you're not careful with what you capture.

## Best Practices
- Keep closures small and give them a clear single purpose — a closure that captures many outer variables becomes hard to reason about.
- When in doubt about loop-variable capture across Go versions, use the `i := i` shadowing pattern — it is correct and explicit either way.
- Prefer named functions over anonymous ones once the logic grows beyond a couple of lines; closures shine for small, focused behavior.

## Real-World Example
Web frameworks in Go commonly use closures for **middleware** — a function that wraps another function to add logging, authentication, or timing without changing the original function's code. The `withLogging` function in [main.go](main.go) does exactly this: it takes a function and a name, and returns a new function that logs before and after calling the original — the same shape used by real HTTP middleware chains.

## Exercise
Write a closure factory `makeMultiplierChain(start int) func(int) int` that returns a function which, each time it's called with a number `n`, multiplies its running total by `n` and returns the new total (starting from `start`). Call it several times in a row and print the running total after each call.

## Mini Project
Build a simple "bank account" simulator using only closures (no structs yet): write `makeAccount(initialBalance int) (func(int), func(int) bool, func() int)` that returns three closures sharing one captured balance — a `deposit(amount int)`, a `withdraw(amount int) bool` (returns `false` if funds are insufficient), and a `balance() int`. Use them together in `main` to simulate a few deposits and withdrawals, printing the balance after each.

## Summary
Anonymous functions let you define behavior inline; closures let that behavior capture and remember variables from its creation scope, even after the enclosing function has returned. This is how Go achieves private, persistent state and powers patterns like counters, generators, and middleware — all without needing a class.

## What to Learn Next
Previous: [11 – Scope](../11Scope/README.md)
Next: [13 – Variadic Functions](../13Variadic/README.md)
