package main

import (
    "bufio"
    "bytes"
    "strings"
    "testing"
)

func TestRun(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {
            name:     "Admin user gets special message",
            input:    "admin\n",
            expected: "Enter your name:\nWelcome back, boss 😎\n",
        },
        {
            name:     "Regular user gets welcome message",
            input:    "John\n",
            expected: "Enter your name:\nHello, John! Welcome to Go CLI 🚀\n",
        },
        {
            name:     "Empty name",
            input:    "\n",
            expected: "Enter your name:\nHello, ! Welcome to Go CLI 🚀\n",
        },
        {
            name:     "Name with spaces",
            input:    "Jane Doe\n",
            expected: "Enter your name:\nHello, Jane Doe! Welcome to Go CLI 🚀\n",
        },
        {
            name:     "Case sensitivity test",
            input:    "ADMIN\n",
            expected: "Enter your name:\nHello, ADMIN! Welcome to Go CLI 🚀\n",
        },
        {
            name:     "Name with special characters",
            input:    "user@123\n",
            expected: "Enter your name:\nHello, user@123! Welcome to Go CLI 🚀\n",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Setup input
            input := bufio.NewReader(strings.NewReader(tt.input))
            
            // Setup output buffer
            output := &bytes.Buffer{}
            bufioWriter := bufio.NewWriter(output)
            
            // Run the function (using your original Run implementation)
            Run(input, bufioWriter)
            bufioWriter.Flush()
            
            // Check the output
            if output.String() != tt.expected {
                t.Errorf("Expected:\n%sGot:\n%s", tt.expected, output.String())
            }
        })
    }
}

// Test for the greetUser function (simpler approach)
func TestGreetUser(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {
            name:     "Admin user",
            input:    "admin",
            expected: "Welcome back, boss 😎",
        },
        {
            name:     "Regular user",
            input:    "John",
            expected: "Hello, John! Welcome to Go CLI 🚀",
        },
        {
            name:     "Empty name",
            input:    "",
            expected: "Hello, ! Welcome to Go CLI 🚀",
        },
        {
            name:     "Case sensitive - ADMIN",
            input:    "ADMIN",
            expected: "Hello, ADMIN! Welcome to Go CLI 🚀",
        },
        {
            name:     "Name with spaces",
            input:    "Jane Doe",
            expected: "Hello, Jane Doe! Welcome to Go CLI 🚀",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := greetUser(tt.input)
            if result != tt.expected {
                t.Errorf("greetUser(%q) = %q, want %q", tt.input, result, tt.expected)
            }
        })
    }
}