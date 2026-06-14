// 24 - Variadic Functions in Go
//
// A variadic function accepts a VARIABLE NUMBER of arguments of the same type.
// Use ...type in the parameter list — the arguments arrive as a SLICE inside the function.
//
// FLOW:
//
//   sum(1, 2, 3, 4, 5)
//        │
//        └─► nums = []int{1, 2, 3, 4, 5}  (Go packs them into a slice)
//                   │
//                   └─► iterate and add → 15
//
// Rules:
//   - The variadic parameter must be the LAST parameter
//   - There can only be ONE variadic parameter per function
//   - You can pass zero arguments — the slice will just be empty
//   - Use ... to spread a slice into variadic args when calling

package main

import "fmt"

// ─── Basic variadic function ───────────────────────────────────────────────────

// sum accepts any number of ints — nums is a []int slice inside the function
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// ─── Mixed parameters (regular + variadic) ────────────────────────────────────

// The variadic part must come LAST
func greetAll(greeting string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}

// ─── Variadic with interface{} — accept any type ──────────────────────────────

// This is how fmt.Println is implemented internally
func printAll(values ...interface{}) {
	for i, v := range values {
		fmt.Printf("  [%d] %v (%T)\n", i, v, v)
	}
}

// ─── Passing variadic args to another variadic function ───────────────────────

// When forwarding variadic args, use ... to spread the slice
func sumAndDouble(nums ...int) int {
	return sum(nums...) * 2 // spread nums slice into sum's variadic param
}

// ─── Variadic with a non-int type ─────────────────────────────────────────────

func joinStrings(sep string, parts ...string) string {
	result := ""
	for i, p := range parts {
		if i > 0 {
			result += sep
		}
		result += p
	}
	return result
}

func main() {

	// ─── Calling with different numbers of arguments ───────────────────────────

	fmt.Println("sum():", sum())           // 0 — zero args, nums is empty slice
	fmt.Println("sum(5):", sum(5))         // 5
	fmt.Println("sum(1,2,3):", sum(1, 2, 3)) // 6
	fmt.Println("sum(1..10):", sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // 55

	// ─── Spreading a slice into variadic args with ... ─────────────────────────

	// You already have a slice and want to pass it to a variadic function
	numbers := []int{10, 20, 30, 40, 50}

	// WITHOUT ...: would be a compile error — []int is not int
	// sum(numbers) ← ERROR

	// WITH ...: spreads the slice elements as individual arguments
	fmt.Println("sum(numbers...):", sum(numbers...)) // 150

	// ─── Mixed regular + variadic ──────────────────────────────────────────────

	greetAll("Hello", "Alice", "Bob", "Charlie")
	greetAll("Hi")             // zero names — nothing prints
	greetAll("Hey", "Dave")    // one name

	// ─── interface{} variadic ─────────────────────────────────────────────────

	fmt.Println("printAll demo:")
	printAll(42, "hello", true, 3.14)

	// ─── Forwarding variadic args ─────────────────────────────────────────────

	fmt.Println("sumAndDouble(1,2,3):", sumAndDouble(1, 2, 3)) // (1+2+3)*2 = 12

	// ─── joinStrings ──────────────────────────────────────────────────────────

	fmt.Println(joinStrings(", ", "apple", "banana", "cherry"))
	fmt.Println(joinStrings(" | ", "Go", "Rust", "Python"))
	fmt.Println(joinStrings("-"))  // empty — no parts

	// ─── Variadic inside a function (local) ───────────────────────────────────

	// You can also assign variadic functions to variables
	max := func(nums ...int) int {
		if len(nums) == 0 {
			return 0
		}
		m := nums[0]
		for _, n := range nums[1:] {
			if n > m {
				m = n
			}
		}
		return m
	}

	fmt.Println("max(3,1,4,1,5,9,2,6):", max(3, 1, 4, 1, 5, 9, 2, 6)) // 9

	// ─── Real-world pattern: options / config ──────────────────────────────────

	// A common Go pattern is "functional options" using variadic funcs
	type ServerConfig struct {
		Host    string
		Port    int
		Timeout int
	}

	type Option func(*ServerConfig)

	withPort := func(port int) Option {
		return func(cfg *ServerConfig) { cfg.Port = port }
	}
	withTimeout := func(t int) Option {
		return func(cfg *ServerConfig) { cfg.Timeout = t }
	}

	newServer := func(host string, opts ...Option) ServerConfig {
		cfg := ServerConfig{Host: host, Port: 8080, Timeout: 30} // defaults
		for _, opt := range opts {
			opt(&cfg) // apply each option
		}
		return cfg
	}

	s1 := newServer("localhost")                           // all defaults
	s2 := newServer("prod.example.com", withPort(443), withTimeout(60)) // custom

	fmt.Printf("Server1: %+v\n", s1)
	fmt.Printf("Server2: %+v\n", s2)
}
