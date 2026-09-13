# 31 – Context

## What You Will Learn
- What problem `context.Context` solves: cancelling and timing out in-progress work
- `context.Background()` / `context.TODO()` as root contexts
- `context.WithCancel` — manual cancellation
- `context.WithTimeout` / `context.WithDeadline` — automatic cancellation after a duration or at a fixed time
- `context.WithValue` — carrying request-scoped data through a call chain
- Why you must always call the returned `cancel()` function

## Why This Matters
By now you can start goroutines (lesson 27), coordinate them with channels (28), wait on several channels at once (29), and protect shared state (30). One problem remains: once several layers of function calls and goroutines are involved, how does a *cancellation* or *timeout* signal reach all of them consistently? If a user closes their browser tab mid-request, or a request should give up after 2 seconds, every goroutine working on that request needs to find out and stop — without every function along the way inventing its own ad-hoc "stop channel" parameter. `context.Context` is Go's standard, uniform answer to that problem, and it is the mechanism nearly every real Go HTTP server and client uses.

## Concept Explanation
A `context.Context` is a value you pass as the first parameter into functions that do cancellable or time-bounded work. It flows **down** through a call chain, and a single cancellation at the top can reach every goroutine below it:

```
context.Background()
       │
       └─► WithCancel(parent) ──► childCtx + cancel()
                 │
                 ├─► goroutine 1 watches ctx.Done()
                 ├─► goroutine 2 watches ctx.Done()
                 └─► goroutine 3 watches ctx.Done()

call cancel() ──► ctx.Done() closes ──► all goroutines receive the signal and stop
```

Every context has a `Done()` method returning a channel that's closed exactly when the context is cancelled or its deadline passes — closing a channel is how Go broadcasts "this happened" to any number of goroutines at once (every receiver on a closed channel gets a zero value immediately, forever). `ctx.Err()` then tells you *why*: `context.Canceled` or `context.DeadlineExceeded`.

Contexts form a tree: `context.WithTimeout(parent, d)` and friends create a **child** context that inherits the parent's cancellation (if the parent is cancelled, so is every child) but can also add its own, tighter deadline.

## Simple Example
A worker that stops as soon as its context says to, from [main.go](main.go):

```go
func worker(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("  worker %d stopped: %v\n", id, ctx.Err())
            return
        default:
            fmt.Printf("  worker %d working...\n", id)
            time.Sleep(100 * time.Millisecond)
        }
    }
}

ctx, cancel := context.WithCancel(context.Background())
go worker(ctx, 1)
// ... later ...
cancel() // tells worker to stop
```

Racing real work against a deadline:

```go
ctx2, cancel2 := context.WithTimeout(context.Background(), 500*time.Millisecond)
defer cancel2()

result, err := fetchData(ctx2, "api.example.com")
if err != nil {
    fmt.Println("  error:", err) // context.DeadlineExceeded, if it took too long
}
```

## How It Works
`context.WithCancel` returns a context plus a `cancel()` function; calling `cancel()` closes that context's `Done()` channel, which every goroutine holding that context sees via `select` immediately, regardless of how many goroutines are watching.

`context.WithTimeout(parent, d)` is really `WithDeadline(parent, time.Now().Add(d))` under the hood — both close `Done()` automatically once the deadline passes, no manual `cancel()` call required for that to happen (though you must still call the returned `cancel()` yourself — see below).

`fetchData` demonstrates the idiomatic pattern for "do work that might be cancelled": start the real work in a goroutine writing to a result channel, then `select` between that channel and `ctx.Done()`, so whichever happens first — success or cancellation — wins.

The call chain example (`serviceA` → `serviceB` → `fetchData`) shows context flowing through ordinary function parameters — it is just a regular value, passed explicitly, not hidden global state. `serviceB` even creates its own child context with a *shorter* timeout than the one it received, showing that inner layers can tighten (but never loosen) an overall deadline.

`context.WithValue` is different in kind: instead of carrying cancellation, it attaches a key-value pair that any function holding that context can read back with `ctx.Value(key)`. The example defines a private `contextKey` type specifically so its keys can never collide with a key defined in another package (a plain `string` key like `"requestID"` could clash with someone else's identical string).

## Common Mistakes
- **Not calling the returned `cancel()` function.** Every `WithCancel`/`WithTimeout`/`WithDeadline` call returns a `cancel` function, and Go's own documentation says to call it even if the context's work finishes normally — otherwise the context (and any timer backing it) stays alive, leaking resources, until its parent is cancelled or the process exits. The idiomatic fix is `defer cancel()` immediately after creating the context, exactly as this lesson's code does every time.
- **Passing `nil` instead of a context.** Every function that needs a context should require one explicitly; use `context.Background()` at the top of `main`/tests, or `context.TODO()` as a placeholder while you're still deciding, never `nil`.
- **Using `WithValue` for things that should be ordinary parameters.** Context values are for cross-cutting, request-scoped metadata (a request ID, an auth token, a trace ID) that many layers need without each one explicitly threading it through — not a way to avoid writing normal function arguments. If a value is required for a function to do its job correctly, make it a real parameter; hiding it in the context makes the function's dependencies invisible from its signature.
- **Storing large or mutable data in context.** Context values should be small, immutable, and safe to read concurrently — it's not a general-purpose data bag.

## Best Practices
- Accept `context.Context` as the **first parameter** of any function that does I/O, might block, or might run for a while — this is a strong, near-universal Go convention.
- Always pair a cancellable/timeout context with `defer cancel()`.
- Propagate the context you were given rather than creating a fresh `context.Background()` deeper in a call chain — that's how a top-level timeout or cancellation actually reaches the bottom.
- Define a private, unexported type for context keys (as shown here) instead of using a raw `string` or other built-in type, to guarantee no collisions with keys from other packages.

## Real-World Example
An HTTP server using `net/http` automatically gives every incoming request a `context.Context` (`r.Context()`) that is cancelled the moment the client disconnects or the request's timeout expires. A handler that kicks off a slow database query or an outbound API call should pass that same context down so the query/call is abandoned the instant it's no longer needed — instead of wastefully running to completion for a client that already gave up.

## Exercise
Write a function `countTo(ctx context.Context, n int)` that prints `1, 2, 3, ...` up to `n`, sleeping 100ms between numbers, but returns early (printing why, via `ctx.Err()`) if the context is cancelled first. Call it once with `context.WithTimeout(context.Background(), 350*time.Millisecond)` and `n = 10`, and confirm it stops around 3 or 4 instead of reaching 10.

## Mini Project
Build a small "fetch with cancellation" helper: a function `fetch(ctx context.Context, delayMs int) (string, error)` that simulates a network call (sleep `delayMs`, then return a canned string) but respects `ctx.Done()` via `select`, exactly like `fetchData` in this lesson. In `main`, call it three times with a shared 300ms-timeout context and delays of 100ms, 250ms, and 500ms, and print whether each call succeeded or timed out.

## Summary
`context.Context` is how cancellation, deadlines, and small pieces of request-scoped data flow consistently through a call chain and across goroutines — `Done()` gives every interested goroutine a channel to watch, `Err()` explains why it closed, and `WithCancel`/`WithTimeout`/`WithDeadline` create contexts that close it under different conditions, while always requiring you to call the returned `cancel()`. This closes out the Intermediate tier's concurrency toolkit: goroutines, channels, select, sync, and now context. From here, the course moves into using these tools to build real things — file I/O, JSON, and HTTP.

## What to Learn Next
Previous: [30Sync](../30Sync/README.md)

Next: put concurrency aside for a moment and work with the filesystem in [32FileIO](../32FileIO/README.md).
