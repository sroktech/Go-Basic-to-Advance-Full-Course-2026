// 16 - Range in Go
//
// 'range' is a keyword used with 'for' to iterate over data structures.
// It gives you both the INDEX and the VALUE for each iteration.
//
// You can use it on:
//   - Arrays and Slices  → index, value
//   - Strings            → byte index, rune (Unicode character)
//   - Maps               → key, value
//   - Channels           → value (advanced — not covered here)
//
// Syntax:
//   for index, value := range collection { }
//
// Use _ to discard either the index or the value if you don't need it:
//   for _, value := range collection { }   // ignore index
//   for index := range collection { }      // ignore value (just the index)

package main

import "fmt"

func main() {

	// ─── Range over a slice ───────────────────────────────────────────────────

	fruits := []string{"apple", "banana", "cherry", "date"}

	// Both index and value
	fmt.Println("All fruits:")
	for i, fruit := range fruits {
		fmt.Printf("  [%d] = %s\n", i, fruit)
	}

	// Only value — discard index with _
	fmt.Println("Fruit names only:")
	for _, fruit := range fruits {
		fmt.Println(" ", fruit)
	}

	// Only index — useful for checking positions, modifying in place
	fmt.Println("Indices only:")
	for i := range fruits {
		fmt.Printf("  index %d\n", i)
	}

	// ─── Range over an array ──────────────────────────────────────────────────

	primes := [5]int{2, 3, 5, 7, 11}
	sum := 0
	for _, v := range primes {
		sum += v // accumulate sum of all elements
	}
	fmt.Println("Sum of primes:", sum) // 28

	// ─── Range over a string ──────────────────────────────────────────────────

	// Ranging over a string gives you: byte index + rune (Unicode code point)
	// This handles multi-byte characters correctly (unlike indexing with s[i])
	word := "Go🚀"
	fmt.Println("Characters in 'Go🚀':")
	for i, ch := range word {
		fmt.Printf("  byte index %d: %c (Unicode: U+%04X)\n", i, ch, ch)
	}
	// Output shows 🚀 starts at byte index 2 (not character index 2)
	// because 🚀 is 4 bytes in UTF-8

	// ─── Range over a map ─────────────────────────────────────────────────────

	capitals := map[string]string{
		"USA":     "Washington D.C.",
		"France":  "Paris",
		"Japan":   "Tokyo",
	}

	// Map iteration order is RANDOM — Go intentionally randomizes it each run
	fmt.Println("Capitals:")
	for country, capital := range capitals {
		fmt.Printf("  %s → %s\n", country, capital)
	}

	// Only keys
	fmt.Println("Countries:")
	for country := range capitals {
		fmt.Println(" ", country)
	}

	// ─── Range with modification ──────────────────────────────────────────────

	// 'value' in range is a COPY — modifying it does NOT change the original slice
	// To modify elements in place, use the index
	numbers := []int{1, 2, 3, 4, 5}

	// Wrong way — this modifies a copy, original unchanged
	for _, v := range numbers {
		v *= 2 // only changes local copy
	}
	fmt.Println("After wrong modify:", numbers) // [1 2 3 4 5] — unchanged

	// Correct way — use the index to modify the slice directly
	for i := range numbers {
		numbers[i] *= 2 // modifies the actual element
	}
	fmt.Println("After correct modify:", numbers) // [2 4 6 8 10]

	// ─── Range to build a new slice ───────────────────────────────────────────

	// Common pattern: filter elements into a new slice
	evens := []int{}
	for _, n := range numbers {
		if n%4 == 0 { // pick multiples of 4 from the doubled numbers
			evens = append(evens, n)
		}
	}
	fmt.Println("Multiples of 4:", evens)
}
