# 22 – Interfaces

## What You Will Learn
- What an interface is: a set of method signatures with no implementation
- Why Go's interfaces are satisfied *implicitly* — no `implements` keyword
- Writing functions that accept an interface instead of a concrete type
- Polymorphism: treating different concrete types uniformly through a shared interface
- The empty interface (`interface{}` / `any`) and when it's appropriate
- Type assertions (`v, ok := x.(T)`) and type switches (`switch v := x.(type)`)
- Interface composition (embedding one interface inside another)

## Why This Matters
Back in the Switch lesson, we said "we'll explain this properly later" about interfaces — this is that promise kept. Every function you've written so far takes a specific concrete type: a `Circle`, an `int`, a `string`. That's fine until you want one function to work with *several* different types that all share some behavior — printing shape info for a circle, a rectangle, and a triangle, say, without writing three near-identical functions. Interfaces solve exactly this: they let you write code against *behavior* ("can this compute its own area?") rather than against a specific concrete type. This is also the foundation of easy testing — a function that depends on an interface can be tested with a fake/mock implementation instead of a real database or network call.

## Concept Explanation
An interface declares method signatures only:
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```
Any type that has methods matching *all* of an interface's signatures automatically satisfies that interface — there's no `implements Shape` keyword anywhere, unlike Java or C#. This is called **structural typing** (or implicit implementation): what matters is the shape of a type's method set, not any declared relationship to the interface.

Why is this considered good design, rather than just a shortcut?
- **Decoupling.** A function that accepts a `Shape` doesn't need to know (or import) the package that defines `Circle` or `Rectangle` — it only needs the `Shape` interface. Concrete types and the code that uses them abstractly can live independently.
- **Retroactive satisfaction.** A type you didn't write — even one from a third-party library — automatically satisfies your interface the moment its methods happen to match. You never need to modify that type or wrap it just to make it "implement" something.
- **Easy testing.** Code that depends on an interface (say, a `Storage` interface) can be tested against a small in-memory fake instead of a real database, with zero changes to the interface itself.

The **empty interface** `interface{}` (or its modern alias `any`) has no methods at all — so *every* type satisfies it trivially. It can hold a value of any type, which is useful for things like a generic "print anything" helper, but it throws away all compile-time type safety, so it should be used sparingly.

Once you have a value stored in an interface, you can recover the concrete type with a **type assertion** (`v, ok := x.(Circle)`) — the `ok` form never panics, it just tells you whether the assertion succeeded. A **type switch** (`switch v := x.(type) { case Circle: ... }`) handles several possible concrete types cleanly in one block.

## Simple Example
From [main.go](main.go):

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

func printShapeInfo(s Shape) {
    fmt.Printf("  Area:      %.2f\n", s.Area())
    fmt.Printf("  Perimeter: %.2f\n", s.Perimeter())
}
```

## How It Works
- `Circle`, `Rectangle`, and `Triangle` never mention `Shape` anywhere in their own declarations — they satisfy it purely because each defines an `Area() float64` and a `Perimeter() float64` method. That's structural typing in action.
- `printShapeInfo(s Shape)` can be called with a `Circle`, a `Rectangle`, or a `Triangle` — the function body only ever calls `s.Area()` and `s.Perimeter()`, never caring which concrete type is actually behind `s`.
- `shapes := []Shape{c, r, t}` stores three *different* concrete types in one slice, because all three satisfy the same interface — this is polymorphism: uniform treatment of different types through a shared behavior contract.
- `Person` satisfies a `Stringer` interface (mirroring the standard library's `fmt.Stringer`) just by defining `String() string` — which is also why `fmt.Println(p)` in the file automatically prints `"Alice (age 30)"` instead of the struct's raw fields.
- `printAnything(v interface{})` accepts absolutely anything, since the empty interface has no methods to satisfy.
- `if circle, ok := shape.(Circle); ok` safely recovers the concrete `Circle` value from the `Shape` interface variable without risking a panic if the assertion were wrong.
- The type switch inside `describe` handles `Circle`, `Rectangle`, and `Triangle` differently in one readable block, based on whatever concrete type is actually stored in the `Shape` argument.

## Common Mistakes
- **Satisfying an interface by accident.** Because implementation is structural, a type can satisfy an interface it was never intended to satisfy, just by happening to have matching method names and signatures. This is usually harmless, but it means interface satisfaction isn't always a deliberate, visible decision the way `implements Shape` would be in Java — pay attention to method names when designing interfaces to avoid unintentional or confusing overlaps.
- **Confusing a nil interface with an interface holding a nil pointer.** `var s Shape` is a true nil interface — comparing it to `nil` is `true`. But if you assign a nil pointer of some concrete type into an interface variable (e.g. a function returns a `*Circle` that is `nil`, stored in a `Shape`), the interface is *not* equal to `nil` anymore — it now holds a concrete type (`*Circle`) paired with a nil value. This surprises almost everyone the first time they hit it, typically when a function that "returns nil on success" for an interface-typed result unexpectedly compares as non-nil to its caller.
- **Over-abstracting too early.** Reaching for an interface before you have a second real implementation just adds a layer of indirection with no benefit. A common guideline is "accept interfaces, return structs" — write functions that *accept* an interface parameter (flexible for callers), but *return* a concrete struct type (concrete and easy to reason about) until a second implementation genuinely exists.

## Best Practices
- Keep interfaces small — often just one or two methods (the standard library's `io.Reader`, with a single `Read` method, is the classic example). Small interfaces are easier for many types to satisfy naturally.
- Follow "accept interfaces, return structs": function parameters can be interfaces for flexibility; return values are usually clearer as concrete types.
- Don't introduce an interface until you actually need to swap implementations (production vs. test, memory vs. database) — a single, obvious implementation rarely needs one yet.
- Use the comma-ok form of a type assertion (`v, ok := x.(T)`) instead of the panicking form (`v := x.(T)`) unless you're certain of the underlying type.

## Real-World Example
A `Storage` interface with methods like `Get(id string) (Record, error)` and `Save(r Record) error` lets application code work against "something that can store records" without caring whether the real implementation is an in-memory map (for fast unit tests), a Postgres database (in production), or a mock (for testing error handling). Swapping implementations means passing a different concrete value in — the calling code never changes. This pattern (interfaces for swappable dependencies) is one of the most common uses of interfaces in real Go codebases.

## Exercise
Define a `Named` interface with a single method `Name() string`. Create two different struct types (e.g. `Dog` and `Robot`), each with its own `Name()` method returning something different. Write a function `greet(n Named)` that prints `"Hello, " + n.Name()`, and call it with both types.

## Mini Project
Build on this lesson's `Shape` interface: add a `Square` type that also implements `Area()` and `Perimeter()`, then write a function `totalArea(shapes []Shape) float64` that sums the area of every shape in a slice, regardless of its concrete type. Print the total area for a mixed slice containing at least three different shape types.

## Summary
An interface is a set of method signatures describing behavior, and in Go any type that has matching methods satisfies it automatically — no `implements` keyword required. This structural typing decouples code from concrete types, enables polymorphism (treating different types uniformly), and makes testing easier by allowing fake implementations to stand in for real ones. The empty interface (`any`) accepts anything but gives up compile-time safety; type assertions and type switches let you recover a concrete type from an interface value when you need it. Watch for accidental satisfaction, the nil-interface-vs-nil-pointer trap, and resist introducing interfaces before a second implementation actually exists.

## What to Learn Next
Previous: [21 Maps](../21Maps/README.md)

Next: [23 Type Casting](../23TypeCasting/README.md)
