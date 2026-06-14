// math_test.go — tests for math.go
//
// Rules for test files:
//   - File name MUST end in _test.go
//   - Same package as the code being tested (or packagename_test for black-box tests)
//   - Test functions MUST be named TestXxx (capital T, then anything)
//   - Test function signature: func TestXxx(t *testing.T)
//
// t *testing.T methods:
//   t.Error("msg")        — mark test as failed, continue running
//   t.Errorf("fmt", ...)  — like Error but with formatting
//   t.Fatal("msg")        — mark as failed and STOP this test immediately
//   t.Fatalf("fmt", ...)  — Fatal with formatting
//   t.Log("msg")          — log info (only shown on failure or with -v)
//   t.Skip("reason")      — skip this test
//
// FLOW — what happens when you run: go test -v ./...
//
//   TestAdd        ── calls Add(a,b), checks result ──► PASS / FAIL
//   TestSubtract   ── ...
//   TestDivide     ── ...
//   TestFactorial  ── ...

package main

import (
	"fmt"
	"testing"
)

// ─── Simple test ──────────────────────────────────────────────────────────────

func TestAdd(t *testing.T) {
	result := Add(3, 4)
	expected := 7
	if result != expected {
		t.Errorf("Add(3, 4) = %d; want %d", result, expected)
	}
}

// ─── Table-driven tests — the standard Go pattern ─────────────────────────────
//
// Instead of writing one test per case, define a slice of test cases.
// This makes it easy to add new cases without new functions.
//
// FLOW:
//   tests := []struct{ ... }{ {case1}, {case2}, ... }
//   for _, tt := range tests {
//       t.Run(tt.name, func(t *testing.T) { ... })
//   }

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive result", 10, 3, 7},
		{"zero result", 5, 5, 0},
		{"negative result", 3, 10, -7},
		{"both zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive numbers", 3, 4, 12},
		{"multiply by zero", 5, 0, 0},
		{"negative numbers", -3, -4, 12},
		{"mixed signs", -3, 4, -12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// ─── Testing error cases ──────────────────────────────────────────────────────

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		want      float64
		wantError bool // true if we expect an error
	}{
		{"normal division", 10, 2, 5, false},
		{"decimal result", 7, 2, 3.5, false},
		{"divide by zero", 10, 0, 0, true}, // should return an error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			if tt.wantError {
				// We EXPECT an error — fail if there isn't one
				if err == nil {
					t.Errorf("Divide(%v, %v) expected error but got none", tt.a, tt.b)
				}
				return
			}

			// We do NOT expect an error
			if err != nil {
				t.Fatalf("Divide(%v, %v) unexpected error: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{10, 3628800},
		{-1, -1}, // invalid input, expect -1
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Factorial(%d)", tt.input), func(t *testing.T) {
			got := Factorial(tt.input)
			if got != tt.want {
				t.Errorf("Factorial(%d) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"racecar", true},
		{"hello", false},
		{"", true},   // empty string is a palindrome
		{"a", true},  // single char is a palindrome
		{"aba", true},
		{"abc", false},
		{"level", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

// ─── Benchmark tests ──────────────────────────────────────────────────────────
//
// Benchmarks measure performance. Run with: go test -bench=.
// b.N is the number of iterations — Go adjusts it automatically for stable results.
//
// FLOW:
//   for i := 0; i < b.N; i++ {
//       // code to benchmark
//   }

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(3, 4)
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10)
	}
}

// ─── Setup and teardown with TestMain ────────────────────────────────────────
//
// TestMain runs before any tests in the package.
// Use it to set up shared resources (DB connection, temp files, etc.)
//
// func TestMain(m *testing.M) {
//     // setup
//     setup()
//
//     code := m.Run() // run all tests
//
//     // teardown
//     teardown()
//     os.Exit(code)
// }
