# 27 – Goroutines

## What You Will Learn
- What a goroutine is, and how it differs from an operating-system thread
- How to start a goroutine with the `go` keyword
- Why `main()` exiting kills every goroutine immediately, with no automatic waiting
- How to safely capture loop variables inside goroutines
- The "fan-out" pattern: starting many goroutines and collecting their results
- What a goroutine leak is, and a first glimpse of `select` used for timeouts

## Why This Matters
Every program you've written so far in this course has run one instruction at a time. That's fine for a script, but real software — web servers, data pipelines, chat apps — needs to do many things *at once*: handle 10,000 client connections, fetch from three APIs in parallel, or process a queue while still responding to new requests.

Other languages give you OS threads for this, but OS threads are expensive: each one reserves megabytes of stack memory and costs the kernel time to schedule. Spawning thousands of them will bring a machine to its knees. Go's answer is the **goroutine** — a unit of concurrency managed entirely by the Go runtime, not the OS. Because goroutines start with a tiny (~2KB) stack that grows and shrinks as needed, you can realistically run hundreds of thousands of them on ordinary hardware. This changes how you're allowed to *design* programs: instead of avoiding concurrency because it's expensive, in Go you reach for it whenever a task is independent — "handle this request in its own goroutine" is a completely normal, cheap decision.

## Concept Explanation
A goroutine is created by putting `go` before a function call:

```go
go printMessage(1, "goroutine A")
```

This tells the Go runtime "start running this function concurrently" and immediately returns control to the caller — it does **not** wait for the function to finish. The calling code and the new goroutine now run independently, scheduled by Go's runtime across the available OS threads (controlled by `GOMAXPROCS`, which defaults to your number of CPUs).

The critical mental model to internalize:

```
Sequential (normal):               Concurrent (goroutines):
─────────────────────              ──────────────────────────────────
main ──► task1 (wait)              main ──► go task1() ──► [runs in background]
         │                                 │
         ▼ (done)                          ──► go task2() ──► [runs in background]
main ──► task2 (wait)                      │
         │                                 ──► go task3() ──► [runs in background]
         ▼ (done)                          │
main ──► task3 (wait)                      └──► all three run AT THE SAME TIME
```

The **main goroutine** (the one running `main()`) is not special in how it executes, but it is special in one crucial way: when `main()` returns, the whole program exits — instantly, without waiting for any other goroutine to finish. There is no implicit "join all threads" step like some other languages provide. If you don't explicitly wait, your goroutines may never get to run at all.

In this lesson, the file uses `time.Sleep` as a crude way to wait ("crude wait (use WaitGroup in real code)"). That's intentional for teaching purposes — real code uses `sync.WaitGroup` (lesson 30) or channels (lesson 28), both of which wait for actual completion instead of guessing a duration.

## Simple Example
Starting two goroutines and giving them time to run ([main.go](main.go)):

```go
go printMessage(1, "goroutine A")
go printMessage(2, "goroutine B")
// We need to wait — otherwise main() exits and kills both goroutines
time.Sleep(300 * time.Millisecond) // crude wait (use WaitGroup in real code)
```

The fan-out pattern — many goroutines feeding results into one channel:

```go
results := make(chan string, 5) // buffered channel — can hold 5 items
for i := 1; i <= 5; i++ {
    go fetchData(i, results)
}
for i := 0; i < 5; i++ {
    fmt.Println(" ", <-results)
}
```

## How It Works
`go fetchData(i, results)` launches five independent goroutines almost instantly; Go's scheduler interleaves their execution (and may run several truly in parallel across CPU cores). Each one sleeps for a different simulated duration and then sends its result into the shared buffered channel. The main goroutine then reads exactly five values from `results` — this receive loop is itself a form of synchronization: it blocks until each result exists, so by the time all five prints happen, we know every worker has finished.

Notice the loop-variable capture in section 3 of `main.go`:

```go
for i := 1; i <= 3; i++ {
    i := i // shadow i — each goroutine gets its own copy
    go func() {
        fmt.Printf("  worker %d started\n", i)
    }()
}
```

This file demonstrates the manual-shadowing style (`i := i`), which was the required idiom in Go versions before 1.22. As of Go 1.22+, the language itself changed loop semantics so each iteration gets its own fresh variable automatically, making the shadow line technically redundant on modern Go — but it's harmless, self-documenting, and still correct on any Go version, which is exactly why the example keeps it.

