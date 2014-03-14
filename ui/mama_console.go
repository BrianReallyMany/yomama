package ui

import (
    "fmt"
    "log"
    "os/exec"
)

// The yomama console class
type MamaConsole struct {
	ctl MamaController
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

// For the lost
func (c *MamaConsole) Help() string {
	return "Available commands: echo, system, exit"
}

// Execute some yomama commands
func (c *MamaConsole) Execute(cmd string, args []string, line []byte) bool {
    switch cmd {
	    // TODO default (unrecognized command) case
    case "":
	    fmt.Println(c.Help())

    case "help":
	    fmt.Println(c.Help())

    case "echo":
        fmt.Println(string(line))
        break

    case "system":
        out, err := exec.Command(args[0], args[1:]...).Output()
	    if err != nil {
		    log.Fatal(err)
	    }
        fmt.Print(string(out))
        break

    case "yomama":
        fmt.Print(c.ctl.Dozens())

	case "prepfiles":
		c.ctl.PrepFiles(args)

    case "exit":
        return false
    }

    return true
}
