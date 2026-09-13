# 13 – Variadic Functions

## What You Will Learn
- How to write a function that accepts any number of arguments with `...type`
- How to spread an existing slice into a variadic call with `...`
- How to mix regular parameters with a variadic one, and forward variadic args to another function

## Why This Matters
Sometimes you don't know in advance how many values a function needs to handle — think of `fmt.Println`, which can print one value or twenty. Without variadic functions, you'd need a different function for every possible argument count, or force callers to build a collection just to pass a handful of values. Variadic parameters solve this cleanly: the caller writes plain comma-separated arguments, and Go packs them into a collection for you inside the function.

## Concept Explanation
A variadic parameter is written as `...type` and must be the **last** parameter in the list. Inside the function, it behaves like an ordered collection of that type called a slice (we'll cover slices properly in a later lesson) — you can loop over it with `for _, v := range nums`. A function can have at most one variadic parameter, and it's fine to call it with zero arguments — you'll just get an empty collection.

If you already have a collection of values and want to pass all of them into a variadic function, you can't just hand the collection over directly — Go requires the `...` spread operator after it: `sum(numbers...)`. This tells Go "unpack these values as individual arguments" rather than treating the whole collection as one argument.

## Simple Example
From [main.go](main.go):

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

```go
sum()           // 0
sum(5)          // 5
sum(1, 2, 3)    // 6

numbers := []int{10, 20, 30, 40, 50}
sum(numbers...) // 150 — spread with ...
```

## How It Works
When you call `sum(1, 2, 3)`, Go collects `1, 2, 3` into `nums` automatically before the function body runs; when you call `sum()`, `nums` is simply empty, so the loop never executes and `total` stays `0`. When you already have a slice like `numbers` and want to forward all of its values, writing `sum(numbers...)` spreads each element as a separate argument — the same technique is used inside the function itself in `sumAndDouble`, which calls `sum(nums...)` to forward its own variadic parameter into another variadic function. A variadic parameter can mix with regular ones as long as it comes last, as in `greetAll(greeting string, names ...string)`.

## Common Mistakes
- **Passing a slice without the `...` spread.** Writing `sum(numbers)` where `numbers` is `[]int` is a compile error — Go does not automatically unpack a collection into a variadic call; you must write `sum(numbers...)`.
- **Trying to put the variadic parameter anywhere but last**, or trying to declare more than one variadic parameter — both are compile errors.
- **Assuming a variadic parameter is required.** Calling with zero arguments is completely valid; always handle the empty case (as `max` does in the example by checking `len(nums) == 0` first).

## Best Practices
- Use variadic parameters for genuinely optional, same-type, "zero or more" values — not as a workaround for what should really be a single collection parameter passed in directly.
- When forwarding variadic arguments to another variadic function, always remember the `...` spread — it's easy to forget and get a confusing type error.
- For truly "any type" variadic parameters (like `printAll(values ...interface{})` in the example), know this uses Go's "any type" mechanism (`interface{}`) — a more advanced topic we'll cover properly later; for now it's enough to know it's how `fmt.Println` accepts anything.

## Real-World Example
A very common Go pattern built on variadic functions is **functional options**, shown in [main.go](main.go): a constructor like `newServer(host string, opts ...Option)` takes a required value plus any number of optional configuration functions. Each option (`withPort(443)`, `withTimeout(60)`) is itself a function that adjusts a configuration value passed to it. This pattern relies on a couple of ideas — a data grouping called a struct, and pointers — that we'll cover properly in later lessons; for now, just notice how variadic parameters make "zero or more optional settings" a clean, readable API instead of a constructor with a dozen parameters.

## Exercise
Write a variadic function `average(nums ...float64) float64` that returns the average of its arguments (return `0` for zero arguments to avoid dividing by zero). Call it with no arguments, one argument, and several arguments, printing each result.

## Mini Project
Build a small "report line" builder: write `buildLine(label string, values ...int) string` that returns a string like `"Sales: 10, 20, 30 (total 60)"` — joining the values with `", "` and appending their sum. Call it for a few different labels and sets of values (including zero values) and print each resulting line.

## Summary
Variadic parameters (`...type`) let a function accept any number of same-type arguments, packed into a collection inside the function body. Use `...` to spread an existing collection into a variadic call — you cannot pass it directly. Variadic parameters must come last and there can only be one per function, but this simple mechanism is powerful enough to underpin patterns like `fmt.Println` and functional options.

## What to Learn Next
Previous: [12 – Closures](../12Closures/README.md)
Next: [14 – Recursion](../14Recursion/README.md)
