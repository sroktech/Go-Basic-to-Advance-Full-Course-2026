# 08 – Switch

## What You Will Learn
- Why `switch` exists as a cleaner alternative to long `if/else if` chains
- Basic `switch` on a value, with multiple values per case
- Condition-less `switch` (acts like `if/else if`)
- `switch` with an initializer statement
- `fallthrough` and why Go doesn't fall through by default
- A first, light look at type switches

## Why This Matters
You already know `if`/`else if`/`else` from the last lesson. That works fine, but once you're comparing *one* value against many possible options — a day of the week, a status code, a menu choice — a long `else if` chain gets repetitive and harder to scan. `switch` is Go's cleaner way to write exactly that kind of branching. That's also why this lesson comes right after if/else and before loops: switch isn't a new, harder concept that needs loops first — it's simply a tidier form of the branching you just learned.

## Concept Explanation
A `switch` compares a value against a series of `case`s, top to bottom, and runs the block for the first match. Go's `switch` has a few differences from what you may have heard about switch statements in other languages:
- **No automatic fall-through.** Once a matching case's block finishes, Go exits the switch — it does *not* keep running the next case's code. This is a deliberate design choice: automatic fall-through (like in C) is a common source of bugs because forgetting a `break` silently runs code you didn't intend to. Go flips the default so the safe behavior requires no extra keyword.
- **Multiple values per case.** `case "Saturday", "Sunday":` matches either value.
- **A condition-less switch** (just `switch {`) lets each `case` be a full boolean expression, making it behave like an `if/else if` chain but often read more cleanly when there are many branches.
- **A switch with an initializer**, `switch hour := time.Now().Hour(); {`, runs a short statement first (like `if` can), scoping the new variable to the switch block only.
- **`fallthrough`** is the explicit escape hatch when you *do* want to run into the next case's block regardless of its own condition — it's rare, and usually a sign the logic could be restructured.

Go also lets a `switch` check the underlying *kind* of a value stored in a special box called `interface{}` (or `any`) — this is called a **type switch**. We'll explain that box fully in lesson 22 (Interfaces); for now just know a type switch can ask "what kind of value is this — an `int`? a `string`? something else?" and branch accordingly.

## Simple Example
From [main.go](main.go):

```go
day := "Monday"
switch day {
case "Monday":
	fmt.Println("Start of the work week — coffee time!")
case "Saturday", "Sunday": // multiple values, one case
	fmt.Println("Weekend — relax!")
default:
	fmt.Println("Mid-week grind:", day)
}
```

A condition-less switch, doing the same job an `if/else if` chain would:

```go
switch {
case score >= 90:
	fmt.Println("Grade: A")
case score >= 80:
	fmt.Println("Grade: B") // matches here
default:
	fmt.Println("Grade: F")
}
```

## How It Works
- In the first example, Go checks `day` against each `case` top to bottom. `"Monday"` matches the first case, so that block runs and the switch exits immediately — no `break` needed.
- `case "Saturday", "Sunday":` matches if `day` equals *either* value — a single case can bundle several options.
- The condition-less form (`switch {`) has no value after `switch`, so each `case` is its own boolean condition, evaluated top to bottom just like `if`/`else if` — the first `true` one wins.
- `switch hour := time.Now().Hour(); { ... }` declares `hour` right in the switch statement, keeping that variable's scope limited to the switch block, just like `if x := foo(); x > 0 { }` would.
- The `fallthrough` example forces case `1`'s block to also run case `2`'s block unconditionally — even though case `2`'s own condition (`n == 2`) isn't checked at all.

## Common Mistakes
- Expecting fall-through by habit from C/Java/JavaScript and being surprised each case exits automatically — in Go, no `break` is needed, and adding one is redundant (though harmless).
- Using `fallthrough` casually — it skips the next case's condition entirely, which can produce confusing bugs if that case was meant to be checked independently.
- Forgetting that case order matters in a condition-less switch — like `else if`, the first matching (true) case wins, so broader conditions placed too early can shadow more specific ones below them.

## Best Practices
- Reach for `switch` once you're comparing one value against three or more known options — it reads more clearly than a long `else if` chain.
- Keep using `if`/`else if` when your branches involve complex boolean logic across different, unrelated variables — that's still `if`'s job.
- Avoid `fallthrough` unless you have a specific, well-understood reason; prefer combining values in one `case` (`case "Saturday", "Sunday":`) instead.

## Real-World Example
Switch statements show up constantly in real code: mapping an HTTP status code to a user-facing message, handling different commands in a CLI tool, or converting a numeric score into a letter grade (as in this lesson and the last one). Type switches specifically are common in Go programs that accept loosely-typed input — like a JSON parser or a CLI argument parser — where a value's shape isn't known until you inspect it at runtime.

## Exercise
Write a `switch` that takes a `month` variable (an `int` from 1–12) and prints which season it falls in ("Winter", "Spring", "Summer", "Fall"), grouping multiple month numbers into each case using commas.

## Mini Project
Build a simple grade calculator combining this lesson and the last one: take a `score` variable, use a condition-less `switch` to print a letter grade (A/B/C/D/F), and inside the case for a borderline grade (e.g. `case score >= 80:`), nest an `if`/`else` to further split it into a "+" or plain grade — same idea as the nested-if example from lesson 07, rewritten with switch.

## Summary
`switch` is a cleaner way to write many of the same branching decisions `if`/`else if` handles — no automatic fall-through, support for multiple values per case, an optional initializer statement, and a condition-less form that behaves like `if`/`else if`. `fallthrough` exists for the rare case you want to force the next block to run anyway. Type switches let you ask "what kind of value is this?" for values held in an `interface{}`/`any` — a topic covered fully in lesson 22.

## What to Learn Next
Previous: [07 Control Flow (if/else)](../07Control_Flow_if_else/README.md)

Next you'll learn how to repeat code with Go's single loop keyword, `for`.

Next: [09 Control Flow (Loops)](../09Control_Flow_Loops/README.md)
