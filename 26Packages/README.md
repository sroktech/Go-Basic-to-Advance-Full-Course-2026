# 26 – Packages

## What You Will Learn
- What a package is, and why Go organizes code into packages instead of one giant file
- How exported vs. unexported identifiers work — and why it's capitalization, not a keyword
- How Go's module system ties packages together (`go.mod`, `go mod init`, import paths)
- How to import and use a package you wrote yourself, in a real two-package example
- Import aliases, blank imports, and `init()`, at a glance

## Why This Matters
Every language eventually needs a way to split code across files and give it a namespace — otherwise a real project becomes an unmanageable single file, or a pile of files that quietly collide on names. Go's answer is the **package**: a directory of `.go` files that share one purpose and one namespace. This lesson closes out the Beginner tier because organizing code into packages is a foundational skill for writing *any* real Go project, however small — and the lessons coming up (concurrency, JSON, HTTP) will all assume you're comfortable splitting code across files and packages rather than writing everything in `main`.

Go also makes an unusual choice for visibility: instead of `public`/`private` keywords (as in Java, C#, or TypeScript), visibility is just capitalization. `Add` is exported (callable from other packages); `add` is not. This is baked into the language, not a convention layered on top — the compiler enforces it.

## Concept Explanation
A **module** is a collection of packages versioned and distributed together, declared by a `go.mod` file at its root (you briefly saw `go mod init` back in [01Setup](../01Setup/README.md)). This lesson's `go.mod` declares:

```
module learngo/packages

go 1.24.0
```

That first line, `module learngo/packages`, sets the **import path prefix** for every package inside this folder. A package living in `mathutils/` is imported as `"learngo/packages/mathutils"` — module path + the sub-directory path, as seen in [main/main.go](main/main.go):

```go
import "learngo/packages/mathutils"
```

Inside a module, each directory is one package:
- `main/main.go` declares `package main` — the special package name that marks an executable program with an entry-point `func main()`.
- `mathutils/math.go` declares `package mathutils` — an ordinary *library* package, meant to be imported, not run directly.

All `.go` files in the same directory must declare the same package name — one package per directory is a hard rule, not just a convention.

**Exported vs. unexported**: a top-level name (function, type, variable, constant) starting with an uppercase letter is visible to other packages that import this one; a name starting lowercase is private to its own package. [mathutils/math.go](mathutils/math.go) uses both:

```go
func Add(a, b int) int { return a + b }        // exported — callers outside the package can use it

func max(a, b int) int { ... }                  // unexported — only usable inside mathutils

func Max(a, b int) int { return max(a, b) }     // exported wrapper around the unexported helper
```

## Simple Example
From [main/main.go](main/main.go), calling into the `mathutils` package:

```go
import "learngo/packages/mathutils"

result, err := mathutils.Divide(10, 3)
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Printf("Divide(10, 3): %.4f\n", result)
}

fmt.Printf("Pi: %.5f\n", mathutils.Pi) // an exported constant
```

## How It Works
When `main.go` writes `import "learngo/packages/mathutils"`, the Go toolchain resolves that path using the current module's `go.mod` (`module learngo/packages`) plus the remaining path segment (`mathutils`) to find the directory `mathutils/` on disk. Everything exported from that package becomes available as `mathutils.Name`.

The `max` function in `mathutils/math.go` cannot be called as `mathutils.max(...)` from `main.go` — that line would be a compile error, not a runtime one. The compiler checks capitalization at compile time, the same way it checks that a variable exists at all. This is why `Max` exists as a thin exported wrapper: it's a common pattern to keep a helper unexported (so you're free to change its behavior later without breaking anyone) while exposing a stable, exported entry point.

## Common Mistakes
- **Import cycles**: if package A imports package B, package B cannot import package A (directly or transitively) — Go refuses to compile it. Design your packages so dependencies flow one direction (e.g., a low-level `mathutils` never needs to import the `main` package that uses it).
- **Forgetting that lowercase = unexported**: coming from languages with explicit `public`/`private` keywords, it's easy to write a lowercase helper and be confused why another package can't see it. There's no keyword to search for — check the first letter.
- **Mixing package names in one directory**: every `.go` file in a folder must declare the same `package` name. A stray `package mathutil` (typo) in one file of the `mathutils/` folder won't compile.

## Best Practices
- One package per directory, and name the package to match the directory (`mathutils/` → `package mathutils`).
- Keep exported surface area small and deliberate — export only what callers actually need, and keep implementation helpers unexported so you can change them freely later.
- Use `Max`-style exported wrappers around unexported helpers when you want a stable public API but the freedom to refactor internals.
- Give packages short, lowercase, no-underscore names (`mathutils`, not `MathUtils` or `math_utils`), matching the standard library's own convention.

## Real-World Example
Real Go projects extend this same idea to a whole directory layout: an `internal/` folder for packages that should never be imported from outside the module (enforced by the compiler, not just convention), a `pkg/` folder for code meant to be reused by other projects, and feature-oriented packages (`billing/`, `auth/`, `notifications/`) instead of one flat pile of files. The two-package split you just built — `main` importing `mathutils` — is the smallest possible version of that same structure: separate the thing that *runs* from the libraries that *do the work*, so the libraries stay testable and reusable on their own.

## Exercise
Create a new package `stringutils` (its own directory, its own `package stringutils` declaration) with one exported function, `Reverse(s string) string`, that reverses a string (remember from [15Strings](../15Strings/README.md) that you'll likely need to convert to `[]rune` first to handle multi-byte characters correctly, as in lesson 23's Unicode example). Import it from `main` and print a few reversed strings.

## Mini Project
Extend the `mathutils` package with an exported `Average(nums ...float64) (float64, error)` function (using variadic parameters from [13Variadic](../13Variadic/README.md)) that returns an error if given zero arguments, and an unexported helper `sum(nums []float64) float64` that it calls internally. Update `main.go` to call `Average` with a few sample slices, handling the returned error the way lesson 25 taught, and confirm that trying to call `mathutils.sum(...)` directly from `main` fails to compile.

## Summary
A package is a directory of `.go` files sharing one package declaration — Go's unit of code organization and reuse — and a module (declared by `go.mod`, with its own import-path prefix) is a versioned collection of packages. Visibility between packages is controlled purely by capitalization: uppercase-first is exported, lowercase-first is private to the package, enforced by the compiler rather than a keyword. This closes out the Beginner tier of the course — from here the course moves into concurrency (the Intermediate tier), starting with goroutines.

## What to Learn Next
Previous: [25ErrorHandling](../25ErrorHandling/README.md)
Next: [27Goroutines](../27Goroutines/README.md)
