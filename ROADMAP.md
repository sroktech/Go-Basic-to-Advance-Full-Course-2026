# Roadmap — Beyond Lesson 36

This course currently takes a learner from zero Go knowledge through the **Beginner** tier
(00–26) and the concurrency/real-world **Intermediate** tier (27–36): goroutines, channels,
file I/O, JSON, HTTP, generics, and testing.

That is a solid, complete foundation — genuinely enough to build small real programs and
APIs. It is **not yet** enough to call yourself a professional backend Go engineer. This
document tracks the planned continuation, so the gap between "finished the course" and
"production-ready Go engineer" is visible and plannable, even before it's written.

This is a plan, not yet written content. Treat lesson numbers below as provisional.

---

## Phase 2 — Backend Fundamentals (Intermediate, continued)

Builds directly on lessons 32–34 (File I/O, JSON, HTTP) to turn "an HTTP handler" into
"a real backend service."

| Topic | Why it's needed |
|---|---|
| REST API design | Structuring routes/handlers beyond a single-file demo; request validation; pagination, filtering, versioning |
| Databases & SQL basics | Nearly every real backend needs persistent storage |
| PostgreSQL + `database/sql` / a driver (e.g. `pgx`) | The most common production relational database for Go services |
| Transactions | Correctness under concurrent writes; commit/rollback; isolation levels |
| Structured logging (`log/slog`) | Debugging production systems requires structured, searchable logs, not `fmt.Println` |
| Configuration management | Env vars, config files, secrets — twelve-factor app principles |
| Dependency management in practice | Go modules at project scale: versioning, replace directives, vendoring, private modules |

**Project for this phase:** a REST API backed by PostgreSQL (e.g. a task/todo API or a
simple blog API) with proper config, logging, and migrations.

---

## Phase 3 — Advanced Engineering

Where "it works" becomes "it's engineered well."

| Topic | Why it's needed |
|---|---|
| Advanced concurrency patterns | Worker pools, pipelines, fan-in/fan-out, rate limiting, backpressure |
| Clean architecture / layered design | Separating handlers, business logic, and storage so the codebase scales past one file |
| SOLID principles in a Go idiom | Which OOP principles translate to Go (and which don't — Go has no inheritance) |
| Dependency injection (Go-style) | Constructor injection, interfaces for testability — without a DI framework |
| Design patterns in Go | Which classic patterns are idiomatic in Go (functional options, strategy via interfaces) vs. which are anti-patterns here |
| Performance optimization & profiling | `pprof`, benchmarking, escape analysis, avoiding needless allocations |
| Memory management | How Go's GC works, when it matters, how to reason about it |
| Advanced testing | Mocking via interfaces, integration tests, fuzzing, test containers |
| Security | Input validation, secrets handling, common vulnerability classes in Go services |
| Observability | Metrics, tracing, structured logs working together (not just logging alone) |
| Distributed systems basics | Why single-node assumptions break, CAP-adjacent tradeoffs |
| Message queues | Async processing with something like Kafka/NATS/RabbitMQ |
| Microservices | Splitting a service, service boundaries, inter-service communication |
| Docker | Containerizing a Go service |
| CI/CD | Automating build/test/deploy for a Go project |

**Projects for this phase:** an authentication service; an e-commerce/order service; an
event-driven application; a small microservice split of an earlier monolith project.

---

## Phase 4 — Senior / Production-Level Go

Where correctness and cleanliness meet running real systems at scale, and reviewing others'
code.

| Topic | Why it's needed |
|---|---|
| Designing production-ready services from scratch | Everything above, applied together |
| Scalable API architecture | Handling growth in traffic and team size |
| High-concurrency systems | Beyond patterns — real load, real failure modes |
| Database architecture & optimization | Indexing, query performance, connection pooling at scale |
| Distributed transactions | Sagas, outbox pattern, eventual consistency |
| Event-driven architecture | Designing systems around events, not just requests |
| Resilience & fault tolerance | Retries, circuit breakers, timeouts, graceful degradation |
| Caching | Where, when, and how to cache safely |
| Monitoring & observability at scale | SLOs, alerting, on-call-readiness |
| Security architecture | Defense in depth, not just input sanitization |
| Performance & scalability | Capacity planning, load testing |
| Production troubleshooting | Debugging systems you can't attach a debugger to |
| System design | Designing a whole system on a whiteboard, tradeoffs and all |
| Code review & engineering standards | How to review Go code, what a senior reviewer looks for |

**Capstone project:** a production-shaped distributed system — e.g. a payment service or
order-processing platform composed of multiple services, a message queue, a database,
caching, observability, and a CI/CD pipeline.

---

## Project Ladder (across all phases)

Each project should be buildable with only what's been taught up to that point, and each
should introduce a bounded set of new concepts rather than everything at once:

1. CLI application *(buildable today, with lessons 00–26)*
2. Simple REST API *(buildable today, with lessons 27–36)*
3. REST API with PostgreSQL *(Phase 2)*
4. Authentication service *(Phase 3)*
5. E-commerce / order service *(Phase 3)*
6. Payment service *(Phase 3–4)*
7. Event-driven application *(Phase 3–4)*
8. Microservice system *(Phase 3–4)*
9. Production-ready distributed system *(Phase 4 capstone)*

---

## How this roadmap will be used

As each phase is written, its lessons will be added to this repository following the same
per-lesson structure used for lessons 00–36 (What You Will Learn, Why This Matters, Concept
Explanation, Simple Example, How It Works, Common Mistakes, Best Practices, Real-World
Example, Exercise, Mini Project, Summary, What to Learn Next), and this file will shrink as
topics move from "planned" to "written" in the main README.
