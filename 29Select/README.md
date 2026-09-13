# 29 – Select Statement

## What You Will Learn
- Why you need `select` once a goroutine has more than one channel to watch
- How `select` picks a ready case (and what happens when several are ready)
- Non-blocking channel operations using `default`
- The timeout pattern with `time.After`
- How a `nil` channel behaves inside a `select` (and why that's useful)

## Why This Matters
In lesson 28, every example dealt with exactly one channel at a time: send on it, receive from it, range over it. But real programs quickly end up needing to wait on *several* channels simultaneously — a data channel and a cancellation channel, two independent data sources, a ticker and a stop signal. A plain `<-ch` can only wait on one channel; if you tried calling it on channel A and then channel B in sequence, you'd be stuck waiting for A even if B became ready first.

`select` solves exactly this problem: it lets a goroutine block on multiple channel operations at once and proceed with whichever one becomes ready first. You already saw a preview of this in lesson 27, racing a result channel against `time.After` for a timeout. This lesson gives you the full picture: how `select` chooses among several ready channels, how to make it non-blocking, and a couple of easy-to-hit pitfalls with `nil` channels and timers.

## Concept Explanation
`select` looks like a `switch`, but every case is a channel operation, not a value comparison:

```
select {
   │
   ├─── case <-ch1:  ──► ch1 has data? run this
   │
   ├─── case <-ch2:  ──► ch2 has data? run this
   │
   ├─── case ch3<-v: ──► ch3 ready to receive? run this
   │
   └─── default:     ──► nothing ready? run immediately (non-blocking)
```

The rules:
- If exactly one case's channel is ready, that case runs.
- If **multiple** cases are ready at the same time, Go picks **one at random** — this is deliberate, so you can't accidentally write code that always favors one channel over another.
- If **no** case is ready and there's **no** `default`, `select` **blocks** until one becomes ready.
- If there's a `default` case, `select` never blocks: it runs `default` immediately whenever no other case is ready yet.

## Simple Example
Waiting on whichever of two channels responds first, from [main.go](main.go):

```go
select {
case msg1 := <-ch1:
    fmt.Println("received:", msg1)
case msg2 := <-ch2:
    fmt.Println("received:", msg2) // prints this — ch2 is ready first
}
```

Non-blocking check with `default`:

```go
select {
case v := <-ch:
    fmt.Println("  received:", v)
default:
    fmt.Println("  nothing to receive — moving on")
}
```

The timeout pattern:

```go
select {
case res := <-result:
    fmt.Println("  success:", res)
case <-time.After(100 * time.Millisecond):
    fmt.Println("  timed out waiting for operation") // fires because 300ms > 100ms
}
```

## How It Works
In the first example, two goroutines each sleep for a different duration and then send on their own channel. `select` is watching both `ch1` and `ch2` at once; since `ch2`'s goroutine sleeps for less time (50ms vs 100ms), its send becomes ready first, so `select` runs that case and the `ch1` goroutine's eventual send is simply never received in this particular select (in this program it has nowhere left to be read, which is fine here since the program moves on).

`time.After(d)` returns a channel that automatically receives a value once `d` has elapsed — it doesn't require you to build your own timer goroutine. Racing a real operation's result channel against `time.After`'s channel inside one `select` gives you a timeout for free: whichever fires first wins.

The loop-based example combines `select` with `time.NewTicker`, which repeatedly delivers values on a fixed interval, and a `stop` channel from `time.After` to end the loop after a fixed duration:

```
loop
  │
  └─► select
           ├─ ch1 ready? process
           ├─ ch2 ready? process
           └─ done ready? exit loop
```

Each pass through the `for` loop re-evaluates the `select`, so ticks from either ticker (or the stop signal) get handled as they arrive, in whatever order they actually occur.

Finally, the nil-channel example shows a more advanced but genuinely useful trick: assigning `nil` to a channel variable used in a `select` case disables that case *permanently* (a `nil` channel is never ready to send or receive, so its case can never be chosen). This lets you "turn off" a channel dynamically — for example, after you've already received its one expected value — without restructuring your `select` statement.

## Common Mistakes
- **A `select` with no ready case and no `default` blocks forever.** If none of the channels you're watching will ever become ready, and you didn't add a `default`, the goroutine (and possibly your whole program, if it's the main goroutine) hangs. Always make sure at least one case is guaranteed to eventually fire, or include a `default`/timeout as an escape hatch.
- **Confusing "non-blocking" with "no cost."** Adding `default` makes `select` return immediately when nothing's ready, but that's easy to accidentally spin-loop on (calling `select` with `default` repeatedly in a tight loop burns CPU). The done-channel example in lesson 28 mixes `default` with a `time.Sleep` in the loop body for exactly this reason.
- **Timer leaks in loops.** Calling `time.After(d)` inside a loop creates a brand new timer *every iteration*, and that timer isn't garbage-collected until it fires — if the loop runs many times before each timer expires, you accumulate live timers in memory. For a one-shot timeout like the examples in this lesson, that's fine. But inside a long-running loop, prefer creating one `time.NewTimer` (or `time.NewTicker`, as used in the loop example here) outside the loop and resetting/reusing it, rather than calling `time.After` on every pass.

## Best Practices
- Always ask "what guarantees this select eventually unblocks?" before writing it — either a channel that's certain to fire, or a timeout/default.
- Use `time.After` for simple, one-off timeouts; use `time.NewTicker`/`time.NewTimer` (with `defer ticker.Stop()`) when the same timing logic runs repeatedly in a loop.
- When disabling a `select` case with a `nil` channel, make sure you actually intend "never again" — it's a one-way trip until you reassign a real channel to that variable.
- Keep `select` cases short; if a case needs to do significant work, consider spinning that off into its own function call so the `select` itself stays easy to read.

## Real-World Example
Imagine a service that needs a single answer from three redundant backend replicas, but only wants to wait so long: it fires off a request to each replica (each writing its response into its own channel) and then uses one `select` to grab whichever reply comes back first, with a `time.After` case as an overall deadline. This "fan-in with timeout" is a very common shape in backend systems — call several sources concurrently, take the fastest (or first-good) answer, and never let a slow or dead dependency hang the whole request.

## Exercise
Write a program with two channels, `evens` and `odds`. Start two goroutines: one sends the numbers 2, 4, 6 on `evens` (with a short sleep between each), the other sends 1, 3, 5 on `odds`. In `main`, use a `select` inside a loop to print whichever number arrives, from either channel, until you've printed 6 numbers total.

## Mini Project
Build a "race a slow computation against a timeout" utility: write a function `computeSlowly(ms int, result chan<- int)` that sleeps for `ms` milliseconds, then sends a computed value (e.g. `42`). In `main`, call it with a channel and use `select` with `time.After(200 * time.Millisecond)` to either print the successful result or print "computation timed out" — try it once with `ms` under 200 and once with `ms` over 200 to see both branches run.

## Summary
`select` is what makes it possible to wait on multiple channels at once, choosing randomly among any that are simultaneously ready, blocking if none are ready (unless a `default` case is present), and enabling the extremely common timeout pattern via `time.After`. You also saw that a `nil` channel is a safe, permanent way to disable a `select` case. Combined with the channels from lesson 28, `select` gives you full control over how a goroutine reacts to multiple sources of events — which is exactly the foundation the `context` package (lesson 31) builds its cancellation model on top of.

## What to Learn Next
Previous: [28Channels](../28Channels/README.md)

Next: learn how to safely share data between goroutines with the `sync` package in [30Sync](../30Sync/README.md).
