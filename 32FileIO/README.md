# 32 – File I/O

## What You Will Learn
- Writing files with `os.Create`, `os.WriteFile`, and `os.OpenFile`
- Reading files whole (`os.ReadFile`) versus streaming line-by-line (`bufio.Scanner`)
- Appending to an existing file safely
- Checking whether a file exists before treating an error as something else
- File permission bits (`0644`, `0755`) and what they mean on Unix systems

## Why This Matters
Every real program eventually touches the filesystem: config files, logs, CSV exports, uploaded documents, cache files. Go deliberately does not hide this behind magic — you open a handle, you read or write through it, and you close it yourself. That explicitness is a feature, not friction: it means you always know exactly when a file is opened, how much memory a read will use, and when data actually hits disk. Frameworks and higher-level languages that "just work" with files often hide resource leaks and silent truncation until they show up in production under load.

## Concept Explanation
Go splits file work across two packages. `os` gives you the low-level primitives — `Open`, `Create`, `OpenFile`, `Stat`, `Remove` — that map closely to the underlying system calls and return an `*os.File`, which implements `io.Reader` and `io.Writer`. `bufio` wraps a reader or writer to make small, frequent operations (like reading one line at a time) efficient by batching the actual system calls internally.

The two directions you need are:

```
os.Open(path) ──► *os.File ──► bufio.NewScanner ──► scan line by line
os.Create(path) ──► *os.File ──► fmt.Fprintln / file.Write
```

`os.Create` always creates a **new, empty** file — if one already exists at that path, its contents are truncated. `os.OpenFile` gives you full control via flags (`os.O_APPEND`, `os.O_WRONLY`, `os.O_CREATE`) when you need something other than "start fresh" or "just read."

Reading has the same two-tier choice as writing: `os.ReadFile` loads everything into a `[]byte` in one call — simple, but the whole file lives in memory at once. `bufio.NewScanner` reads incrementally, which matters once a file is bigger than you'd want to hold in RAM (logs, large exports).

## Simple Example
Writing lines and then reading them back line by line ([main.go](main.go)):

```go
file, err := os.Create("example.txt")
if err != nil {
    fmt.Println("Error creating file:", err)
    return
}
defer file.Close()

for _, line := range lines {
    fmt.Fprintln(file, line)
}
```

```go
f, err := os.Open("example.txt")
...
defer f.Close()

scanner := bufio.NewScanner(f)
for scanner.Scan() {
    line := scanner.Text()
    ...
}
if err := scanner.Err(); err != nil {
    fmt.Println("Scanner error:", err)
}
```

## How It Works
`os.Create` returns a file handle and an error — check the error before touching the handle. `defer file.Close()` runs when `main` returns, guaranteeing the file is closed even if a later line returns early. `fmt.Fprintln` writes to any `io.Writer`, so the same call works whether the destination is a file, a buffer, or the network. The scanner's `Scan()` method advances one line at a time and returns `false` once it hits EOF or an error — which is why the loop condition alone can't tell you *why* it stopped; you must call `scanner.Err()` afterward to distinguish "clean EOF" from "read failure." The file also demonstrates checking existence with `os.Stat` plus `os.IsNotExist(err)`, rather than assuming any error from `Stat` means "does not exist" — a `Stat` call can also fail due to permissions or a bad path, which is a different problem entirely.

## Common Mistakes
- Forgetting `defer f.Close()` — leaked file descriptors accumulate and can eventually exhaust the OS limit in a long-running process.
- Ignoring the error returned by `Close()`. On a write-heavy path, a failed `Close()` can mean buffered data never made it to disk — silently dropping it is a real bug in production code.
- Treating *any* error from `os.Open`/`os.Stat` as "file doesn't exist" instead of checking `os.IsNotExist(err)` specifically — permission errors and disk errors look like generic errors too.
- Loading an entire large file into memory with `os.ReadFile` when only a stream (`bufio.Scanner`/`bufio.Reader`) is actually needed.

## Best Practices
- Always check the error from every `os` call — `Create`, `Open`, `OpenFile`, `WriteFile`, `ReadFile` can all fail (missing directory, permissions, disk full).
- Pair every successful open with `defer f.Close()` immediately, before anything else can go wrong.
- Choose the read strategy based on file size expectations: whole-file reads for small, bounded files (configs); scanners/streams for anything that could grow (logs, user uploads).
- Use `os.WriteFile`/`os.ReadFile` for simple one-shot cases where you don't need fine-grained control — they handle open/write-or-read/close for you correctly.

## Real-World Example
Configuration loading is the most common production use of file I/O: a service reads a `config.yaml` or `.env`-style file at startup with `os.ReadFile`, parses it, and fails fast with a clear error if the file is missing or malformed — much better than a service that starts and then fails mysteriously later. Log rotation and CSV/report exports use the append and streaming patterns shown here directly.

## Exercise
Write a small program that reads a text file line by line and counts how many lines contain a given word (case-insensitive), printing the total. Use only `bufio.Scanner` and `strings` — no external libraries.

## Mini Project
Build a tiny file-backed key-value store: `Set(key, value string)` appends a `key=value` line to a file, and `Get(key string)` scans the file and returns the value of the *last* matching key (so later writes override earlier ones). Test it by setting a few keys, overwriting one, and retrieving each.

## What to Learn Next
Continue to [33 – JSON](../33JSON/README.md), where you'll take the raw bytes you now know how to read and write and turn them into structured Go data — the format almost every API and config file actually uses.
