// 20 - Interfaces in Go
//
// An interface defines a SET OF METHOD SIGNATURES (behaviors) without implementation.
// Any type that implements all the methods of an interface AUTOMATICALLY satisfies it —
// no "implements" keyword needed. This is called IMPLICIT implementation.
//
// Interface syntax:
//   type InterfaceName interface {
//       MethodName(params) returnType
//       ...
//   }
//
// Why use interfaces?
//   - Write functions that work with ANY type that has certain behaviors
//   - Decouple code — depend on behavior, not concrete types (easier to test/extend)
//   - Polymorphism — treat different types uniformly through a shared interface

package main

import (
	"fmt"
	"math"
)

// ─── Define an interface ──────────────────────────────────────────────────────

// Shape defines the behavior that any shape must have
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ─── Implement the interface with concrete types ───────────────────────────────

// Circle — satisfies Shape because it has Area() and Perimeter() methods
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle — also satisfies Shape (two methods match the interface)
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Triangle — also satisfies Shape
type Triangle struct {
	A, B, C float64 // three side lengths
}

func (t Triangle) Area() float64 {
	// Heron's formula
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// ─── Function that accepts ANY Shape ─────────────────────────────────────────

// printShapeInfo takes a Shape interface — works with Circle, Rectangle, Triangle, etc.
// We don't care WHAT the shape is, only that it can tell us its Area and Perimeter.
func printShapeInfo(s Shape) {
	fmt.Printf("Type: %T\n", s)
	fmt.Printf("  Area:      %.2f\n", s.Area())
	fmt.Printf("  Perimeter: %.2f\n", s.Perimeter())
}

// ─── Another interface example: Stringer ─────────────────────────────────────

// Stringer is similar to Go's built-in fmt.Stringer interface.
// Any type with String() string is "printable" by fmt package automatically.
type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

// Person satisfies Stringer — fmt.Println will call this automatically
func (p Person) String() string {
	return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

// ─── Empty interface: interface{} / any ───────────────────────────────────────

// interface{} has NO methods — every type satisfies it.
// It can hold any value. Use 'any' (Go 1.18+) as a cleaner alias.
func printAnything(v interface{}) {
	fmt.Printf("value: %v, type: %T\n", v, v)
}

func main() {

	// ─── Polymorphism through interfaces ─────────────────────────────────────

	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}
	t := Triangle{A: 3, B: 4, C: 5}

	// Each concrete type is passed as the Shape interface
	shapes := []Shape{c, r, t}
	for _, s := range shapes {
		printShapeInfo(s)
		fmt.Println()
	}

	// ─── Interface variable ───────────────────────────────────────────────────

	// An interface variable can hold any value that satisfies the interface
	var s Shape
	s = Circle{Radius: 3}
	fmt.Printf("Circle area: %.2f\n", s.Area())

	s = Rectangle{Width: 2, Height: 8}
	fmt.Printf("Rectangle area: %.2f\n", s.Area())

	// ─── Stringer interface ───────────────────────────────────────────────────

	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p) // calls p.String() automatically → "Alice (age 30)"

	// ─── Empty interface ──────────────────────────────────────────────────────

	printAnything(42)
	printAnything("hello")
	printAnything(true)
	printAnything(Circle{Radius: 1})

	// ─── Type assertion on interface ──────────────────────────────────────────

	// When you have an interface value, you can get the concrete type back
	var shape Shape = Circle{Radius: 7}

	// Safe assertion — won't panic if wrong type
	if circle, ok := shape.(Circle); ok {
		fmt.Printf("It's a circle with radius %.0f\n", circle.Radius)
	}

	// Type switch — handle multiple types cleanly
	describe := func(s Shape) {
		switch v := s.(type) {
		case Circle:
			fmt.Printf("Circle radius: %.0f\n", v.Radius)
		case Rectangle:
			fmt.Printf("Rectangle %0.fx%.0f\n", v.Width, v.Height)
		case Triangle:
			fmt.Printf("Triangle sides: %.0f, %.0f, %.0f\n", v.A, v.B, v.C)
		}
	}

	for _, sh := range shapes {
		describe(sh)
	}

	// ─── Interface composition ────────────────────────────────────────────────

	// Interfaces can embed other interfaces to combine behaviors
	type ReadWriter interface {
		Read() string
		Write(s string)
	}
	// (not implemented here — just showing the concept)
	_ = (*ReadWriter)(nil) // suppress unused warning
}
