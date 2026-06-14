// 02 - Basic Syntax in Go
//
// Every Go source file follows this structure:
//   1. package declaration  — what package this file belongs to
//   2. import statement     — which packages this file uses
//   3. functions            — the actual code
//
// 'package main' is special: it tells Go this file is an executable program.
// Every Go program must have exactly one 'main' package with one 'main()' function —
// that is the entry point where execution begins.

// Package declaration — must be the very first non-comment line
package main

// Import statement — brings in the "fmt" package from Go's standard library.
// "fmt" stands for "format" and provides functions for printing and formatting text.
// fmt.Println, fmt.Printf, fmt.Print all live here.
import "fmt"

// main() is the entry point of every Go program.
// When you run 'go run main.go', Go calls main() automatically.
// It takes no parameters and returns nothing.
func main() {

	// ─── Single-line comment ──────────────────────────────────────────────────
	// Anything after // on a line is ignored by the compiler. Use these for short notes.

	/*
		Multi-line comment.
		Everything between slash-star and star-slash is ignored.
		Use for longer explanations that span several lines.
	*/

	// ─── Printing to the console ──────────────────────────────────────────────

	// println() is a built-in Go function — works without any import.
	// It prints to stderr and is intended for debugging only.
	// Avoid it in real programs — use fmt.Println instead.
	println("Hello world")

	// fmt.Println — prints to stdout with an automatic newline at the end.
	// This is the standard way to print in Go.
	fmt.Println("Hello, Go!")

	// fmt.Print — prints WITHOUT an automatic newline.
	// Output stays on the same line until you add \n yourself.
	fmt.Print("Hello ")
	fmt.Print("World!\n") // \n moves to the next line manually

	// fmt.Printf — formatted print using format verbs (placeholders):
	//   %s = string    %d = integer    %f = float
	//   %t = bool      %T = type name  %v = any value (default format)
	name := "Gopher"
	age := 5
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// ─── Key syntax rules in Go ───────────────────────────────────────────────
	// 1. No semicolons — Go inserts them automatically; never write them yourself.
	// 2. Opening brace { must be on the SAME line as if/for/func — never the next line.
	// 3. Go is case-sensitive: fmt.Println (capital P) is different from fmt.println.
	// 4. Uppercase first letter = exported (public); lowercase = unexported (private).
	//    That is why Println is capitalized — it is exported from the fmt package.
	// 5. Every imported package MUST be used — unused imports are a compile error.
	// 6. Every declared variable MUST be used — unused variables are a compile error.
}
