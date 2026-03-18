# Go 30-Day Mastery Plan for Shawon

**Purpose:** Build real backend intuition in Go by moving from language fundamentals to a small wallet/remittance-style backend project.

**How to use this plan**
- Spend **1.5 to 2.5 hours per day**.
- For each day, do the lesson, write code, and then write **3-5 lines of notes** on what felt confusing.
- Keep all code in your `go-mastery` repo.
- Create one folder per day if helpful, or keep evolving the same project depending on the day.

---

## Week 1 — Core Go Thinking

### Day 01 — Go Setup, Modules, and Program Execution
**Goal**
Understand how Go programs are structured and executed.

**You will learn**
- `package main`
- `func main()`
- `go run` vs `go build`
- `go mod init`
- why Go modules matter

**Build**
- Create a simple CLI app that prints:
  - app name
  - current learning day
  - a sample wallet balance

**Suggested files**
- `main.go`
- `go.mod`

**Focus point**
You are not just learning syntax. You are learning how a compiled backend language organizes projects.

**Checkpoint**
At the end of the day, you should be able to explain:
- what `go.mod` does
- why `main` is special
- why `go build` produces a binary

---

### Day 02 — Wallet Struct with Pointer Receivers
**Goal**
Understand state mutation in Go.

**You will learn**
- structs
- pointer receivers
- value receivers
- zero-value safety
- basic encapsulation

**Build**
Create a `Wallet` with:
- private `balance`
- `Deposit(amount int)`
- `Balance() int`
- optional `Withdraw(amount int) error`

**Suggested files**
- `wallet.go`
- `main.go`

**Why this matters**
This is the foundation of all backend domain modeling. Wallets, accounts, carts, ledgers, users — they all depend on correct state mutation.

**Checkpoint**
You should be able to explain:
- why `*Wallet` is needed in `Deposit`
- why `Balance()` can use a value receiver
- how Go automatically takes the address when calling pointer receiver methods on a value

---

### Day 03 — Methods, Constructors, and Encapsulation
**Goal**
Learn how to make your types safer and harder to misuse.

**You will learn**
- exported vs unexported fields
- constructor functions like `NewWallet()`
- enforcing invariants
- keeping unsafe writes out of `main`

**Build**
Refactor wallet creation so that code uses:
- `NewWallet()`
- `Deposit()`
- `Withdraw()`
- `Balance()`

**Suggested files**
- `wallet.go`
- `main.go`

**Exercise**
Prevent negative deposits and negative withdrawals.

**Checkpoint**
You should be able to explain why hiding `balance` is good design.

---

### Day 04 — Arrays, Slices, and Transaction History
**Goal**
Understand how Go collections work in memory.

**You will learn**
- arrays vs slices
- length vs capacity
- `append`
- slice reallocation behavior

**Build**
Add transaction history to wallet:
- store deposits and withdrawals in a slice
- print all transactions at the end

**Suggested files**
- `wallet.go`
- `transaction.go`
- `main.go`

**Exercise**
Print:
- total transaction count
- latest transaction
- all transaction history

**Checkpoint**
You should understand that a slice is a small descriptor over underlying storage, not the raw array itself.

---

### Day 05 — Maps and In-Memory Data Modeling
**Goal**
Use maps to model real backend state.

**You will learn**
- map creation
- existence checks
- map lookup patterns
- storing pointers in maps

**Build**
Create an in-memory wallet registry:
- `map[string]*Wallet`
- keys as user IDs

**Exercise**
- create 3 wallets
- deposit into 2 wallets
- look up one wallet by ID
- handle missing wallet safely

**Checkpoint**
You should be able to explain when to use a map instead of a slice.

---

### Day 06 — Interfaces and Dependency Inversion
**Goal**
Learn how Go achieves abstraction.

**You will learn**
- interface definition
- implicit implementation
- depending on behavior, not concrete types

**Build**
Create a storage interface:

```go
 type WalletStore interface {
     Save(id string, wallet *Wallet) error
     Get(id string) (*Wallet, error)
 }
```

Then implement:
- `MemoryWalletStore`

**Suggested files**
- `store.go`
- `memory_store.go`
- `wallet.go`
- `main.go`

**Why this matters**
This is the first step toward replaceable infrastructure.

**Checkpoint**
You should understand why interfaces help testing and future database integration.

---

### Day 07 — Mini Project 1: CLI Wallet Manager
**Goal**
Consolidate Week 1 into a small but meaningful app.

**Build**
A CLI app that can:
- create wallet
- deposit
- withdraw
- print balance
- print transactions

**Stretch goal**
Use command-line arguments or a small menu loop.

**Checkpoint**
By the end of Week 1, you should feel comfortable with structs, methods, slices, maps, and interfaces.

---

## Week 2 — Real Backend Foundations

### Day 08 — Error Handling Like a Go Developer
**Goal**
Learn idiomatic error handling.

