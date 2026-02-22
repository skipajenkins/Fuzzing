---

# 🐹 Go Reverse + Fuzz Testing - Learning Go Testing & Fuzzing

This repository explores:
- Writing reusable string utilities  
- Validating correctness using unit tests  
- Discovering edge cases using fuzz testing  
- Ensuring UTF-8 correctness  

---

## 🎯 Learning Objectives

By working through this project, you will learn:

- How to write reverse string  functions
- How to implement unit testing in Go
- How fuzz testing works in Go
- How to detect unexpected bugs automatically
- How to ensure UTF-8 correctness in string operations

---

## 📁 Project Structure

- Fuzz/
  - main.go
  - reverse_test.go
  - go.mod
  - testdata/
    - fuzz/
      - FuzzReverse/
        - beed24892637985f
        - 1ffc28f7538e29d7
    
---

## ⚙️ Prerequisites

| Tool | Purpose |
|------|----------|
| Go   | Compiler & runtime |
| Git  | Version control |

---

## 🦫 Installing Go

### Step 1: Check if Go is installed

```bash
go version
```
Step 2: Install Go (if missing)

Linux / macOS:
```bash
curl -OL https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go1.22.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```
Windows:
```bash
https://go.dev/dl/
```

---

🚀 Getting Started
Step 1: Clone the Repository
```bash
git clone https://github.com/skipajenkins/Fuzzing.git
cd morestrings
```
Step 2: Install Dependencies
```
go mod tidy
```
Step 3: Run the Application
```
go run ./hello
```
---

## 🧠 Key Concepts

| Concept            | Explanation                                   |
|--------------------|-----------------------------------------------|
| Go Modules         | Dependency management using `go.mod`          |
| Packages           | Logical grouping of Go source files           |
| Exported Functions | Capitalized functions (e.g., `Reverse`)       |
| Unit Testing       | Validation using `_test.go` files             |
| Fuzz Testing       | Automatic random input testing                |
| UTF-8 Safety       | Ensuring valid Unicode output                 |

## 🧪 Testing & Fuzzing
- Run Unit Tests
- go test ./...
- Run Fuzz Tests
- go test -fuzz=FuzzReverse

This automatically generates:
``` bash 
testdata/fuzz/FuzzReverse/
```
Containing seed corpus inputs such as:
```bash
beed24892637985f

1ffc28f7538e29d7
```
These inputs help the fuzz engine discover edge cases and prevent regressions.

---

## 🏗️ How It Works

1️⃣ Reverse Function
```bash
Reverse(string) → (string, error)
```
Safely reverses UTF-8 encoded strings.

2️⃣ Unit Testing
```bash
reverse_test.go
```
Validates correctness using deterministic test cases.

3️⃣ Fuzz Testing
```bash
FuzzReverse()
````
Automatically generates thousands of random inputs to detect edge cases.

4️⃣ UTF-8 Validation
``` bash
utf8.ValidString(...)
```
Ensures output remains valid Unicode.

---

## 🧩 Example Output
```bash
go test -fuzz=FuzzReverse

fuzz: elapsed: 10s, execs: 2213270 (221327/sec)
PASS
ok      example/morestrings   10.512s
```
(Exact output varies per run.)

---

# 🧱 Built With

- Go (Golang)

- Go Modules

- Go Testing Framework

- Go Fuzzing Engine

---

## 📜 License

This project is licensed under the MIT License — free to use, modify, and distribute.

---
