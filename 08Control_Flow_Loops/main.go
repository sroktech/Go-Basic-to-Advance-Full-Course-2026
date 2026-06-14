// Control Flow: Loops
//
// Go has only ONE loop keyword: 'for'.
// Unlike many languages (while, do-while, foreach), Go uses 'for' for all loop styles.
// This keeps the language simple — you just write 'for' differently depending on what you need.
//
// Loop styles in Go:
//   1. Classic for loop   — for init; condition; post { }
//   2. While-style loop   — for condition { }         (omit init and post)
//   3. Infinite loop      — for { }                   (omit everything)
//   4. Nested loop        — a loop inside another loop
//
// Loop control statements:
//   break    — immediately exits the loop
//   continue — skips the rest of the current iteration, jumps to next
//   goto     — jumps to a labeled line anywhere in the function (use sparingly)

package main

import "fmt"

func main() {
	// ─── 1. Classic for loop ─────────────────────────────────────────────────────
	// Three parts separated by semicolons:
	//   init:      i := 0   — runs ONCE before the loop starts, sets up the counter
	//   condition: i < 10   — checked BEFORE each iteration; loop stops when false
	//   post:      i++      — runs AFTER each iteration; updates the counter
	//
	// Execution order: init → [condition → body → post] → [condition → body → post] → ...
	// Prints: 0 1 2 3 4 5 6 7 8 9  (stops before 10 because condition i < 10 becomes false)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// ─── 2. Nested Loop — Multiplication Table ───────────────────────────────────
	// A loop inside another loop. The inner loop runs COMPLETELY for every single
	// step of the outer loop.
	//
	// Outer loop controls the ROW (i goes 1 → 5)
	// Inner loop controls the COLUMN (j goes 1 → 5) for each row
	//
	// Total iterations = 5 * 5 = 25
	// %d = integer placeholder, %d * %d = %d prints "i * j = result"
	// \t adds a tab for spacing between columns
	// fmt.Println() after the inner loop ends the row and moves to the next line
	//
	// Output (first two rows):
	//   1 * 1 = 1   1 * 2 = 2   1 * 3 = 3   1 * 4 = 4   1 * 5 = 5
	//   2 * 1 = 2   2 * 2 = 4   ...
	// for i := 1; i <= 5; i++ {
	// 	for j := 1; j <= 5; j++ {
	// 		fmt.Printf("%d * %d = %d\t", i, j, i*j)
	// 	}
	// 	fmt.Println() // move to next line after each row
	// }

	// ─── 3. Loop Control Statements ─────────────────────────────────────────────

	// A. break — exit the loop immediately, skip remaining iterations
	// When i reaches 5, 'break' stops the loop right away.
	// Numbers printed: 0 1 2 3 4  (5 is never printed because break happens first)
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break // jumps OUT of the loop entirely
	// 	}
	// 	fmt.Println(i)
	// }

	// B. continue — skip the rest of THIS iteration, go to the next one
	// When i%2 == 0 (i is even), 'continue' skips fmt.Println and moves to i++.
	// Only odd numbers reach fmt.Println.
	// Numbers printed: 1 3 5 7 9
	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue // skip even numbers, jump to next iteration
	// 	}
	// 	fmt.Println(i) // only reached when i is odd
	// }

	// C. goto — jump directly to a labeled line in the function
	// 'goto' is rarely used in modern Go code because it makes logic hard to follow.
	// Here it jumps to the label 'end:' when i == 5, stopping before i reaches 10.
	// Numbers printed: 0 1 2 3 4 5
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// 	if i == 5 {
	// 		goto end // jump past the loop to the 'end' label below
	// 	}
	// }
	// end:
	// 	fmt.Println("Loop ended")

	// goto can also replace a while loop (not recommended — shown for learning only):
	// i := 0
	// start:
	// 	if i < 10 {
	// 		fmt.Println(i)
	// 		i++
	// 		goto start // jump back to 'start' — behaves like: while i < 10
	// 	}

	// ─── 4. Infinite Loop ────────────────────────────────────────────────────────
	// 'for { }' with no condition runs forever — never stops on its own.
	// In real programs you'd put a 'break' or 'return' inside to eventually exit.
	// Common use: servers, game loops, background workers that run until shutdown.
	// for {
	// 	fmt.Println("Infinite Loop")
	// }

	// ─── goto infinite loop (classic BASIC style) ─────────────────────────────
	// This is the active code in this file — it loops forever using goto.
	// 'start:' is a label — a named point in the code you can jump to.
	// After printing, 'goto start' jumps back up to the label, repeating endlessly.
	// This is equivalent to: for { fmt.Println("This is an infinite loop!") }
	// WARNING: this will run forever — stop it with Ctrl+C in the terminal.
start:
	fmt.Println("This is an infinite loop!")
	goto start
}

// Historical note: this pattern comes from early BASIC programming (1960s–80s):
//   10 PRINT "This is an infinite loop!"
//   20 GOTO 10
// Line numbers were how BASIC identified where to jump. 'goto' in Go is the modern equivalent.
