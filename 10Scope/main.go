// 10 - Scope in Go
//
// Scope defines WHERE in your code a variable can be seen and used.
// If you try to use a variable outside its scope, Go gives a compile error.
//
// Go has three levels of scope:
//   1. Package scope  — declared outside all functions; visible to the whole package
//   2. Function scope — declared inside a function; visible only inside that function
//   3. Block scope    — declared inside { }; visible only inside that block (if, for, etc.)
//
// Shadowing: when an inner scope declares a variable with the SAME name as an outer scope,
// the inner one "shadows" (hides) the outer one inside that block.

package main

import "fmt"

// ─── Package-level (global) variables ────────────────────────────────────────
// Declared outside any function using 'var' (cannot use := at package level).
// Accessible from ANY function in this file (and package).
// Zero values are assigned automatically: 0 for int, "" for string, false for bool.
var g int      // zero value = 0
var x int = 25 // package-level x = 25
var y int = 55 // package-level y = 55

func main() {
	// ─── 1. Local (function-scope) variables ──────────────────────────────────
	// Declared inside main() — only visible within this function.
	// They disappear when the function returns.
	var a, b int // local integers, zero value = 0

	// Variable initialization
	a = 20
	b = 20
	g = a + b // 'g' is the package-level variable — we can read and write it here

	// Displaying the values of the variables
	fmt.Printf("a = %d, b = %d, global variable g = %d\n", a, b, g) // g = 40

	// ─── 2. Shadowing ─────────────────────────────────────────────────────────
	// We declare NEW local 'x' and 'y' inside main().
	// These SHADOW the package-level x (25) and y (55).
	// Inside main(), any reference to x or y now refers to the LOCAL ones.
	// The package-level x (25) and y (55) still exist but are hidden in this scope.
	var x int = 100  // shadows package-level x = 25
	var y int = 1000 // shadows package-level y = 55

	fmt.Printf("x = %d\n", x) // prints 100 (local), NOT 25 (package-level)
	fmt.Printf("This y local variable shadows the global variable = %d\n", y) // 1000

	// ─── 3. Block scope ───────────────────────────────────────────────────────
	// Variables declared inside an if/for/switch block exist ONLY inside that block.
	// Once the closing } is reached, the variable is gone.
	if true {
		blockVar := "I only exist inside this if block"
		fmt.Println(blockVar)
	}
	// fmt.Println(blockVar) // ← would be a compile error: undefined: blockVar

	// ─── 4. Short variable declaration scope with shadowing ───────────────────
	// := always creates a new variable in the CURRENT scope.
	// This means you can accidentally shadow an outer variable — watch out!
	z := "outer z"
	fmt.Println(z) // "outer z"
	{
		z := "inner z" // new variable — shadows outer z inside this block only
		fmt.Println(z) // "inner z"
	}
	fmt.Println(z) // "outer z" — outer z is unchanged after the block ends
}
