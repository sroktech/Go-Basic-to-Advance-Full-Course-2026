// 32 - File I/O in Go
//
// Go's 'os' and 'bufio' packages handle file reading and writing.
// Always close files after opening — use defer to guarantee it.
//
// FLOW — read a file:
//
//   os.Open(path) ──► *os.File ──► bufio.NewScanner ──► scan line by line
//        │                                                      │
//        └─► error if file missing                    file.Close() via defer
//
// FLOW — write a file:
//
//   os.Create(path) ──► *os.File ──► fmt.Fprintln / file.Write
//        │                                  │
//        └─► creates or truncates           └─► file.Close() via defer
//
// File permission bits (Unix octal):
//   0644 — owner read/write, group read, others read (common for data files)
//   0755 — owner read/write/execute, others read/execute (for executables)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// ─── 1. Write a file ─────────────────────────────────────────────────────

	fmt.Println("=== 1. Write file ===")

	// os.Create creates or TRUNCATES the file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // always close, even if we return early

	// Write using fmt.Fprintln — writes to any io.Writer (file, buffer, etc.)
	lines := []string{
		"Hello, Go File I/O!",
		"Line 2: Learning Go in 2026",
		"Line 3: Reading and writing files",
		"Line 4: bufio for efficiency",
		"Line 5: always close your files",
	}
	for _, line := range lines {
		fmt.Fprintln(file, line)
	}
	fmt.Println("  wrote", len(lines), "lines to example.txt")

	// ─── 2. Read entire file at once ──────────────────────────────────────────

	fmt.Println("\n=== 2. Read entire file (os.ReadFile) ===")

	// os.ReadFile reads the whole file into memory — simple but uses more memory
	// Good for small files, avoid for large files
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("  file size: %d bytes\n", len(data))
	fmt.Println("  content:")
	fmt.Println(string(data))

	// ─── 3. Read line by line with bufio.Scanner ──────────────────────────────
	//
	// FLOW:
	//   os.Open ──► bufio.NewScanner ──► scanner.Scan() [loop] ──► scanner.Text()
	//
	// bufio.Scanner reads efficiently in chunks — good for large files

	fmt.Println("=== 3. Read line by line (bufio.Scanner) ===")

	f, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNum := 0
	for scanner.Scan() { // Scan() returns false at EOF or error
		lineNum++
		line := scanner.Text() // current line (without \n)
		fmt.Printf("  line %d: %s\n", lineNum, line)
	}
	if err := scanner.Err(); err != nil { // check for read errors
		fmt.Println("Scanner error:", err)
	}

	// ─── 4. Append to a file ─────────────────────────────────────────────────

	fmt.Println("\n=== 4. Append to file ===")

	// os.OpenFile with flags:
	//   os.O_APPEND — seek to end before each write
	//   os.O_WRONLY — write-only access
	//   os.O_CREATE — create if not exists
	appendFile, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer appendFile.Close()

	fmt.Fprintln(appendFile, "Line 6: appended later!")
	fmt.Println("  appended a new line")

	// ─── 5. Write with os.WriteFile ───────────────────────────────────────────

	fmt.Println("\n=== 5. os.WriteFile (simple one-shot write) ===")

	content := "This is a one-shot file write.\nNo need to open/close manually.\n"
	err = os.WriteFile("simple.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("  wrote simple.txt")

	// ─── 6. Check if file exists ──────────────────────────────────────────────

	fmt.Println("\n=== 6. Check file existence ===")

	checkFile := func(path string) {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			fmt.Printf("  %s: does NOT exist\n", path)
		} else if err != nil {
			fmt.Printf("  %s: error checking: %v\n", path, err)
		} else {
			fmt.Printf("  %s: EXISTS\n", path)
		}
	}

	checkFile("example.txt")
	checkFile("nonexistent.txt")

	// ─── 7. Read CSV-style data (split lines) ────────────────────────────────

	fmt.Println("\n=== 7. Parse structured text ===")

	csvData := "Alice,30,Engineer\nBob,25,Designer\nCharlie,35,Manager\n"
	err = os.WriteFile("data.csv", []byte(csvData), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	csvFile, _ := os.Open("data.csv")
	defer csvFile.Close()

	csvScanner := bufio.NewScanner(csvFile)
	fmt.Printf("  %-10s %-5s %-10s\n", "Name", "Age", "Role")
	fmt.Println(" ", strings.Repeat("-", 28))
	for csvScanner.Scan() {
		parts := strings.Split(csvScanner.Text(), ",")
		if len(parts) == 3 {
			fmt.Printf("  %-10s %-5s %-10s\n", parts[0], parts[1], parts[2])
		}
	}

	// ─── Cleanup ──────────────────────────────────────────────────────────────
	os.Remove("example.txt")
	os.Remove("simple.txt")
	os.Remove("data.csv")
	fmt.Println("\n  cleaned up temp files")
}
