# Go Programming — Basic to Advanced (2026)

A complete, self-contained Go course for beginners who want to reach an advanced level.
Every lesson is a runnable Go program with detailed comments explaining the **what**, **why**, and **how** — including ASCII flow diagrams for the harder concepts. Every lesson folder also has its own `README.md` with the full teaching material: why the topic matters, common mistakes, best practices, a real-world example, an exercise, and a mini project.

> **Powered by [SrokTech](https://github.com/sroktech)**

---

## Who This Is For

- You are new to Go and want to learn from zero
- You know another language (Python, JavaScript, Java, etc.) and want to pick up Go
- You want to go all the way from basic syntax to goroutines, generics, and production patterns

No prior Go knowledge required. Basic programming concepts (variables, loops, functions) are helpful but each topic is explained from scratch.

---

## Course Structure

The course is organized into **37 numbered lessons** (00–36) covering **Beginner** and **Intermediate** Go. Lessons were reordered in the latest revision so every topic builds only on what came before it — see [CHANGELOG.md](CHANGELOG.md) for exactly what moved and why.

A [ROADMAP.md](ROADMAP.md) describes what comes after lesson 36 — databases, REST APIs, architecture, and production/senior-level topics — for when you're ready to keep going.

```
Go-Basic-to-Advance-Full-Course-2026/
├── 00Introduction/       ← What is Go?
├── 01Setup/              ← Install and configure your environment
├── 02Basic_Syntax/       ← Your first Go program
│   ...
├── 26Packages/           ← End of Beginner tier
├── 27Goroutines/         ← Start of Intermediate tier (concurrency)
│   ...
└── 36Testing/            ← End of current course content
```

---

## Lessons

### 🟢 Beginner

| # | Folder | Topics Covered |
|---|--------|----------------|
| 00 | `00Introduction` | What is Go, history, key features, why choose Go |
| 01 | `01Setup` | Installation, go.mod, workspace, CLI commands |
| 02 | `02Basic_Syntax` | package, import, main(), print functions, syntax rules |
| 03 | `03Data_Types` | int, float32/64, bool, complex64/128, string — sizes and ranges |
| 04 | `04Variables` | var, :=, type inference, multi-declaration, shadowing |
| 05 | `05Constants` | const, typed/untyped, grouped const, iota, compile-time expressions |
| 06 | `06Operators` | arithmetic, relational, logical, bitwise (with binary diagrams), assignment |
| 07 | `07Control_Flow_if_else` | if, else, else if chains, flow traces for each path |
| 08 | `08Switch` | basic switch, expression switch, type switch, fallthrough, switch with initializer |
| 09 | `09Control_Flow_Loops` | for (all 4 forms), break, continue, goto, infinite loop |
| 10 | `10Functions` | declaration, parameters, return values, multiple returns, call by value/reference |
| 11 | `11Scope` | package scope, function scope, block scope, shadowing rules |
| 12 | `12Closures` | anonymous functions, closures capturing variables, higher-order functions, loop variable gotcha |
| 13 | `13Variadic` | ...T parameter, spreading slices, mixing with regular params, functional options pattern |
| 14 | `14Recursion` | base case, recursive case, factorial, Fibonacci, call stack traces |
| 15 | `15Strings` | immutable bytes vs runes, escape sequences, strings package, []byte conversion |
| 16 | `16Arrays` | fixed-size, zero values, indexing, iteration, value-copy, 2D arrays |
| 17 | `17Pointers` | &, *, nil pointer, new(), call-by-reference, pointer to struct |
| 18 | `18Structures` | struct definition, fields, methods, value vs pointer receivers, composition vs. embedding, anonymous structs |
| 19 | `19Slice` | dynamic arrays, make, append, slicing, shared backing array, copy(), 2D slices |
| 20 | `20Range` | range over slice/array/string/map, discard with _, modifying via index |
| 21 | `21Maps` | create, read, update, delete, existence check (ok idiom), reference type, struct values |
| 22 | `22Interfaces` | implicit satisfaction, polymorphism, Stringer, empty interface, type assertion |
| 23 | `23TypeCasting` | numeric conversion, truncation/overflow, strconv, type assertion, type switch |
| 24 | `24Defer` | LIFO order, eager argument evaluation, named returns + defer, defer in loops |
| 25 | `25ErrorHandling` | error return pattern, custom error types, fmt.Errorf wrapping, errors.Is/As, panic/recover |
| 26 | `26Packages` | exported vs unexported, module system, init(), standard library overview |

### 🟡 Intermediate

| # | Folder | Topics Covered |
|---|--------|----------------|
| 27 | `27Goroutines` | go keyword, sequential vs concurrent flow, goroutine leaks, fan-out pattern |
| 28 | `28Channels` | unbuffered vs buffered, producer/consumer, pipeline, done channel, directional channels |
| 29 | `29Select` | multi-channel wait, non-blocking with default, timeout pattern, nil channel trick |
| 30 | `30Sync` | WaitGroup, Mutex, RWMutex, sync.Once, sync/atomic, race detector |
| 31 | `31Context` | WithCancel, WithTimeout, WithDeadline, WithValue, cancellation propagation |
| 32 | `32FileIO` | os.ReadFile, os.WriteFile, bufio.Scanner, append mode, file existence check |
| 33 | `33JSON` | Marshal/Unmarshal, struct tags, omitempty, json:"-", streaming encoder/decoder |
| 34 | `34HTTP` | HTTP server with mux, handler functions, middleware, HTTP client with timeouts |
| 35 | `35Generics` | type parameters, constraints, generic Map/Filter/Reduce, generic Stack and Pair |
| 36 | `36Testing` | TestXxx, table-driven tests, t.Run, error cases, benchmarks, TestMain |

**What's next?** Lesson 36 completes the course's current content. [ROADMAP.md](ROADMAP.md) lays out the planned continuation: databases (PostgreSQL), REST APIs, logging/config, clean architecture, microservices, and production/senior-level Go engineering — plus the project ladder (CLI → REST API → REST+DB → auth service → ... → production system).

---

## How to Run a Lesson

Each lesson (except `00Introduction`, `01Setup`, `26Packages`) is a standalone Go module.

```bash
# Navigate to any lesson
cd 03Data_Types

# Run the program
go run main.go

# For 36Testing — run the tests
cd 36Testing
go test -v ./...

# Run benchmarks
go test -bench=. ./...

# Check for race conditions
go run -race main.go
```

---

## Prerequisites

- Go 1.21 or later installed (lesson 35 uses Go 1.18+ generics; lesson 26 uses 1.21+ stdlib)
- A code editor — [VS Code](https://code.visualstudio.com/) with the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) is recommended
- A terminal

Verify your Go version:
```bash
go version
# go version go1.24.x ...
```

---

## Learning Path

```
START
  │
  ▼
[00-02] Understand what Go is and write your first program
  │
  ▼
[03-09] Learn the building blocks: types, variables, constants,
        operators, conditionals (if/else, switch), loops
  │
  ▼
[10-14] Master functions: declaration, scope, closures, variadic
        functions, recursion
  │
  ▼
[15-21] Work with data: strings, arrays, pointers, structs, slices,
        range, maps
  │
  ▼
[22-26] Tie it together: interfaces, type casting/assertion, defer,
        error handling, packages & modules — end of Beginner tier
  │
  ▼
[27-31] Concurrency: goroutines, channels, select, sync, context
  │
  ▼
[32-36] Real-world Go: file I/O, JSON, HTTP APIs, generics, testing
  │
  ▼
DONE — ready for REST APIs, databases, and production Go (see ROADMAP.md)
```

---

## Go 2026 Highlights Covered

This course is current as of **Go 1.24** and covers features introduced in recent versions:

| Feature | Since | Lesson |
|---------|-------|--------|
| Generics (type parameters) | Go 1.18 | 35 |
| `any` alias for `interface{}` | Go 1.18 | 22, 23 |
| `errors.Join` | Go 1.20 | 25 |
| `slices` and `maps` packages | Go 1.21 | 26 |
| `log/slog` structured logging | Go 1.21 | 26 |
| `range` over integers (`for i := range 10`) | Go 1.22 | 09 |
| Enhanced loop variable scoping | Go 1.22 | 12 |

---

## Recommended Next Steps After This Course

1. **Build something** — a CLI tool, REST API, or file processor
2. **Keep going** — see [ROADMAP.md](ROADMAP.md) for the planned Advanced/Senior continuation of this course
3. **Explore popular frameworks** — [Gin](https://github.com/gin-gonic/gin) (HTTP), [GORM](https://gorm.io/) (database)
4. **Read the Go standard library** — `go doc fmt`, `go doc os`, etc.
5. **Tour of Go** — https://go.dev/tour (official interactive tutorial)
6. **Effective Go** — https://go.dev/doc/effective_go (best practices)
7. **Go by Example** — https://gobyexample.com (quick reference)

---

## License

This course material is free to use for personal learning.

---

<div align="center">

**Powered by [SrokTech](https://github.com/sroktech)**

*Learning Go, one lesson at a time.*

</div>
