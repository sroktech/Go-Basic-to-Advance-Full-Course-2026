# 30 – Sync Package

## What You Will Learn
- What a data race is and why it produces wrong, unpredictable results
- `sync.WaitGroup` — waiting for a group of goroutines to finish
- `sync.Mutex` — protecting shared data so only one goroutine touches it at a time
- `sync.RWMutex` — allowing many concurrent readers but only one writer
- `sync.Once` — running initialization code exactly once, no matter how many goroutines ask for it
- `sync/atomic` — lock-free updates for simple counters
- How to catch races automatically with `go run -race`

## Why This Matters
`select` (lesson 29) let you coordinate *which* channel a goroutine reacts to. This lesson is about a different but equally common problem: several goroutines reading and writing the **same variable** at the same time, with no channel involved at all. Unlike some languages, Go does not protect shared memory for you — if two goroutines write to the same variable without coordination, the result is a **data race**: the final value depends on timing, and it can silently be wrong. The `sync` package is the standard toolkit for making shared-memory access safe.

## Concept Explanation
The clearest way to see the problem is the classic "lost increment": two goroutines both read a counter's current value, both compute value+1 based on what they read, and both write it back. If they overlap just right, one increment is silently lost:

```
goroutine 1: reads count=5          goroutine 2: reads count=5
goroutine 1: count = 5+1 = 6        goroutine 2: count = 5+1 = 6
goroutine 1: writes count=6         goroutine 2: writes count=6
RESULT: count=6 (expected 7!)  ← one increment was LOST
```

A **mutex** ("mutual exclusion lock") fixes this by making sure only one goroutine can be inside the critical section — the read-modify-write — at a time:

```
goroutine 1: Lock() ──► reads count=5, writes count=6, Unlock()
goroutine 2:                                             Lock() ──► reads count=6, writes count=7, Unlock()
RESULT: count=7 ✓
```

`sync.WaitGroup` solves a related but different problem: not "who can touch this data," but "how do I know when a group of goroutines has finished?" You tell it how many goroutines to expect (`Add`), each one reports in when done (`Done`), and the caller blocks until they all have (`Wait`).

## Simple Example
Waiting for several goroutines to finish, from [main.go](main.go):

```go
var wg sync.WaitGroup
for i := 1; i <= 5; i++ {
    wg.Add(1) // one more goroutine to wait for
    go func() {
        defer wg.Done() // always defer Done, right after Add
        // ... work ...
    }()
}
wg.Wait() // blocks until all 5 have called Done
```

Protecting shared data with a mutex:

```go
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++ // safe — only one goroutine can be here at a time
}
```

## How It Works
`demoWaitGroup` starts 5 goroutines, calling `wg.Add(1)` once per goroutine *before* starting it (not inside it — that avoids a race on the counter itself). Each goroutine defers `wg.Done()`, so it reports in no matter how it exits. `wg.Wait()` blocks the calling goroutine until the internal counter drops back to zero.

`demoMutex` wraps an `int` and a `sync.Mutex` together in a `SafeCounter` struct, with `Lock`/`Unlock` inside its methods rather than scattered at call sites — this keeps the locking rule attached to the data it protects, so callers can't forget it. 1000 goroutines all call `Increment()`; because each call locks before touching `count`, the final value is reliably 1000, not some smaller number lost to overlapping reads.

`demoRWMutex` shows `sync.RWMutex`, which distinguishes reads from writes: `RLock`/`RUnlock` for reads (many goroutines can hold a read lock simultaneously), `Lock`/`Unlock` for writes (exclusive — blocks everyone else). This is a real optimization for data that's read far more often than it's written, like a cache.

`demoOnce` uses `sync.Once` to build a simple singleton: no matter how many goroutines call `getInstance()` concurrently, the initialization closure passed to `once.Do` runs exactly one time; every other call just returns the already-created value.

