# 24 – Defer

## What You Will Learn
- What `defer` does and exactly when a deferred call runs
- Why `defer` exists — the cleanup problem it solves
- LIFO ordering when multiple `defer` statements are stacked
- Why deferred arguments are evaluated immediately, not when the call actually fires
- How `defer` interacts with named return values
- How `defer` + `recover()` catches a panic (previewed here, covered in full in the next lesson)

## Why This Matters
Every time you open a file, acquire a lock, or open a database connection, something has to close it again — on every possible exit path, including early returns and even panics. In languages with exceptions, this is what `try/finally` is for. Go doesn't have exceptions, so it gives you `defer`: a statement that says "run this when the surrounding function returns, no matter how it returns." Writing the cleanup call right next to the acquisition (`f, _ := os.Open(...); defer f.Close()`) means you can never forget it, even as the function grows more complex later.

## Concept Explanation
`defer someCall(args)` schedules `someCall` to run when the *enclosing function* returns — not at the end of the current block, and not right away. It runs whether the function returns normally, returns early, or is unwinding from a panic.

Two behaviors are essential to internalize:
1. **LIFO order** — if you defer multiple calls, they run in reverse order of registration, like a stack: the last `defer` registered is the first to run.
2. **Eager argument evaluation** — the *arguments* to a deferred call are evaluated immediately, at the point of the `defer` statement. Only the *call itself* is postponed.

## Simple Example
LIFO ordering, from [main.go](main.go):

```go
func demoLIFO() {
	defer fmt.Println("  defer 1 — registered first, runs LAST")
	defer fmt.Println("  defer 2")
	defer fmt.Println("  defer 3 — registered last, runs FIRST")
	fmt.Println("demoLIFO: end of function body")
}
```

Eager argument evaluation — a classic gotcha:

```go
func demoEagerArgs() {
	i := 0
	defer fmt.Println("  deferred i (captured at defer time):", i) // captures 0 NOW
	i = 10
	fmt.Println("  i at end of function:", i) // 10
	// the deferred call still prints 0, not 10
}
```

## How It Works
Think of `defer` as pushing a fully-prepared call onto a stack. "Fully-prepared" is the key detail: Go evaluates the arguments right then, freezes them, and pushes the *call with those fixed values* onto the function's defer stack. It does not push "call this function later and figure out the arguments then." That's why `i` in the example above is captured as `0`, even though `i` becomes `10` before the function returns.

When the function is about to return — for any reason — Go pops the stack and runs each deferred call in turn, which is why the most-recently-deferred call runs first (LIFO).

Named return values add one more twist, used in [main.go](main.go)'s `doSomething`:

```go
func doSomething() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("doSomething failed: %w", err)
		}
	}()
	err = fmt.Errorf("connection timeout")
	return // defer runs after this sets err, and CAN still modify it
}
```

Because `err` is a named return variable, a deferred closure can read and reassign it after `return` has "set" it but before the caller actually receives it — this is exactly how `defer` + `recover()` (next lesson's topic) is able to turn a panic into a returned error.

## Common Mistakes
- **Assuming deferred arguments are evaluated later.** `defer fmt.Println(i)` captures the value of `i` at the moment of the `defer` statement, not at function exit. If you need the final value, defer a closure (`defer func() { fmt.Println(i) }()`) instead.
- **Deferring inside a loop for per-iteration cleanup.** `defer file.Close()` inside a `for` loop does not close each file at the end of that iteration — it piles up all the closes until the *whole function* returns, potentially holding many resources open simultaneously. Wrap the loop body in its own function (or call cleanup directly) when you need per-iteration cleanup.

## Best Practices
- Defer the cleanup call immediately after the resource is successfully acquired (`f, err := os.Open(...); if err != nil { return }; defer f.Close()`), so it's impossible to add code later that forgets it.
- Use a deferred closure (not a bare deferred call) whenever you need the deferred logic to see a value's *final* state or to modify a named return.
- Keep deferred cleanup simple and side-effect-focused; don't bury important business logic inside a `defer`.

## Real-World Example
Production Go code leans on `defer` constantly: closing an HTTP response body (`defer resp.Body.Close()`), unlocking a `sync.Mutex` right after locking it (`mu.Lock(); defer mu.Unlock()`), closing a database transaction, or releasing a semaphore slot in a worker pool. The pattern is always the same: acquire, then immediately defer the release, so the release can never be missed regardless of how many exit paths the function later grows.

## Exercise
Write a function `withLogging(name string)` that prints `"start: <name>"` and returns a function which, when deferred and called, prints `"end: <name>"`. Then write a second function that calls `defer withLogging("task")()` at its top and does a couple of `fmt.Println` calls in its body. Confirm the "start" line prints immediately and the "end" line prints last, after the body runs.

## Mini Project
Build a small "resource pool" simulation using only `defer`, functions, and error handling from lessons so far: a `openResource(name string) (string, error)` function that returns a fake handle and an accompanying `closeResource(name string)` function. Write a `useResources(names []string)` function that opens each resource in turn and defers its close — but do it two ways: once with `defer` directly in the outer loop (observe that all resources stay open until `useResources` returns) and once wrapping each iteration's body in an anonymous function (observe that each resource closes before the next iteration begins). Print timestamps or step labels to make the ordering difference visible.

## Summary
`defer` schedules a call to run when the enclosing function returns, covering every exit path including panics — making it Go's answer to reliable cleanup without exceptions or `finally` blocks. Multiple defers run in LIFO order, and a deferred call's arguments are locked in at the moment of the `defer` statement, not when it actually executes. Combined with named return values, `defer` can inspect and modify what a function ultimately returns — the mechanism the next lesson uses for `recover()`.

## What to Learn Next
Previous: [23TypeCasting](../23TypeCasting/README.md)
Next: [25ErrorHandling](../25ErrorHandling/README.md)
