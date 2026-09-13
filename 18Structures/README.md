# 18 – Structures

## What You Will Learn
- How to define a `struct` type and create instances of it
- How to attach methods to a struct with value vs. pointer receivers
- How structs behave as value types (copies) and when to use a pointer instead
- Anonymous structs, for quick one-off data shapes
- The difference between **composition** (a named struct field) and **embedding** (an anonymous
  struct field) — and why this distinction matters

## Why This Matters
Structs are Go's way of modeling real-world "things" that have multiple properties — a person, a
product, an employee. Go deliberately has no classes and no inheritance, so structs plus
composition/embedding are how Go programmers build up complex types from simpler ones. This is a
foundational pattern you'll use in almost every non-trivial Go program from here on.

## Concept Explanation
A **struct** groups fields of different types under one name, letting you define your own custom
type. Unlike a class in Java or Python, a Go struct has no constructors and no inheritance
hierarchy — it's just a data grouping. Behavior is added separately via **methods**: a function
with a special *receiver* parameter tying it to a struct type, e.g.
`func (r Rectangle) Area() float64 { ... }`.

Go gives you a choice of receiver:
- A **value receiver** (`func (r Rectangle) Area()`) works on a *copy* of the struct — safe, but
  any modification inside the method is lost when the method returns.
- A **pointer receiver** (`func (r *Rectangle) Scale(factor float64)`) works on the *original*
  struct — use it whenever the method needs to modify the struct, or when the struct is large
  enough that copying it on every call would be wasteful.

Because Go has no classes, it has no "extends" keyword either. Instead it offers two related but
distinct ways to build one struct out of another:

**Composition** means one struct has a *named field* whose type is another struct — the outer
struct simply *contains* the inner one, reached via the field name (`emp.Address.City`).

**Embedding** means the inner struct is listed as an *anonymous field* — just the type name, no
field name. Go still names that field after the type (`Address`) behind the scenes, but it also
**promotes** the embedded struct's fields (and methods) up to the outer struct, so `mgr.City`
works directly (both spellings are valid). This promotion is Go's substitute for inheritance: a
`Manager` gets `Address`'s fields "for free" with no parent/child relationship — still one
struct containing another, just with sugar for reaching through it.

| | Composition (named field) | Embedding (anonymous field) |
|---|---|---|
| Declared as | `Address Address` | `Address` (type only, no field name) |
| Access | `emp.Address.City` (required) | `mgr.City` (promoted) or `mgr.Address.City` (still valid) |
| Relationship | "has-a", explicit | "has-a", with field/method promotion |
| Use when | You want a clearly named sub-object | You want the inner type's members to feel like part of the outer type |

## Simple Example
From [main.go](main.go):
```go
// Composition — a named field
type Employee struct {
    Name    string
    Age     int
    Address Address // a named field whose type happens to be a struct
}
fmt.Println("Employee:", emp.Name, "lives in", emp.Address.City)

// Embedding — an anonymous field
type Manager struct {
    Name string
    Age  int
    Address // embedded — no field name, just the type
}
fmt.Println("Manager:", mgr.Name, "lives in", mgr.City) // promoted field
```

## How It Works
`Employee` declares `Address Address` — a field named `Address` of type `Address`. Nothing is
special about the name matching the type here; you must always write `emp.Address.City` to
reach the city.

`Manager` instead writes just `Address` with no field name — the anonymous-field syntax that
triggers embedding. Go still stores it under the field name `Address`, but the compiler also
lets you skip straight to `mgr.City` because `City` is promoted from the embedded `Address`. If
`Address` had methods, those would be promoted too, letting `Manager` "inherit" behavior without
any class hierarchy — this is why embedding is Go's idiomatic alternative to inheritance.

## Common Mistakes
- Choosing a value receiver when the method needs to modify the struct — the change silently
  vanishes because the method only touched a copy.
- Confusing composition with embedding: writing `Address Address` and expecting `mgr.City` to
  work — it won't, since that's a named field, not an embedded one. Only anonymous fields get
  promotion.
- Assuming two structs can always be compared with `==`. Every field must itself be comparable
  (numbers, strings, bools, arrays, other comparable structs); a struct containing a slice or
  map cannot be compared with `==` and causes a compile error.

## Best Practices
- Use a pointer receiver consistently across all of a type's methods once any one needs to
  mutate the struct, to stay consistent and avoid accidental copies.
- Reach for composition when the relationship should read explicitly (`emp.Address.City`); reach
  for embedding when the outer type should feel like a natural extension of the inner one.
- Keep struct literals using named fields (`Person{Name: "Alice", ...}`) rather than positional
  values, so adding or reordering fields later doesn't silently break call sites.

## Real-World Example
Structs plus embedding model domain entities without classical inheritance — e.g., a `BaseEvent`
struct (`ID`, `Timestamp`) embedded into `UserCreatedEvent` or `OrderPlacedEvent` so every event
gets the shared fields promoted, while composition (a named field) fits a clear "owns/contains"
relationship, like an `Order` with a named `ShippingAddress Address` field.

## Exercise
Define a `Book` struct with `Title`, `Author`, and `Pages` fields. Write a value-receiver method
`Summary()` that returns a formatted string describing the book, and a pointer-receiver method
`AddPages(n int)` that increases `Pages` by `n`. Create a `Book`, call `AddPages`, then print its
`Summary()`.

## Mini Project
Model a simple contact card: a `ContactInfo` struct (`Phone`, `Email`) and a `Contact` struct
that *embeds* `ContactInfo` alongside its own `Name` field. Add a method
`func (c Contact) Describe() string` that returns a formatted string using the promoted `Phone`
and `Email` fields directly (`c.Phone`, not `c.ContactInfo.Phone`). Create two or three contacts
and print each one's description.

## Summary
A struct groups related fields into a custom type, and methods (via value or pointer receivers)
attach behavior to it. Go has no inheritance; instead, a struct can contain another struct either
as a named field (**composition**, requiring `.FieldName`) or an anonymous field
(**embedding**, which promotes fields/methods onto the outer type). Structs are value types,
copied on assignment, and comparable with `==` only when every field itself is comparable.

## What to Learn Next
Previous: [17Pointers](../17Pointers/README.md)
Next: [19Slice](../19Slice/README.md)
