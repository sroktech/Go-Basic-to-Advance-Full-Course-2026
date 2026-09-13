# 07 – Control Flow (if / else)

## What You Will Learn
- What "control flow" means and why it matters
- The `if`, `if/else`, and `if/else if/else` forms in Go
- Go's specific syntax rules for conditionals
- How to nest an `if` inside an `else if` branch for finer-grained checks

## Why This Matters
So far your programs have run top to bottom, one line after another, with no decisions. Real programs need to react differently depending on the situation: is the user old enough, did the request succeed, which grade band does a score fall into? Conditionals are how a program branches — they let you say "do this, but only if that's true," which is the foundation almost every piece of program logic is built on.

## Concept Explanation
An `if` statement runs a block of code only when its condition evaluates to `true`. Adding `else` gives you a fallback block that runs when the condition is `false` — exactly one of the two blocks always runs. Chaining `else if` lets you test several conditions in order, top to bottom, stopping at the first one that's `true`; a final `else` acts as a catch-all if none of them matched.

Go enforces a few rules that differ from many other languages:
- The condition must already be a `bool` expression. There's no "truthy" value like `0` or an empty string being treated as false — you must write an explicit comparison (e.g. `age >= 18`), never just `if age`.
- Curly braces `{ }` are mandatory, even for a single-line body — you can never write `if x > 0 fmt.Println(x)` without braces.
- The opening brace `{` must be on the *same line* as `if`/`else if`/`else`, not on the next line — Go's automatic semicolon insertion would otherwise misinterpret a brace placed on its own line as ending the statement.

## Simple Example
From [main.go](main.go):

```go
score := 95
if score >= 90 {
	fmt.Println("Grade:A")
} else if score >= 80 {
	if score >= 85 {
		fmt.Println("Grade:B+") // 85–89
	} else {
		fmt.Println("Grade:B") // 80–84
	}
} else {
	fmt.Println("Grade:C")
}
```

## How It Works
- Go checks `score >= 90` first. For `score = 95`, this is `true`, so it prints `"Grade:A"` and skips every other branch entirely — it never even looks at the `else if`.
- If that first condition were `false`, Go would check `score >= 80` next. Reaching this branch already tells us `score < 90`, so a nested `if score >= 85` inside it splits the 80–89 range further into `"B+"` (85–89) and `"B"` (80–84).
- If both `score >= 90` and `score >= 80` are `false`, the final `else` catches everything else and prints `"Grade:C"`.
- Only one printed line ever results from this whole structure — Go stops as soon as it finds a matching branch.

## Common Mistakes
- Writing `if (score >= 90)` with parentheses out of habit from C/Java/JavaScript — Go doesn't need or want them around the condition (`if score >= 90` is correct and idiomatic).
- Trying to use a non-boolean value as a condition, like `if score` — Go requires an explicit boolean expression such as `if score != 0`.
- Placing the opening `{` on its own line — Go's formatting rules require it on the same line as `if`/`else`.
- Reaching for a ternary operator (`condition ? a : b`) — Go intentionally has none. This is a deliberate design choice: the language's authors felt ternaries often get abused to cram complex logic into unreadable one-liners, so Go asks you to write a full `if/else` instead, keeping branching logic explicit and readable.

## Best Practices
- Order `else if` conditions from most specific/restrictive to least, especially when ranges overlap (like the grade example above), so the first true one is the one you actually intend.
- Keep condition expressions simple and readable; if a condition needs several `&&`/`||` combined, consider naming it with a `bool` variable first for clarity.
- Avoid deeply nested `if` chains where possible — as you'll see in the next lesson, `switch` is often a cleaner way to express "match one value against several options."

## Real-World Example
Conditional logic like this appears anywhere software makes a decision based on data: an e-commerce app deciding shipping cost tiers based on order total, an authentication system checking permission levels, or a grading system (exactly like this example) that converts a numeric score into a letter grade. Nearly every business rule ("if the customer is a premium member, apply a discount") is expressed with conditionals like these.

## Exercise
Write a program that assigns a `temperature` variable (an `int`, in Celsius) and prints `"Freezing"` if it's at or below 0, `"Cold"` if it's between 1 and 15, `"Mild"` if it's between 16 and 25, and `"Hot"` if it's above 25 — using only `if`/`else if`/`else`.

## Mini Project
Build a simple "ticket price calculator": given an `age` variable, print `"Child price: $5"` for ages under 12, `"Senior price: $8"` for ages 65 and up, and `"Adult price: $12"` for everyone else. Then extend it with a nested `if` that gives a further discount if the person is also a `student` (a `bool` variable) in the adult price band.

## Summary
Control flow lets a program take different paths depending on conditions. Go's `if`/`else if`/`else` requires a genuine boolean condition, mandatory braces, and a same-line opening brace — and deliberately has no ternary operator, favoring explicit, readable branches instead. Chained conditions are checked top to bottom, and only the first matching branch runs.

## What to Learn Next
Previous: [06 Operators](../06Operators/README.md)

You now know how to branch with if/else. Next you'll learn `switch`, which is often a cleaner way to write the same kind of branching when you're comparing one value against several possibilities.

Next: [08 Switch](../08Switch/README.md)
