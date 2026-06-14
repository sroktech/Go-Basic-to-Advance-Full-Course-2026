// 17 - Maps in Go
//
// A map is an unordered collection of KEY → VALUE pairs.
// Keys are unique — each key maps to exactly one value.
// Think of it like a dictionary, hash table, or object in other languages.
//
// Map syntax:
//   map[KeyType]ValueType
//
// Key rules:
//   - Keys must be a comparable type (string, int, bool — NOT slices or maps)
//   - Values can be any type
//   - Maps are REFERENCE types (like slices) — assigning copies the reference, not data
//   - Iteration order is RANDOM — Go deliberately randomizes it every run
//   - Reading a missing key returns the zero value — it does NOT panic

package main

import "fmt"

func main() {

	// ─── Creating maps ────────────────────────────────────────────────────────

	// Map literal — declare and initialize at once
	capitals := map[string]string{
		"USA":    "Washington D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	fmt.Println("Capitals:", capitals)

	// make() — creates an empty map ready to use
	// Always use make() or a literal — a nil map will panic on write
	ages := make(map[string]int)
	fmt.Println("Empty map:", ages) // map[]

	// var nilMap map[string]int  ← nil map, DO NOT write to it (panic!)
	// nilMap["Alice"] = 30       ← this would panic at runtime

	// ─── Adding and updating entries ──────────────────────────────────────────

	// Add new key-value pairs using assignment
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35
	fmt.Println("Ages:", ages)

	// Update an existing key — same syntax as adding
	ages["Bob"] = 26 // Bob's age updated from 25 to 26
	fmt.Println("Updated Bob:", ages["Bob"]) // 26

	// ─── Reading values ───────────────────────────────────────────────────────

	// Access a value by key
	fmt.Println("Alice's age:", ages["Alice"]) // 30

	// Reading a missing key returns the ZERO VALUE — no error, no panic
	fmt.Println("Missing key:", ages["Nobody"]) // 0 (zero value for int)

	// ─── Checking if a key exists ─────────────────────────────────────────────

	// The two-value form returns: (value, ok)
	// 'ok' is true if the key exists, false if it doesn't
	// This is the correct way to distinguish "key missing" from "value is zero"
	age, ok := ages["Alice"]
	fmt.Printf("Alice: age=%d, exists=%t\n", age, ok) // age=30, exists=true

	age, ok = ages["Nobody"]
	fmt.Printf("Nobody: age=%d, exists=%t\n", age, ok) // age=0, exists=false

	// Common pattern — check before using
	if val, exists := capitals["France"]; exists {
		fmt.Println("Capital of France:", val) // Paris
	} else {
		fmt.Println("France not found")
	}

	// ─── Deleting entries ─────────────────────────────────────────────────────

	// delete(map, key) removes the key-value pair
	// Deleting a non-existent key is safe — it does nothing
	fmt.Println("Before delete:", ages)
	delete(ages, "Charlie")
	fmt.Println("After delete Charlie:", ages)

	// ─── Iterating over a map ─────────────────────────────────────────────────

	// Use for-range — order is RANDOM every run
	fmt.Println("All capitals:")
	for country, capital := range capitals {
		fmt.Printf("  %s → %s\n", country, capital)
	}

	// ─── Map length ───────────────────────────────────────────────────────────

	fmt.Println("Number of capitals:", len(capitals)) // 3

	// ─── Maps are reference types ─────────────────────────────────────────────

	// Assigning a map to another variable makes them POINT to the same map
	// Changes through either variable affect the same underlying data
	original := map[string]int{"x": 1, "y": 2}
	reference := original   // NOT a copy — both point to the same map
	reference["x"] = 999

	fmt.Println("original x:", original["x"]) // 999 — changed through reference!

	// To make a true copy, copy key by key
	copyMap := make(map[string]int)
	for k, v := range original {
		copyMap[k] = v
	}
	copyMap["x"] = 1 // only affects copyMap
	fmt.Println("original x after copy modify:", original["x"]) // 999 — still unchanged

	// ─── Map with struct values ───────────────────────────────────────────────

	// Maps can hold structs as values — common for grouping related data
	type Student struct {
		Grade int
		Score float64
	}

	students := map[string]Student{
		"Alice": {Grade: 10, Score: 95.5},
		"Bob":   {Grade: 11, Score: 88.0},
	}

	fmt.Println("Alice's score:", students["Alice"].Score) // 95.5

	// To update a struct field in a map, replace the whole struct
	s := students["Bob"]
	s.Score = 90.0
	students["Bob"] = s
	fmt.Println("Bob's updated score:", students["Bob"].Score) // 90.0
}