At the very end, `main.go` includes a short preview of `select`:

```go
select {
case <-done:
    fmt.Println("  goroutine finished in time")
case <-time.After(500 * time.Millisecond):
    fmt.Println("  goroutine timed out")
}
```

This races "did the goroutine finish?" against "did 500ms pass?" — a very common pattern for adding a timeout to concurrent work. Don't worry about fully understanding `select`'s rules yet; lesson 29 (`29Select`) is dedicated entirely to it. For now, just recognize the shape: it's a way to wait on more than one channel at once.

## Common Mistakes
- **Assuming goroutines are awaited automatically.** They are not. If `main()` reaches its end, every running goroutine is terminated mid-flight, even ones that were "almost done."
- **Goroutine leaks.** A goroutine that blocks forever — waiting on a channel nobody sends to, or looping without an exit condition — never gets cleaned up. It just sits there consuming memory for the life of the program. Always give a goroutine a way out (a stop channel, a `context`, or a timeout).
- **Capturing a loop variable incorrectly.** On Go versions before 1.22, writing `go func() { use(i) }()` inside a `for i := ...` loop without shadowing `i` meant every goroutine could see the *same*, final value of `i` by the time it actually ran. The fix shown in this lesson (`i := i`) creates a fresh copy per iteration. On Go 1.22+ this is fixed at the language level, but writing the shadow copy (or passing `i` as a function parameter, e.g. `go func(n int) { use(n) }(i)`) is still good, portable practice.
- **Using `time.Sleep` as a substitute for real synchronization** in production code. It's used here only as a simple teaching stand-in; it's fragile because it guesses how long work takes instead of actually waiting for it.

## Best Practices
- Always have an explicit plan for how the main goroutine will wait for others: `sync.WaitGroup`, a channel receive, or `context` cancellation.
- Give every goroutine a clear termination condition before you write the code that starts it.
- Prefer passing values into a goroutine's closure explicitly (as a parameter, or via a shadowed local) rather than relying on capturing outer variables.
- Keep goroutine bodies focused on one job; use channels to communicate results back rather than shared variables (more on why in lesson 28).

## Real-World Example
A typical Go web server (built with `net/http`) spins up a **new goroutine for every incoming HTTP request** automatically — that's how the standard library's server achieves high concurrency without threads-per-request being prohibitively expensive. Thousands of simultaneous requests can each get their own goroutine, do their work (query a database, call another service), and finish, all without the memory overhead that thousands of OS threads would require. This is precisely the property — cheap, numerous goroutines — that makes Go a popular choice for backend services.

## Exercise
Write a program that starts 4 goroutines, each printing its own worker number and then sleeping for `(workerNumber * 25) * time.Millisecond`. Use the loop-variable-capture technique shown in this lesson so each goroutine prints the correct number. Use `time.Sleep` in `main` to give them enough time to finish before the program exits (you'll replace this crude approach with `sync.WaitGroup` in lesson 30).

## Mini Project
Build a tiny "concurrent downloader simulator": write a function `simulateDownload(id int, ms int, result chan<- string)` that sleeps for `ms` milliseconds and then sends a message like `"file 3 downloaded"` into `result`. In `main`, launch 6 simulated downloads with different millisecond values (using a buffered channel sized 6), then collect and print all 6 results. Optionally, add the timeout `select` pattern from this lesson around the whole batch, so if it takes longer than expected you print "download batch timed out" instead of hanging forever.

## Summary
Goroutines are lightweight, runtime-managed units of concurrency — cheap enough to start by the thousands, which is why Go programs default to "just spawn a goroutine" for independent work instead of avoiding concurrency. But cheap doesn't mean automatic: nothing waits for a goroutine unless you tell it to, and `main()` exiting kills everything instantly. You saw the `go` keyword, safe loop-variable capture, the fan-out-and-collect pattern, and a first peek at `select` for timeouts. The channel used to collect results (`results := make(chan string, 5)`) is itself the subject of the next lesson.

## What to Learn Next
Previous: [26Packages](../26Packages/README.md)

Next: learn how goroutines actually communicate and synchronize with each other in [28Channels](../28Channels/README.md).
