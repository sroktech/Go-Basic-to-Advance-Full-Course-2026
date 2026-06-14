// 13 - Pointers in Go
//
// A pointer is a variable that stores the MEMORY ADDRESS of another variable,
// rather than storing the value itself.
//
// Think of memory as a row of numbered boxes. Each box has an address (like a house number).
// A normal variable stores a VALUE in a box.
// A pointer stores the ADDRESS of a box — it tells you WHERE to find the value.
//
// Two key operators:
//   &  (address-of)  — gives you the memory address of a variable
//   *  (dereference) — follows a pointer to get/set the value at that address
//
// Pointer type syntax:
//   *int   = a pointer to an int
//   *string = a pointer to a string
//
// Why use pointers?
//   1. To modify a variable inside a function (call by reference)
//   2. To share large data without copying it (efficiency)

package main

import "fmt"

func main() {

	// ─── Basic pointer ────────────────────────────────────────────────────────

	x := 42

	// & gives the memory address of x
	// ptr is of type *int — "a pointer to an int"
	ptr := &x
	fmt.Println("Value of x:      ", x)    // 42
	fmt.Println("Address of x:    ", &x)   // e.g. 0xc000018090 (changes each run)
	fmt.Println("ptr holds:       ", ptr)   // same address as &x
	fmt.Printf("Type of ptr:     %T\n", ptr) // *int

	// * (dereference) — follow the pointer to read the value at that address
	fmt.Println("Value at *ptr:   ", *ptr) // 42 — same as x

	// ─── Modifying via pointer ────────────────────────────────────────────────

	// Writing through a pointer changes the ORIGINAL variable
	*ptr = 100           // set the value at ptr's address to 100
	fmt.Println("x after *ptr = 100:", x) // x is now 100 — both ptr and x point to the same location

	// ─── new() — allocate a pointer with zero value ───────────────────────────

	// new(T) allocates memory for a value of type T, sets it to zero, and returns a *T
	// Useful when you want a pointer without an existing variable
	p := new(int)   // p is *int, *p is 0
	fmt.Println("new int pointer:", p)   // memory address
	fmt.Println("zero value at p:", *p)  // 0
	*p = 77
	fmt.Println("after *p = 77:", *p)   // 77

	// ─── Nil pointer ──────────────────────────────────────────────────────────

	// A pointer that hasn't been assigned an address holds the value nil.
	// Dereferencing a nil pointer causes a runtime panic — always check before use.
	var nilPtr *int // nil pointer — doesn't point to anything
	fmt.Println("nil pointer:", nilPtr) // <nil>

	if nilPtr == nil {
		fmt.Println("nilPtr is nil — safe, we did not dereference it")
	}

	// ─── Passing pointer to a function (call by reference) ────────────────────

	// When you pass a regular variable, the function gets a COPY — changes don't propagate.
	// When you pass a pointer, the function can modify the ORIGINAL variable.
	a := 10
	fmt.Println("Before double:", a) // 10
	double(&a)                       // pass the address of a
	fmt.Println("After double:", a)  // 20 — original was changed

	// ─── Pointer to a struct (common in real Go code) ─────────────────────────

	type Point struct {
		X, Y int
	}

	// Create a struct and a pointer to it
	pt := Point{X: 3, Y: 4}
	ptPtr := &pt

	// Go lets you access struct fields through a pointer WITHOUT writing (*ptPtr).X
	// Both are equivalent: ptPtr.X and (*ptPtr).X
	fmt.Println("Point via pointer:", ptPtr.X, ptPtr.Y) // 3 4

	// Modify through pointer
	ptPtr.X = 99
	fmt.Println("After ptPtr.X = 99:", pt.X) // 99 — original struct changed
}

// double takes a pointer to an int and multiplies the value by 2
// *n means "the value at the address n" — we're modifying the original
func double(n *int) {
	*n = *n * 2
}