**You will learn**
- `error` type
- `errors.New`
- `fmt.Errorf`
- wrapped errors
- returning early

**Build**
Create meaningful domain errors:
- insufficient balance
- invalid amount
- wallet not found

**Exercise**
Return clear errors from `Withdraw`, `Deposit`, and store lookups.

**Checkpoint**
You should stop thinking in exceptions and start thinking in explicit control flow.

---

### Day 09 — Project Structure and Package Layout
**Goal**
Move from toy code to organized code.

**You will learn**
- package boundaries
- keeping `main` thin
- separating domain from app startup

**Refactor into**
- `cmd/`
- `internal/wallet/`
- `internal/store/`

**Example target structure**
```text
cmd/
  app/
    main.go
internal/
  wallet/
    wallet.go
    transaction.go
  store/
    memory_store.go
```

**Checkpoint**
You should understand why backend projects should not put everything in `main.go`.

---

### Day 10 — JSON and DTO Thinking
**Goal**
Prepare your app for HTTP APIs.

**You will learn**
- struct tags
- marshaling and unmarshaling
- separating domain model from response model

**Build**
Create JSON-friendly response structs for:
- wallet response
- transaction response

**Exercise**
Convert wallet state to JSON and print it.

**Checkpoint**
You should understand that API output shape and internal domain shape do not always have to match.

---

### Day 11 — HTTP Server Basics
**Goal**
Enter backend runtime land.

**You will learn**
- `net/http`
- handlers
- response writer
- request object
- status codes

**Build**
Create endpoints:
- `GET /health`
- `GET /hello`

**Checkpoint**
You should be able to start a Go HTTP server and hit it from browser or curl.

---

### Day 12 — Building RESTful Wallet Endpoints
**Goal**
Turn your wallet logic into an API.

**Build**
Add endpoints such as:
- `POST /wallet`
- `GET /wallet?id=abc123`
- `POST /wallet/deposit`
- `POST /wallet/withdraw`

**You will learn**
- request parsing
- JSON decoding
- writing JSON responses
- basic route design

**Checkpoint**
You should see your wallet logic as a service, not just local code.

---

### Day 13 — Middleware and Request Logging
**Goal**
Understand cross-cutting concerns.

**You will learn**
- middleware pattern
- wrapping handlers
- request timing
- logging incoming requests

**Build**
Create middleware that logs:
- method
- path
- duration

**Checkpoint**
You should understand how frameworks are often optional because `net/http` is already powerful.

---

### Day 14 — Mini Project 2: REST Wallet Service
**Goal**
Finish a real backend milestone.

**Build**
A minimal wallet API with:
- in-memory storage
- deposit/withdraw logic
- transaction history
- logging middleware
- JSON responses

**Checkpoint**
At this point, you already have a small backend project worth keeping in GitHub.

---

## Week 3 — Concurrency and Production Readiness

### Day 15 — Goroutines and Concurrent Thinking
**Goal**
Understand Go's concurrency model conceptually.

**You will learn**
- goroutines
- blocking work vs concurrent work
- why concurrency is useful for servers

**Build**
Simulate async transaction processing:
- launch goroutines for fake transfer tasks
- print when each starts and finishes

**Checkpoint**
You should understand that concurrency is about structure, not guaranteed speed.

---

### Day 16 — Channels for Communication
**Goal**
Learn safe communication between goroutines.

**You will learn**
- unbuffered channels
- buffered channels
- send/receive patterns
- range over channels

**Build**
Create a transaction job queue with channels.

**Checkpoint**
You should feel how channels can model workflows and pipelines.

---

### Day 17 — Worker Pool Pattern
**Goal**
Use concurrency in a practical way.

**Build**
Create 3 workers that process transaction jobs from a channel.

**You will learn**
- fan-out processing
- job distribution
- basic throughput scaling pattern

**Checkpoint**
You should see how background jobs in real systems can be modeled.

---

### Day 18 — Context and Cancellation
**Goal**
Learn how Go manages request lifecycles.

**You will learn**
- `context.Context`
- deadlines
- cancellation
- timeouts

**Build**
Add timeout-aware processing to a fake long-running transfer.

**Checkpoint**
You should understand why context is essential in APIs and database calls.

---

### Day 19 — Structured Logging and Debuggability
**Goal**
Think like an engineer supporting production systems.

**You will learn**
- why logs need structure
- request IDs
- operation names
- useful log fields

**Build**
Improve logs so each request includes:
- endpoint
- duration
- wallet ID if present

**Checkpoint**
You should understand the difference between random printing and operational logging.

---

### Day 20 — Configuration Management
**Goal**
Separate config from code.

**You will learn**
- environment variables
- default values
- config structs
- keeping secrets out of code

**Build**
Create config loading for:
- app port
- environment name
- API key

**Checkpoint**
You should understand why production apps avoid hardcoded config.

---

