# 01 - Setting Up Go

## Step 1: Install Go

Download the installer from https://go.dev/dl/ and choose your OS:
- **macOS**: Download the `.pkg` file and run it
- **Windows**: Download the `.msi` file and run it
- **Linux**: Follow the tarball instructions on the download page

After installation, verify it worked:
```bash
go version
# Expected output: go version go1.2x.x darwin/amd64 (or your OS)
```

## Step 2: Understand the Go workspace

Go doesn't require a strict workspace anymore (since Go modules in v1.11).
You can create a project anywhere on your machine.

```
my-project/
├── go.mod       ← module definition file (like package.json in Node.js)
└── main.go      ← your Go source code
```

## Step 3: Create your first module

```bash
mkdir my-project
cd my-project
go mod init my-project    # creates go.mod
```

The `go.mod` file tracks the module name and Go version:
```
module my-project

go 1.21.0
```

## Step 4: Write and run code

Create `main.go`:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Run it:
```bash
go run main.go       # compile + run in one step (for development)
go build             # compile to a binary (for production)
./my-project         # run the compiled binary
```

## Key CLI commands

| Command | What it does |
|---------|-------------|
| `go run main.go` | Compile and run immediately |
| `go build` | Compile to a binary |
| `go fmt ./...` | Auto-format all Go files |
| `go vet ./...` | Check for common mistakes |
| `go mod tidy` | Remove unused dependencies |
| `go test ./...` | Run all tests |

## Recommended Editor

**VS Code** with the official **Go extension** (`golang.go`) gives you:
- Auto-complete
- Inline error checking
- Auto-format on save
- Jump to definition

## Go environment variables

```bash
go env GOPATH    # where Go tools are installed (~/.go or ~/go)
go env GOROOT    # where Go itself is installed
```

You rarely need to change these — Go sets them automatically.
