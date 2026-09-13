# 28 – Channels

## What You Will Learn
- What a channel is, and the Go proverb behind it: "Do not communicate by sharing memory; share memory by communicating"
- The difference between unbuffered and buffered channels, and why that difference matters
- Channel directions (`chan<- T`, `<-chan T`) and why they exist
- The producer/consumer pattern with `range` and `close`
- How to chain channels together into a pipeline
- How to use a "done" channel to tell a goroutine to stop

## Why This Matters
In lesson 27 you learned that goroutines run independently — but independent workers are only useful if they can hand results back to each other and coordinate timing. The naive approach (have two goroutines read and write the same shared variable) is exactly what causes the race conditions and coordination nightmares that gave concurrent programming its bad reputation in other languages.

Go's core answer is the **channel**: a typed, thread-safe pipe that goroutines use to pass data between each other. Crucially, an **unbuffered** channel doesn't just transfer data — sending and receiving on it forces both sides to rendezvous at the same moment in time. That means an unbuffered channel operation is *also* a synchronization point, for free: if you receive a value, you now know for certain that the sender reached that exact line of code. This is why channels replace not just "how do I pass data" but also "how do I know the other goroutine got there yet" — one tool, two jobs. A **buffered** channel relaxes this: it lets a sender put values in without an immediate matching receiver (up to its capacity), trading that instant hand-off guarantee for decoupling and throughput.

## Concept Explanation
A channel is created with `make`:

```go
make(chan int)     // unbuffered — synchronous rendezvous
make(chan int, 5)  // buffered with capacity 5 — asynchronous up to 5 items
```

Unbuffered channel flow — both sides must be present at the same moment:

```
Sender goroutine          Channel          Receiver goroutine
────────────────     ───────────────     ────────────────────
ch <- value    ───►  [  value  ]   ───►  value := <-ch
(blocks until               │            (blocks until
 receiver is ready)         │             sender sends)
```

Buffered channel flow (capacity 3) — sends succeed without a receiver, until the buffer fills:

```
send 1 ──► [1][ ][ ]   ← doesn't block (space available)
send 2 ──► [1][2][ ]   ← doesn't block
send 3 ──► [1][2][3]   ← doesn't block
send 4 ──► BLOCKS       ← buffer full, waits for receiver

recv   ──► [2][3][ ]   ← makes space, send 4 can proceed
```

Channel **direction** types restrict how a channel can be used inside a function signature:

```go
chan T      // bidirectional
chan<- T    // send-only  (can only put values in)
<-chan T    // receive-only (can only take values out)
```

This is purely a compile-time safety feature: if a function only needs to *send*, declaring its parameter as `chan<- int` means the compiler will reject any attempt inside that function to accidentally receive from it — the type system documents and enforces intent.

## Simple Example
Producer and consumer from [main.go](main.go):

```go
func producer(ch chan<- int, count int) {
    for i := 1; i <= count; i++ {
        ch <- i // send value into channel
    }
    close(ch) // signal that no more values will be sent — IMPORTANT
}

func consumer(ch <-chan int) {
    for v := range ch { // loops until the channel is closed
        fmt.Printf("  consumer: received %d\n", v)
    }
    fmt.Println("  consumer: channel closed, done")
}
```

A two-stage pipeline, where each stage is its own goroutine connected by a channel:

```go
generate := func(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}
```

## How It Works
`producer` sends five values one at a time into `ch`, then calls `close(ch)` to announce "nothing more is coming." `consumer` uses `for v := range ch`, which keeps pulling values out until it detects the channel is closed *and* drained, at which point the loop ends by itself — no manual bookkeeping needed. This `close` + `range` combination is the standard way to signal "stream finished" in Go.

The pipeline pattern (`generate` → `double`) chains this same idea: each stage is a function that spins up its own goroutine, reads from an input channel, transforms each value, writes to an output channel, and closes that output channel when its input closes. Because `range` propagates the "closed" signal forward automatically, closing the very first channel (inside `generate`) eventually causes the entire chain to wind down cleanly.

```
generate  ──► ch1 ──► double ──► ch2 ──► print
(1,2,3,4,5)          (x*2)               (2,4,6,8,10)
```

