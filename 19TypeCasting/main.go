// 19 - Type Casting (Type Conversion) in Go
//
// Go is STRICTLY typed — it will NEVER automatically convert one type to another.
// You must convert EXPLICITLY when you want to use a value as a different type.
//
// Two things to know:
//   1. Type Conversion — converting between compatible types (int → float64, etc.)
//      Syntax: targetType(value)
//
//   2. Type Assertion — extracting the concrete type from an interface{}
//      Syntax: value.(ConcreteType)
//
// This is different from many languages (JS, Python) where types are coerced silently.
// Go's strictness prevents bugs caused by unexpected automatic conversions.

package main

import (
	"fmt"
	"strconv" // standard library for string ↔ number conversions
)

func main() {

	// ─── Numeric conversions ──────────────────────────────────────────────────

	var i int = 42
	var f float64 = float64(i) // int → float64: explicit conversion required
	var u uint = uint(f)       // float64 → uint: truncates decimal part

	fmt.Printf("int: %d  →  float64: %.1f  →  uint: %d\n", i, f, u)

	// Why is explicit conversion required?
	// Without it, this would be a compile error:
	// var f float64 = i  ← ERROR: cannot use i (int) as float64

	// Precision loss — converting float to int truncates (does NOT round)
	pi := 3.99
	truncated := int(pi) // truncates to 3, not 4
	fmt.Printf("float64 %.2f → int %d (truncated, not rounded)\n", pi, truncated)

	// Overflow — converting a large value to a smaller type wraps around
	var big int = 300
	small := int8(big) // int8 max is 127; 300 overflows
	fmt.Printf("int %d → int8 %d (overflow wraps around)\n", big, small)

	// ─── Integer size conversions ──────────────────────────────────────────────

	var a int32 = 100
	var b int64 = int64(a) // int32 → int64 (widening — safe, no data loss)
	var c int16 = int16(a) // int32 → int16 (narrowing — possible data loss if > 32767)

	fmt.Printf("int32: %d → int64: %d → int16: %d\n", a, b, c)

	// ─── float32 ↔ float64 ────────────────────────────────────────────────────

	var f32 float32 = 3.14159265358979 // loses precision — float32 has ~7 digits
	var f64 float64 = float64(f32)     // convert back — but precision is already lost

	fmt.Printf("float32: %v\n", f32) // 3.1415927 (rounded)
	fmt.Printf("float64 from f32: %v\n", f64) // still only float32 precision

	// ─── String ↔ byte slice and rune slice ───────────────────────────────────

	s := "Hello, 世界" // string with ASCII and Chinese characters

	// string → []byte (each byte of UTF-8 encoding)
	bytes := []byte(s)
	fmt.Println("As bytes:", bytes)

	// []byte → string
	backToString := string(bytes)
	fmt.Println("Back to string:", backToString)

	// string → []rune (each Unicode code point / character)
	runes := []rune(s)
	fmt.Println("Rune count:", len(runes)) // 9 characters (not 13 bytes)

	// int → string gives the character with that Unicode code point, NOT the digit
	ch := string(rune(65)) // 65 = Unicode 'A'
	fmt.Println("string(65):", ch) // A

	// ─── String ↔ number with strconv ────────────────────────────────────────

	// strconv.Itoa — int to string (decimal representation)
	num := 42
	str := strconv.Itoa(num)
	fmt.Printf("int %d → string %q (type: %T)\n", num, str, str)

	// strconv.Atoi — string to int (returns value + error)
	// Always check the error — the string might not be a valid number
	parsed, err := strconv.Atoi("123")
	if err == nil {
		fmt.Printf("string \"123\" → int %d\n", parsed)
	}

	// Bad input
	_, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("Atoi error:", err) // strconv.Atoi: parsing "abc": invalid syntax
	}

	// strconv.ParseFloat — string to float64
	fval, err := strconv.ParseFloat("3.14", 64) // 64 = float64 precision
	if err == nil {
		fmt.Printf("string \"3.14\" → float64 %.2f\n", fval)
	}

	// strconv.FormatFloat — float64 to string
	fstr := strconv.FormatFloat(3.14159, 'f', 2, 64) // 'f' = fixed, 2 = 2 decimal places
	fmt.Printf("float64 3.14159 → string %q\n", fstr) // "3.14"

	// ─── Type assertion on interface{} ────────────────────────────────────────

	// interface{} (or 'any' in Go 1.18+) can hold any type.
	// Type assertion extracts the concrete value with its real type.
	var val interface{} = "Hello, Go!"

	// Safe assertion — two-value form: (value, ok)
	// If wrong type, ok=false and value is the zero value — no panic
	str2, ok := val.(string)
	fmt.Printf("string assertion: %q, ok=%t\n", str2, ok) // "Hello, Go!", true

	num2, ok := val.(int)
	fmt.Printf("int assertion: %d, ok=%t\n", num2, ok) // 0, false

	// Type switch — cleanest way to handle multiple possible types
	describe := func(v interface{}) {
		switch t := v.(type) {
		case int:
			fmt.Printf("int: %d\n", t)
		case string:
			fmt.Printf("string: %q\n", t)
		case bool:
			fmt.Printf("bool: %t\n", t)
		default:
			fmt.Printf("unknown type: %T\n", t)
		}
	}

	describe(42)
	describe("hello")
	describe(true)
	describe(3.14)
}
