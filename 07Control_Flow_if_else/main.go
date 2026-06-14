// Control Flow: Conditionals (if / else if / else)
//
// Control flow determines the ORDER in which code is executed.
// By default Go runs code top to bottom — conditionals let you BRANCH:
// run different code depending on whether a condition is true or false.
//
// Structure:
//   if <condition> {
//       // runs when condition is true
//   } else if <another condition> {
//       // runs when the first is false but this one is true
//   } else {
//       // runs when ALL conditions above are false
//   }
//
// Rules in Go:
//   - The condition must be a bool expression (no implicit "truthy" like in JS/Python)
//   - Curly braces { } are always required, even for a single line
//   - The opening { must be on the SAME line as the if/else — not the next line

package main

import "fmt"

func main() {
	// ─── 1. if statement ────────────────────────────────────────────────────────
	// The simplest form — runs the block only when the condition is true.
	// If the condition is false, the block is skipped entirely.
	// age := 18
	// if age >= 18 {
	// 	fmt.Println("You are eligible to vote") // only prints when age >= 18
	// }

	// ─── 2. if / else ───────────────────────────────────────────────────────────
	// Adds a fallback block that runs when the condition is false.
	// Exactly ONE of the two blocks will always execute — never both, never neither.
	// age := 18
	// if age >= 18 {
	// 	fmt.Println("You are eligible to vote")   // runs when age >= 18
	// } else {
	// 	fmt.Println("You are not eligible to vote") // runs when age < 18
	// }

	// ─── 3. if / else if / else (chained conditions) ────────────────────────────
	// Use 'else if' to check multiple conditions in sequence.
	// Go tests each condition TOP TO BOTTOM and stops at the FIRST true one.
	// The final 'else' is a catch-all — it runs only if every condition above was false.
	//
	// Flow for score = 95:
	//   score >= 90? → true  ✓  → prints "Grade:A", skips all else if / else blocks
	//
	// Flow for score = 82:
	//   score >= 90? → false
	//   score >= 80? → true  ✓  → enters else if block
	//     score >= 85? → false  → prints "Grade:B"
	//
	// Flow for score = 60:
	//   score >= 90? → false
	//   score >= 80? → false
	//   else         → prints "Grade:C"
	score := 95
	if score >= 90 {
		// Reached only when score is 90 or above
		fmt.Println("Grade:A")
	} else if score >= 80 {
		// Reached only when score is 80–89 (because >= 90 was already false)
		// Nested if inside else if — adds a finer split within the 80–89 range
		if score >= 85 {
			fmt.Println("Grade:B+") // 85–89
		} else {
			fmt.Println("Grade:B") // 80–84
		}
	} else {
		// Catch-all — reached only when score is below 80
		fmt.Println("Grade:C")
	}
}
