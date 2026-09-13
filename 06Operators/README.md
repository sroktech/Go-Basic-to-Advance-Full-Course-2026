# 06 – Operators

## What You Will Learn
- Arithmetic operators and Go's integer division/truncation behavior
- Why `++` and `--` are statements in Go, not expressions
- Relational (comparison) operators that always return a `bool`
- Logical operators for combining boolean values
- Assignment shorthand operators (`+=`, `-=`, etc.)
- Bitwise operators and what they do to the binary bits of a number
- The address-of (`&`) and dereference (`*`) operators and what a pointer is

## Why This Matters
Operators are the small symbols that do the actual work in almost every line of code you write: adding prices, comparing ages, checking multiple conditions before letting a user in, or toggling a feature flag. Go keeps its operator set small and predictable — no operator overloading, no surprising automatic type conversions — which means code that uses operators behaves the same way everywhere you see it. Understanding a few Go-specific quirks (like integer division) early prevents subtle bugs later.

## Concept Explanation
Go groups operators into families. **Arithmetic operators** (`+ - * / % ++ --`) do math but require both operands to be the same type — Go will not silently convert an `int` to a `float64` for you, unlike many other languages. **Relational operators** (`== != > < >= <=`) compare two values and always produce a `bool`, which is exactly what an `if` condition or loop condition expects. **Logical operators** (`&& || !`) combine or invert `bool` values — `&&` needs both sides true, `||` needs at least one side true, `!` flips a value. **Assignment operators** (`+= -= *= /= %=` and the bitwise equivalents) are shorthand for "take the current value, apply an operation, store it back" — `A += B` means exactly `A = A + B`. **Bitwise operators** (`& | ^ << >>`) work on the individual binary digits of an integer rather than the number as a whole — used for flags, permissions, and low-level encoding. Finally, `&` and `*` also serve a second purpose beyond arithmetic and bitwise use: `&` takes the memory address of a variable (producing a **pointer**), and `*` reads the value stored at an address a pointer holds.

## Simple Example
Every operator family in [main.go](main.go) runs and prints its own result. Arithmetic:

```go
A, B, C, D := 10, 20, 11, 3
fmt.Println("A / B =", A/B) // Division: 10 / 20 = 0 (integer division truncates)
fmt.Println("C % D =", C%D) // Modulus: 11 % 3 = 2
A++                          // Only postfix (A++) is valid in Go, not ++A
```

Pointers, at the end of the file:
```go
A := 10
ptr := &A                           // ptr holds the memory address of A
fmt.Println("Address of A:", ptr)   // e.g. 0xc000018090
fmt.Println("Value of *ptr:", *ptr) // 10 — the value stored at that address
```

## How It Works
- `A / B` with two integers performs **integer division**: it throws away any remainder rather than producing a decimal. `10 / 20` is `0`, not `0.5`.
- `A++` and `A--` are **statements**, not expressions — they don't return a value, so you cannot write `B = A++` or use them inside a larger expression. Go also only allows postfix (`A++`), never prefix (`++A`).
- `ptr := &A` creates a variable of type `*int` ("pointer to int") holding the address where `A` lives in memory.
- `*ptr` dereferences that pointer — it means "go to the address `ptr` holds and give me the value stored there." You could also write `*ptr = 99` to change `A`'s value indirectly, through the pointer.

## Common Mistakes
- Expecting `10 / 20` to give `0.5` — with integer operands, Go truncates. Convert to `float64` first if you need a fractional result.
- Writing `B = A++` or `if A++ > 5` — this doesn't compile because `++`/`--` are standalone statements in Go, unlike C or JavaScript.
- Mixing an `int` and a `float64` in the same arithmetic expression without an explicit conversion — Go requires both operands to match exactly.
- Comparing values of different types with `==` — Go requires the same (or compatible comparable) types on both sides.

## Best Practices
- Reach for assignment shorthand (`+=`, `-=`, etc.) instead of writing out the full expression — it's shorter and the conventional Go style.
- Use bitwise operators only when you actually need bit-level control (flags, masks); for ordinary math, arithmetic operators are clearer.
- Be deliberate about integer vs. floating-point division — convert explicitly with `float64(x)` when you need a precise, non-truncated result.
- Treat pointers with care at first: only use `&`/`*` when you specifically need to share or modify a value across function boundaries (covered more in later lessons on functions).

## Real-World Example
Integer division truncation matters constantly in real systems — e.g. splitting a total price evenly among N people, or computing "pages needed" for pagination (`totalItems / pageSize` silently drops a partial page unless you round up deliberately). Bitwise operators show up in permission systems (Unix file permissions, feature-flag bitmasks) where each bit represents an independent on/off setting. Pointers are the foundation of how Go functions can modify a caller's data or avoid copying large structures — a pattern you'll see everywhere once functions are introduced.

## Exercise
Declare two `int` variables and print the results of all six arithmetic operators on them (`+ - * / % ` plus one `++`). Then declare two `bool` variables and print the results of `&&`, `||`, and `!` on them. Pay attention to the integer division result.

## Mini Project
Write a simple "bill splitter": given a total bill amount (`int`, in cents) and a number of people, use integer division and the modulus operator (`%`) to compute how much each person owes and how many cents are left over that don't divide evenly, then print both values.

## Summary
Go's operators are grouped into arithmetic, relational, logical, assignment, bitwise, and address/pointer operators. Two Go-specific behaviors are worth remembering early: integer division truncates rather than producing a decimal, and `++`/`--` are statements you can't embed in expressions. The `&` and `*` operators additionally let you work directly with memory addresses through pointers.

## What to Learn Next
Previous: [05 Constants](../05Constants/README.md)

Now that you can compute and compare values, the next step is controlling which code runs based on those results.

Next: [07 Control Flow (if/else)](../07Control_Flow_if_else/README.md)
