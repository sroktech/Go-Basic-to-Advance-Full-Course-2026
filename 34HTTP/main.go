// 34 - HTTP Server & Client in Go
//
// Go's standard library 'net/http' has everything needed to build production HTTP servers.
// No framework required — but frameworks like Gin, Echo, Chi add convenience.
//
// FLOW — HTTP Server:
//
//   Register handlers        Start server          Handle request
//   ─────────────────        ────────────          ──────────────
//   http.HandleFunc(         http.ListenAndServe   w http.ResponseWriter ← write response
//     "/path", handler)  ──► (":8080", nil)   ──►  r *http.Request      ← read request
//
// FLOW — HTTP Client:
//
//   http.Get(url) ──► *http.Response ──► io.ReadAll(resp.Body) ──► []byte
//                           │
//                     resp.Body.Close()  ← always defer this!
//
// HTTP methods:
//   GET    — read data (no body)
//   POST   — create data (with body)
//   PUT    — replace data (with body)
//   PATCH  — update part of data
//   DELETE — delete data

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// ─── Data model ───────────────────────────────────────────────────────────────

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// In-memory "database" — in production, use a real DB
var tasks = []Task{
	{ID: 1, Title: "Learn Go basics", Done: true},
	{ID: 2, Title: "Learn goroutines", Done: false},
	{ID: 3, Title: "Build an HTTP API", Done: false},
}

// ─── Handler helpers ─────────────────────────────────────────────────────────

// writeJSON sends a JSON response with the correct Content-Type header
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ─── Route handlers ───────────────────────────────────────────────────────────

// GET /tasks — return all tasks
func handleGetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	writeJSON(w, http.StatusOK, tasks)
}

// POST /tasks — create a new task
func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	writeJSON(w, http.StatusCreated, newTask)
}

// GET /health — simple health check endpoint
func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

// ─── Middleware ───────────────────────────────────────────────────────────────
//
// Middleware wraps a handler to add behavior (logging, auth, CORS, etc.)
// FLOW:
//   request ──► middleware ──► actual handler ──► response
//                  │
//                  └── runs before AND after the handler

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("[%s] %s %s", time.Now().Format("15:04:05"), r.Method, r.URL.Path)
		next(w, r) // call the actual handler
		fmt.Printf(" (%v)\n", time.Since(start))
	}
}

func main() {

	// ─── Register routes ─────────────────────────────────────────────────────

	mux := http.NewServeMux() // custom mux — avoids polluting DefaultServeMux

	// Wrap handlers with logging middleware
	mux.HandleFunc("/tasks", loggingMiddleware(handleGetTasks))
	mux.HandleFunc("/tasks/create", loggingMiddleware(handleCreateTask))
	mux.HandleFunc("/health", loggingMiddleware(handleHealth))

	// ─── Start server in background (so we can demo the client too) ───────────

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		fmt.Println("Server listening on http://localhost:8080")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal("Server error:", err)
		}
	}()

	// Give server time to start
	time.Sleep(100 * time.Millisecond)

	// ─── HTTP Client demos ────────────────────────────────────────────────────
	//
	// FLOW:
	//   client.Do(req) ──► check status ──► read body ──► parse JSON

	client := &http.Client{Timeout: 5 * time.Second}

	// GET /health
	fmt.Println("\n--- GET /health ---")
	resp, err := client.Get("http://localhost:8080/health")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Body:", string(body))

	// GET /tasks
	fmt.Println("\n--- GET /tasks ---")
	resp2, _ := client.Get("http://localhost:8080/tasks")
	defer resp2.Body.Close()
	var fetchedTasks []Task
	json.NewDecoder(resp2.Body).Decode(&fetchedTasks)
	for _, t := range fetchedTasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("  %s %d. %s\n", status, t.ID, t.Title)
	}

	// POST /tasks/create
	fmt.Println("\n--- POST /tasks/create ---")
	payload := `{"title": "Learn HTTP in Go", "done": false}`
	resp3, _ := client.Post(
		"http://localhost:8080/tasks/create",
		"application/json",
		strings.NewReader(payload),
	)
	defer resp3.Body.Close()
	var created Task
	json.NewDecoder(resp3.Body).Decode(&created)
	fmt.Printf("  Created: ID=%d, Title=%s\n", created.ID, created.Title)

	// ─── Request with custom headers ──────────────────────────────────────────

	fmt.Println("\n--- Custom headers ---")
	req, _ := http.NewRequest("GET", "http://localhost:8080/health", nil)
	req.Header.Set("X-Request-ID", "demo-001")
	req.Header.Set("Accept", "application/json")

	resp4, _ := client.Do(req)
	defer resp4.Body.Close()
	fmt.Println("  X-Request-ID sent, status:", resp4.Status)

	// ─── Shut down gracefully ─────────────────────────────────────────────────
	server.Close()
	fmt.Println("\nServer stopped.")

	// ─── Quick reference ─────────────────────────────────────────────────────
	fmt.Println(`
HTTP Status codes:
  200 OK           — success (GET, PUT, PATCH)
  201 Created      — new resource created (POST)
  400 Bad Request  — client sent invalid data
  401 Unauthorized — not authenticated
  403 Forbidden    — authenticated but not allowed
  404 Not Found    — resource doesn't exist
  500 Internal Server Error — server bug`)
}
