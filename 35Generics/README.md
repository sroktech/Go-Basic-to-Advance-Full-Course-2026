# 35 – Generics

## What You Will Learn
- Writing generic functions with type parameters (`func Min[T Ordered](a, b T) T`)
- Defining custom type constraints (`Integer`, `Number`, `Ordered`) as interfaces
- The built-in `comparable` and `any` constraints
- Generic data structures: a type-parameterized `Stack[T]` and `Pair[A, B]`
- `Map`, `Filter`, and `Reduce` — generic slice utilities
- When generics are the right tool, and when they aren't

## Why This Matters
Before Go 1.18 (2022), you had exactly two options for code that needed to work across multiple types: duplicate the function for every type (`MinInt`, `MinFloat64`, `MinString`, ...), or take `interface{}` and lose type safety, forcing type assertions and reflection at every call site with no compile-time guarantee you got it right. Generics solve a specific, narrow problem: letting you write **one function body** that the compiler turns into type-safe code for many concrete types, with no runtime type assertions and no duplicated source. That's it — generics are not a general replacement for interfaces, and knowing the boundary between the two is exactly what separates idiomatic Go generics use from over-engineered code.

## Concept Explanation
A type parameter list in square brackets introduces one or more placeholder types, constrained by an interface that lists which concrete types are allowed:

```
func Min[T int | float64](a, b T) T
      │ │
      │ └─ T is constrained: only int or float64 allowed
      └─── [T ...] is the type parameter list

Call: Min[int](3, 5)       ← explicit type
      Min(3, 5)            ← type inferred from arguments (preferred)
```

A constraint is just an interface whose method set (or, for generics, whose allowed underlying types via a union like `int | float64`) defines what operations `T` supports. `comparable` is a built-in constraint meaning `T` supports `==`/`!=` — needed for anything that compares values, like `Contains`. `any` is an alias for `interface{}` used as a constraint meaning "no restriction at all" — appropriate when the function never operates *on* the values directly, only stores or passes them through (like `Stack[T]` or `Map`'s input/output types).

Interfaces and generics solve *different* problems: an interface lets one function work with many types **because they share behavior** (a `Stringer` has `String()`, regardless of underlying type) — the function calls that shared method and doesn't need to know the concrete type. Generics let one function work with many types **because the logic is identical regardless of type** — there's no shared method to call; the same arithmetic or comparison just needs to happen on whatever concrete type was passed in.

## Simple Example
A constraint and a generic function using it, from [main.go](main.go):

```go
type Ordered interface {
    int | int8 | int16 | int32 | int64 |
        float32 | float64 | string
}

func Min[T Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

A generic data structure:

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}
```

## How It Works
`Min(3, 7)` and `Min(3.14, 2.71)` call the *same* function definition, but the compiler generates a version specialized for `int` and one for `float64` — there's no runtime type switch and no boxing into `interface{}`. Type inference means you almost never need to write `Min[int](3, 7)` explicitly; the compiler works out `T` from the argument types.

`Sum[T Number](nums []T) T` needs a wider constraint than `Ordered` restricts to arithmetic-capable types (`+=` isn't valid on strings the way `Sum` uses it), which is why `Number` and `Ordered` are defined as separate, purpose-built constraints rather than one catch-all.

`Contains[T comparable]` only needs `==`, so `comparable` — not a custom union — is the right, minimal constraint; it works for `int`, `string`, or any comparable struct without listing types explicitly.

`Map[T, U any]`, `Filter[T any]`, and `Reduce[T, U any]` never perform any operation *on* `T` or `U` themselves — they only call a caller-supplied function (`fn func(T) U`, `predicate func(T) bool`) — so `any` is correct: the constraint carries no requirement because the generic code itself makes no assumptions about the type.

`Stack[T]` and `Pair[A, B]` show that generics apply to types, not just functions — `Stack[int]{}` and `Stack[string]{}` are two different concrete types generated from one generic definition, each fully type-safe (`Pop()` on a `Stack[string]` returns a `string`, not `interface{}`).

## Common Mistakes
- Reaching for a generic function when a concrete type or a plain interface would do — if you only ever call `Min` with `int`, a generic version adds indirection for no benefit.
- Writing an overly broad constraint (`any`) on a function that actually needs specific operations (`<`, `+`), which only surfaces as a compile error at the call site instead of at the function definition where it belongs.
- Treating generics as an interface replacement — if the goal is "many types share *behavior* via a method," that's an interface; generics are for "the same code should run against different types with no shared method."
- Forgetting `comparable` when a generic function needs `==`/`!=` — `any` does not support comparison operators.

## Best Practices
- Design the constraint to be exactly as wide as the operations you need — no wider. `Ordered` for comparisons, `Number` for arithmetic, `comparable` for equality, `any` for pure pass-through.
- Prefer type inference (`Min(3, 7)`) over explicit instantiation (`Min[int](3, 7)`) — only specify the type parameter when inference genuinely can't determine it.
- Don't generalize prematurely: write the concrete version first, and promote it to a generic only once you actually need it for more than one type.
- Use generics for utility functions and containers (Map/Filter/Reduce, Stack, Pair, Set) where the logic is identical regardless of type — that's the sweet spot Go 1.18 was built for.

## Real-World Example
A generic utility library (`Map`, `Filter`, `Reduce`, `Contains`) is common in real Go codebases as a small internal package used across many services, replacing what used to be either copy-pasted per-type helpers or `interface{}`-based utilities with runtime type assertions. Generic data structures like a type-safe `Set[T]` or a request/response cache keyed by a `comparable` type also show up directly in backend infrastructure code.

## Exercise
Write a generic `Unique[T comparable](slice []T) []T` function that returns a new slice with duplicate values removed, preserving order of first appearance. Test it with both an `[]int` and an `[]string`.

## Mini Project
Build a small generic utility package with `Map`, `Filter`, and a new `Find[T any](slice []T, predicate func(T) bool) (T, bool)` function (returns the first matching element and whether one was found). Demonstrate all three working against at least two different concrete types (e.g., a slice of `int` and a slice of a small `Person` struct).

## What to Learn Next
Continue to [36 – Testing](../36Testing/README.md), the final lesson of this course, where you'll learn to verify code like the generic functions and HTTP handlers you've just built using Go's built-in `testing` package.
