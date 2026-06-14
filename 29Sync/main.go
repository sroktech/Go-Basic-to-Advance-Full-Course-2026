// 29 - Sync Package in Go (WaitGroup, Mutex, Once, atomic)
//
// When multiple goroutines share data, you need synchronization to prevent
// RACE CONDITIONS — bugs where the result depends on which goroutine runs first.
//
// The 'sync' package provides:
//   WaitGroup — wait for a group of goroutines to finish
//   Mutex     — lock/unlock to protect shared data (only one goroutine at a time)
//   RWMutex   — multiple readers OR one writer at a time
//   Once      — run a function exactly once (e.g., initialization)
//
// 'sync/atomic' provides lock-free atomic operations on integers (faster than mutex).
//
// FLOW — race condition (WITHOUT mutex):
//
//   goroutine 1: reads count=5          goroutine 2: reads count=5
//   goroutine 1: count = 5+1 = 6       goroutine 2: count = 5+1 = 6
//   goroutine 1: writes count=6         goroutine 2: writes count=6
//   RESULT: count=6 (expected 7!)  ← one increment was LOST
//
// FLOW — with mutex:
//
//   goroutine 1: Lock() ──► reads count=5, writes count=6, Unlock()
//   goroutine 2:                                             Lock() ──► reads count=6, writes count=7, Unlock()
//   RESULT: count=7 ✓

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ─── WaitGroup ────────────────────────────────────────────────────────────────

// FLOW:
//   wg.Add(N)        ← tell WaitGroup: N goroutines are starting
//   go func() {
//       defer wg.Done()  ← goroutine signals it's done (decrements counter)
//   }()
//   wg.Wait()        ← blocks until counter reaches 0

func demoWaitGroup() {
	fmt.Println("=== WaitGroup ===")

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // increment counter BEFORE starting goroutine
		i := i
		go func() {
			defer wg.Done() // decrement counter when this goroutine finishes
			time.Sleep(time.Duration(i*50) * time.Millisecond)
			fmt.Printf("  worker %d done\n", i)
		}()
	}

	wg.Wait() // blocks here until all 5 workers call Done()
	fmt.Println("  all workers finished")
}

// ─── Mutex — protect shared data ─────────────────────────────────────────────

type SafeCounter struct {
	mu    sync.Mutex // the lock — zero value is unlocked
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()         // acquire lock — only one goroutine proceeds
	defer c.mu.Unlock() // release lock when function returns
	c.count++           // safe to modify — no other goroutine can be here
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func demoMutex() {
	fmt.Println("\n=== Mutex ===")

	counter := SafeCounter{}
	var wg sync.WaitGroup

	// 1000 goroutines all increment the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("  final count (should be 1000):", counter.Value())
}

// ─── RWMutex — multiple readers, single writer ────────────────────────────────
//
// FLOW:
//   Read:  RLock() / RUnlock()  — multiple goroutines can read simultaneously
//   Write: Lock()  / Unlock()   — exclusive, blocks all readers and writers
//
// Use RWMutex when reads are far more frequent than writes (e.g., a cache).

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Set(key, val string) {
	c.mu.Lock() // exclusive write lock
	defer c.mu.Unlock()
	c.data[key] = val
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock() // shared read lock — multiple goroutines can read at once
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

func demoRWMutex() {
	fmt.Println("\n=== RWMutex ===")

	cache := Cache{data: make(map[string]string)}
	cache.Set("language", "Go")
	cache.Set("version", "1.24")

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			if val, ok := cache.Get("language"); ok {
				fmt.Printf("  reader %d: language = %s\n", i, val)
			}
		}()
	}
	wg.Wait()
}

// ─── Once — initialize exactly once ──────────────────────────────────────────
//
// FLOW:
//   first call:  Once.Do(fn) ──► fn() runs
//   second call: Once.Do(fn) ──► fn() is SKIPPED
//   third call:  Once.Do(fn) ──► fn() is SKIPPED

var (
	instance *expensiveService
	once     sync.Once
)

type expensiveService struct {
	name string
}

func getInstance() *expensiveService {
	once.Do(func() {
		fmt.Println("  initializing service (should only print once!)")
		instance = &expensiveService{name: "MyService"}
	})
	return instance
}

func demoOnce() {
	fmt.Println("\n=== sync.Once (singleton pattern) ===")

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := getInstance()
			fmt.Printf("  got service: %s\n", s.name)
		}()
	}
	wg.Wait()
}

// ─── atomic — lock-free counter ───────────────────────────────────────────────
//
// For simple integer operations, atomic is faster than a mutex.
// Operations are guaranteed to be atomic (indivisible) at the CPU level.

func demoAtomic() {
	fmt.Println("\n=== sync/atomic ===")

	var counter int64 // must use int32, int64, etc.

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1) // atomic increment — no race condition
		}()
	}
	wg.Wait()

	// Load reads the current value atomically
	fmt.Println("  atomic counter (should be 1000):", atomic.LoadInt64(&counter))
}

func main() {
	demoWaitGroup()
	demoMutex()
	demoRWMutex()
	demoOnce()
	demoAtomic()

	// ─── Race detector ────────────────────────────────────────────────────────
	// Run with: go run -race main.go
	// The race detector finds race conditions at runtime — use it during development.
	fmt.Println("\nTip: run with 'go run -race main.go' to detect race conditions")
}
