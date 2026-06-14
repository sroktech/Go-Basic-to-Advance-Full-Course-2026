// 15 - Slices in Go
//
// A slice is a DYNAMIC, flexible view into an underlying array.
// Unlike arrays (fixed size), slices can grow and shrink.
// Slices are the most commonly used data structure in Go.
//
// Internally, a slice has three parts:
//   - Pointer  → points to the first element in the underlying array
//   - Length   → number of elements currently in the slice (len)
//   - Capacity → total space in the underlying array from the pointer (cap)
//
// Slices are REFERENCE types — copying a slice copies the pointer, NOT the data.
// Two slices can share the same underlying array.

package main

import "fmt"

func main() {

	// ─── Creating slices ──────────────────────────────────────────────────────

	// Slice literal — like an array but without a size between the brackets
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Fruits:", fruits)
	fmt.Println("Length:", len(fruits))  // 3
	fmt.Println("Capacity:", cap(fruits)) // 3

	// make([]type, length, capacity) — creates a slice with given len and cap
	// All elements start at zero value
	nums := make([]int, 3, 5) // length=3, capacity=5
	fmt.Println("make slice:", nums)      // [0 0 0]
	fmt.Println("len:", len(nums), "cap:", cap(nums)) // 3, 5

	// nil slice — declared but not initialized; len=0, cap=0
	var empty []int
	fmt.Println("nil slice:", empty)           // []
	fmt.Println("nil slice == nil:", empty == nil) // true

	// ─── Append — adding elements ──────────────────────────────────────────────

	// append() adds elements and returns the NEW slice (always capture the return)
	// If capacity is exceeded, Go allocates a larger underlying array automatically
	scores := []int{10, 20, 30}
	scores = append(scores, 40)        // add one element
	scores = append(scores, 50, 60)    // add multiple elements
	fmt.Println("After append:", scores) // [10 20 30 40 50 60]

	// Append one slice to another using ... to expand the second slice
	extra := []int{70, 80}
	scores = append(scores, extra...) // ... spreads the slice into individual args
	fmt.Println("After append slice:", scores) // [10 20 30 40 50 60 70 80]

	// ─── Slicing — creating sub-slices ────────────────────────────────────────

	// slice[low:high] — elements from index low UP TO (not including) high
	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Println("letters[1:3]:", letters[1:3]) // [b c]       — index 1 and 2
	fmt.Println("letters[:2]:", letters[:2])   // [a b]       — from start to index 1
	fmt.Println("letters[3:]:", letters[3:])   // [d e]       — from index 3 to end
	fmt.Println("letters[:]:", letters[:])     // [a b c d e] — full copy reference

	// ─── Slices share the underlying array ────────────────────────────────────

	// sub is a view into the same array as letters — NOT a copy
	sub := letters[1:4] // [b c d]
	sub[0] = "B"        // modifies sub[0] which is also letters[1]

	fmt.Println("sub:", sub)         // [B c d]
	fmt.Println("letters:", letters) // [a B c d e] — original changed too!

	// To get a truly independent copy, use copy()
	original := []int{1, 2, 3, 4, 5}
	cloned := make([]int, len(original))
	copy(cloned, original) // copy(dst, src)
	cloned[0] = 999

	fmt.Println("original:", original) // [1 2 3 4 5] — unchanged
	fmt.Println("cloned:", cloned)     // [999 2 3 4 5]

	// ─── Iterating over a slice ────────────────────────────────────────────────

	colors := []string{"red", "green", "blue"}
	for i, color := range colors {
		fmt.Printf("  [%d] = %s\n", i, color)
	}

	// ─── Removing an element ──────────────────────────────────────────────────

	// Go has no built-in remove — use append to join the parts before and after
	items := []int{1, 2, 3, 4, 5}
	removeIndex := 2 // remove element at index 2 (value: 3)
	items = append(items[:removeIndex], items[removeIndex+1:]...)
	fmt.Println("After remove index 2:", items) // [1 2 4 5]

	// ─── 2D slices ────────────────────────────────────────────────────────────

	// A slice of slices — each inner slice can have a different length
	matrix := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9},
	}
	for _, row := range matrix {
		fmt.Println(row)
	}
}
