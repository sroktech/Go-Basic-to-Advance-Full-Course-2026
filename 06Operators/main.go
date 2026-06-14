// Lesson6: Operators in Go Language
// An operator is a symbol that tells the compiler to perform a specific operation on values.
// Go supports the following categories of operators:
//   - Arithmetic Operators  — math operations (+, -, *, /, %, ++, --)
//   - Relational Operators  — compare two values and return a bool (==, !=, >, <, >=, <=)
//   - Logical Operators     — combine boolean values (&&, ||, !)
//   - Bitwise Operators     — operate directly on binary bits (&, |, ^, <<, >>)
//   - Assignment Operators  — assign and update values (=, +=, -=, etc.)
//   - Miscellaneous         — address-of (&) and pointer dereference (*)

package main

import "fmt"

func main() {
	// ─── Arithmetic Operators ───────────────────────────────────────────────────
	// Used to perform basic math. Operands must be the same type (Go won't auto-convert).
	// A := 10
	// B := 20
	// C := 11
	// D := 3
	// fmt.Println("A + B =", A+B) // Addition:       10 + 20 = 30
	// fmt.Println("A - B =", A-B) // Subtraction:    10 - 20 = -10
	// fmt.Println("A * B =", A*B) // Multiplication: 10 * 20 = 200
	// fmt.Println("A / B =", A/B) // Division:       10 / 20 = 0 (integer division truncates, no remainder)
	// fmt.Println("C % D =", C%D) // Modulus:        11 % 3  = 2 (remainder after division)
	// A++                          // Increment: adds 1 to A (A = 10 → 11). Only postfix (A++) is valid in Go, not ++A.
	// fmt.Println("A++ =", A)      // 11
	// A--                          // Decrement: subtracts 1 from A (A = 11 → 10). Same rule — only A--, not --A.
	// fmt.Println("A-- =", A)      // 10

	// ─── Relational (Comparison) Operators ─────────────────────────────────────
	// Compare two values and always return a bool (true or false).
	// Used heavily in if statements, loops, and conditions.
	// A := 10
	// B := 20
	// fmt.Println("A == B:", A == B) // Equal to:              10 == 20 → false
	// fmt.Println("A != B:", A != B) // Not equal to:          10 != 20 → true
	// fmt.Println("A > B:", A > B)   // Greater than:          10 > 20  → false
	// fmt.Println("A < B:", A < B)   // Less than:             10 < 20  → true
	// fmt.Println("A >= B:", A >= B) // Greater than or equal: 10 >= 20 → false
	// fmt.Println("A <= B:", A <= B) // Less than or equal:    10 <= 20 → true

	// ─── Logical Operators ──────────────────────────────────────────────────────
	// Combine or invert boolean values. Operands must be bool type.
	// A := true
	// B := false
	// fmt.Println("A && B:", A && B) // AND: true only if BOTH are true  → false
	// fmt.Println("A || B:", A || B) // OR:  true if AT LEAST ONE is true → true
	// fmt.Println("!A:", !A)         // NOT: flips the value              → false (true becomes false)
	// fmt.Println("!B:", !B)         // NOT: flips the value              → true  (false becomes true)

	// ─── Assignment Operators ───────────────────────────────────────────────────
	// Shorthand to update a variable by applying an operation to its current value.
	// These are equivalent to writing the full expression: A += B is the same as A = A + B.
	// A := 10
	// B := 20
	// A += B                    // A = A + B → A becomes 30
	// fmt.Println("A += B:", A) // 30
	// A -= B                    // A = A - B → A becomes 10 again
	// fmt.Println("A -= B:", A) // 10
	// Other assignment operators: *=, /=, %=, &=, |=, ^=, <<=, >>= (same pattern)

	// ─── Bitwise Operators ──────────────────────────────────────────────────────
	// Operate on the individual binary (0/1) bits of an integer, not the whole number.
	// Useful for low-level tasks: flags, permissions, encoding, and performance-critical code.
	//
	// C := 30  →  binary: 11110
	//          2  →  binary: 00010
	//
	// & (AND): bit is 1 only if BOTH bits are 1
	//   11110 & 00010 = 00010 → 2
	// fmt.Println("C & 2 =", C&2) // 2
	//
	// | (OR): bit is 1 if AT LEAST ONE bit is 1
	//   11110 | 00010 = 11110 → 30
	// fmt.Println("C | 2 =", C|2) // 30
	//
	// ^ (XOR): bit is 1 only if the bits are DIFFERENT
	//   11110 ^ 00010 = 11100 → 28
	// fmt.Println("C ^ 2 =", C^2) // 28
	//
	// ^ (NOT / complement): flips all bits. In Go, ^ is also used as unary NOT.
	//   ^30 = -(30+1) = -31  (because Go uses two's complement for negative numbers)
	// fmt.Println("^C =", ^C) // -31
	//
	// << (left shift): shifts bits left by N positions, equivalent to multiplying by 2^N
	//   11110 << 1 = 111100 → 60  (30 * 2 = 60)
	// fmt.Println("C << 1 =", C<<1) // 60
	//
	// >> (right shift): shifts bits right by N positions, equivalent to dividing by 2^N
	//   11110 >> 1 = 01111 → 15  (30 / 2 = 15)
	// fmt.Println("C >> 1 =", C>>1) // 15

	// ─── Miscellaneous Operators ────────────────────────────────────────────────
	// Go has two operators related to pointers — a core concept for working with memory directly.
	A := 10

	// & (address-of operator): returns the memory address where variable A is stored.
	// 'ptr' is a pointer — its type is *int, meaning "a pointer to an int".
	// The actual address value looks like 0xc000018090 (changes each run).
	ptr := &A
	fmt.Println("Address of A:", ptr) // prints the memory address of A

	// * (dereference operator): reads the value stored at the memory address the pointer holds.
	// *ptr means "go to the address in ptr and give me the value there" → 10.
	// You can also write *ptr = 99 to change the value of A through the pointer.
	fmt.Println("Value of *ptr:", *ptr) // 10
}
