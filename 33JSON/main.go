// 33 - JSON in Go
//
// JSON (JavaScript Object Notation) is the most common data format for APIs.
// Go's 'encoding/json' package handles encoding (Go → JSON) and decoding (JSON → Go).
//
// FLOW — encoding (Marshal):
//
//   Go struct/map ──► json.Marshal() ──► []byte (JSON string)
//
// FLOW — decoding (Unmarshal):
//
//   []byte (JSON string) ──► json.Unmarshal() ──► Go struct/map
//
// Struct tags control how fields map to JSON keys:
//   `json:"name"`           — use "name" as the JSON key (not the field name)
//   `json:"name,omitempty"` — omit the field if it's the zero value
//   `json:"-"`              — always omit this field from JSON
//
// Key rules:
//   - Only EXPORTED fields (uppercase) are included in JSON
//   - Lowercase fields are silently ignored

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ─── Structs with JSON tags ───────────────────────────────────────────────────

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	Zip    string `json:"zip"`
}

type Person struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Email    string  `json:"email,omitempty"` // omit if empty string
	Password string  `json:"-"`               // NEVER include in JSON
	Address  Address `json:"address"`
	Tags     []string `json:"tags,omitempty"`  // omit if nil/empty
}

// ─── Custom JSON marshaling ───────────────────────────────────────────────────

// Temperature wraps a float and always marshals as Celsius string
type Temperature struct {
	Celsius float64
}

func (t Temperature) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%.1f°C", t.Celsius))
}

func (t *Temperature) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	_, err := fmt.Sscanf(strings.TrimSuffix(s, "°C"), "%f", &t.Celsius)
	return err
}

func main() {

	// ─── 1. Encoding: Go struct → JSON ────────────────────────────────────────

	fmt.Println("=== 1. Marshal (struct → JSON) ===")

	person := Person{
		Name:     "Alice",
		Age:      30,
		Email:    "alice@example.com",
		Password: "secret123",  // this will NOT appear in JSON (json:"-")
		Address: Address{
			Street: "123 Main St",
			City:   "Springfield",
			Zip:    "12345",
		},
		Tags: []string{"go", "programming"},
	}

	// json.Marshal — compact JSON (no indentation)
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Compact:", string(data))

	// json.MarshalIndent — pretty-printed JSON
	pretty, _ := json.MarshalIndent(person, "", "  ")
	fmt.Println("Pretty:")
	fmt.Println(string(pretty))

	// Person with zero-value Email — omitempty kicks in
	personNoEmail := Person{Name: "Bob", Age: 25, Address: Address{City: "Austin"}}
	compactNoEmail, _ := json.Marshal(personNoEmail)
	fmt.Println("No email (omitempty):", string(compactNoEmail))

	// ─── 2. Decoding: JSON → Go struct ────────────────────────────────────────

	fmt.Println("\n=== 2. Unmarshal (JSON → struct) ===")

	jsonInput := `{
		"name": "Charlie",
		"age": 35,
		"email": "charlie@example.com",
		"address": {
			"street": "456 Oak Ave",
			"city": "Portland",
			"zip": "97201"
		},
		"tags": ["backend", "devops"],
		"unknownField": "ignored silently"
	}`

	var decoded Person
	err = json.Unmarshal([]byte(jsonInput), &decoded)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Name: %s, Age: %d, City: %s\n", decoded.Name, decoded.Age, decoded.Address.City)
	fmt.Printf("Tags: %v\n", decoded.Tags)
	// unknownField is silently ignored — Go won't error on extra JSON fields

	// ─── 3. JSON array ────────────────────────────────────────────────────────

	fmt.Println("\n=== 3. JSON array ===")

	jsonArray := `[
		{"name": "Alice", "age": 30},
		{"name": "Bob",   "age": 25},
		{"name": "Carol", "age": 28}
	]`

	var people []Person
	json.Unmarshal([]byte(jsonArray), &people)
	for _, p := range people {
		fmt.Printf("  %s (age %d)\n", p.Name, p.Age)
	}

	// ─── 4. map[string]interface{} — dynamic/unknown JSON ─────────────────────
	//
	// When you don't know the JSON shape ahead of time, decode into a map.
	// All values become interface{} — you must type-assert to use them.

	fmt.Println("\n=== 4. Dynamic JSON with map ===")

	dynamicJSON := `{"event": "click", "x": 42.5, "y": 100.0, "active": true}`

	var dynamic map[string]interface{}
	json.Unmarshal([]byte(dynamicJSON), &dynamic)

	for key, val := range dynamic {
		fmt.Printf("  %s: %v (%T)\n", key, val, val)
	}
	// Numbers decode as float64 by default
	x := dynamic["x"].(float64)
	fmt.Printf("  x as float64: %.1f\n", x)

	// ─── 5. Streaming: json.Encoder / json.Decoder ────────────────────────────
	//
	// For large data sets, stream JSON rather than loading it all into memory.
	// Encoder/Decoder work with any io.Writer/io.Reader (files, network, etc.)

	fmt.Println("\n=== 5. json.Encoder (stream to writer) ===")

	var sb strings.Builder
	encoder := json.NewEncoder(&sb)
	encoder.SetIndent("", "  ")

	records := []Person{
		{Name: "Dave", Age: 40},
		{Name: "Eve", Age: 22},
	}
	for _, r := range records {
		encoder.Encode(r) // writes one JSON object per line
	}
	fmt.Print(sb.String())

	// ─── 6. Custom marshal / unmarshal ───────────────────────────────────────

	fmt.Println("=== 6. Custom marshaling ===")

	temp := Temperature{Celsius: 36.6}
	tempJSON, _ := json.Marshal(temp)
	fmt.Println("Marshaled:", string(tempJSON)) // "36.6°C"

	var temp2 Temperature
	json.Unmarshal(tempJSON, &temp2)
	fmt.Printf("Unmarshaled: %.1f°C\n", temp2.Celsius)
}
