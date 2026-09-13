# 02 – Basic Syntax

## What You Will Learn
The anatomy of every Go source file (package, imports, functions), how comments work, the different ways to print output, and the strict syntax rules the Go compiler enforces.

## Why This Matters
Before you can write any logic, you need to know the "shape" every Go file must take and the ground rules the compiler will not bend on. Go is deliberately less permissive than languages like Python or JavaScript here — fewer ways to do the same thing, enforced by the compiler itself — which keeps codebases consistent across an entire team or company, not just your own project.

## Concept Explanation
Every Go source file has the same three-part structure: a **package declaration** (which package this file belongs to), an **import statement** (which other packages this file needs), and then the actual **functions**. `package main` is special — it marks a file as a runnable program rather than a reusable library, and it must contain a `func main()`, which is the entry point: the first code that runs when you execute the program.

To print anything, you need the `fmt` ("format") package from Go's standard library — the set of packages that ship with Go itself, so no installation is required. Go offers a few different printing functions because the standard library favors being explicit over having one function try to do everything.

## Simple Example
From [main.go](main.go) in this folder:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")

    name := "Gopher"
    age := 5
    fmt.Printf("My name is %s and I am %d years old.\n", name, age)
}
```

## How It Works
- `package main` + `import "fmt"` must come before any other code, in that order.
- `fmt.Println` prints its arguments and automatically adds a newline at the end — the most common way to print in Go.
- `fmt.Print` prints with no automatic newline, so consecutive calls stay on the same line unless you add `\n` yourself.
- `fmt.Printf` is "formatted print" — it takes a template string with **verbs** (placeholders) like `%s` (string), `%d` (integer), `%f` (float), `%t` (bool), `%T` (type name), and `%v` (any value, default format), then fills them in with the arguments that follow.
- The built-in `println()` (lowercase, no import needed) also exists, but it writes to stderr and is meant only for compiler-level debugging — real Go code uses `fmt.Println` instead.

## Common Mistakes
- **Unused imports and unused variables are compile errors in Go**, not warnings. If you `import "fmt"` but never call anything from it, or declare a variable you never read, your program simply will not build. This is one of the most common early frustrations, and it's a deliberate design choice to keep code free of dead weight.
- Putting the opening curly brace `{` on its own line. In Go it must be on the same line as `func`, `if`, `for`, etc. — this isn't just style, the compiler actually requires it because of how Go auto-inserts semicolons.
- Mixing up case: `fmt.Println` (capital P) works; `fmt.println` does not exist. Capitalization also has meaning in Go — a capitalized name is **exported** (accessible from other packages), lowercase is not. That's why `Println` is capitalized in the `fmt` package.
- Manually adding semicolons at the end of lines — Go inserts them automatically, and writing your own is redundant and non-idiomatic.

## Best Practices
- Prefer `fmt.Println` for simple output and `fmt.Printf` when you need to control formatting or mix in variables.
- Run `go fmt ./...` regularly so your code always matches Go's one official style — there's no debate to have about formatting in a Go codebase.
- Use `go vet ./...` to catch subtle mistakes (like a mismatched `Printf` verb) that compile fine but behave wrong at runtime.

## Real-World Example
This strictness is not pedantry — in a codebase with thousands of files and dozens of engineers, an unused import or variable is very often a leftover from a refactor, a sign of a bug, or dead code nobody noticed. By making these compile errors instead of warnings, Go forces every file that reaches production to be clean, which is a big part of why large companies rely on it for infrastructure code that many people touch over time.

## Exercise
Modify `main.go` (or write a new small file) to print your own name and favorite number using `fmt.Printf` with the `%s` and `%d` verbs, in a single formatted string.

## Mini Project
Write a tiny "personal info printer": declare a name (string) and an age (int) using `:=`, then use `fmt.Printf` to print a sentence combining both, followed by a separate `fmt.Println` call that prints a one-line fun fact about yourself.

## Summary
Every Go file starts with a package declaration and imports, followed by functions, with `main()` as the required entry point in `package main`. Go gives you `Println`, `Print`, and `Printf` for output, and its compiler strictly enforces things like brace placement, unused imports, and unused variables — rules that keep every Go codebase uniformly clean.

## What to Learn Next
Now that you can structure a file and print values, the next lesson covers the actual types those values can have — integers, floats, booleans, and more.

Prev: [../01Setup/README.md](../01Setup/README.md)
Next: [../03Data_Types/README.md](../03Data_Types/README.md)
