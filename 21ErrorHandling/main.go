// 21 - Error Handling in Go
//
// Go handles errors differently from most languages.
// There are NO exceptions (try/catch). Instead, errors are RETURN VALUES.
// Functions that can fail return an 'error' as their last return value.
// The caller is responsible for checking it — Go makes this explicit by design.
//
// The built-in 'error' interface:
//   type error interface {
//       Error() string
//   }
// Any type with an Error() string method IS an error.
//
// Three mechanisms:
//   1. error return value  — the standard way to signal failure
//   2. panic / recover     — for truly unexpected, unrecoverable situations
//   3. defer               — runs cleanup code when a function exits (any reason)

package main

import (
	"errors"  // standard library for creating simple errors
	"fmt"
)

// ─── Returning errors from functions ──────────────────────────────────────────

// divide returns (result, error). If divisor is 0, it returns an error.
// Convention: return the error as the LAST return value.
// Convention: return zero value for other return values when there's an error.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// errors.New creates a simple error with a message string
		return 0, errors.New("division by zero")
	}
	return a / b, nil // nil means "no error" — everything went fine
}

// ─── Custom error types ───────────────────────────────────────────────────────

// A custom error type lets you attach extra information to an error.
// It must implement the error interface: Error() string
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// validateAge returns a custom error if age is out of range
func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "cannot be negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically large"}
	}
	return nil // valid
}

// ─── fmt.Errorf — formatted error messages ────────────────────────────────────

// %w verb wraps an error — the wrapped error can be extracted later with errors.Is / errors.As
func openFile(filename string) error {
	// Simulating a file-not-found situation
	baseErr := errors.New("no such file or directory")
	return fmt.Errorf("openFile %q: %w", filename, baseErr) // wraps base error
}

// ─── defer ────────────────────────────────────────────────────────────────────

// defer schedules a function call to run WHEN the surrounding function returns.
// Useful for cleanup: closing files, releasing locks, logging.
// Multiple defers run in LIFO order (last in, first out).
func demoDefer() {
	fmt.Println("demoDefer: start")
	defer fmt.Println("demoDefer: first defer (runs last)")
	defer fmt.Println("demoDefer: second defer (runs second)")
	defer fmt.Println("demoDefer: third defer (runs first)")
	fmt.Println("demoDefer: end")
}

// ─── panic and recover ────────────────────────────────────────────────────────

// panic stops normal execution and begins unwinding the call stack.
// Use it ONLY for truly unrecoverable situations (programming bugs, not user errors).
//
// recover() can catch a panic inside a deferred function.
// After recover, execution continues normally in the caller.
func safeDivide(a, b int) (result int, err error) {
	// defer with recover — catches any panic in this function
	defer func() {
		if r := recover(); r != nil {
			// r is whatever was passed to panic()
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	if b == 0 {
		panic("attempted to divide by zero") // triggers panic
	}
	return a / b, nil
}

func main() {

	// ─── Checking errors — the standard pattern ───────────────────────────────

	// Always check: if err != nil
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.1f\n", result) // 5.0
	}

	// Error case
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Error: division by zero
	} else {
		fmt.Printf("Result: %.1f\n", result)
	}

	// ─── Custom error ─────────────────────────────────────────────────────────

	err = validateAge(-5)
	if err != nil {
		fmt.Println("Validation failed:", err)

		// errors.As — check if error is a specific type and extract it
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("  Field: %s, Message: %s\n", ve.Field, ve.Message)
		}
	}

	err = validateAge(25)
	if err == nil {
		fmt.Println("Age 25 is valid")
	}

	// ─── Wrapped errors ───────────────────────────────────────────────────────

	err = openFile("config.yaml")
	if err != nil {
		fmt.Println("Error:", err) // openFile "config.yaml": no such file or directory

		// errors.Is — checks if the error chain contains a specific target error
		target := errors.New("no such file or directory")
		fmt.Println("Is file-not-found:", errors.Is(err, target)) // false (different instance)
		// For Is() to work, use a package-level sentinel error and wrap it with %w
	}

	// ─── defer demo ───────────────────────────────────────────────────────────

	fmt.Println()
	demoDefer()
	// Output order:
	//   demoDefer: start
	//   demoDefer: end
	//   demoDefer: third defer  (runs first — LIFO)
	//   demoDefer: second defer
	//   demoDefer: first defer  (runs last)

	// ─── panic and recover demo ───────────────────────────────────────────────

	fmt.Println()
	res, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("safeDivide error:", err)
	} else {
		fmt.Println("safeDivide 10/2 =", res) // 5
	}

	res, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("safeDivide error:", err) // recovered from panic: attempted to divide by zero
	} else {
		fmt.Println("safeDivide 10/0 =", res)
	}

	// ─── Best practices summary ───────────────────────────────────────────────
	//
	// DO:
	//   - Always handle errors — never ignore with _
	//   - Return errors as the last value
	//   - Add context with fmt.Errorf("what I was doing: %w", err)
	//   - Use custom error types when callers need to inspect error details
	//   - Use defer for cleanup (closing files, unlocking mutexes)
	//
	// DON'T:
	//   - Use panic for normal error conditions (use error returns instead)
	//   - Use panic/recover as a substitute for exceptions
	fmt.Println()
	fmt.Println("Done — all error handling examples complete.")
}
