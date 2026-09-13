# 01 – Setting Up Go

## What You Will Learn
How to install Go on your computer, what a Go "module" is, and how to create and run your very first Go program from the command line.

## Why This Matters
You can't write Go without a working Go installation, and you can't run a Go program without understanding modules — the way Go organizes and names projects. Getting this right at the start avoids confusing errors later (like "no go.mod file found") that have nothing to do with your actual code.

## Concept Explanation
Installing Go gives you the `go` command-line tool, which does far more than just run your code — it compiles, formats, tests, and manages dependencies for your project. See [setup.md](setup.md) in this folder for the full step-by-step install instructions per operating system, the exact commands, and a table of the most useful `go` CLI commands.

A couple of ideas worth calling out explicitly:

- **Module** — since Go 1.11, every Go project is a "module": a folder containing a `go.mod` file that names the project and records which version of Go it targets. Think of `go.mod` as similar to `package.json` in Node.js — it's the file that makes a folder into a real, buildable project rather than just loose files.
- **`go run` vs `go build`** — `go run main.go` compiles your code into a temporary binary and immediately executes it, which is convenient while developing. `go build` instead compiles it into a permanent, standalone binary file that you (or anyone else) can run later without needing Go installed at all — this is what you'd actually ship to a server or a user.

## Common Mistakes
- Trying to `go run` a file inside a folder that has no `go.mod` yet — Go will complain because it doesn't know what module the file belongs to. Always run `go mod init <name>` first.
- Forgetting to restart your terminal (or editor) after installing Go, so the `go` command isn't found yet.
- Confusing `GOPATH` and `GOROOT` — `GOROOT` is where Go itself lives, `GOPATH` is where Go tooling and downloaded packages are cached. As a beginner you will almost never need to touch either manually.

## Best Practices
- Use VS Code with the official Go extension (`golang.go`) — it gives you inline error checking and auto-formatting on save, which catches mistakes before you even try to run the code.
- Run `go version` right after installing to confirm everything worked before moving on.
- Get in the habit of running `go fmt ./...` — Go has one official formatting style, and the tool applies it automatically, so there's never a debate about code style on a team.

## Exercise
Install Go from https://go.dev/dl/ for your operating system, then open a terminal and run:
```bash
go version
```
Confirm it prints a version number (e.g., `go version go1.2x.x darwin/arm64`) instead of a "command not found" error.

## Mini Project
Create your first Go module and program from scratch:
```bash
mkdir hello-go
cd hello-go
go mod init hello-go
```
Then create a file named `main.go` inside that folder with this content:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```
Run it with:
```bash
go run main.go
```
You should see `Hello, Go!` printed to your terminal. That's your first working Go program.

## Summary
Go is installed as a single toolchain that gives you the `go` command for running, building, formatting, and testing code. Every project lives inside a module, defined by a `go.mod` file, and `go run` is your go-to command while learning and experimenting.

## What to Learn Next
Now that Go is installed and you've run your first program, it's time to learn the building blocks of a Go source file itself: packages, imports, and printing.

Prev: [../00Introduction/README.md](../00Introduction/README.md)
Next: [../02Basic_Syntax/README.md](../02Basic_Syntax/README.md)
