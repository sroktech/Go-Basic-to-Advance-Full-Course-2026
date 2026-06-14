// Lesson5: Constants in Go programming language
// Constants are immutable values known at compile time and cannot be changed after declaration.
// Unlike variables, constants cannot use := and their value must be a fixed literal or constant expression.
// Constants can be declared at the package level (outside main) or inside functions.
// Go constants are "untyped" by default, which gives them more flexibility (explained below).

package main

import "fmt"

func main() {
	// Typed constants — the type is explicitly written (string, int, etc.).
	// A typed constant is bound to that specific type and cannot be used where another type is expected.
	// Convention: constants are often written in UPPER_SNAKE_CASE to distinguish them from variables.
	const NAME string = "John Doe"
	const PRICE int = 100

	fmt.Println("Name is : ", NAME)
	fmt.Println("Price is :", PRICE)

	// Integer literals — Go supports three numeral bases for integer constants.
	// All three below represent the same value (255) written in different bases:
	//   Decimal     — base 10, the normal everyday number system (no prefix)
	//   Octal       — base 8,  prefixed with 0  (digits 0–7 only)
	//   Hexadecimal — base 16, prefixed with 0x (digits 0–9 and a–f)
	// Grouped const block: multiple constants declared together with a single 'const' keyword.
	const (
		DECIMAL     = 255  // 255 in base 10
		OCTAL       = 0377 // 3*64 + 7*8 + 7 = 255 in base 8
		HEXADECIMAL = 0xff // 15*16 + 15 = 255 in base 16
	)
	// All three print as 255 — Go converts to decimal for display.
	fmt.Println("Decimal:", DECIMAL, "Octal:", OCTAL, "Hexadecimal:", HEXADECIMAL)

	// Floating-point literals
	// float64 is the default floating-point type in Go (higher precision than float32).
	// PI is a typed constant — explicitly declared as float64.
	const PI float64 = 3.141
	fmt.Println("PI value is : ", PI)

	// Scientific notation: 6.022e23 means 6.022 × 10²³ (Avogadro's number).
	// AVOGADRO is an untyped constant — Go will infer the type when it is used in an expression.
	// Untyped constants are more flexible because they can be used with float32, float64, etc.
	const AVOGADRO = 6.022e23
	fmt.Println("AVOGADRO value is : ", AVOGADRO)

	// String literals with escape sequences
	// Escape sequences are special character combinations starting with \ that represent
	// characters that are hard to type directly in source code.

	// \n — newline: moves the cursor to the next line
	const GREETING = "Hello, Earth!\n"

	// \" — escaped double quote: lets you include " inside a double-quoted string
	const QUOTE = "\"GO is simple!\" - A programmer!"

	fmt.Printf(GREETING)
	fmt.Println(QUOTE)

	// \a — alert/bell: triggers a beep sound on supported terminals (may be silent in modern terminals)
	const BELL = "Bell is \a"
	fmt.Println(BELL)

	// \n used multiple times to split a single string across multiple output lines
	const LB = "I\nAM\nBATMAN\n!"
	fmt.Println(LB)

	// Multi-line string literal using + concatenation at compile time.
	// Go does not allow a string literal to span multiple lines with a raw line break —
	// use + to join them. The + must be at the end of the line, not the beginning.
	// Note: there is no space before "span" — the strings join directly, so add one if needed.
	const MULTILINE = "The Eiffel tower is so long that it needs to" +
		"span multiple clouds for better\nphotographing in France!"

	// Concatenation of two string constants — joined at compile time into one string.
	// This is different from runtime string building; no performance cost at runtime.
	const CONCATENATED = "Concatenated " + "string"

	fmt.Println(MULTILINE)
	fmt.Println(CONCATENATED)

	// Boolean constants — only two possible values: true or false.
	// Useful for feature flags, configuration, or default states that should never change.
	const ACTIVE = true
	const READY = false
	fmt.Println("ACTIVE:", ACTIVE, " READY:", READY)

	// Constants in expressions — Go evaluates constant expressions at compile time, not at runtime.
	// AREA is computed from LENGTH * WIDTH before the program even runs.
	// This is more efficient than computing the same value repeatedly as a variable.
	const LENGTH = 50
	const WIDTH = 5
	const AREA = LENGTH * WIDTH // computed at compile time: 50 * 5 = 250
	fmt.Println("AREA:", AREA)
}
