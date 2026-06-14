// 11 - Strings in Go
//
// A string in Go is an immutable sequence of bytes, typically UTF-8 encoded text.
// "Immutable" means once created, the string content cannot be changed in place —
// you always create a new string when you modify one.
//
// Two string literal forms:
//   "interpreted string"  — supports escape sequences like \n, \t, \"
//   `raw string literal`  — everything between backticks is literal, no escape processing

package main

import (
	"fmt"
	"strings" // standard library package for string operations
	"unicode/utf8" // for counting Unicode characters (runes), not just bytes
)

func main() {

	// ─── Creating strings ─────────────────────────────────────────────────────

	// Interpreted string — backslash sequences are processed
	greeting := "Hello, Go!\n"
	fmt.Print(greeting) // \n causes a newline

	// Raw string literal — backticks, no escape processing
	// Useful for file paths, regex patterns, multi-line text
	raw := `Hello, Go!\n` // \n is printed literally, not as a newline
	fmt.Println(raw)

	// Multi-line raw string
	poem := `Roses are red,
Violets are blue,
Go is simple,
And fast too!`
	fmt.Println(poem)

	// ─── String length ────────────────────────────────────────────────────────

	word := "Hello"
	fmt.Println("Byte length:", len(word)) // len() counts BYTES, not characters

	// For strings with non-ASCII characters (e.g., emoji, Chinese), bytes ≠ characters
	emoji := "Go🚀"
	fmt.Println("Byte length of Go🚀:", len(emoji))                     // 6 (🚀 = 4 bytes)
	fmt.Println("Rune count of Go🚀:", utf8.RuneCountInString(emoji)) // 3 (3 characters)

	// ─── Accessing characters ─────────────────────────────────────────────────

	// Indexing a string gives a BYTE (uint8), not a character
	// This works fine for ASCII but can give wrong results for multi-byte characters
	s := "Golang"
	fmt.Printf("First byte: %c (decimal: %d)\n", s[0], s[0]) // 'G' = 71

	// To iterate over characters (runes) properly, use a for-range loop
	fmt.Println("Characters in 'Golang':")
	for i, ch := range "Golang" {
		fmt.Printf("  index %d: %c\n", i, ch) // ch is a rune (Unicode code point)
	}

	// ─── String concatenation ─────────────────────────────────────────────────

	// + joins two strings — creates a new string (original unchanged)
	first := "Hello"
	second := "World"
	combined := first + ", " + second + "!"
	fmt.Println(combined) // Hello, World!

	// For building strings in a loop, use strings.Builder (more efficient than repeated +)
	var builder strings.Builder
	for i := 0; i < 3; i++ {
		builder.WriteString("Go! ")
	}
	fmt.Println(builder.String()) // Go! Go! Go!

	// ─── Common strings package functions ─────────────────────────────────────

	str := "  Hello, Go Programming!  "

	fmt.Println(strings.ToUpper(str))         // ALL UPPERCASE
	fmt.Println(strings.ToLower(str))         // all lowercase
	fmt.Println(strings.TrimSpace(str))       // remove leading/trailing spaces
	fmt.Println(strings.Contains(str, "Go"))  // true — does str contain "Go"?
	fmt.Println(strings.HasPrefix(str, "  Hello")) // true — starts with?
	fmt.Println(strings.HasSuffix(str, "!  "))     // true — ends with?
	fmt.Println(strings.Count(str, "o"))      // count occurrences of "o"
	fmt.Println(strings.Replace(str, "Go", "Golang", 1)) // replace first occurrence
	fmt.Println(strings.ReplaceAll(str, " ", "_"))        // replace ALL spaces

	// Split — breaks a string into a slice of substrings by a separator
	csv := "apple,banana,cherry"
	fruits := strings.Split(csv, ",")
	fmt.Println(fruits)        // [apple banana cherry]
	fmt.Println(fruits[0])     // apple
	fmt.Println(len(fruits))   // 3

	// Join — opposite of Split, joins a slice into a single string
	joined := strings.Join(fruits, " | ")
	fmt.Println(joined) // apple | banana | cherry

	// ─── String formatting with fmt.Sprintf ───────────────────────────────────

	// Sprintf returns a formatted string instead of printing it
	// Useful when you need to build a string to store or pass around
	name := "Alice"
	age := 30
	profile := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(profile) // Name: Alice, Age: 30

	// ─── Converting between string and []byte ─────────────────────────────────

	// Strings and byte slices are closely related.
	// []byte(s) converts a string to a byte slice — allows mutation
	// string(b) converts a byte slice back to a string
	original := "hello"
	bytes := []byte(original)   // convert to byte slice
	bytes[0] = 'H'              // modify the byte slice (strings are immutable, slices are not)
	modified := string(bytes)   // convert back to string
	fmt.Println(modified)       // Hello
}
