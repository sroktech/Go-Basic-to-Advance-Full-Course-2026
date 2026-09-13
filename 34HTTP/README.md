# 34 – HTTP Server & Client

## What You Will Learn
- Building an HTTP server with `net/http`, `http.ServeMux`, and handler functions
- Writing JSON responses and decoding JSON request bodies
- Writing middleware (a function that wraps a handler)
- Making HTTP client requests with `http.Client`, `client.Get`, `client.Post`, and `client.Do`
- Setting timeouts and always checking errors from network calls
- Graceful server shutdown with `server.Close()`

## Why This Matters
`net/http` is a full, production-capable HTTP toolkit built into the standard library — no framework is required to build a real API. Understanding it directly (rather than only ever going through Gin or Echo) means you understand what those frameworks are actually built on top of: a `Handler` interface, a `ServeMux` for routing, and `http.Client`/`http.Server` for the network boundary itself. A framework mainly adds convenience — route parameters and groups, built-in middleware chains, request binding/validation helpers, and often better ergonomics for large route tables — but it doesn't replace understanding what happens to a request underneath.

The other reason this lesson matters more than most: network calls are the single most common source of unhandled runtime failures in real backend systems, because a network call can fail in ways local code rarely does — timeouts, connection refused, DNS failure, a dropped connection mid-response. Code that looks fine in a demo because "the server is right there on localhost" behaves completely differently against a real, unreliable network.

## Concept Explanation
A server has two phases: register handlers, then run the loop that dispatches incoming requests to them.

```
Register handlers        Start server          Handle request
─────────────────        ────────────          ──────────────
http.HandleFunc(         http.ListenAndServe   w http.ResponseWriter ← write response
  "/path", handler)  ──► (":8080", nil)   ──►  r *http.Request      ← read request
```

A client call is the mirror image — send a request, get a response, and the response body must be read (and closed) explicitly:

```
http.Get(url) ──► *http.Response ──► io.ReadAll(resp.Body) ──► []byte
                        │
                  resp.Body.Close()  ← always defer this!
```

Middleware is just a function that takes a handler and returns a new handler that does something before and/or after calling the original — no special framework mechanism needed, because functions are values in Go.

## Simple Example
A handler and a route registration from [main.go](main.go):

```go
func handleHealth(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusOK, map[string]string{
        "status": "ok",
        "time":   time.Now().Format(time.RFC3339),
    })
}
```

```go
mux := http.NewServeMux()
mux.HandleFunc("/health", loggingMiddleware(handleHealth))
```

A client call, with the error checked immediately:

```go
resp, err := client.Get("http://localhost:8080/health")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
```

## How It Works
`http.NewServeMux()` creates a router that maps URL paths to handler functions, kept separate from the package-level `http.DefaultServeMux` so this program doesn't leak routes into anything else that might also register on the default mux. Each handler receives an `http.ResponseWriter` (to write the response) and an `*http.Request` (to read the incoming request); `writeJSON` is a small helper that sets the `Content-Type` header, writes the status code, and encodes the response body as JSON in one place so every handler shares the same response shape.

The server itself runs inside a goroutine (`go func() { ... }()`) so `main` can continue and act as an HTTP *client* against its own server in the same program — this file was recently fixed for two related, important reasons:

1. **Every error from the HTTP client calls is now checked.** Earlier, some of these calls discarded their error, which is dangerous even in a demo: a nil check that never happens means a nil `*http.Response` or a partially-read body can be used as if the call had succeeded, causing a confusing panic or wrong output far from the real cause.
2. **`log.Fatal` is never called from inside the background goroutine anymore.** `log.Fatal` calls `os.Exit`, which terminates the *entire process* immediately — from inside a goroutine, that means `main()` gets no chance to run any cleanup, close resources, or even print a useful message, because the whole program just vanishes. The code now uses `log.Printf` inside the goroutine to report the error and let that goroutine end normally, while `main()` keeps control and can call `server.Close()` on its own terms.

This is worth internalizing: `log.Fatal`/`os.Exit` should only ever be called from `main()` (or made conditional very deliberately), never from a spawned goroutine, because a goroutine has no way to signal "something went wrong, please shut down cleanly" — it can only kill everything immediately or report and continue.

## Common Mistakes
- Discarding the error from `client.Get`, `client.Post`, or `client.Do` — a network call can fail for many reasons (connection refused, DNS failure, timeout), and code that skips the check will panic on a nil response instead of failing with a clear message.
- Forgetting `defer resp.Body.Close()` after every successful response — this leaks the underlying TCP connection and can exhaust the client's connection pool under load.
- Using `http.Get`/`http.Post` (the package-level helpers) or a bare `&http.Client{}` with no `Timeout` set — a hung server or dead connection can then block forever with no way to recover.
- Calling `log.Fatal` (or otherwise calling `os.Exit`) from inside a goroutine that isn't `main` — it kills the whole program with no chance for anything else to clean up.

## Best Practices
- Always check the error from every HTTP client call, exactly like every other error in Go — a demo habit of skipping "because it's just a tutorial" is a habit that carries straight into production code.
- Always `defer resp.Body.Close()` immediately after confirming `err == nil`.
- Construct `http.Client` with an explicit `Timeout` (as this file does: `&http.Client{Timeout: 5 * time.Second}`), and set `ReadTimeout`/`WriteTimeout` on `http.Server` too — both protect against a hung peer.
- Keep `log.Fatal`/`os.Exit` calls in `main()` only; inside any other goroutine, report the error (`log.Printf`) and let the caller decide what to do.

## Real-World Example
This is essentially a miniature version of any backend service: a JSON REST API (`/tasks`, `/health`) with middleware for logging (in production: also auth, CORS, rate limiting), served over `net/http` directly or via a thin framework layer, consumed by other services or a frontend using exactly the same `http.Client` patterns shown here — with the same required discipline around timeouts, error checks, and closing response bodies.

## Exercise
Add a `DELETE /tasks/{id}`-style handler (you can parse the ID from a query parameter, e.g. `/tasks/delete?id=2`, since this lesson doesn't cover path parameters) that removes a task from the in-memory slice and returns the updated list as JSON. Make sure it checks the HTTP method and returns `404` if the ID doesn't exist.

## Mini Project
Build a single-endpoint REST API for a simple "notes" resource: `GET /notes` returns all notes as JSON, `POST /notes` accepts a JSON body and appends a new note (in-memory, no database yet). Write a small client function in the same program that calls both endpoints using `http.Client` with a timeout, checks every error, and closes every response body — then print the results.

## What to Learn Next
Continue to [35 – Generics](../35Generics/README.md), where you'll learn how to write utility code (like `Map`/`Filter`/`Contains`) once, for any type, instead of duplicating it or falling back to untyped `interface{}`.
