# 🧠 4-Week Golang Learning Roadmap – For C++/Python Developers

## 🎯 Objectives
- Quickly get up to speed with Golang for developers with a C++/Python background
- Build backend services for a gateway system that receives and forwards data to other nodes
- Focus on REST APIs, TCP/HTTP handling, concurrency, and clean architecture

---

## 📅 Week 1: Master Go Syntax & Data Types

**Goal:** Understand the basics of Go syntax, data handling, and structure.

### Topics:
- Install Go + VSCode or GoLand
- Hello world, packages, and modules
- Variables, constants, structs, slices, maps, pointers
- Functions and method receivers
- Interfaces – how Go replaces traditional inheritance

### ✅ Exercise:
> Build a CLI tool that takes input from the user, calculates a sum, sorts values, and prints results.

---

## 📅 Week 2: Concurrency & HTTP Service

**Goal:** Learn Go’s concurrency model and build a basic HTTP server.

### Topics:
- Goroutines and channels
- `select` statement and channel timeouts
- `context.Context` – controlling goroutines with timeout/cancellation
- Creating a REST API using `net/http`
- Logging and error handling

### ✅ Exercise:
> Build a REST API with two endpoints:  
> `POST /data` – receive and process data  
> `GET /status` – return system status (mocked forwarding state)

---

## 📅 Week 3: Communication & Data Forwarding

**Goal:** Work with HTTP/TCP clients to forward data between services or nodes.

### Topics:
- Send HTTP requests using Go’s `net/http` client
- Parse and encode JSON using `encoding/json`
- TCP communication using `net.Dial`, `net.Listen`
- Implement
