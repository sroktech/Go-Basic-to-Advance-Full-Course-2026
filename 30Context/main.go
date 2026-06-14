// 30 - Context in Go
//
// The 'context' package carries deadlines, cancellation signals, and request-scoped
// values across API boundaries and goroutines.
//
// Context solves: "how do I cancel work that's in progress?" — e.g.,
//   - User cancels an HTTP request
//   - A timeout expires
//   - A parent operation is done and all child work should stop
//
// FLOW — cancellation propagates DOWN the tree:
//
//   context.Background()
//          │
//          └─► WithCancel(parent) ──► childCtx + cancel()
//                    │
//                    ├─► goroutine 1 watches ctx.Done()
//                    ├─► goroutine 2 watches ctx.Done()
//                    └─► goroutine 3 watches ctx.Done()
//
//   call cancel() ──► ctx.Done() closes ──► all goroutines receive signal and stop
//
// Rule: always call cancel() — it releases resources even if work finished normally.
// Best practice: defer cancel() immediately after creating a cancellable context.

package main

import (
	"context"
	"fmt"
	"time"
)

// ─── Worker that respects context cancellation ────────────────────────────────

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			// ctx.Err() explains WHY it was cancelled: context.Canceled or context.DeadlineExceeded
			fmt.Printf("  worker %d stopped: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("  worker %d working...\n", id)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// ─── Simulated HTTP call that can be cancelled ────────────────────────────────

func fetchData(ctx context.Context, url string) (string, error) {
	// Simulate network call with a channel
	resultCh := make(chan string, 1)

	go func() {
		time.Sleep(200 * time.Millisecond) // simulate latency
		resultCh <- "data from " + url
	}()

	select {
	case result := <-resultCh:
		return result, nil
	case <-ctx.Done():
		return "", fmt.Errorf("fetch cancelled: %w", ctx.Err())
	}
}

// ─── Function that passes context through a call chain ────────────────────────

func serviceA(ctx context.Context) (string, error) {
	fmt.Println("  serviceA: calling serviceB")
	return serviceB(ctx)
}

func serviceB(ctx context.Context) (string, error) {
	// Create a child context with its own shorter timeout
	childCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	return fetchData(childCtx, "https://api.example.com/data")
}

func main() {

	// ─── 1. context.Background() — the root context ───────────────────────────
	//
	// Always start here. Background is never cancelled, has no deadline, no values.
	// Used at the top of main() or in tests.

	fmt.Println("=== 1. WithCancel — manual cancellation ===")
	//
	// FLOW:
	//   ctx, cancel := WithCancel(parent)
	//   go worker(ctx)  ← worker watches ctx.Done()
	//   cancel()        ← closes ctx.Done() channel → worker exits

	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(250 * time.Millisecond)
	fmt.Println("  main: cancelling all workers")
	cancel() // signal all workers to stop
	time.Sleep(50 * time.Millisecond)

	// ─── 2. WithTimeout — automatic cancellation after duration ───────────────
	//
	// FLOW:
	//   ctx automatically cancelled after N milliseconds
	//   OR when cancel() is called — whichever comes first

	fmt.Println("\n=== 2. WithTimeout ===")

	// Success case — timeout is generous
	ctx2, cancel2 := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel2()

	result, err := fetchData(ctx2, "api.example.com")
	if err != nil {
		fmt.Println("  error:", err)
	} else {
		fmt.Println("  success:", result)
	}

	// Timeout case — too short
	ctx3, cancel3 := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel3()

	result, err = fetchData(ctx3, "api.slow.com")
	if err != nil {
		fmt.Println("  timed out:", err) // context.DeadlineExceeded
	} else {
		fmt.Println("  success:", result)
	}

	// ─── 3. WithDeadline — cancel at absolute time ────────────────────────────

	fmt.Println("\n=== 3. WithDeadline ===")

	deadline := time.Now().Add(100 * time.Millisecond)
	ctx4, cancel4 := context.WithDeadline(context.Background(), deadline)
	defer cancel4()

	fmt.Println("  deadline in:", time.Until(deadline))
	<-ctx4.Done() // wait for deadline to expire
	fmt.Println("  deadline reached:", ctx4.Err())

	// ─── 4. Passing context through a call chain ─────────────────────────────
	//
	// FLOW:
	//
	//   main ──► serviceA(ctx) ──► serviceB(ctx) ──► fetchData(childCtx)
	//
	//   Context flows DOWN through every function call.
	//   Each level can add tighter timeouts or values.
	//   Cancellation from the top propagates to all children.

	fmt.Println("\n=== 4. Context through call chain ===")

	topCtx, topCancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer topCancel()

	data, err := serviceA(topCtx)
	if err != nil {
		fmt.Println("  serviceA error:", err)
	} else {
		fmt.Println("  serviceA result:", data)
	}

	// ─── 5. WithValue — carry request-scoped values ───────────────────────────
	//
	// context.WithValue attaches a key-value pair to the context.
	// Retrieve it anywhere in the call chain with ctx.Value(key).
	//
	// Use for: request IDs, auth tokens, trace IDs — NOT for function parameters!
	// Never store mutable values or large data in context.

	fmt.Println("\n=== 5. WithValue ===")

	type contextKey string // define a custom key type to avoid collisions
	const requestIDKey contextKey = "requestID"

	ctx5 := context.WithValue(context.Background(), requestIDKey, "req-abc-123")

	// Retrieve anywhere down the call chain
	handleRequest := func(ctx context.Context) {
		requestID := ctx.Value(requestIDKey)
		fmt.Printf("  handling request: %v\n", requestID)
	}
	handleRequest(ctx5) // req-abc-123

	// ─── Summary: when to use what ────────────────────────────────────────────
	fmt.Println(`
Context quick reference:
  context.Background()         — root, use in main/tests
  context.TODO()               — placeholder when unsure (same as Background)
  WithCancel(parent)           — manual cancel()
  WithTimeout(parent, dur)     — auto-cancel after duration
  WithDeadline(parent, time)   — auto-cancel at absolute time
  WithValue(parent, key, val)  — attach request-scoped data`)
}
