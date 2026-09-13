// Package mathutils provides basic math helper functions.
//
// Package naming rules:
//   - Lowercase, short, no underscores: 'mathutils' not 'MathUtils' or 'math_utils'
//   - The package name is how callers reference it: mathutils.Add(...)
//   - The folder name should match the package name
//
// Exported vs Unexported:
//   - Uppercase first letter = EXPORTED (public) — visible outside the package
//   - Lowercase first letter = unexported (private) — only usable inside this package
package mathutils

import "fmt"

// Add returns the sum of two integers. (Exported — starts with uppercase A)
func Add(a, b int) int {
	return a + b
}

// Subtract returns a minus b. (Exported)
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns a times b. (Exported)
func Multiply(a, b int) int {
	return a * b
}

// Divide returns a divided by b and an error if b is zero. (Exported)
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide: cannot divide by zero")
	}
	return a / b, nil
}

// max is unexported — only usable inside this package (lowercase m)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Max wraps the unexported max — exported version for callers
func Max(a, b int) int {
	return max(a, b)
}

// Pi is an exported constant
const Pi = 3.14159265358979
