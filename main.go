package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // For the refactored version
    input := bufio.NewReader(os.Stdin)
    output := bufio.NewWriter(os.Stdout)
    Run(input, output)
    
    // Or use the simple version:
    // SimpleMain()
}

// Refactored version with dependency injection (for testing)
func Run(input *bufio.Reader, output *bufio.Writer) {
    defer output.Flush()
    
    fmt.Fprintln(output, "Enter your name:")
    output.Flush()
    
    name, _ := input.ReadString('\n')
    name = strings.TrimSpace(name)
    
    if name == "admin" {
        fmt.Fprintln(output, "Welcome back, boss 😎")
    } else {
        fmt.Fprintf(output, "Hello, %s! Welcome to Go CLI 🚀\n", name)
    }
}

// Simple version with extracted logic (easier to test)
func SimpleMain() {
    var name string
    
    fmt.Println("Enter your name:")
    fmt.Scanln(&name)
    
    result := greetUser(name)
    fmt.Println(result)
}

// Extracted logic for easy testing
func greetUser(name string) string {
    if name == "admin" {
        return "Welcome back, boss 😎"
    }
    return fmt.Sprintf("Hello, %s! Welcome to Go CLI 🚀", name)
}