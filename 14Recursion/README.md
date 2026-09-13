# 14 – Recursion

## What You Will Learn
- What recursion is and why every recursive function needs a **base case**
- How to trace a recursive call stack by hand
- Which everyday problems are naturally suited to a recursive solution

## Why This Matters
Some problems are naturally defined in terms of smaller versions of themselves — a factorial, a Fibonacci number, or summing a list. Recursion lets you express the solution the same way you'd describe the problem in plain language ("the sum of a list is its first element plus the sum of the rest"), often making the code shorter and easier to follow than an equivalent loop. But recursion is also the first place in this course where a simple mistake (forgetting to stop) can crash your program, so understanding exactly how and when it stops is essential.

## Concept Explanation
A recursive function calls itself to solve a smaller version of the same problem. It needs exactly two parts:
1. **Base case** – a condition that stops the recursion and returns a value directly, with no further self-call.
2. **Recursive case** – the function calling itself with a simpler or smaller input, moving it closer to the base case.

If a function has no base case (or the base case can never be reached), it calls itself forever. Since Go allocates a new stack frame for every call and does **not** perform tail-call optimization (an optimization some languages use to reuse the same stack frame for a self-call), this eventually exhausts the call stack and crashes the program — this is the "stack overflow" mentioned in the code comments.

## Simple Example
From [main.go](main.go):

```go
func factorial(n int) int {
    if n == 0 || n == 1 { // base case
        return 1
    }
    return n * factorial(n-1) // recursive case
}
```

## How It Works
Calling `factorial(4)` doesn't compute an answer immediately — it calls `factorial(3)`, which calls `factorial(2)`, which calls `factorial(1)`, which finally hits the base case and returns `1` without any further call. Only then does the chain "unwind": `factorial(2)` computes `2 * 1 = 2`, `factorial(3)` computes `3 * 2 = 6`, and `factorial(4)` computes `4 * 6 = 24`. Each pending call waits on the stack until the one it depends on returns. The same shape appears in `sumSlice`, which recurses on a shorter slice (an ordered collection covered properly in a later lesson) each time — `nums[0] + sumSlice(nums[1:])` — until the base case of an empty slice returns `0`.

## Common Mistakes
- **Missing or unreachable base case.** If the condition that should stop the recursion is wrong, or the recursive call doesn't actually move toward it (e.g., calling `factorial(n)` instead of `factorial(n-1)`), the function recurses until the stack overflows and the program crashes.
- **Assuming Go optimizes deep recursion for free.** Go has no tail-call optimization, so every recursive call adds a real stack frame with real memory cost — deep recursion (thousands of levels) is meaningfully more expensive than an equivalent loop, and can crash where a loop wouldn't.
- **Recomputing the same work repeatedly**, as the naive `fibonacci` in the example does — `fibonacci(n-1)` and `fibonacci(n-2)` each recompute overlapping subproblems, making it exponentially slow for larger `n`. (Techniques like memoization or an iterative rewrite fix this, but need more tools than we've covered yet.)

## Best Practices
- Always write the base case first and make sure it's actually reachable from every recursive path.
- Make sure each recursive call moves strictly closer to the base case (a smaller number, a shorter slice, etc.) — never the same or a larger input.
- For simple counting or accumulation problems, prefer a loop unless recursion genuinely makes the logic clearer — it's easier to reason about performance and stack usage with loops until you've mastered recursion.

## Real-World Example
Recursion is the natural fit for traversing tree-shaped or nested data — walking a filesystem's folders and subfolders, or recursively walking nested JSON/config data (an object containing objects containing objects). Each "go one level deeper" step is a recursive call, and the base case is reaching a plain value with nothing left to descend into.

## Exercise
Write a recursive function `countDown(n int, step int)` that prints `n`, then `n - step`, then `n - 2*step`, and so on, stopping (base case) once the value would go to zero or below. Call it with a couple of different starting values and steps.

## Mini Project
Build a small "digit sum" tool using recursion only: write `digitSum(n int) int` that adds up the digits of a positive integer (e.g., `digitSum(1234)` returns `1+2+3+4 = 10`), using `n % 10` to get the last digit and `n / 10` to shrink the number toward the base case of `0`. Extend it with `repeatedDigitSum(n int) int` that keeps applying `digitSum` (recursively or in a loop) until the result is a single digit, and print a few examples.

## Summary
Recursion solves a problem by having a function call itself on a smaller version of that problem, always anchored by a base case that stops the calls. Go builds a real stack frame for every call and does not optimize away deep recursion chains, so a missing or unreachable base case is a genuine crash risk, not just inefficient code. Used well, recursion is a natural, readable fit for self-similar problems like factorials, sequences, and traversing nested data.

## What to Learn Next
Previous: [13 – Variadic Functions](../13Variadic/README.md)
Next: [15 – Strings](../15Strings/README.md)
