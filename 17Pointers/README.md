# 17 – Pointers

## What You Will Learn
- What a pointer is and how it relates to a variable's memory address
- The `&` (address-of) and `*` (dereference) operators
- How to allocate a pointer with `new()`
- What a `nil` pointer is and why dereferencing one panics
- How to pass a pointer to a function so the function can modify the caller's variable
- How pointers to structs work, including Go's automatic dereferencing for field access

## Why This Matters
So far, every value you've worked with (ints, strings, arrays, structs) has been passed around
by copying. That's simple and safe, but it has two costs: a function can't modify the caller's
original variable, and copying large data repeatedly wastes memory and time. Pointers solve both
problems, and they appear constantly in real Go code — especially once you start writing methods
on structs (next lesson) and, later, working with slices and maps.

## Concept Explanation
Think of a running program's memory as a long row of numbered boxes, each holding a value. Every
box has an address — like a house number. A normal variable is a box that holds a *value*. A
**pointer** is a variable that holds an *address* — it tells you which box to go look in, rather
than holding the value itself.

Two operators make this work:
- `&x` — "address-of": gives you the memory address where `x` is stored.
- `*p` — "dereference": follows a pointer `p` to read or write the value at that address.

A pointer's type reflects what it points to: `*int` is "a pointer to an int," `*string` is "a
pointer to a string," and so on.

Go chose **explicit pointers** instead of making everything reference-like (as some languages
do) so that, just by reading a function's signature, you can tell whether it might modify your
variable. A function taking `int` cannot change your original value; a function taking `*int`
can — a deliberate design choice for clarity and predictability.

## Simple Example
From [main.go](main.go):
```go
a := 10
fmt.Println("Before double:", a) // 10
double(&a)                       // pass the address of a
fmt.Println("After double:", a)  // 20 — original was changed
```
```go
// double takes a pointer to an int and multiplies the value by 2
func double(n *int) {
    *n = *n * 2
}
```

## How It Works
`&a` produces the memory address of `a` and passes it into `double`. Inside `double`, the
parameter `n` is a `*int` holding that same address — not a copy of `a`'s value. The line
`*n = *n * 2` dereferences `n` to read the current value at that address, doubles it, and writes
the result back to that same address. Because `n` points at `a`'s actual memory location, the
change is visible in `main` after the call returns — this is "call by reference" achieved
explicitly through a pointer, rather than being every function call's default behavior.

The same idea applies to structs: given `ptPtr := &pt`, Go lets you write `ptPtr.X` instead of
the more verbose `(*ptPtr).X` — the compiler auto-dereferences for field access, but the effect
is the same: you're reading and writing the original struct's memory.

## Common Mistakes
- Dereferencing a `nil` pointer (`*p` when `p` is `nil`) — this causes a runtime panic. Always
  check `if p != nil` before dereferencing a pointer that might not have been assigned.
- Confusing `&` and `*`: `&x` takes a value and gives you its address (going value → pointer);
  `*p` takes a pointer and gives you the value at that address (going pointer → value). They are
  opposites.
- Overusing pointers for small values (like a single `int` or `bool`) where passing by value
  would be simpler, just as safe, and avoids the risk of unintended mutation.

## Best Practices
- Use a pointer when a function genuinely needs to modify the caller's variable, or when the
  data being passed is large enough that copying it would be wasteful.
- Prefer plain values for small, simple data (ints, small structs used read-only) — pointers add
  a layer of indirection that isn't free to reason about.
- Always treat an uninitialized pointer as potentially `nil` and guard against dereferencing it.

## Real-World Example
Pointers are how Go avoids copying large data structures on every function call — for example,
passing a pointer to a large struct (like a database record or a big configuration object)
instead of copying it every time it's passed to a function. They're also essential whenever a
function needs to update something in place, such as a method that mutates a struct's fields
(you'll use this constantly starting in the next lesson on structs).

## Exercise
Write a function `increment(n *int)` that adds 1 to the int at the given address. Call it three
times in a row on the same variable and print the value after each call to confirm it keeps
increasing.

## Mini Project
Write a function `swap(a, b *int)` that swaps the values of two integers using pointers (no
extra return values needed — modify both through their pointers). Demonstrate it in `main` with
two variables, printing their values before and after the swap.

## Summary
A pointer stores a memory address rather than a value; `&` gets an address, `*` follows it to
read or write the value there. Pointers let functions modify a caller's original variable and let
large data be shared without copying — but a `nil` pointer must never be dereferenced. Struct
field access through a pointer is automatically dereferenced by Go, which is why pointers to
structs (`&SomeStruct{...}`) are so common in idiomatic Go code, as you'll see extensively in the
next lesson.

## What to Learn Next
Previous: [16Arrays](../16Arrays/README.md)
Next: [18Structures](../18Structures/README.md)
