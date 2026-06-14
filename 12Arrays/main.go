// 12 - Arrays in Go
//
// An array is a FIXED-SIZE, ordered collection of elements of the SAME type.
// "Fixed-size" means the length is set when you declare it and can NEVER change.
// If you need a resizable collection, use a Slice (lesson 15) instead.
//
// Array syntax:
//   var name [size]type
//   name := [size]type{value1, value2, ...}
//
// Key property: arrays in Go are VALUES, not references.
// Assigning one array to another copies ALL the elements.

package main

import "fmt"

func main() {

	// ─── Declaring arrays ─────────────────────────────────────────────────────

	// Declare with zero values — all elements default to the zero value of the type
	// int zero = 0, string zero = "", bool zero = false
	var scores [5]int
	fmt.Println("Zero-value array:", scores) // [0 0 0 0 0]

	// Declare and initialize with values
	colors := [3]string{"red", "green", "blue"}
	fmt.Println("Colors:", colors) // [red green blue]

	// Let Go count the size automatically using ... — Go infers length = 4
	primes := [...]int{2, 3, 5, 7}
	fmt.Println("Primes:", primes)       // [2 3 5 7]
	fmt.Println("Length:", len(primes))  // 4

	// ─── Accessing and modifying elements ─────────────────────────────────────

	// Arrays are zero-indexed: first element is index 0, last is index len-1
	fmt.Println("First prime:", primes[0]) // 2
	fmt.Println("Last prime:", primes[3])  // 7

	// Modify an element by assigning to its index
	scores[0] = 95
	scores[1] = 82
	scores[2] = 78
	scores[3] = 91
	scores[4] = 88
	fmt.Println("Scores:", scores) // [95 82 78 91 88]

	// ─── Iterating over an array ──────────────────────────────────────────────

	// Classic for loop with index
	fmt.Println("Scores with index:")
	for i := 0; i < len(scores); i++ {
		fmt.Printf("  scores[%d] = %d\n", i, scores[i])
	}

	// for-range loop — cleaner, gives both index and value
	// Use _ to discard the index if you don't need it
	fmt.Println("Colors with range:")
	for index, value := range colors {
		fmt.Printf("  [%d] = %s\n", index, value)
	}

	// ─── Arrays are value types ───────────────────────────────────────────────

	// Assigning an array to another variable creates a COMPLETE COPY.
	// Changes to the copy do NOT affect the original.
	original := [3]int{1, 2, 3}
	copyArr := original    // full copy of all 3 elements
	copyArr[0] = 999       // modify the copy

	fmt.Println("Original:", original) // [1 2 3] — unchanged
	fmt.Println("Copy:", copyArr)      // [999 2 3]

	// ─── Multi-dimensional arrays ─────────────────────────────────────────────

	// A 2D array is an array of arrays — think of it as a grid/table.
	// [rows][cols]type
	var matrix [3][3]int // 3x3 grid of integers, all zeros

	// Fill the diagonal (top-left to bottom-right) with 1s
	matrix[0][0] = 1
	matrix[1][1] = 1
	matrix[2][2] = 1

	// Print the matrix row by row
	fmt.Println("3x3 matrix:")
	for _, row := range matrix {
		fmt.Println(" ", row)
	}

	// Initialize a 2D array with values directly
	grid := [2][3]string{
		{"a", "b", "c"}, // row 0
		{"d", "e", "f"}, // row 1
	}
	fmt.Println("Grid row 0:", grid[0]) // [a b c]
	fmt.Println("Grid[1][2]:", grid[1][2]) // f

	// ─── Important limitation ─────────────────────────────────────────────────

	// The size is PART of the type: [3]int and [4]int are different types.
	// You CANNOT pass a [3]int where a [4]int is expected.
	// This is why slices (lesson 15) are used far more often in real Go code.
	fmt.Printf("Type of scores: %T\n", scores)  // [5]int
	fmt.Printf("Type of primes: %T\n", primes)  // [4]int
}
