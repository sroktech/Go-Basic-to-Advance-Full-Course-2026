// 25 - Defer in Go
//
// 'defer' schedules a function call to run WHEN THE SURROUNDING FUNCTION RETURNS —
// whether it returns normally, hits a panic, or returns early via 'return'.
//
// This makes defer perfect for cleanup tasks:
//   - Closing files after opening them
//   - Releasing locks after acquiring them
//   - Closing database connections
//   - Logging function entry/exit
//
// FLOW — single defer:
//
//   func doWork() {
//       defer cleanup()   ← registered, but NOT called yet
//       step1()
//       step2()
//       return            ← cleanup() fires HERE, just before returning
//   }
//
// FLOW — multiple defers (LIFO order: last in, first out):
//
//   defer A()   ← registered first, runs LAST
//   defer B()   ← registered second, runs second
//   defer C()   ← registered last, runs FIRST
//
//   Return order: C → B → A
//   (like a stack — each defer is pushed, then popped on exit)

package main

import (
	"fmt"
	"os"
)

// ─── Basic defer — cleanup pattern ────────────────────────────────────────────

func openAndRead(filename string) {
	fmt.Println("Opening file:", filename)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening:", err)
		return
	}
	// defer close IMMEDIATELY after open — never forget to close
	// Even if the function panics or returns early, this will run
	defer f.Close()

	fmt.Println("File opened successfully, doing work...")
	// ... read from file ...
	fmt.Println("Done with file — defer will close it now")
}

// ─── Defer order: LIFO ────────────────────────────────────────────────────────

func demoLIFO() {
	fmt.Println("demoLIFO: start")
	defer fmt.Println("  defer 1 — registered first, runs LAST")
	defer fmt.Println("  defer 2")
	defer fmt.Println("  defer 3 — registered last, runs FIRST")
	fmt.Println("demoLIFO: end of function body")
}

// ─── Defer evaluates arguments eagerly ───────────────────────────────────────
//
// The ARGUMENTS to a deferred function are evaluated IMMEDIATELY at the defer statement,
// not when the deferred function actually runs.
//
// FLOW:
//   i := 0
//   defer fmt.Println(i)  ← i is evaluated NOW (captures 0)
//   i = 10
//   return                ← deferred call runs with i=0, not 10

func demoEagerArgs() {
	i := 0
	defer fmt.Println("  deferred i (captured at defer time):", i) // captures 0 NOW
	i = 10
	fmt.Println("  i at end of function:", i) // 10
	// defer will print 0, not 10
}

// ─── Defer with named return values ──────────────────────────────────────────
//
// If a function uses NAMED return values, a deferred function CAN modify them.
// This is used for "deferred error wrapping" patterns.
//
// FLOW:
//   err is the named return
//   defer runs AFTER return sets err, but BEFORE caller receives it
//   defer can change err's final value

func doSomething() (err error) {
	defer func() {
		if err != nil {
			// wrap the error with extra context before returning to caller
			err = fmt.Errorf("doSomething failed: %w", err)
		}
	}()

	// Simulate an error
	err = fmt.Errorf("connection timeout")
	return // named return — defer runs and wraps the error
}

// ─── Defer for function tracing ───────────────────────────────────────────────

func trace(name string) func() {
	fmt.Println("  ENTER:", name)
	return func() {
		fmt.Println("  EXIT:", name)
	}
}

func tracedFunction() {
	defer trace("tracedFunction")() // call trace() now, defer the returned func
	fmt.Println("  doing work inside tracedFunction")
}

// ─── Defer + panic + recover ──────────────────────────────────────────────────
//
// FLOW:
//   panic fires
//       │
//       └─► stack unwinds, running defers
//                │
//                └─► deferred recover() catches the panic
//                         │
//                         └─► function returns normally with an error

func safeOperation(shouldPanic bool) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered panic: %v", r)
		}
	}()

	if shouldPanic {
		panic("something went terribly wrong")
	}
	return nil
}

func main() {

	// ─── File cleanup pattern ─────────────────────────────────────────────────
	openAndRead("nonexistent.txt") // will show error, still defers correctly
	fmt.Println()

	// ─── LIFO order ──────────────────────────────────────────────────────────
	fmt.Println("LIFO demo:")
	demoLIFO()
	fmt.Println()

	// ─── Eager argument evaluation ────────────────────────────────────────────
	fmt.Println("Eager args demo:")
	demoEagerArgs()
	fmt.Println()

	// ─── Named return modification ────────────────────────────────────────────
	err := doSomething()
	fmt.Println("Error with context:", err)
	fmt.Println()

	// ─── Function tracing ─────────────────────────────────────────────────────
	fmt.Println("Trace demo:")
	tracedFunction()
	fmt.Println()

	// ─── Defer + recover ──────────────────────────────────────────────────────
	err = safeOperation(false)
	fmt.Println("No panic:", err) // <nil>

	err = safeOperation(true)
	fmt.Println("With panic:", err) // recovered panic: something went terribly wrong

	// ─── Common defer mistake ─────────────────────────────────────────────────
	//
	// Defer inside a loop — defers accumulate until the FUNCTION returns, not loop iteration.
	// For cleanup per-iteration (e.g., closing a file each loop), call the cleanup directly.
	//
	// BAD (defers pile up until function returns):
	//   for _, f := range files {
	//       file, _ := os.Open(f)
	//       defer file.Close()  ← all files stay open until function exits!
	//   }
	//
	// GOOD (close each file immediately after use):
	//   for _, f := range files {
	//       func() {
	//           file, _ := os.Open(f)
	//           defer file.Close()  ← closes when the anonymous func returns
	//       }()
	//   }
}
