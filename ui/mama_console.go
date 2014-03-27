package ui

import (
	"fmt"
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

func (c *MamaConsole) Echo_help() string {
	return "Usage: echo <text>\nPrints supplied text to standard out\n"
}

// Execute some yomama commands
func (c *MamaConsole) Execute(cmd string, args []string, line []byte) bool {
	switch cmd {
	// TODO default (unrecognized command) case
	case "":
		fmt.Println(c.Help())
		break

	case "help":
		fmt.Println(c.Help())
		break

	case "echo":
		fmt.Println(string(line))
		break

	case "system":
		fmt.Print(c.ctl.System(args))
		break

	case "yomama":
		fmt.Print(c.ctl.Dozens())
		break

	case "prepfiles":
		channel := make(chan string)
		go c.ctl.PrepFiles(args, channel)
		for {
			output, more := <- channel
			if !more {
				break
			}
			fmt.Println(output)
		}
		break

	case "exit":
		return false
	}

	return true
}
