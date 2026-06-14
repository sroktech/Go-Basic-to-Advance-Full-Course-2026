// Functions in Go
//
// A function is a named, reusable block of code that performs a specific task.
// Instead of repeating the same code in multiple places, you write it once in a
// function and call it whenever you need it.
//
// General form:
//   func functionName(parameterName parameterType) returnType {
//       // body
//   }
//
// Key rules in Go:
//   - Functions are declared with the 'func' keyword
//   - Parameters must have explicit types
//   - If there is no return value, omit the return type entirely
//   - Go supports MULTIPLE return values (unique compared to many languages)
//   - Functions defined outside main() are available to all code in the same package

package main

import "fmt"

// ─── No return value (void) functions ────────────────────────────────────────
// These functions receive a string argument and print it.
// They have no return type — they just DO something and finish.
// 's string' means the parameter is named 's' and must be of type string.
// %s in Printf is the placeholder for a string value.
func sayCapital(s string) {
	fmt.Printf("%s, is the Capital \n", s)
}

func sayCountry(s string) {
	fmt.Printf("%s, is the Country  \n", s)
}

// ─── Single return value ──────────────────────────────────────────────────────
// max receives two int parameters and returns one int.
// When two parameters share the same type, Go lets you write: (n1, n2 int)
// instead of: (n1 int, n2 int) — both are equivalent.
// 'return' sends a value back to the caller and exits the function immediately.
func max(n1, n2 int) int {
	if n1 > n2 {
		return n1 // return n1 if it is the larger value
	} else {
		return n2 // otherwise return n2
	}
}

// ─── Multiple return values ───────────────────────────────────────────────────
// Go functions can return MORE than one value — listed in parentheses after the params.
// swap returns two strings: (string, string)
// The caller must capture both values: firstName, lastName := swap("John", "Doe")
// This is commonly used to return a result AND an error: (value, error)
func swap(x, y string) (string, string) {
	return y, x // return y first, then x — the positions are swapped
}

// ─── Call by Value ────────────────────────────────────────────────────────────
// When you pass a variable to a function in Go, the function receives a COPY.
// Any changes made to the parameter inside the function do NOT affect the original.
//
// Here, 'number' is a copy of whatever was passed in.
// Incrementing 'number' only changes the local copy — the original variable in main() stays the same.
func increment(number int) {
	number++
	fmt.Println("Inside increment: ", number) // prints the local copy (original + 1)
}

// ─── Call by Reference (using a slice) ───────────────────────────────────────
// Slices in Go are reference types — they contain a pointer to the underlying array.
// When you pass a slice to a function, both the caller and the function share the SAME array.
// So changes made inside the function ARE visible to the caller.
//
// This is NOT the same as passing a pointer explicitly (*[]int) — slices already
// carry a reference internally, which is why modification affects the original.
func modify(slice []int) {
	slice[0] = 999                          // changes the first element of the shared array
	fmt.Println("Inside modify: ", slice)   // prints [999 2000 3000]
}

func main() {
	// Calling void functions — no return value to capture
	sayCountry("USA")
	sayCapital("Washington")

	// Calling max — capture the single returned int into 'result'
	a, b := 100, 200
	result := max(a, b)
	fmt.Printf("Max value is : %d\n", result) // 200

	// Calling swap — capture BOTH return values at once using multi-assignment
	// Go requires you to receive all return values (use _ to discard one you don't need)
	firstName, LastName := swap("John", "Doe")
	fmt.Printf("Swapped names: %s %s\n", firstName, LastName) // Doe John

	// Call by value demo:
	// x is 10 before the call. Inside increment, the copy becomes 11.
	// But x in main() is still 10 after the call — the original is untouched.
	x := 10
	increment(x)
	fmt.Println("In main after increment:", x) // still 10 — original unchanged

	// Call by reference demo:
	// mySlice points to the same underlying array as 'slice' inside modify().
	// After the call, mySlice[0] is 999 — the change made inside modify() persists.
	mySlice := []int{1000, 2000, 3000}
	modify(mySlice)
	fmt.Println("In main after modify:", mySlice) // [999 2000 3000] — original changed
}
