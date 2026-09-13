# 09 – Control Flow (Loops)

## What You Will Learn
- Why Go has only one loop keyword, `for`, instead of `while`/`do-while`/`foreach`
- The classic three-part `for` loop, the while-style `for`, and the infinite `for`
- Nested loops
- Loop control statements: `break`, `continue`, and `goto`
- A well-known gotcha with loop variables and closures (and how Go 1.22+ changed it)

## Why This Matters
Almost everything a computer does usefully involves repeating something: processing every item in a list, retrying a failed request, running a game loop, or printing a table. Instead of giving you several different loop keywords for different situations (like many languages do with `while`, `do-while`, and `for`), Go deliberately gives you just one — `for` — and lets you shape it differently depending on what you need. Fewer keywords means less to memorize and one consistent mental model for every kind of repetition.

## Concept Explanation
Go's `for` loop has three optional parts, separated by semicolons: `for init; condition; post { }`. The **init** statement runs once before the loop starts (usually setting up a counter), the **condition** is checked before every iteration and stops the loop once it's `false`, and the **post** statement runs after every iteration (usually incrementing the counter).

Because all three parts are optional, the same `for` keyword covers every loop style:
- Leaving out `init` and `post` and keeping only the condition — `for condition { }` — gives you what other languages call a `while` loop.
- Leaving out all three — `for { }` — gives you an infinite loop, which only stops when something inside it (like `break` or `return`) says to stop.

Inside a loop body, two statements change its flow: `break` exits the loop immediately, and `continue` skips the rest of the current iteration and jumps straight to the next one. Go also has `goto`, which jumps to a labeled line anywhere in the function — it's rarely used in modern code because it can make control flow hard to follow, but it's worth understanding since it's the mechanism behind the active example in this lesson's file.

## Simple Example
Every loop style in [main.go](main.go) actually runs, so you can see its real output. The classic three-part form:

```go
for i := 0; i < 10; i++ {
    fmt.Print(i, " ")
}
// 0 1 2 3 4 5 6 7 8 9
```

An infinite `for {}`, made safe with a `break`:

```go
count := 0
for {
    fmt.Print(count, " ")
    count++
    if count >= 5 {
        break // without this, the loop would never stop
    }
}
```

## How It Works
- In the classic form, execution order is: `init` runs once → `condition` is checked → if true, the body runs → `post` runs → `condition` is checked again → repeat until `condition` is `false`.
- The while-style loop (`for n < 32 { ... }`) has no `init`/`post` at all — it just repeats while its condition holds, doubling `n` each time until it's no longer under 32.
- The bare `for { }` in section 5 has no condition, so it would run forever on its own; the `if count >= 5 { break }` inside it is what makes it stop after 5 iterations. This is the standard shape for things that are conceptually endless (servers, workers) but still need a real exit path.
- `start:` in section 6 is a **label** — a named point in the function. `goto start` jumps execution back to that exact line, creating a loop without using the `for` keyword at all — this is shown only as a historical curiosity; real Go code uses the while-style `for` from section 4 instead.
- In the `break` example, when `i == 5`, the loop exits immediately — the numbers 5 through 9 never print.
- In the `continue` example, when `i` is even, the rest of that iteration is skipped — only odd numbers reach `fmt.Print`.

## Common Mistakes
- Writing `for (i := 0; i < 10; i++)` with parentheses out of old habit — Go's `for` never uses them.
- Forgetting a `for { }` infinite loop needs its own `break`/`return` somewhere inside, or the program will genuinely never stop — this is exactly why section 5 in this lesson's code pairs its `for {}` with a `break` instead of leaving it truly unbounded.
- The classic **loop-variable-capture gotcha**: in older Go versions, a loop variable (like `i` in `for i := 0; i < 3; i++`) was a single shared variable across all iterations, so a closure or goroutine created inside the loop and run later would often see only the *final* value of `i`, not the value it had "at that iteration." **Go 1.22 changed this** — each iteration now gets its own fresh copy of the loop variable, so this gotcha no longer bites by default. If you use an older Go version, or read older code, it's worth knowing this history.
- Relying on `goto` for everyday looping — it works, but a plain `for` is almost always clearer; `goto` is best reserved for rare, specific cases.

## Best Practices
- Prefer the plain `for condition { }` (while-style) or classic three-part `for` for almost all everyday loops — reserve `goto` for the rare case nothing else expresses cleanly.
- Use `continue` to skip uninteresting cases early rather than wrapping the rest of the loop body in a big `if`.
- Use `break` to stop as soon as you've found what you're looking for, instead of letting a loop run needlessly to completion.
- Always give an infinite `for { }` a clear exit condition (a `break`, `return`, or an external signal) unless it's truly meant to run forever, like a server's main loop.

## Real-World Example
Loops are everywhere in real software: a web server's main loop accepting connections forever until shutdown, a retry loop that keeps attempting a network request with `continue`/`break` controlling when to give up, or a nested loop processing rows and columns of a spreadsheet or image. The historical BASIC-style `10 PRINT ... / 20 GOTO 10` pattern noted in this lesson's file is the direct ancestor of the modern infinite `for { }` loop you'll use far more often in practice.

## Exercise
Write a classic `for` loop that prints the numbers from 1 to 20, but uses `continue` to skip any number divisible by 3, and `break` to stop entirely once the loop reaches 15.

## Mini Project
Build a simple number-guessing loop: declare a fixed `secret := 7` and a slice of guesses like `guesses := []int{3, 9, 7, 5}`. Loop over the guesses with a classic `for i := 0; i < len(guesses); i++`, and for each one print whether it's "Too low", "Too high", or "Correct!" — using `break` to stop the loop as soon as the correct guess is found.

## Summary
Go uses a single `for` keyword for every kind of loop: classic counted loops, while-style loops, and infinite loops, depending on which of its three optional parts you include. `break` exits a loop early, `continue` skips to the next iteration, and `goto` jumps to a label but is rarely needed. A well-known historical gotcha — loop variables being shared across iterations when captured by a closure — was fixed in Go 1.22, which now gives each iteration its own copy.

## What to Learn Next
Previous: [08 Switch](../08Switch/README.md)

Next you'll learn how to package code into reusable, named blocks: functions.

Next: [10 Functions](../10Functions/README.md)
