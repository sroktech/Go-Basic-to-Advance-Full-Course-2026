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
	fmt.Println("=== 1. Classic for loop ===")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

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
	fmt.Println("\n=== 2. Nested loop — multiplication table ===")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("%d*%d=%-3d", i, j, i*j)
		}
		fmt.Println() // move to next line after each row
	}

	// ─── 3. Loop Control Statements ─────────────────────────────────────────────

	// A. break — exit the loop immediately, skip remaining iterations
	// When i reaches 5, 'break' stops the loop right away.
	// Numbers printed: 0 1 2 3 4  (5 is never printed because break happens first)
	fmt.Println("\n=== 3a. break ===")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // jumps OUT of the loop entirely
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// B. continue — skip the rest of THIS iteration, go to the next one
	// When i%2 == 0 (i is even), 'continue' skips fmt.Println and moves to i++.
	// Only odd numbers reach fmt.Println.
	// Numbers printed: 1 3 5 7 9
	fmt.Println("\n=== 3b. continue ===")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skip even numbers, jump to next iteration
		}
		fmt.Print(i, " ") // only reached when i is odd
	}
	fmt.Println()

	// C. goto — jump directly to a labeled line in the function
	// 'goto' is rarely used in modern Go code because it makes logic hard to follow.
	// Here it jumps to the label 'end:' when i == 5, stopping before i reaches 10.
	// Numbers printed: 0 1 2 3 4 5
	fmt.Println("\n=== 3c. goto (rare — shown for completeness) ===")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
		if i == 5 {
			goto end // jump past the loop to the 'end' label below
		}
	}
end:
	fmt.Println("\n  loop ended via goto")

	// ─── 4. While-style loop ─────────────────────────────────────────────────────
	// 'for condition { }' — omit init and post, and 'for' behaves like a while loop
	// in other languages. Runs as long as the condition is true.
	fmt.Println("\n=== 4. While-style loop ===")
	n := 1
	for n < 32 {
		fmt.Print(n, " ")
		n *= 2
	}
	fmt.Println()

	// ─── 5. Infinite loop, safely bounded with break ─────────────────────────────
	// 'for { }' with no condition runs forever — never stops on its own.
	// In real programs you'd put a 'break' or 'return' inside to eventually exit.
	// Common use: servers, game loops, background workers that run until shutdown.
	// A bare 'for {}' would run forever and never let this program finish, so here
	// a counter + break demonstrates the shape safely.
	fmt.Println("\n=== 5. Infinite loop (for {}), stopped with break ===")
	count := 0
	for {
		fmt.Print(count, " ")
		count++
		if count >= 5 {
			break // without this, the loop would never stop
		}
	}
	fmt.Println()

	// ─── 6. goto as a while-loop replacement (historical curiosity) ─────────────
	// goto can replace a while loop entirely — this is how loops worked in old
	// BASIC-style languages before 'while'/'for' existed:
	//   10 PRINT i
	//   20 GOTO 10
	// Modern Go code should prefer the 'for condition {}' form from section 4
	// instead — this is shown only so you can recognize the pattern if you ever
	// see it in older code.
	fmt.Println("\n=== 6. goto as a while-loop replacement (for reference only) ===")
	i := 0
start:
	if i < 5 {
		fmt.Print(i, " ")
		i++
		goto start // jump back to 'start' — behaves like: while i < 5
	}
	fmt.Println()
}
