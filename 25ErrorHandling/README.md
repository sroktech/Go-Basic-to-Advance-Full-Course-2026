# 25 – Error Handling

## What You Will Learn
- Why Go represents failures as ordinary return values instead of exceptions
- How to define and return `error` values, including custom error types
- How to add context to an error with `fmt.Errorf` and `%w`, and inspect error chains with `errors.Is`/`errors.As`
- When (rarely) to use `panic`, and how `recover` turns a panic back into a normal error
- How `defer` (covered in the previous lesson) is the mechanism that makes `recover` work

## Why This Matters
In languages with exceptions, any line of code can secretly transfer control somewhere far away, and a caller has to remember to wrap things in `try/catch` to find out that something could fail at all. Go's designers chose the opposite trade-off: a function that can fail says so directly in its signature by returning an `error` as its last value. This makes failure a visible, ordinary part of the function's contract — the compiler won't stop you from ignoring it, but the convention and tooling constantly nudge you to check `if err != nil`. `panic`/`recover` still exist, but they're deliberately reserved for programming errors and truly unrecoverable situations — not for "the file wasn't there" or "the input was invalid," which are just error returns.

## Concept Explanation
The built-in `error` type is just an interface:

```go
type error interface {
	Error() string
}
```

Anything with an `Error() string` method satisfies it. Go gives you three tools, in order of how often you should reach for them:

1. **Return an `error`** — the standard way to signal that something went wrong. The caller decides what to do.
2. **`defer`** — runs cleanup (as covered in [the previous lesson](../24Defer/README.md)) regardless of how a function exits.
3. **`panic`/`recover`** — for truly exceptional situations (e.g., an invariant your own code violated), not for routine failures a caller should be expected to handle.

## Simple Example
A function returning a plain error, from [main.go](main.go):

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil // nil means "no error"
}
```

A custom error type carrying structured data, and wrapping with `%w`:

```go
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

func openFile(filename string) error {
	baseErr := errors.New("no such file or directory")
	return fmt.Errorf("openFile %q: %w", filename, baseErr) // wraps baseErr
}
```

`panic`/`recover`, using `defer` exactly as introduced in lesson 24 — here it's used to run `recover()` so a panic becomes a normal returned error:

```go
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	if b == 0 {
		panic("attempted to divide by zero")
	}
	return a / b, nil
}
```

## How It Works
`errors.New` and `fmt.Errorf` both produce a value satisfying the `error` interface — usually just a message string underneath. A **custom error type** like `*ValidationError` also satisfies `error` (it has an `Error() string` method), but it carries extra fields (`Field`, `Message`) that a caller can recover with `errors.As(err, &target)` when it needs more than just the message text.

`fmt.Errorf("...: %w", err)` is different from `%v`: `%w` wraps the original error, keeping a link to it inside the new error value, forming a chain. `errors.Is(err, target)` walks that chain looking for a specific sentinel error; `errors.As(err, &target)` walks it looking for a specific *type*. Using `%v` instead of `%w` breaks that chain — the wrapped error becomes just text, invisible to `errors.Is`/`errors.As`.

`panic` immediately stops normal execution and starts unwinding the call stack, running any deferred calls along the way. If one of those deferred calls invokes `recover()`, the panic stops unwinding right there, and the function returns normally — using the named-return-value trick you saw in the previous lesson to turn the recovered value into a proper `error`.

## Common Mistakes
- **Ignoring errors** — writing `_ = err` or, worse, simply not capturing the returned error at all. An unchecked error is a silent failure waiting to surface somewhere confusing.
- **Using panic/recover as normal control flow** — reaching for `panic` to signal "invalid input" or "not found" instead of returning an `error`. Reserve `panic` for situations where continuing would be actively unsafe (e.g., a violated invariant), not for anything a caller could reasonably be expected to check.
- **Using `%v` where `%w` was needed** — `fmt.Errorf("...: %v", err)` loses the chain, so `errors.Is`/`errors.As` can no longer find the original error inside it. Use `%w` whenever callers might need to inspect what's underneath.

## Best Practices
- Always check `if err != nil` immediately after a call that returns one.
- Return errors as the last value, and return zero values for everything else on the error path.
- Add context while wrapping: `fmt.Errorf("doing X: %w", err)`, so error messages read like a trace of what was happening.
- Use a custom error type when a caller needs to programmatically inspect *why* something failed, not just read a message.
- Keep `panic`/`recover` at the edges (e.g., recovering inside a request handler so one failure doesn't crash a whole server) rather than sprinkling it through business logic.

## Real-World Example
Production services commonly build deep error chains: a database call fails, a repository layer wraps it with `%w` and context ("fetching user 42"), a service layer wraps that again ("processing order"), and the top-level handler logs the full chain while using `errors.Is`/`errors.As` to decide the right HTTP status code (e.g., "was this a not-found error, so return 404?"). `recover` shows up at these same top-level boundaries — an HTTP middleware recovers from any panic in a handler, logs it, and returns a 500, so one buggy request can't take down the whole server.

## Exercise
Write a function `parseAge(s string) (int, error)` that uses `strconv.Atoi` (from lesson 23) to parse a string into an int, returning a wrapped error (`fmt.Errorf("parseAge: %w", err)`) if parsing fails, and a `*ValidationError`-style custom error if the parsed value is negative or over 130. Call it with a few valid and invalid strings and print the results.

## Mini Project
Build a small "safe divide" utility package-free function set: a `SafeDivide(a, b float64) (result float64, err error)` that returns a normal error for division by zero (no panic needed for this case — it's an ordinary, expected failure), plus a separate `MustPositive(n float64) float64` that `panic`s if `n` is negative (representing a genuine programmer-error invariant), wrapped by a `SafeMustPositive` that uses `defer`+`recover` to turn that panic into a returned error instead of crashing. This mirrors the real distinction between ordinary errors and truly exceptional situations.

## Summary
Go treats failure as data: functions return an `error` value, and callers are expected to check it explicitly rather than relying on exceptions. Custom error types let you attach structured detail, and `fmt.Errorf` with `%w` builds inspectable error chains via `errors.Is`/`errors.As`. `panic`/`recover` exist for truly exceptional situations — and `defer` is used here to run `recover()`, as covered in the previous lesson — but they are not a substitute for ordinary error handling.

## What to Learn Next
Previous: [24Defer](../24Defer/README.md)
Next: [26Packages](../26Packages/README.md)
