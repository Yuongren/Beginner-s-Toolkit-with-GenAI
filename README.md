# Getting Started with Go: Building a Simple CLI Application

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## 🎯 Objective

This project introduces Go programming language (Golang) by building a simple Command Line Interface (CLI) application that interacts with user input.

## ❓ Why I Chose This Technology

I chose Go because:

- It is beginner-friendly and easy to set up
- It is widely used in backend systems and DevOps tools
- It emphasizes clean, readable, and efficient code

## 🏁 End Goal

To create a CLI program that:

- Accepts user input
- Processes conditional logic
- Displays a personalized message

## 📖 Quick Summary of the Technology

**Go (Golang)** is a statically typed, compiled programming language developed by Google.

### 🔍 Where It Is Used

- Backend web services
- Cloud infrastructure tools
- Command-line applications

### 🌍 Real-World Example

Tools like **Docker** and **Kubernetes** are built using Go.

---

## 💻 System Requirements

| Requirement | Specification |
|------------|---------------|
| **OS** | Windows / Linux / macOS |
| **Go** | Latest version (1.21+) |
| **Code Editor** | VS Code (recommended) |

---

## ⚙️ Installation & Setup

### Step 1: Install Go

sudo apt install golang-go

### Step 2: Verify Installation

go version

## 🧩 Minimal Working Example

## 📌 Description
This program:

    * Prompts the user for their name

    * Displays a custom greeting

    * Includes conditional logic for a special user ("admin")


### ▶️ Run the Program

 go run main.go


### 🧪 Run Tests

 go test -v -cover

## 🧠 AI Prompt Journal

### 🧾 Prompt 1: Testing & Behavior Understanding

#### Prompt:

"I'm learning how to test this Go CLI program… Ask me questions about its behavior, help identify missed cases, and guide me on edge cases and test prioritization."

- What the AI Helped With:

Encouraged thinking about program behavior

Helped identify test scenarios

Identified edge cases like empty input, case sensitivity, trailing spaces

- What I Learned:

Testing focuses on behavior, not just output

Edge cases are important for robust software

Table-driven tests are idiomatic in Go

- Evaluation: Very effective for deep understanding

### 🧾 Prompt 2: Code Readability & Maintainability

#### Prompt:

"I want to improve readability and maintainability… identify unclear parts, suggest better names, and improve structure."

- What the AI Helped With:

Improved variable naming

Suggested breaking code into smaller functions

Removed magic strings by using constants

- What I Learned:

Clean code improves maintainability

Smaller functions are easier to test

Single responsibility principle applies to functions

- Evaluation: Highly useful for writing professional code

### 🧾 Prompt 3: Writing Idiomatic Go

#### Prompt:

"I'm learning to write more idiomatic Go… suggest improvements, explain best practices, and compare versions."

- What the AI Helped With:

Introduced idiomatic patterns

Highlighted proper error handling

Showed dependency injection benefits

- What I Learned:

Idiomatic Go emphasizes simplicity

Avoid ignoring errors with _

Use bufio for better I/O control

- Evaluation: Improved code quality significantly

### 🧾 Prompt 4: Understanding Concepts (Dependency Injection & I/O)
#### Prompt:

"I want to understand dependency injection and input/output handling in Go… explain, give use cases, and suggest practice ideas."

- What the AI Helped With:

Explained design patterns in simple terms

Connected concepts to real-world use

Provided refactored examples

- What I Learned:

Passing dependencies improves testability

Modular design is important

Interfaces enable flexible I/O handling

- Evaluation: Helped bridge theory and practice

## 📁 Project Structure
```
go-cli-app/
├── main.go          # Main application code
├── main_test.go     # Unit tests
├── go.mod           # Go module file
└── README.md        # This file
```

## ▶️ Quick Start

Clone the repositoy.from [github](https://github.com/Yuongren/Beginner-s-Toolkit-with-GenAI).
Then Initialize module 
```
- go mod init cli-app
```

