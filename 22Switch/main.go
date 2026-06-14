// 22 - Switch Statement in Go
//
// Switch is a cleaner alternative to a long chain of if/else if.
// Go's switch is more powerful than in C/Java:
//   - No automatic fall-through (no need for 'break' after each case)
//   - Cases can have multiple values
//   - Cases can be expressions, not just constants
//   - Can switch on ANY comparable type
//   - Can switch with no condition (acts like if/else)
//
// FLOW:
//
//   switch value {
//        │
//        ├─── case A: ──► [run A block] ──► exit switch
//        │
//        ├─── case B: ──► [run B block] ──► exit switch
//        │
//        ├─── case C,D: ► [run C/D block] ─► exit switch
//        │
//        └─── default: ─► [run default] ──► exit switch
//

package main

import (
	"fmt"
	"time"
)

func main() {

	// ─── Basic switch ─────────────────────────────────────────────────────────

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the work week — coffee time!")
	case "Friday":
		fmt.Println("End of the work week — almost there!")
	case "Saturday", "Sunday": // multiple values in one case
		fmt.Println("Weekend — relax!")
	default: // runs when no case matches
		fmt.Println("Mid-week grind:", day)
	}

	// ─── Switch on integer ────────────────────────────────────────────────────

	score := 85

	switch {
	// No condition — each case is a full boolean expression (like if/else if)
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B") // matches here
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// ─── Switch with initializer ──────────────────────────────────────────────

	// Like if, switch can run a short statement before evaluating the value.
	// 'hour' is scoped to the switch block only.
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// ─── Type switch ──────────────────────────────────────────────────────────

	// Used to discover the concrete type of an interface{} value.
	// Very common when working with interface{} / any parameters.
	//
	// FLOW:
	//   interface{} value
	//        │
	//        ├─── case int    ──► handle as int
	//        ├─── case string ──► handle as string
	//        ├─── case bool   ──► handle as bool
	//        └─── default     ──► unknown type

	values := []interface{}{42, "hello", true, 3.14, nil}

	for _, v := range values {
		switch t := v.(type) {
		case int:
			fmt.Printf("int: %d (doubled: %d)\n", t, t*2)
		case string:
			fmt.Printf("string: %q (length: %d)\n", t, len(t))
		case bool:
			fmt.Printf("bool: %t\n", t)
		case float64:
			fmt.Printf("float64: %.2f\n", t)
		case nil:
			fmt.Println("nil value")
		default:
			fmt.Printf("unknown type: %T\n", t)
		}
	}

	// ─── fallthrough — explicit fall-through ──────────────────────────────────

	// By default Go does NOT fall through. Use 'fallthrough' to force it.
	// The next case block runs unconditionally — its condition is NOT checked.
	// Rarely used — usually a sign that your logic can be restructured.

	n := 1
	switch n {
	case 1:
		fmt.Println("case 1")
		fallthrough // forces execution into case 2 regardless
	case 2:
		fmt.Println("case 2 (fell through from 1)")
	case 3:
		fmt.Println("case 3")
	}

	// ─── Switch vs if/else — when to use which ────────────────────────────────

	// Use switch when: matching one value against many known options (cleaner)
	// Use if/else when: complex boolean logic with different variables per branch

	// Switch on type name (common pattern in Go CLI tools / parsers)
	printType := func(v interface{}) string {
		switch v.(type) {
		case int:
			return "integer"
		case float64:
			return "float"
		case string:
			return "text"
		case bool:
			return "boolean"
		default:
			return "other"
		}
	}

	fmt.Println(printType(10))      // integer
	fmt.Println(printType("hi"))    // text
	fmt.Println(printType(true))    // boolean
	fmt.Println(printType([]int{})) // other
}
