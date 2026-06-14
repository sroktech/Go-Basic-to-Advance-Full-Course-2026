// 31 - Packages in Go
//
// A package is a directory of Go source files that share a package declaration.
// Packages are Go's unit of code organization and reuse.
//
// FLOW — how packages work:
//
//   my-project/           ← module root (go.mod lives here)
//   ├── go.mod            ← module name: learngo/packages
//   ├── main/
//   │   └── main.go       ← package main (entry point)
//   └── mathutils/
//       └── math.go       ← package mathutils (library)
//
//   import "learngo/packages/mathutils"  ← module path + package directory
//   mathutils.Add(1, 2)                  ← packagename.ExportedName
//
// Key concepts:
//   - package main = executable program (has main() function)
//   - any other name = library package (imported by others)
//   - all .go files in a folder must share the same package name
//   - exported = uppercase first letter (public)
//   - unexported = lowercase first letter (private to package)

package main

import (
	"fmt"
	// Import using: module-name/path-to-package
	"learngo/packages/mathutils"
)

func main() {

	// Using exported functions from mathutils
	fmt.Println("Add(3, 4):", mathutils.Add(3, 4))           // 7
	fmt.Println("Subtract(10, 3):", mathutils.Subtract(10, 3)) // 7
	fmt.Println("Multiply(6, 7):", mathutils.Multiply(6, 7))  // 42

	result, err := mathutils.Divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Divide(10, 3): %.4f\n", result) // 3.3333
	}

	_, err = mathutils.Divide(5, 0)
	if err != nil {
		fmt.Println("Divide by zero:", err)
	}

	fmt.Println("Max(8, 12):", mathutils.Max(8, 12)) // 12
	fmt.Printf("Pi: %.5f\n", mathutils.Pi)            // 3.14159

	// mathutils.max(1,2) would be a compile error — unexported, not visible here

	// ─── Import aliases ───────────────────────────────────────────────────────
	// If two packages have the same name, alias one:
	//   import (
	//     myu "learngo/packages/mathutils"
	//     stdu "math"
	//   )
	//   myu.Add(1, 2)
	//   stdu.Sqrt(4)

	// ─── Blank import (side effects only) ────────────────────────────────────
	// import _ "some/package"   ← runs package init() but doesn't use any names
	// Used for: database drivers, image format decoders, plugin registration

	// ─── init() function ─────────────────────────────────────────────────────
	// Each package can have one or more init() functions.
	// They run automatically when the package is imported, before main().
	// Order: imported packages init first, then the importing package.

	// ─── Standard library overview ───────────────────────────────────────────
	fmt.Println("\nCommonly used standard library packages:")
	fmt.Println("  fmt        — formatted I/O (Println, Printf, Sprintf)")
	fmt.Println("  os         — OS operations (files, env vars, exit)")
	fmt.Println("  io         — basic I/O primitives")
	fmt.Println("  bufio      — buffered I/O (read line by line)")
	fmt.Println("  strings    — string manipulation")
	fmt.Println("  strconv    — string <-> number conversions")
	fmt.Println("  math       — math functions (Sqrt, Abs, Floor...)")
	fmt.Println("  sort       — sorting slices and custom collections")
	fmt.Println("  time       — time, duration, timers, tickers")
	fmt.Println("  encoding/json — JSON encode/decode")
	fmt.Println("  net/http   — HTTP client and server")
	fmt.Println("  context    — cancellation and deadlines")
	fmt.Println("  sync       — WaitGroup, Mutex, Once")
	fmt.Println("  errors     — error creation and wrapping")
	fmt.Println("  log/slog   — structured logging (Go 1.21+)")
	fmt.Println("  slices     — slice helpers: Sort, Contains, Index (Go 1.21+)")
	fmt.Println("  maps       — map helpers: Keys, Values, Clone (Go 1.21+)")
}
