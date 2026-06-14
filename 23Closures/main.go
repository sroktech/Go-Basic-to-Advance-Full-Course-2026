// 23 - Closures & Anonymous Functions in Go
//
// In Go, functions are FIRST-CLASS VALUES — you can:
//   - Assign a function to a variable
//   - Pass a function as an argument
//   - Return a function from another function
//
// A CLOSURE is a function that "closes over" (captures) variables from
// its surrounding scope. Even after the outer function returns, the closure
// still has access to those captured variables.
//
// FLOW — how a closure captures a variable:
//
//   outer function
//       │
//       ├── creates variable: count = 0
//       │
//       └── returns inner function (closure)
//                │
//                └── closure HOLDS A REFERENCE to count
//                         │
//                         ├── call 1: count++ → count = 1
//                         ├── call 2: count++ → count = 2
//                         └── call 3: count++ → count = 3
//
// The variable 'count' lives on as long as the closure exists.

package main

import "fmt"

// ─── Anonymous functions ──────────────────────────────────────────────────────

// An anonymous function has no name — defined inline where it's needed.
// Used when the function is short and only needed in one place.

// ─── Closure factory — returns a function ─────────────────────────────────────

// makeCounter returns a new closure each time it's called.
// Each closure has its OWN independent 'count' variable.
func makeCounter() func() int {
	count := 0 // captured by the closure below
	return func() int {
		count++ // modifies the captured variable
		return count
	}
}

// makeAdder returns a function that adds 'x' to any number
func makeAdder(x int) func(int) int {
	return func(n int) int {
		return x + n // x is captured from makeAdder's scope
	}
}

// makeMultiplier returns a function that multiplies by 'factor'
func makeMultiplier(factor float64) func(float64) float64 {
	return func(n float64) float64 {
		return n * factor
	}
}

func main() {

	// ─── Anonymous function — called immediately ───────────────────────────────

	// Define and call at once using ()
	result := func(a, b int) int {
		return a + b
	}(3, 4) // immediately invoked with args 3 and 4
	fmt.Println("Immediate call result:", result) // 7

	// ─── Assign function to variable ──────────────────────────────────────────

	greet := func(name string) string {
		return "Hello, " + name + "!"
	}
	fmt.Println(greet("Alice")) // Hello, Alice!
	fmt.Println(greet("Bob"))   // Hello, Bob!
	fmt.Printf("Type of greet: %T\n", greet) // func(string) string

	// ─── Closure capturing a variable ─────────────────────────────────────────

	counterA := makeCounter()
	counterB := makeCounter() // completely independent counter

	fmt.Println("Counter A:", counterA()) // 1
	fmt.Println("Counter A:", counterA()) // 2
	fmt.Println("Counter A:", counterA()) // 3
	fmt.Println("Counter B:", counterB()) // 1 — B has its own count, starts fresh
	fmt.Println("Counter A:", counterA()) // 4 — A continues from where it left off

	// ─── Adder closures ───────────────────────────────────────────────────────

	add5 := makeAdder(5)
	add10 := makeAdder(10)

	fmt.Println("add5(3):", add5(3))   // 8
	fmt.Println("add10(3):", add10(3)) // 13
	fmt.Println("add5(7):", add5(7))   // 12 — add5 still remembers x=5

	// ─── Functions as arguments (higher-order functions) ──────────────────────

	// A function that accepts another function as a parameter
	// This is the basis for map/filter/reduce patterns in Go

	// apply runs operation on each number and returns new slice
	apply := func(nums []int, operation func(int) int) []int {
		result := make([]int, len(nums))
		for i, n := range nums {
			result[i] = operation(n)
		}
		return result
	}

	nums := []int{1, 2, 3, 4, 5}

	doubled := apply(nums, func(n int) int { return n * 2 })
	squared := apply(nums, func(n int) int { return n * n })

	fmt.Println("Original:", nums)   // [1 2 3 4 5]
	fmt.Println("Doubled:", doubled) // [2 4 6 8 10]
	fmt.Println("Squared:", squared) // [1 4 9 16 25]

	// filter returns elements that satisfy a condition
	filter := func(nums []int, predicate func(int) bool) []int {
		var result []int
		for _, n := range nums {
			if predicate(n) {
				result = append(result, n)
			}
		}
		return result
	}

	evens := filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("Evens:", evens) // [2 4]

	// ─── Closure capturing loop variable — common gotcha ──────────────────────

	// WRONG: all closures capture the SAME 'i' variable (its final value)
	funcsWrong := make([]func(), 3)
	for i := 0; i < 3; i++ {
		funcsWrong[i] = func() {
			fmt.Print(i, " ") // all print 3 — captures reference to i, not the value
		}
	}
	fmt.Print("Wrong loop capture: ")
	for _, f := range funcsWrong {
		f()
	}
	fmt.Println()

	// CORRECT: pass i as a parameter to create a new copy each iteration
	funcsRight := make([]func(), 3)
	for i := 0; i < 3; i++ {
		i := i // new variable per iteration (shadows outer i)
		funcsRight[i] = func() {
			fmt.Print(i, " ") // each closure captures its own copy
		}
	}
	fmt.Print("Right loop capture: ")
	for _, f := range funcsRight {
		f()
	}
	fmt.Println()

	// NOTE: In Go 1.22+, for loop variables are re-created each iteration,
	// so the "wrong" example above actually prints 0 1 2 correctly. But
	// the explicit 'i := i' pattern still works and is clearer about intent.

	// ─── Practical example: middleware-style wrapping ─────────────────────────

	// Closures are used heavily for wrapping behavior around functions
	withLogging := func(fn func(int) int, name string) func(int) int {
		return func(n int) int {
			fmt.Printf("calling %s(%d)\n", name, n)
			result := fn(n)
			fmt.Printf("%s(%d) = %d\n", name, n, result)
			return result
		}
	}

	double := func(n int) int { return n * 2 }
	loggedDouble := withLogging(double, "double")
	loggedDouble(21) // logs input and output automatically
}
