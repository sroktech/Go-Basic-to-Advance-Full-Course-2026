# Changelog

## 2026-09-13 — Course restructuring (Phase 1)

A full lesson-by-lesson review was done against the goal of a smooth, beginner-friendly
learning path. This phase focused on **fixing and documenting the existing 37 lessons**;
new Intermediate/Advanced/Senior content and the project ladder are tracked separately in
[ROADMAP.md](ROADMAP.md).

### Reordered lessons

The lesson numbering changed to remove three real "taught out of order" problems and one
awkward placement. Every lesson's code content is unchanged except where noted below —
only folder numbers and the numeric headers inside file comments moved.

| Old # | Lesson | New # | Why it moved |
|---|---|---|---|
| 22 | Switch | **08** | Switch is a control-flow construct like if/else — it was previously taught 14 lessons after if/else, well after structs, interfaces, and error handling. It now sits right after if/else, before loops, matching how it's actually used (an alternative to if/else chains). |
| 18 | Recursion | **14** | Grouped with the other function-related lessons (Functions, Scope, Closures, Variadic) instead of sitting after Maps. |
| 20 | Interfaces | **22** | Moved before TypeCasting (see below) — interfaces are now taught before a lesson that relies on them. |
| 19 | TypeCasting | **23** | This lesson's type-assertion and type-switch sections operate on `interface{}`/`any` — a genuine dependency on understanding interfaces, which used to be taught in the *next* lesson. It now comes right after Interfaces. |
| 25 | Defer | **24** | ErrorHandling's panic/recover section depends on understanding `defer` first. Defer now comes immediately before ErrorHandling instead of four lessons after it, removing a "concept used before it's taught" gap (and the resulting duplicate defer explanation inside ErrorHandling). |
| 21 | ErrorHandling | **25** | See above — now follows Defer. |
| 31 | Packages | **26** | Packages/modules is foundational to structuring any real Go project and was previously taught very late (after all of concurrency). It now closes out the Beginner tier, right before concurrency begins — matching how foundational it actually is. |

All other lessons kept their relative order; folder numbers shifted only to close the gaps
left by the moves above. See the updated table in [README.md](README.md) for the full final
order.

### Content fixes

- **05Constants** — `fmt.Printf(GREETING)` (passing a variable as a format string) replaced
  with `fmt.Print(GREETING)`, with a comment explaining why passing a variable to `Printf`
  is a latent bug risk (a stray `%` in the string would be misread as a format verb).
- **18Structures** — the "nested structs" example was mislabeled as struct embedding when it
  was actually a named field (composition). Added a corrected explanation plus a new, real
  embedding example (`Manager` embeds `Address` anonymously) demonstrating actual field
  promotion, so the composition-vs-embedding distinction is taught correctly.
- **27Goroutines** — the lesson used a `select`-based timeout pattern before `select` is
  formally taught (lesson 29). Rather than removing a legitimate, common pattern, it's now
  flagged as a deliberate preview with a comment pointing to the full explanation later.
  Stale lesson-number references in comments (to WaitGroup, Channels, Context) were also
  updated to match the new numbering.
- **06Operators** — every section except the final pointer example was commented-out
  reference text rather than code that actually ran. Uncommented all sections (arithmetic,
  relational, logical, assignment, bitwise), wrapping each in its own `{ }` block so they
  can reuse simple variable names (`A`, `B`, `C`, `D`) without redeclaration errors. The
  file's own output was verified against every value already documented in its comments.
- **09Control_Flow_Loops** — the *only* code that actually ran was a bare `goto`-based
  infinite loop with no exit condition, meaning `go run main.go` never terminated on its
  own; every other loop form (classic `for`, nested loop, `break`, `continue`) was
  commented out. Rewrote the file so all six sections actually execute: classic `for`,
  a nested multiplication table, `break`, `continue`, a while-style `for`, and a genuinely
  infinite `for {}` — now paired with a `break` so it terminates after 5 iterations instead
  of hanging forever. The old unbounded `goto` loop is kept as a clearly-labeled historical
  curiosity, also now bounded.
- **34HTTP** — several HTTP client calls silently discarded errors (`_ = err`), which
  contradicted the strict error-checking discipline taught in the ErrorHandling lesson;
  all client calls now check their errors. `log.Fatal` was also being called from inside a
  background goroutine on server startup failure — this calls `os.Exit` with no chance for
  `main()` to clean up, which is a risky pattern to model for beginners. Replaced with
  `log.Printf` plus a comment explaining why.

- **36Testing** — `go test ./...` didn't actually work: `go.mod` declared `module main`,
  which makes the package's import path exactly `main`. Go's toolchain refuses to import
  any package named `main` — including the internal import the `go test` harness generates
  to call the package's own test functions — so every test run failed at build time with
  `cannot import "main"`, before a single test executed. Renamed the module to
  `module mathtest`; all tests and benchmarks now run and pass correctly.

### Housekeeping

- Ran `gofmt -w` across the whole repository. Nineteen files had drifted from `gofmt`
  formatting (mostly misaligned trailing comments) — none of it changed behavior, only
  whitespace.

- Removed 6 compiled binaries (`main`) that had been accidentally committed to git despite
  `.gitignore` already listing `main`/`*.exe` (the ignore rule only prevents *new* untracked
  binaries — files already tracked before the rule existed stay tracked until explicitly
  removed). All 37 lessons were rebuilt (`go build ./...`) after the reorder to confirm
  nothing broke.
- Added a `README.md` to every one of the 37 lesson folders, following a consistent
  structure: What You Will Learn, Why This Matters, Concept Explanation, Simple Example,
  How It Works, Common Mistakes, Best Practices, Real-World Example, Exercise, Mini Project,
  Summary, What to Learn Next.
