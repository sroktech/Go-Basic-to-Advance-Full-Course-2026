// 14 - Structures (Structs) in Go
//
// A struct is a composite data type that groups together fields of different types
// under one name. It's Go's way of defining custom types — similar to a class in
// other languages, but without inheritance.
//
// Use structs when you want to model a "thing" with multiple properties:
//   Person  → name, age, email
//   Product → id, name, price, stock
//   Point   → x, y
//
// Struct syntax:
//   type StructName struct {
//       FieldName FieldType
//       ...
//   }

package main

import "fmt"

// ─── Define structs at package level ─────────────────────────────────────────

// Person groups related fields that describe a person
type Person struct {
	Name  string
	Age   int
	Email string
}

// Rectangle has two fields for its dimensions
type Rectangle struct {
	Width  float64
	Height float64
}

// ─── Methods on structs ───────────────────────────────────────────────────────

// A method is a function with a RECEIVER — the struct type it belongs to.
// Syntax: func (receiverName ReceiverType) MethodName() ReturnType
// This attaches the function to the Rectangle type.

// Value receiver — works on a COPY of the struct, does not modify original
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Pointer receiver — works on the ORIGINAL struct, can modify it
// Use a pointer receiver when: (1) you need to modify the struct, or (2) the struct is large
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor  // modifies original
	r.Height *= factor // modifies original
}

// ─── Stringer method ──────────────────────────────────────────────────────────

// If a type has a String() string method, fmt.Println will call it automatically.
// This is Go's way of customizing how a type prints.
func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d, Email: %s}", p.Name, p.Age, p.Email)
}

func main() {

	// ─── Creating structs ─────────────────────────────────────────────────────

	// Named field initialization — order doesn't matter, most readable
	p1 := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}

	// Positional initialization — order must match the struct definition
	// Less readable, avoid for structs with many fields
	p2 := Person{"Bob", 25, "bob@example.com"}

	// Zero-value struct — all fields get their zero values (0, "", false)
	var p3 Person
	fmt.Println("Zero Person:", p3) // { 0 }

	// ─── Accessing fields ─────────────────────────────────────────────────────

	// Use dot notation to read or write individual fields
	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)
	p2.Age = 26 // update a field
	fmt.Println("Updated Bob age:", p2.Age)

	// fmt.Println calls our String() method automatically
	fmt.Println(p1) // Person{Name: Alice, Age: 30, Email: alice@example.com}

	// %+v prints field names and values — great for debugging without a String() method
	fmt.Printf("%+v\n", p2) // {Name:Bob Age:26 Email:bob@example.com}

	// ─── Using methods ────────────────────────────────────────────────────────

	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("Area: %.2f\n", rect.Area())        // 15.00
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter()) // 16.00

	// Scale uses a pointer receiver — it modifies rect directly
	rect.Scale(2) // Go automatically takes the address: (&rect).Scale(2)
	fmt.Printf("After Scale(2): Width=%.0f Height=%.0f\n", rect.Width, rect.Height) // 10 6

	// ─── Structs are value types ──────────────────────────────────────────────

	// Assigning a struct copies ALL its fields — changes to the copy don't affect original
	original := Person{Name: "Charlie", Age: 40, Email: "charlie@example.com"}
	copy := original
	copy.Name = "Dave"

	fmt.Println("Original:", original.Name) // Charlie — unchanged
	fmt.Println("Copy:", copy.Name)         // Dave

	// ─── Pointer to struct ────────────────────────────────────────────────────

	// Use a pointer when you want to share and modify the same struct
	pPtr := &Person{Name: "Eve", Age: 28, Email: "eve@example.com"}
	pPtr.Age = 29 // Go auto-dereferences: same as (*pPtr).Age = 29
	fmt.Println("Eve's age:", pPtr.Age) // 29

	// ─── Anonymous structs ────────────────────────────────────────────────────

	// A struct defined and used in one place — no need for a named type
	// Good for one-off data shapes, test fixtures, or JSON decoding
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Server: %s:%d\n", config.Host, config.Port)

	// ─── Nested structs ───────────────────────────────────────────────────────

	type Address struct {
		City    string
		Country string
	}
	type Employee struct {
		Name    string
		Age     int
		Address Address // embed another struct as a field
	}

	emp := Employee{
		Name: "Frank",
		Age:  35,
		Address: Address{
			City:    "New York",
			Country: "USA",
		},
	}
	fmt.Println("Employee:", emp.Name, "lives in", emp.Address.City)
}
