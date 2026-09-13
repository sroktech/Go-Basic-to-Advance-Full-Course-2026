# 36 – Testing

## What You Will Learn
- Writing tests with Go's built-in `testing` package — no external framework required
- The rules `_test.go` files, package placement, and `TestXxx(t *testing.T)` follow
- Table-driven tests — the idiomatic Go pattern for multiple cases
- Subtests with `t.Run` for isolated names and failures
- Testing error cases deliberately, not just the happy path
- Benchmarks (`BenchmarkXxx(b *testing.B)`) and how to run them

## Why This Matters
Go ships a testing framework in the standard library because the language treats testing as a first-class, ordinary part of writing software — not an add-on that requires choosing and configuring a separate library (Jest, Mocha, pytest) before you can even start. `go test` finds, compiles, and runs your tests with zero configuration. That simplicity is deliberate: it removes the excuse to skip testing "until we set up the test framework," and it means every Go codebase you'll ever work in tests the same basic way, which makes onboarding into new projects fast.

Table-driven tests exist because most real bugs live in edge cases, not the one obvious case — dividing by zero, an empty string, a negative number. Writing every case as a new `TestXxx` function doesn't scale; a table of `{name, input, want}` values run through one loop scales to dozens of cases with almost no added code, and each case gets its own name and pass/fail status via `t.Run`.

## Concept Explanation
A Go test file must end in `_test.go`, live in the same package as the code it tests (or use a `_test` package suffix for black-box testing), and contain functions named `TestXxx(t *testing.T)`. `go test` compiles these files only when running tests — they're never part of the normal build.

```
go test ./...
    │
    ├─ finds all *_test.go files
    ├─ compiles them with the package
    ├─ runs functions named TestXxx(t *testing.T)
    └─ reports PASS / FAIL
```

`t.Errorf` marks the test failed but lets it keep running (useful in a loop, so you see every failing case, not just the first). `t.Fatalf` marks it failed and stops immediately — appropriate when continuing would be meaningless, such as after an unexpected error you didn't plan to recover from.

The table-driven shape used throughout [math_test.go](math_test.go):

```
tests := []struct{ ... }{ {case1}, {case2}, ... }
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) { ... })
}
```

`t.Run` creates a named subtest for each table entry. Subtests report individually (`TestSubtract/negative_result`), and a failure in one doesn't stop the others in the same table from running — you get a complete picture of what broke in one `go test -v` run.

## Simple Example
[math.go](math.go) defines the functions under test; [math_test.go](math_test.go) tests them. A minimal single-case test:

```go
func TestAdd(t *testing.T) {
    result := Add(3, 4)
    expected := 7
    if result != expected {
        t.Errorf("Add(3, 4) = %d; want %d", result, expected)
    }
}
```

The table-driven form for the same kind of function:

```go
func TestSubtract(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive result", 10, 3, 7},
        {"negative result", 3, 10, -7},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Subtract(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

## How It Works
`TestDivide` in [math_test.go](math_test.go) is the clearest example of testing error paths deliberately: its table includes a `wantError bool` field, and the divide-by-zero case (`{"divide by zero", 10, 0, 0, true}`) asserts that `Divide` *does* return a non-nil error — not just that the happy-path cases return the right quotient. A test suite that only checks `Add(3, 4) == 7` and never checks what happens on bad input misses exactly the cases most likely to break in production.

`TestFactorial` includes `{-1, -1}` — the function's documented behavior for invalid input (negative `n`) — as a table row alongside normal cases, so the "what should happen when the input is wrong" behavior is pinned down by a test just like everything else, not left as an assumption.

Benchmarks use a different signature, `func BenchmarkXxx(b *testing.B)`, and loop `b.N` times — Go automatically adjusts `b.N` until the timing is statistically stable. They only run with `go test -bench=.` (they're skipped by a plain `go test`), because measuring performance is a different concern from checking correctness.

## Common Mistakes
- Writing only happy-path tests and skipping error/edge cases — `Divide(10, 0)`, `Factorial(-1)`, and an empty-string `IsPalindrome("")` are exactly the inputs bugs hide behind.
- Not using `t.Run`/table-driven structure once there are more than two or three similar cases — copy-pasted `TestXxxCase1`, `TestXxxCase2` functions get unwieldy and hide which specific input failed.
- Using `t.Fatal`/`t.Fatalf` inside a loop over table cases in a way that stops the *entire* test on the first failure — inside a `t.Run` subtest closure this is fine (it only stops that subtest), but calling it directly in the outer loop (not inside `t.Run`) would abort testing every remaining case.
- Forgetting that benchmarks need `-bench=.` — a plain `go test` silently skips every `BenchmarkXxx` function.

## Best Practices
- Default to table-driven tests for any function with more than one interesting case — it's the idiomatic Go style for a reason: low boilerplate, clear per-case naming, easy to extend.
- Give every table row a descriptive `name` field — it becomes the subtest name in `-v` output and in failure reports, so `TestDivide/divide_by_zero` is immediately meaningful.
- Explicitly test error-returning functions in both directions: an input that should succeed, and one that should fail — verify the error is non-nil, not just that the success path returns the right value.
- Run `go test -cover ./...` regularly to see what isn't tested yet, and `go test -race ./...` on anything involving goroutines or shared state (from the concurrency lessons).

## Real-World Example
Every serious Go project runs `go test ./...` in CI on every commit and pull request — this is the single most standard automated gate in the Go ecosystem, usually paired with `go vet` and `golangci-lint`. Table-driven tests are how HTTP handlers get tested too: a table of `{name, method, path, body, wantStatus, wantBody}` driven through `httptest.NewServer` or `httptest.NewRecorder` is the same pattern you just learned, applied to the HTTP lesson's handlers instead of `math.go`'s functions.

## Exercise
Write a table-driven test for `IsPalindrome` (already defined in [math.go](math.go)) that adds at least three new cases not already in [math_test.go](math_test.go) — for example, a phrase with mixed case, a string with spaces, and a numeric string — and note in a comment which of them currently pass or fail given the function's actual (simple, case-sensitive) implementation.

## Mini Project
Pick any one function you wrote in an earlier lesson — the JSON config loader, the file-backed key-value store, or a generic utility like `Unique`/`Find` — and write a complete table-driven test suite for it, covering at least one error/edge case, using `t.Run` for each case. Run it with `go test -v ./...` and `go test -cover ./...` and confirm it passes with good coverage.

## What to Learn Next
This completes the current course content — congratulations on making it from Go's fundamentals all the way through concurrency, file I/O, JSON, HTTP, generics, and testing. That's a genuinely solid, production-relevant foundation in Go.

The next phase, tracked in [`../ROADMAP.md`](../ROADMAP.md), covers building REST APIs with real databases (PostgreSQL), logging, configuration management, clean architecture, microservices, and other production-level topics.

Before moving on, go build something real with what you already know — a CLI tool, a small REST API with an in-memory store, a file-backed utility. Applying these lessons to a project of your own is worth more than any additional lesson right now, and it will make the upcoming database and architecture material land much better when you get there.
