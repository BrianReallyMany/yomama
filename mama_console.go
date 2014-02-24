package main

import (
    "fmt"
)

// The yomama console class
type MamaConsole struct {
}

// Instantiate a new yomama console
func MakeMamaConsole() *MamaConsole {
    return &MamaConsole{}
}

// The yomama console greeting
func (c *MamaConsole) Greet() string {
    return "Welcome to YoMama!"
}

// The yomama console prompt
func (c *MamaConsole) Prompt() string {
    return "Mama> "
}

// Execute some yomama commands
func (c *MamaConsole) Execute(cmd string, args []string, line []byte) bool {
    switch cmd {
    case "echo":
        fmt.Println(string(line))
        break

    case "exit":
        return false
    }

    return true
}
