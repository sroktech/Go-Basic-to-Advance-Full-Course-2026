// 18 - Recursion in Go
//
// Recursion is when a function CALLS ITSELF to solve a smaller version of the same problem.
// Every recursive function needs two things:
//   1. Base case  — a condition that STOPS the recursion (no more calls)
//   2. Recursive case — the function calling itself with a simpler/smaller input
//
// Without a base case, the function calls itself forever → stack overflow (crash).
//
// Recursion is elegant for problems that are naturally self-similar:
//   - Factorial, Fibonacci, tree traversal, directory listing, binary search

package main

import "fmt"

// ─── Factorial ────────────────────────────────────────────────────────────────
//
// Factorial of n (written n!) = n × (n-1) × (n-2) × ... × 1
// Example: 5! = 5 × 4 × 3 × 2 × 1 = 120
//
// Call stack for factorial(4):
//   factorial(4) = 4 * factorial(3)
//                      = 3 * factorial(2)
//                              = 2 * factorial(1)
//                                      = 1   ← base case, unwinds from here
//                              = 2 * 1 = 2
//                      = 3 * 2 = 6
//                  = 4 * 6 = 24
func factorial(n int) int {
	if n == 0 || n == 1 { // base case — stop here
		return 1
	}
	return n * factorial(n-1) // recursive case — call with a smaller n
}

// ─── Fibonacci ────────────────────────────────────────────────────────────────
//
// Fibonacci sequence: 0, 1, 1, 2, 3, 5, 8, 13, 21, ...
// Each number is the SUM of the two before it.
// fib(0)=0, fib(1)=1, fib(n)=fib(n-1)+fib(n-2)
//
// Note: this naive version recomputes values many times (exponential time).
// For large n use memoization or an iterative approach instead.
func fibonacci(n int) int {
	if n == 0 { // base case 1
		return 0
	}
	if n == 1 { // base case 2
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2) // recursive case
}

// ─── Sum of a slice ───────────────────────────────────────────────────────────
//
// Add up all elements in a slice using recursion.
// Base case: empty slice → 0
// Recursive: first element + sum of the rest
func sumSlice(nums []int) int {
	if len(nums) == 0 { // base case — nothing left to add
		return 0
	}
	return nums[0] + sumSlice(nums[1:]) // first + sum of rest
}

// ─── Power function ───────────────────────────────────────────────────────────
//
// Compute base^exponent (base to the power of exponent)
// Example: power(2, 10) = 1024
func power(base, exponent int) int {
	if exponent == 0 { // anything^0 = 1
		return 1
	}
	return base * power(base, exponent-1)
}

// ─── Countdown (simple demo) ──────────────────────────────────────────────────
func countdown(n int) {
	if n == 0 { // base case
		fmt.Println("Go!")
		return
	}
	fmt.Println(n)
	countdown(n - 1) // recursive call with n reduced by 1
}

func main() {

	// Factorial examples
	fmt.Println("─── Factorial ───")
	for _, n := range []int{0, 1, 5, 10} {
		fmt.Printf("  %d! = %d\n", n, factorial(n))
	}

	// Fibonacci sequence
	fmt.Println("─── Fibonacci (first 10) ───")
	for i := 0; i < 10; i++ {
		fmt.Printf("  fib(%d) = %d\n", i, fibonacci(i))
	}

	// Sum of slice
	fmt.Println("─── Sum of slice ───")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("  sum of %v = %d\n", nums, sumSlice(nums)) // 15

	// Power
	fmt.Println("─── Power ───")
	fmt.Printf("  2^10 = %d\n", power(2, 10))   // 1024
	fmt.Printf("  3^4  = %d\n", power(3, 4))    // 81

	// Countdown
	fmt.Println("─── Countdown ───")
	countdown(5) // 5 4 3 2 1 Go!
}
