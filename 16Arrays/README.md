# 16 – Arrays

## What You Will Learn
- How to declare, initialize, and index a fixed-size array
- Why an array's length is part of its type
- Why arrays are value types, and what that means when you assign or copy one
- How to iterate an array with a classic index loop and with `for range`
- How to work with multi-dimensional arrays

## Why This Matters
Arrays are Go's most basic collection type, and understanding their behavior — especially that
they are copied by value — is the foundation for understanding why Go programmers almost always
reach for **slices** instead once they need a resizable or shareable collection. You need to see
the limitation firsthand before the next lesson's solution makes sense.

## Concept Explanation
An array is a **fixed-size, ordered collection of elements of the same type**. "Fixed-size"
means you decide the length when you declare the array, and it can never grow or shrink — the
length is baked into the type itself. `[3]int` and `[4]int` are two completely different types
in Go, even though both hold integers; you cannot assign one to the other or pass a `[3]int`
where a function expects a `[4]int`.

The other key property: **arrays are value types**. In Go, assigning an array to a new variable,
or passing it to a function, copies every single element. The new variable is a fully
independent array — changing it has no effect on the original. This is different from many
other languages where array-like structures are passed around by reference.

This combination — fixed length baked into the type, plus full-copy-on-assignment — is exactly
why arrays are used sparingly in everyday Go code. Needing to know a collection's exact size at
compile time, and paying a copy cost every time you pass it around, is rarely what you want for
things like a list of user records or lines read from a file. The next lesson introduces
**slices**, which wrap an array internally but add flexible length and cheap sharing — for now,
just know they exist as the practical alternative you'll reach for after this lesson.

## Simple Example
From [main.go](main.go):
```go
original := [3]int{1, 2, 3}
copyArr := original // full copy of all 3 elements
copyArr[0] = 999     // modify the copy

fmt.Println("Original:", original) // [1 2 3] — unchanged
fmt.Println("Copy:", copyArr)      // [999 2 3]
```

## How It Works
`copyArr := original` does not make `copyArr` point at the same data as `original` — it
allocates a brand-new `[3]int` and copies all three values into it. That's why modifying
`copyArr[0]` leaves `original` completely untouched: they are two separate arrays in memory
that happened to start out equal. The same copying happens when an array is passed as a
function argument — the function receives its own copy, not access to the caller's array.

## Common Mistakes
- Forgetting that arrays are fixed-size value types, and being surprised when copying an array
  (via assignment or a function call) doesn't share changes back to the original. If you need
  shared, resizable data, that's precisely the problem slices (next lesson) solve.
- Trying to use two arrays of different lengths interchangeably — `[3]int` and `[4]int` are
  different types, so this is a compile error, not a runtime surprise.
- Forgetting that a plain `var arr [5]int` starts filled with zero values, not an empty/undefined
  array — there is no "empty" array state in Go.

## Best Practices
- Use arrays when the size is genuinely fixed and known at compile time (e.g., a 3x3 tic-tac-toe
  board, RGB color components, a fixed lookup table).
- Prefer `[...]T{...}` when initializing with a literal so Go counts the length for you instead
  of you keeping the count and the values in sync manually.
- If you find yourself wanting to add or remove elements, or pass a variable-length collection
  around without copying, that's your signal to use a slice instead (covered next).

## Real-World Example
Fixed-size arrays show up in domains where the size truly never changes: a chessboard (`[8][8]`
grid), the days of the week, a cryptographic hash's fixed byte output, or a small buffer of
sensor readings collected at a known interval. Most everyday "list of things" problems (users,
orders, log lines) use slices instead, precisely because their length isn't known in advance.

## Exercise
Declare a `[7]string` array holding the names of the days of the week. Write a loop that prints
each day with its index, then write a second loop that counts how many days start with the
letter "S" (Saturday, Sunday).

## Mini Project
Build a simple 3x3 tic-tac-toe board using a `[3][3]string` array, initialized with empty
strings. Write code that places `"X"` and `"O"` at a few chosen positions and then prints the
board row by row using `for range`, similar to the `matrix` example in `main.go`.

## Summary
Arrays have a fixed length that is part of their type, and assigning or passing one always
copies every element — they are value types, not references. This makes them predictable but
inflexible: great for genuinely fixed-size data, awkward for collections that grow, shrink, or
need to be shared efficiently. That inflexibility is the exact motivation for slices, which you
will learn next.

## What to Learn Next
Previous: [15Strings](../15Strings/README.md)
Next: [17Pointers](../17Pointers/README.md)
