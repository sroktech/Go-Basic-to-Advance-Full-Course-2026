// Data Types
// Go is a statically typed language. Every variable has a type known at compile time.
// Core types: integers, floating points, booleans, complex numbers, and strings.
package main

import "fmt"

func main() {

	// Section1: Integers
	// Integers are whole numbers with no decimal component.
	// Go provides both signed (can hold negative values) and unsigned (non-negative only) variants.
	// The number suffix indicates the bit size, which determines the range of values the type can hold.

	// Section1-A: Signed integers — can hold negative and positive values
	// int   — platform-dependent: 32-bit on 32-bit systems, 64-bit on 64-bit systems; range: −2^(n-1) to 2^(n-1)−1
	// int8  — 8-bit,  range: −128 to 127
	// int16 — 16-bit, range: −32,768 to 32,767
	// int32 — 32-bit, range: −2,147,483,648 to 2,147,483,647 (also aliased as 'rune' for Unicode code points)
	// int64 — 64-bit, range: −9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	i = -128
	i8 = 127
	i16 = -32768
	i32 = 2147483647
	i64 = -9223372036854775808

	// Section1-B: Unsigned integers — can only hold zero and positive values, giving double the positive range
	// uint   — platform-dependent size; range: 0 to 2^n−1
	// uint8  — 8-bit,  range: 0 to 255 (also aliased as 'byte' for raw binary data)
	// uint16 — 16-bit, range: 0 to 65,535
	// uint32 — 32-bit, range: 0 to 4,294,967,295
	// uint64 — 64-bit, range: 0 to 18,446,744,073,709,551,615
	var u uint
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64

	u = 255
	u8 = 255
	u16 = 65535
	u32 = 4294967295
	u64 = 18446744073709551615

	fmt.Println("Signed integers:", i, i8, i16, i32, i64)
	fmt.Println("Unsigned integers:", u, u8, u16, u32, u64)

	// Section2: Floating Point
	// Floating-point types represent numbers with a decimal component using IEEE 754 standard.
	// A larger bit size means more digits of precision but uses more memory.

	// float32 — 32-bit, ~6–7 significant decimal digits of precision; suitable for graphics or memory-constrained cases
	var f32 float32 = 10.6
	// float64 — 64-bit, ~15–16 significant decimal digits of precision; the default float type in Go, preferred for most math
	var f64 float64 = 10.6

	fmt.Println("FLOAT32:", f32)
	fmt.Println("FLOAT64:", f64)

	// Demonstrating the precision difference: float64 preserves more significant digits than float32
	var HP float64 = 10123456789012345
	var LP float32 = 10123456789012345
	fmt.Println("High precision float64:", HP)
	fmt.Println("Low precision float32:", LP)

	// Section3: Boolean Data Type
	// bool holds one of two values: true or false.
	// Used for conditional logic and control flow. The zero value (default) is false.
	var isActive bool = true
	var isOn bool = false

	fmt.Println("Is Active:", isActive)
	fmt.Println("Is On:", isOn)

	// Section4: Complex Data Type
	// Go has built-in support for complex numbers composed of a real part and an imaginary part.
	// complex64  — real and imaginary parts are each float32 (lower precision, less memory)
	// complex128 — real and imaginary parts are each float64 (higher precision, default complex type)
	// Use complex(real, imag) to construct, and real()/imag() to extract the parts.
	var CN1 complex128 = complex(5, 10) // represents 5 + 10i
	var CN2 complex64 = complex(2, 7)   // represents 2 + 7i
	fmt.Println("CN1: ", CN1)
	fmt.Println("CN2: ", CN2)

	// Section5: String Data Type
	// string is an immutable sequence of bytes, typically representing UTF-8 encoded text.
	// Strings are defined with double quotes. The zero value is an empty string "".
	// len(s) returns the byte count, not the character count (relevant for multi-byte Unicode).
	var name string = "Kermet the frog!"
	fmt.Println("My name is:", name)
}
