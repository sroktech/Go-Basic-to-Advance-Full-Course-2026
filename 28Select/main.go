// 28 - Select Statement in Go
//
// 'select' lets a goroutine wait on MULTIPLE channel operations simultaneously.
// It picks whichever channel is ready first — if multiple are ready, it picks one at RANDOM.
// If none are ready, it blocks until one becomes ready (or hits a default case).
//
// Think of it like a switch statement, but for channels.
//
// FLOW:
//
//   select {
//      │
//      ├─── case <-ch1:  ──► ch1 has data? run this
//      │
//      ├─── case <-ch2:  ──► ch2 has data? run this
//      │
//      ├─── case ch3<-v: ──► ch3 ready to receive? run this
//      │
//      └─── default:     ──► nothing ready? run immediately (non-blocking)
//
// Without default: select BLOCKS until at least one case is ready.
// With default:    select is NON-BLOCKING — falls through if nothing is ready.

package main

import (
	"fmt"
	"time"
)

func main() {

	// ─── Basic select: pick whichever is ready first ───────────────────────────

	fmt.Println("=== 1. Basic Select ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from ch1"
	}()
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "from ch2" // this will be ready first
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("received:", msg2) // prints this — ch2 is ready first
	}

	// ─── Select in a loop: handle multiple channels continuously ──────────────
	//
	// FLOW:
	//
	//   loop
	//     │
	//     └─► select
	//              ├─ ch1 ready? process
	//              ├─ ch2 ready? process
	//              └─ done ready? exit loop

	fmt.Println("\n=== 2. Select in a loop ===")

	ticker1 := time.NewTicker(100 * time.Millisecond)
	ticker2 := time.NewTicker(150 * time.Millisecond)
	stop := time.After(500 * time.Millisecond)

	defer ticker1.Stop()
	defer ticker2.Stop()

	count1, count2 := 0, 0
	for {
		select {
		case <-ticker1.C:
			count1++
			fmt.Printf("  tick1 (total: %d)\n", count1)
		case <-ticker2.C:
			count2++
			fmt.Printf("  tick2 (total: %d)\n", count2)
		case <-stop:
			fmt.Printf("  done — tick1: %d, tick2: %d\n", count1, count2)
			goto afterLoop // exit the for loop
		}
	}
afterLoop:

	// ─── Non-blocking with default ────────────────────────────────────────────
	//
	// FLOW:
	//   select checks all cases
	//      │
	//      ├─ any channel ready? → run that case
	//      └─ none ready?        → run default immediately (no blocking)

	fmt.Println("\n=== 3. Non-blocking (default) ===")

	ch := make(chan int, 1)

	// Try to receive — channel is empty, hits default
	select {
	case v := <-ch:
		fmt.Println("  received:", v)
	default:
		fmt.Println("  nothing to receive — moving on")
	}

	ch <- 42 // put something in

	// Try again — now there's a value
	select {
	case v := <-ch:
		fmt.Println("  received:", v) // 42
	default:
		fmt.Println("  nothing to receive")
	}

	// ─── Timeout pattern ──────────────────────────────────────────────────────
	//
	// FLOW:
	//
	//   select
	//     ├─ operation completes? ──► use result
	//     └─ time.After fires?    ──► report timeout, give up
	//
	// time.After(d) returns a channel that receives after duration d.
	// This is the standard Go pattern for adding timeouts to operations.

	fmt.Println("\n=== 4. Timeout pattern ===")

	slowOperation := func(result chan<- string) {
		time.Sleep(300 * time.Millisecond)
		result <- "operation complete"
	}

	result := make(chan string, 1)
	go slowOperation(result)

	select {
	case res := <-result:
		fmt.Println("  success:", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  timed out waiting for operation") // fires because 300ms > 100ms
	}

	// With longer timeout — operation completes
	result2 := make(chan string, 1)
	go slowOperation(result2)

	select {
	case res := <-result2:
		fmt.Println("  success:", res) // fires because 500ms > 300ms
	case <-time.After(500 * time.Millisecond):
		fmt.Println("  timed out")
	}

	// ─── Nil channel blocks forever ───────────────────────────────────────────
	//
	// A nil channel is NEVER ready — useful to disable a case in select dynamically.
	// Setting a channel to nil in a select is a way to "turn off" a case.

	fmt.Println("\n=== 5. Nil channel disables a case ===")

	a := make(chan int, 1)
	b := make(chan int, 1)

	a <- 1
	b <- 2

	// Disable 'a' case after first receive
	for i := 0; i < 2; i++ {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil // disable this case going forward
			} else {
				fmt.Println("  from a:", v)
				a = nil // disable after reading
			}
		case v := <-b:
			fmt.Println("  from b:", v)
		}
	}
}