`demoAtomic` shows `sync/atomic` as a lighter-weight alternative to a mutex for simple numeric operations: `atomic.AddInt64` performs the read-modify-write as one indivisible CPU-level operation, so 1000 concurrent increments still land on exactly 1000 with no explicit lock at all.

## Common Mistakes
- **Forgetting to unlock a mutex.** If you call `Lock()` and an early return or a panic skips the matching `Unlock()`, every other goroutine waiting on that lock hangs forever. Always write `defer mu.Unlock()` on the line right after `mu.Lock()`, so it's unconditional.
- **Copying a mutex (or WaitGroup) by value.** `sync.Mutex` and `sync.WaitGroup` must never be copied after first use — copying duplicates their internal state, and the copy no longer coordinates with the original. This is why `SafeCounter` embeds `mu sync.Mutex` directly and its methods use a pointer receiver (`*SafeCounter`) — a value receiver would copy the mutex on every call. `go vet` catches many (not all) accidental copies.
- **Calling `wg.Add()` from inside the goroutine instead of before starting it.** If the main goroutine calls `wg.Wait()` before a late `Add()` runs, the count can briefly be zero and `Wait()` returns too early.
- **Reaching for a mutex when a channel would be more idiomatic**, or vice versa. Go's proverb is "share memory by communicating" (channels) rather than "communicate by sharing memory" (mutexes) — but plenty of real Go code correctly uses mutexes for simple shared state like a counter or a cache. Use whichever makes the *ownership* of the data clearest; don't reach for the more complex tool out of habit.

## Best Practices
- Pair every `Lock()` with an immediately-following `defer Unlock()`.
- Keep critical sections (the code between `Lock` and `Unlock`) as small as possible — don't do slow work (network calls, file I/O) while holding a lock.
- Use `sync.RWMutex` only when reads genuinely dominate writes; otherwise a plain `Mutex` is simpler and just as fast in practice.
- Use `sync/atomic` for single counters/flags; reach for a `Mutex` as soon as you're protecting more than one related field, since atomics can't keep multiple fields consistent with each other.
- Run `go run -race` (or `go test -race`) routinely during development — the race detector finds real races that might not show up in ordinary testing, because races are timing-dependent and can pass by luck.

## Real-World Example
A web server handling many concurrent requests often keeps small pieces of shared, in-memory state — a hit counter, a cache of recently computed results, a map of active sessions. Every one of those needs exactly this kind of protection: a `Mutex` (or `RWMutex` for a read-heavy cache) around the shared map, or `sync/atomic` for a simple request counter, so that concurrent requests never corrupt each other's view of the data.

## Exercise
Write a program that starts 100 goroutines, each incrementing a shared `int` counter (protected by a `sync.Mutex`) 10 times, for 1000 total increments. Use a `sync.WaitGroup` to wait for all goroutines to finish, then print the final counter value — it should always print exactly 1000. Then, as an experiment, comment out the `Lock()`/`Unlock()` calls and run it a few times with `go run -race` to see the race detector flag the problem.

## Mini Project
Build a tiny concurrent "visit counter" for a set of web pages: a `PageStats` struct holding a `map[string]int` protected by a `sync.RWMutex`, with a `Hit(page string)` method (write-locked) that increments a page's count, and a `Count(page string) int` method (read-locked) that returns it. Launch several goroutines calling `Hit` on a handful of page names concurrently, then print the final counts once a `sync.WaitGroup` confirms they're all done.

## Summary
Concurrent access to shared data needs explicit coordination in Go — `sync.WaitGroup` waits for goroutines to finish, `sync.Mutex`/`sync.RWMutex` protect shared data from simultaneous access, `sync.Once` guarantees one-time initialization, and `sync/atomic` gives a lightweight option for simple counters. The `-race` flag turns "hope it's fine" into "verified it's fine." With this in place, you're ready for `context` (lesson 31), which builds on goroutines and channels to manage cancellation and timeouts across a whole call chain.

## What to Learn Next
Previous: [29Select](../29Select/README.md)

Next: learn how to cancel and time out work across a whole call chain with [31Context](../31Context/README.md).