The "done channel" example shows a different use for a channel: not to carry data, but purely as a **signal**. `done := make(chan struct{})` uses the empty struct type, which occupies zero bytes — the channel is used only for its closing event, not for any value it carries:

```
main                    worker goroutine
────                    ───────────────────────────────
done := make(chan struct{})
go worker(done)  ──────►  for { select { case <-done: return } }
...
close(done)      ──────►  <-done unblocks, worker exits
```

Calling `close(done)` broadcasts to *every* goroutine listening on it simultaneously — closing a channel, unlike sending a value, wakes up all current and future receivers at once.

## Common Mistakes
- **Deadlock: sending with no one to receive.** An unbuffered `ch <- v` (or a buffered channel with a full buffer) blocks forever if no goroutine is ever going to receive. If that's the *only* thing your program is doing, Go's runtime detects it and panics with "fatal error: all goroutines are asleep - deadlock!"
- **Deadlock: receiving with no one to send.** Symmetric problem — `<-ch` blocks forever if nothing will ever send (or close) that channel.
- **Forgetting to `close` a channel that a receiver ranges over.** `for v := range ch` never exits on its own if `ch` is never closed — the receiving goroutine simply hangs forever (a goroutine leak, as introduced in lesson 27). Only the sender should close a channel, and only once it's truly done sending.
- **Sending on a closed channel panics.** Once you `close(ch)`, any further `ch <- v` from any goroutine crashes the program immediately with `panic: send on closed channel`. Receiving from a closed channel, in contrast, is safe — it returns the zero value and `ok == false` (as seen in the nil-channel example in lesson 29).

## Best Practices
- The sender closes the channel, never the receiver — a receiver generally doesn't know when the last value has been sent.
- Use direction-restricted channel types (`chan<-`, `<-chan`) in function signatures whenever a function only sends or only receives, so mistakes are caught at compile time.
- Reach for a buffered channel when you want to decouple producer speed from consumer speed (like the fan-out collector in lesson 27); reach for an unbuffered channel when you specifically need the rendezvous guarantee.
- Use a `chan struct{}` when a channel exists purely to signal an event, not to carry data — it communicates intent and costs no memory per signal.

## Real-World Example
Picture a background job system: a pool of "worker" goroutines all read job descriptions from one shared channel (the queue), process each job, and write results to another channel that a separate "collector" goroutine drains and persists to a database. This is the **worker pool** pattern, and it's how many real Go services process bounded, high-throughput work — image resizing services, batch email senders, log processors — without needing a heavyweight message broker for in-process work distribution.

## Exercise
Write a program with a function `squareAll(nums []int) <-chan int` that starts a goroutine sending each number's square into a channel, closing the channel when done. In `main`, call it and use `for v := range` to print every squared value. Then modify it into a two-stage pipeline: first a stage that filters out odd numbers, then a stage that squares what's left, connected the way `generate` and `double` are connected in this lesson.

## Mini Project
Build a tiny concurrent word counter: given a slice of strings (sentences), start one goroutine per sentence that counts its words and sends the count into a shared buffered channel sized to the number of sentences. In `main`, receive all the counts and sum them to get a total word count across all sentences. Then extend it with a "done" channel: add a second goroutine that just prints a heartbeat message every 50ms until you close a `done` channel telling it to stop, demonstrating the stop-signal pattern from this lesson.

## Summary
Channels are Go's built-in, type-safe way for goroutines to pass data and to synchronize, replacing shared-memory access with explicit communication. Unbuffered channels rendezvous the sender and receiver at the same instant; buffered channels decouple them up to a capacity. `close` plus `range` is the standard way to signal "no more values are coming," and a `chan struct{}` is a lightweight pure signal. You also saw two new tools in passing — `select` used to watch a `done` channel non-blockingly with `default`, and channel directions for API safety — both of which the upcoming lessons build on directly.

## What to Learn Next
Previous: [27Goroutines](../27Goroutines/README.md)

Next: learn `select` — the tool for waiting on multiple channels at once — in [29Select](../29Select/README.md).
