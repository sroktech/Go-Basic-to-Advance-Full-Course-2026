// 26 - Goroutines in Go
//
// A goroutine is a LIGHTWEIGHT THREAD managed by the Go runtime — not the OS.
// You can run thousands (even millions) of goroutines simultaneously.
// Each goroutine starts with a small stack (~2KB) that grows/shrinks automatically.
//
// Start a goroutine with the 'go' keyword before a function call.
//
// FLOW — sequential vs concurrent:
//
//   Sequential (normal):               Concurrent (goroutines):
//   ─────────────────────              ──────────────────────────────────
//   main ──► task1 (wait)              main ──► go task1() ──► [runs in background]
//            │                                 │
//            ▼ (done)                          ──► go task2() ──► [runs in background]
//   main ──► task2 (wait)                      │
//            │                                 ──► go task3() ──► [runs in background]
//            ▼ (done)                          │
//   main ──► task3 (wait)                      └──► all three run AT THE SAME TIME
//
// Key point: the main goroutine does NOT wait for other goroutines.
// If main() exits, ALL goroutines are killed immediately.
// Use sync.WaitGroup (lesson 29) or channels (lesson 27) to wait for them.

package main

import (
	"fmt"
	"runtime"
	"time"
)

// ─── Simple function to run as goroutine ──────────────────────────────────────

func printMessage(id int, msg string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("  goroutine %d: %s [iteration %d]\n", id, msg, i+1)
		time.Sleep(50 * time.Millisecond) // simulate some work
	}
}

// ─── Goroutine with a result via channel ─────────────────────────────────────
// (covered in depth in lesson 27 — shown briefly here)

func fetchData(id int, result chan<- string) {
	time.Sleep(time.Duration(id*100) * time.Millisecond) // simulate variable latency
	result <- fmt.Sprintf("data from worker %d", id)
}

func main() {

	fmt.Println("=== 1. Sequential vs Goroutine ===")

	// Sequential — each call blocks until done
	fmt.Println("Sequential:")
	start := time.Now()
	printMessage(1, "sequential A")
	printMessage(2, "sequential B")
	fmt.Printf("Sequential took: %v\n\n", time.Since(start))

	// Concurrent — both run at the same time
	fmt.Println("Concurrent with goroutines:")
	start = time.Now()
	go printMessage(1, "goroutine A")
	go printMessage(2, "goroutine B")
	// We need to wait — otherwise main() exits and kills both goroutines
	time.Sleep(300 * time.Millisecond) // crude wait (use WaitGroup in real code)
	fmt.Printf("Concurrent took: %v\n\n", time.Since(start))

	// ─── Anonymous goroutine ──────────────────────────────────────────────────

	fmt.Println("=== 2. Anonymous goroutine ===")
	go func() {
		fmt.Println("  anonymous goroutine running")
	}()
	time.Sleep(10 * time.Millisecond)

	// ─── Goroutine with closure capturing variable ────────────────────────────

	fmt.Println("=== 3. Goroutines with captured values ===")
	// IMPORTANT: pass loop variables as arguments, don't capture directly
	for i := 1; i <= 3; i++ {
		i := i // shadow i — each goroutine gets its own copy
		go func() {
			fmt.Printf("  worker %d started\n", i)
		}()
	}
	time.Sleep(50 * time.Millisecond)

	// ─── Goroutine runtime info ───────────────────────────────────────────────

	fmt.Println("\n=== 4. Runtime info ===")
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines running:", runtime.NumGoroutine())
	// GOMAXPROCS controls how many OS threads run goroutines simultaneously
	// Default is runtime.NumCPU() — uses all available cores
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // 0 = query without changing

	// ─── Fan-out: multiple goroutines, collect results via channel ────────────

	fmt.Println("\n=== 5. Fan-out pattern ===")
	// Start 5 workers concurrently, collect all results
	results := make(chan string, 5) // buffered channel — can hold 5 items

	for i := 1; i <= 5; i++ {
		go fetchData(i, results)
	}

	// Collect all 5 results (blocks until each is ready)
	for i := 0; i < 5; i++ {
		fmt.Println(" ", <-results)
	}

	// ─── Goroutine leak warning ───────────────────────────────────────────────
	//
	// A goroutine leak happens when a goroutine is started but never finishes —
	// it blocks forever, consuming memory.
	//
	// Common causes:
	//   - Waiting on a channel that nobody ever sends to
	//   - Waiting for a mutex that nobody ever unlocks
	//   - Infinite loop with no exit condition
	//
	// Always ensure goroutines have a way to exit:
	//   - via channel signal
	//   - via context cancellation (lesson 30)
	//   - via a timeout

	fmt.Println("\n=== 6. Goroutine with timeout ===")
	done := make(chan bool)
	go func() {
		time.Sleep(50 * time.Millisecond) // simulate work
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("  goroutine finished in time")
	case <-time.After(500 * time.Millisecond):
		fmt.Println("  goroutine timed out")
	}

	fmt.Println("\nMain goroutine done.")
}
