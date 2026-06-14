# Go Programming — Basic to Advanced (2026)

A complete, self-contained Go course for beginners who want to reach an advanced level.
Every lesson is a runnable Go program with detailed comments explaining the **what**, **why**, and **how** — including ASCII flow diagrams for the harder concepts.

> **Powered by [SrokTech](https://github.com/sroktech)**

---

## Who This Is For

- You are new to Go and want to learn from zero
- You know another language (Python, JavaScript, Java, etc.) and want to pick up Go
- You want to go all the way from basic syntax to goroutines, generics, and production patterns

No prior Go knowledge required. Basic programming concepts (variables, loops, functions) are helpful but each topic is explained from scratch.

---

## Course Structure

The course is organized into **37 numbered lessons** across three levels.

```
learn-go/basic-learn/
├── 00Introduction/       ← What is Go?
├── 01Setup/              ← Install and configure your environment
├── 02Basic_Syntax/       ← Your first Go program
│   ...
├── 21ErrorHandling/      ← End of basics
├── 22Switch/             ← Start of intermediate
│   ...
├── 25Defer/              ← End of intermediate
├── 26Goroutines/         ← Start of advanced (concurrency)
│   ...
└── 36Testing/            ← End of course
```

---

## Lessons

### 🟢 Basic

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
| 08 | `08Control_Flow_Loops` | for (all 4 forms), break, continue, goto, infinite loop |
| 09 | `09Functions` | declaration, parameters, return values, multiple returns, call by value/reference |
| 10 | `10Scope` | package scope, function scope, block scope, shadowing rules |

### 🟡 Intermediate

| # | Folder | Topics Covered |
|---|--------|----------------|
| 11 | `11Strings` | immutable bytes vs runes, escape sequences, strings package, []byte conversion |
| 12 | `12Arrays` | fixed-size, zero values, indexing, iteration, value-copy, 2D arrays |
| 13 | `13Pointers` | &, *, nil pointer, new(), call-by-reference, pointer to struct |
| 14 | `14Structures` | struct definition, fields, methods, value vs pointer receivers, embedding, anonymous structs |
| 15 | `15Slice` | dynamic arrays, make, append, slicing, shared backing array, copy(), 2D slices |
| 16 | `16Range` | range over slice/array/string/map, discard with _, modifying via index |
| 17 | `17Maps` | create, read, update, delete, existence check (ok idiom), reference type, struct values |
| 18 | `18Recursion` | base case, recursive case, factorial, Fibonacci, call stack traces |
| 19 | `19TypeCasting` | numeric conversion, truncation/overflow, strconv, type assertion, type switch |
| 20 | `20Interfaces` | implicit satisfaction, polymorphism, Stringer, empty interface, type assertion |
| 21 | `21ErrorHandling` | error return pattern, custom error types, fmt.Errorf wrapping, errors.Is/As |
| 22 | `22Switch` | basic switch, expression switch, type switch, fallthrough, switch with initializer |
| 23 | `23Closures` | anonymous functions, closures capturing variables, higher-order functions, loop variable gotcha |
| 24 | `24Variadic` | ...T parameter, spreading slices, mixing with regular params, functional options pattern |
| 25 | `25Defer` | LIFO order, eager argument evaluation, named returns + defer, defer in loops |

### 🔴 Advanced

| # | Folder | Topics Covered |
|---|--------|----------------|
| 26 | `26Goroutines` | go keyword, sequential vs concurrent flow, goroutine leaks, fan-out pattern |
| 27 | `27Channels` | unbuffered vs buffered, producer/consumer, pipeline, done channel, directional channels |
| 28 | `28Select` | multi-channel wait, non-blocking with default, timeout pattern, nil channel trick |
| 29 | `29Sync` | WaitGroup, Mutex, RWMutex, sync.Once, sync/atomic, race detector |
| 30 | `30Context` | WithCancel, WithTimeout, WithDeadline, WithValue, cancellation propagation |
| 31 | `31Packages` | exported vs unexported, module system, init(), standard library overview |
| 32 | `32FileIO` | os.ReadFile, os.WriteFile, bufio.Scanner, append mode, file existence check |
| 33 | `33JSON` | Marshal/Unmarshal, struct tags, omitempty, json:"-", streaming encoder/decoder |
| 34 | `34HTTP` | HTTP server with mux, handler functions, middleware, HTTP client with timeouts |
| 35 | `35Generics` | type parameters, constraints, generic Map/Filter/Reduce, generic Stack and Pair |
| 36 | `36Testing` | TestXxx, table-driven tests, t.Run, error cases, benchmarks, TestMain |

---

## How to Run a Lesson

Each lesson (except `00Introduction`, `01Setup`, `31Packages`) is a standalone Go module.

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

- Go 1.21 or later installed (lessons 35+ use Go 1.18+ generics; lesson 31 uses 1.21+ stdlib)
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
[03-10] Learn the building blocks: types, variables, constants,
        operators, control flow, functions, scope
  │
  ▼
[11-21] Work with data: strings, arrays, slices, maps, pointers,
        structs, interfaces, error handling
  │
  ▼
[22-25] Deepen function skills: switch, closures, variadic, defer
  │
  ▼
[26-30] Concurrency: goroutines, channels, select, sync, context
  │
  ▼
[31-34] Real-world Go: packages, file I/O, JSON, HTTP APIs
  │
  ▼
[35-36] Modern Go: generics (1.18+) and testing
  │
  ▼
DONE — you're ready to build production Go applications
```

---

## Go 2026 Highlights Covered

This course is current as of **Go 1.24** and covers features introduced in recent versions:

| Feature | Since | Lesson |
|---------|-------|--------|
| Generics (type parameters) | Go 1.18 | 35 |
| `any` alias for `interface{}` | Go 1.18 | 20, 35 |
| `errors.Join` | Go 1.20 | 21 |
| `slices` and `maps` packages | Go 1.21 | 31 |
| `log/slog` structured logging | Go 1.21 | 31 |
| `range` over integers (`for i := range 10`) | Go 1.22 | 08 |
| Enhanced loop variable scoping | Go 1.22 | 23 |

---

## Recommended Next Steps After This Course

1. **Build something** — a CLI tool, REST API, or file processor
2. **Explore popular frameworks** — [Gin](https://github.com/gin-gonic/gin) (HTTP), [GORM](https://gorm.io/) (database)
3. **Read the Go standard library** — `go doc fmt`, `go doc os`, etc.
4. **Tour of Go** — https://go.dev/tour (official interactive tutorial)
5. **Effective Go** — https://go.dev/doc/effective_go (best practices)
6. **Go by Example** — https://gobyexample.com (quick reference)

---

## License

This course material is free to use for personal learning.

---

<div align="center">

**Powered by [SrokTech](https://github.com/sroktech)**

*Learning Go, one lesson at a time.*

</div>