### Day 21 — Safe Concurrent Balance Updates
**Goal**
See the danger of shared mutable state.

**You will learn**
- data races
- race conditions
- basic synchronization ideas
- why in-memory state gets tricky

**Build**
Run concurrent deposits and observe balance consistency problems.
Then protect access using a mutex.

**Checkpoint**
You should understand why financial state needs strong correctness guarantees.

---

## Week 4 — Backend Architecture and Fintech Thinking

### Day 22 — Repository Pattern and Database Readiness
**Goal**
Prepare your code for persistence.

**You will learn**
- repository interfaces
- swapping memory storage for database storage later
- keeping business logic independent from infra

**Build**
Refactor store access behind a repository layer.

**Checkpoint**
You should be able to describe how PostgreSQL could replace in-memory storage.

---

### Day 23 — Ledger Thinking and Transaction Safety
**Goal**
Think beyond simple balance mutation.

**You will learn**
- why ledgers exist
- balance as derived state
- append-only event thinking
- transactional integrity basics

**Build**
Add a transaction record type that captures:
- type
- amount
- timestamp
- reference ID

**Checkpoint**
You should understand why fintech systems often trust the ledger more than a raw balance field.

---

### Day 24 — Rate Limiting and Abuse Prevention
**Goal**
Add a basic protection mechanism.

**You will learn**
- why APIs need limits
- token bucket intuition
- per-client or per-wallet restrictions

**Build**
Add a simple in-memory rate limiter to protect deposit or transfer endpoints.

**Checkpoint**
You should understand why every public API needs guardrails.

---

### Day 25 — Authentication Basics
**Goal**
Add basic access control.

**You will learn**
- API keys
- auth middleware
- separation of authentication and authorization

**Build**
Protect state-changing endpoints with a simple API key check.

**Checkpoint**
You should understand the role of middleware in API security.

---

### Day 26 — Service Boundaries and Microservice Thinking
**Goal**
Learn how to split systems conceptually.

**You will learn**
- wallet service vs user service vs transfer service
- when not to split services too early
- service contracts

**Exercise**
Write a note describing how you would split your system into:
- wallet
- user
- ledger
- transfer
- notification

**Checkpoint**
You should be thinking like a backend architect now, not just a Go learner.

---

### Day 27 — Dockerizing the Service
**Goal**
Prepare your app to run consistently anywhere.

**You will learn**
- Dockerfile basics
- image building
- exposing ports
- why containers matter

**Build**
Create a Dockerfile for your Go API.

**Checkpoint**
You should be able to run your service in a container locally.

---

### Day 28 — Deployment Mental Model
**Goal**
Understand how your app would live in production.

**You will learn**
- stateless services
- horizontal scaling
- load balancers
- logs and config in deployment environments

**Exercise**
Write a short deployment note:
- where the app runs
- where config lives
- where logs go
- where database lives

**Checkpoint**
You should be able to explain your backend like a real system, not just code.

---

### Day 29 — Final Project: Mini Wallet / Remittance Backend
**Goal**
Bring everything together.

**Build**
A mini API that supports:
- wallet creation
- deposit
- withdraw
- transfer between wallets
- transaction history
- middleware logging
- basic auth
- rate limiting
- safe concurrent updates

**Recommended endpoint ideas**
- `POST /wallet`
- `POST /wallet/deposit`
- `POST /wallet/withdraw`
- `POST /wallet/transfer`
- `GET /wallet/history?id=...`

**Checkpoint**
This should become your first meaningful Go backend showcase project.

---

### Day 30 — Reflection, Cleanup, and GitHub Readiness
**Goal**
Turn learning into a durable asset.

**Do these**
- clean folder structure
- improve README
- document what you learned
- note future improvements
- create next-step backlog

**Write down**
- what you now understand well
- what still feels weak
- what you want to learn next: tests, PostgreSQL, Redis, gRPC, message queues

**Final outcome**
By Day 30, you are no longer just "learning Go syntax." You are building backend engineering intuition in Go.

---

## Suggested Repo Structure

```text
go-mastery/
  cmd/
    app/
      main.go
  internal/
    wallet/
      wallet.go
      transaction.go
      service.go
    store/
      memory_store.go
      repository.go
    api/
      handlers.go
      middleware.go
      dto.go
    config/
      config.go
  docs/
    Go_30_Day_Detailed_Plan.md
  README.md
  go.mod
```

---

## Daily Completion Template

Use this at the end of each day:

```md
## Day X Notes
- What I built:
- What I understood clearly:
- What confused me:
- One bug I hit:
- One thing to revise tomorrow:
```

---

## What to Ask ChatGPT Each Day

When you continue on free plan, use prompts like:
- `Start Day 08 from my Go plan`
- `Teach me Day 12 with code and explanation`
- `Review my Day 17 worker pool code`
- `Give me exercises for Day 21 race conditions`

That way, your plan stays usable even without Plus.
