// 36 - Testing in Go
//
// Go has a built-in testing framework in the 'testing' package — no external tools needed.
// Test files end with _test.go and are compiled only when running tests.
//
// FLOW — how tests work:
//
//   go test ./...
//       │
//       ├─ finds all *_test.go files
//       ├─ compiles them with the package
//       ├─ runs functions named TestXxx(t *testing.T)
//       └─ reports PASS / FAIL
//
// This file contains the code being tested.
// math_test.go contains the actual tests.

package main

import "errors"

// ─── Functions to test ────────────────────────────────────────────────────────

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func Factorial(n int) int {
	if n < 0 {
		return -1 // signal: invalid input
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func IsPalindrome(s string) bool {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	// Run: go test ./...          — run all tests
	// Run: go test -v ./...       — verbose output (show each test name)
	// Run: go test -run TestAdd   — run only tests matching "TestAdd"
	// Run: go test -cover ./...   — show test coverage %
	// Run: go test -bench=. ./... — run benchmarks
	// Run: go test -race ./...    — run with race detector
}
