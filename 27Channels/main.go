// 27 - Channels in Go
//
// A channel is a TYPED PIPE through which goroutines communicate and synchronize.
// "Do not communicate by sharing memory; share memory by communicating." — Go proverb
//
// FLOW — basic channel communication:
//
//   Sender goroutine          Channel          Receiver goroutine
//   ────────────────     ───────────────     ────────────────────
//   ch <- value    ───►  [  value  ]   ───►  value := <-ch
//   (blocks until               │            (blocks until
//    receiver is ready)         │             sender sends)
//
// Channel directions:
//   chan T      — bidirectional (send and receive)
//   chan<- T    — send-only    (can only put values in)
//   <-chan T    — receive-only (can only take values out)
//
// Two types:
//   make(chan T)    — unbuffered: sender blocks until receiver is ready (synchronous)
//   make(chan T, N) — buffered:   sender blocks only when buffer is full (async up to N)

package main

import (
	"fmt"
	"time"
)

// ─── Producer: sends-only channel ────────────────────────────────────────────

func producer(ch chan<- int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("  producer: sending %d\n", i)
		ch <- i // send value into channel
		time.Sleep(50 * time.Millisecond)
	}
	close(ch) // signal that no more values will be sent — IMPORTANT
}

// ─── Consumer: receive-only channel ──────────────────────────────────────────

func consumer(ch <-chan int) {
	// range over channel — loops until the channel is closed
	for v := range ch {
		fmt.Printf("  consumer: received %d\n", v)
	}
	fmt.Println("  consumer: channel closed, done")
}

func main() {

	// ─── Unbuffered channel ───────────────────────────────────────────────────
	//
	// FLOW:
	//   goroutine A: ch <- 42  ← BLOCKS until goroutine B reads
	//   goroutine B: <-ch      ← unblocks A, receives 42
	//
	// Both goroutines must be ready at the same time (rendezvous).

	fmt.Println("=== 1. Unbuffered Channel ===")
	ch := make(chan int) // unbuffered

	go func() {
		fmt.Println("  sender: sending 42")
		ch <- 42 // blocks until receiver is ready
		fmt.Println("  sender: sent, continuing")
	}()

	time.Sleep(100 * time.Millisecond)
	val := <-ch // receive — unblocks the sender
	fmt.Println("  receiver: got", val)

	// ─── Buffered channel ─────────────────────────────────────────────────────
	//
	// FLOW (buffer size = 3):
	//
	//   send 1 ──► [1][ ][ ]   ← doesn't block (space available)
	//   send 2 ──► [1][2][ ]   ← doesn't block
	//   send 3 ──► [1][2][3]   ← doesn't block
	//   send 4 ──► BLOCKS       ← buffer full, waits for receiver
	//
	//   recv   ──► [2][3][ ]   ← makes space, send 4 can proceed

	fmt.Println("\n=== 2. Buffered Channel ===")
	buffered := make(chan string, 3)

	// These don't block — buffer has space
	buffered <- "first"
	buffered <- "second"
	buffered <- "third"

	fmt.Println("  buffered len:", len(buffered), "cap:", cap(buffered))

	// Receive all three
	fmt.Println(" ", <-buffered) // first
	fmt.Println(" ", <-buffered) // second
	fmt.Println(" ", <-buffered) // third

	// ─── Producer/Consumer with range ─────────────────────────────────────────
	//
	// FLOW:
	//
	//   producer ──► ch ──► consumer
	//     sends 1..5        range ch loops
	//     close(ch)    ──►  range exits automatically

	fmt.Println("\n=== 3. Producer / Consumer ===")
	pipeline := make(chan int, 5)
	go producer(pipeline, 5)
	consumer(pipeline)

	// ─── Pipeline pattern: chain of channels ─────────────────────────────────
	//
	// FLOW:
	//
	//   generate  ──► ch1 ──► double ──► ch2 ──► print
	//   (1,2,3,4,5)          (x*2)               (2,4,6,8,10)

	fmt.Println("\n=== 4. Pipeline ===")

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

	double := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * 2
			}
			close(out)
		}()
		return out
	}

	// Wire up the pipeline
	nums := generate(1, 2, 3, 4, 5)
	doubled := double(nums)

	for v := range doubled {
		fmt.Printf("  pipeline output: %d\n", v)
	}

	// ─── Done channel: signal goroutine to stop ────────────────────────────────
	//
	// FLOW:
	//
	//   main                    worker goroutine
	//   ────                    ───────────────────────────────
	//   done := make(chan struct{})
	//   go worker(done)  ──────►  for { select { case <-done: return } }
	//   ...
	//   close(done)      ──────►  <-done unblocks, worker exits

	fmt.Println("\n=== 5. Done channel (stop signal) ===")
	done := make(chan struct{}) // struct{} uses zero bytes — it's just a signal

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("  worker: received stop signal, exiting")
				return
			default:
				fmt.Println("  worker: working...")
				time.Sleep(80 * time.Millisecond)
			}
		}
	}()

	time.Sleep(250 * time.Millisecond)
	close(done) // broadcast stop to all goroutines listening on done
	time.Sleep(50 * time.Millisecond)

	// ─── Two-way channel vs directional ───────────────────────────────────────

	fmt.Println("\n=== 6. Channel direction ===")
	// Passing a bidirectional chan as send-only or receive-only to a function
	// prevents accidental misuse inside that function
	twoWay := make(chan int, 1)
	sendTo := func(ch chan<- int, val int) { ch <- val }    // can only send
	recvFrom := func(ch <-chan int) int { return <-ch }     // can only receive

	sendTo(twoWay, 99)
	fmt.Println("  received:", recvFrom(twoWay))
}
