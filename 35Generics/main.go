// 35 - Generics in Go (Go 1.18+)
//
// Generics let you write functions and types that work with ANY type,
// without duplicating code and without losing type safety.
//
// Before generics, you had two bad options:
//   1. Write the same function for every type (int version, string version, ...)
//   2. Use interface{} — loses type safety, requires type assertions everywhere
//
// With generics, you write it ONCE and the compiler generates type-safe versions.
//
// FLOW — generic function:
//
//   func Min[T int | float64](a, b T) T
//         │ │
//         │ └─ T is constrained: only int or float64 allowed
//         └─── [T ...] is the type parameter list
//
//   Call: Min[int](3, 5)       ← explicit type
//         Min(3, 5)            ← type inferred from arguments (preferred)
//         Min(3.14, 2.71)      ← infers float64
//
// Type constraints define what a type parameter is allowed to be.
// Built-in constraints live in the 'constraints' package (or define your own).

package main

import "fmt"

// ─── Type constraints ─────────────────────────────────────────────────────────

// Integer is a constraint — T must be one of these types
type Integer interface {
	int | int8 | int16 | int32 | int64
}

// Number allows any integer or float
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Ordered allows types that support < > == (can be compared/sorted)
type Ordered interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64 | string
}

// ─── Generic functions ────────────────────────────────────────────────────────

// Min returns the smaller of two values — works for any Ordered type
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of two values
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Sum adds all numbers in a slice — works for any Number type
func Sum[T Number](nums []T) T {
	var total T // zero value of T
	for _, n := range nums {
		total += n
	}
	return total
}

// Contains reports whether item is in the slice — works for any comparable type
// 'comparable' is a built-in constraint meaning the type supports == and !=
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Map applies a transform function to each element — like map() in Python/JS
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter returns elements that satisfy a predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce folds a slice into a single value
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// ─── Generic data structures ──────────────────────────────────────────────────

// Stack[T] is a generic LIFO (last in, first out) stack
// Works with any type: Stack[int], Stack[string], Stack[Person], etc.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T // zero value of T
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func (s *Stack[T]) Len() int { return len(s.items) }

// Pair holds two values of potentially different types
type Pair[A, B any] struct {
	First  A
	Second B
}

func NewPair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{First: a, Second: b}
}

func main() {

	// ─── Min / Max ────────────────────────────────────────────────────────────

	fmt.Println("=== Min / Max ===")
	fmt.Println("Min(3, 7):", Min(3, 7))           // int — inferred
	fmt.Println("Min(3.14, 2.71):", Min(3.14, 2.71)) // float64 — inferred
	fmt.Println("Min(\"apple\", \"banana\"):", Min("apple", "banana"))
	fmt.Println("Max(10, 20):", Max(10, 20))

	// ─── Sum ──────────────────────────────────────────────────────────────────

	fmt.Println("\n=== Sum ===")
	ints := []int{1, 2, 3, 4, 5}
	floats := []float64{1.1, 2.2, 3.3}
	fmt.Println("Sum(ints):", Sum(ints))     // 15
	fmt.Println("Sum(floats):", Sum(floats)) // 6.6

	// ─── Contains ─────────────────────────────────────────────────────────────

	fmt.Println("\n=== Contains ===")
	fmt.Println("Contains([1,2,3], 2):", Contains([]int{1, 2, 3}, 2))          // true
	fmt.Println("Contains([1,2,3], 9):", Contains([]int{1, 2, 3}, 9))          // false
	fmt.Println("Contains([\"a\",\"b\"], \"b\"):", Contains([]string{"a", "b"}, "b")) // true

	// ─── Map / Filter / Reduce ────────────────────────────────────────────────

	fmt.Println("\n=== Map / Filter / Reduce ===")

	nums := []int{1, 2, 3, 4, 5, 6}

	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println("Doubled:", doubled) // [2 4 6 8 10 12]

	strs := Map(nums, func(n int) string { return fmt.Sprintf("item%d", n) })
	fmt.Println("As strings:", strs) // [item1 item2 ...]

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("Evens:", evens) // [2 4 6]

	total := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println("Sum via Reduce:", total) // 21

	// ─── Generic Stack ────────────────────────────────────────────────────────

	fmt.Println("\n=== Generic Stack ===")

	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Println("Stack len:", intStack.Len())
	for intStack.Len() > 0 {
		val, _ := intStack.Pop()
		fmt.Println("  popped:", val) // 3, 2, 1
	}

	// Same Stack type, different type parameter
	strStack := Stack[string]{}
	strStack.Push("go")
	strStack.Push("generics")
	top, _ := strStack.Pop()
	fmt.Println("String stack top:", top) // generics

	// ─── Generic Pair ────────────────────────────────────────────────────────

	fmt.Println("\n=== Generic Pair ===")

	p1 := NewPair("hello", 42)
	p2 := NewPair(3.14, true)
	fmt.Printf("Pair 1: (%v, %v)\n", p1.First, p1.Second) // (hello, 42)
	fmt.Printf("Pair 2: (%v, %v)\n", p2.First, p2.Second) // (3.14, true)

	// ─── When to use generics ────────────────────────────────────────────────
	fmt.Println(`
Use generics when:
  ✓ Writing utility functions (Min, Max, Contains, Map, Filter)
  ✓ Building data structures (Stack, Queue, Set, Pair)
  ✓ The logic is identical regardless of type

Don't use generics when:
  ✗ An interface{} with type assertions is simpler
  ✗ The behavior differs per type (use interfaces/methods instead)
  ✗ You only have one type — just write the function directly`)
}
