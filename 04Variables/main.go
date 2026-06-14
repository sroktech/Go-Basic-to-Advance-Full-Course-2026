// Understanding Variables Declarations in Go
// A variable is a named storage location that holds a value.
// In Go, every variable has a specific type that determines what values it can store.
// Variables must be declared before use, and every declared variable must be used — otherwise Go will not compile.
package main

import "fmt"

func main() {
	// Way 1: Declare and assign on 2 separate lines (explicit type, then assign)
	// 'var' keyword declares the variable, followed by the name, then the type.
	// The zero value is assigned first (e.g., "" for string, 0 for int), then overwritten.
	var mango string = "This is a big mango!"
	var weight int = 54

	// Way 2: Declare and assign on the same line (explicit type + value together)
	// This is the most readable form — type is clear at a glance.
	var height int = 23

	fmt.Println("Mango:", mango)
	fmt.Println("weight:", weight)
	fmt.Println("height:", height)

	// Way 3: Shorthand declaration using := (short variable declaration)
	// := is only available inside functions (not at package level).
	// Go infers the type from the value on the right — this is called type inference.
	// 'age' becomes int (Go defaults untyped integers to int)
	// 'city' becomes string
	// Equivalent to: var age int = 54 / var city string = "Washington"
	age := 54
	city := "Washington"
	fmt.Println("My age is:", age)
	fmt.Println("My city is:", city)

	// Multiple variables of the same type declared and initialized on one line.
	// Both 'apples' and 'oranges' are int — the type applies to all names on the left.
	var apples, oranges int = 23, 78
	fmt.Println("I have", apples, "apples and", oranges, "oranges")

	// Type inference with an expression — Go figures out the type from the result.
	// Since apples and oranges are both int, 'fruits' will also be inferred as int.
	var fruits = apples + oranges
	fmt.Println("fruits:", fruits)

	// Multiple string variables declared together using var with explicit type.
	// The \n inside the strings is a newline escape sequence — moves output to next line.
	var windows, mac, linux string = "Windows is ok\n", "Mac is meh\n", "Linux is the GOAT!\n"
	print(windows, mac, linux)

	// Static (explicit) type declaration
	// The type float64 is explicitly written — Go will not allow assigning incompatible types later.
	// %T in Printf prints the Go type of the variable (e.g., float64, int, string).
	var x float64 = 20.5
	fmt.Println(x)
	fmt.Printf("x is of type: %T\n", x)

	// Dynamic (inferred) type declaration using :=
	// Go looks at the value 89 and infers the type as int.
	// The type is still fixed after inference — Go is NOT dynamically typed like Python/JS.
	y := 89
	fmt.Println(y)
	fmt.Printf("y is of type: %T\n", y)

	// Mixed variable declaration — multiple variables with different types in one var statement.
	// Go infers each type independently from its value:
	//   a = 758.52 → float64
	//   b = 8      → int
	//   c = "foobar" → string
	// This works because 'var' with multiple assignments can hold different types, unlike typed multi-var.
	var a, b, c = 758.52, 8, "foobar"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type: %T\n", a)
	fmt.Printf("b is of type: %T\n", b)
	fmt.Printf("c is of type: %T\n", c)
}
