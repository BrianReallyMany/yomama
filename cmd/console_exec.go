package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type ConsoleDoer struct {
	console Console
	reader *bufio.Reader
	writer io.Writer

	history []string
	helps map[string]reflect.Value
}

func NewConsoleDoer(console Console, reader io.Reader, writer io.Writer) *ConsoleDoer {
	c := &ConsoleDoer{}

	c.console = console
	c.reader = bufio.NewReader(reader)
	c.writer = writer
	c.history = make([]string, 0)
	c.helps = make(map[string]reflect.Value)

	cVal := reflect.ValueOf(console)
	cType := cVal.Type()
	numMethods := cVal.NumMethod()
	for i := 0; i < numMethods; i++ {
		mVal := cVal.Method(i)
		mType := mVal.Type()
		mName := cType.Method(i).Name

		// Valid help function tagged by trailing _help in name, takes no parameters, and returns a string
		if strings.HasSuffix(mName, "_help") && mType.NumIn() == 0 && mType.NumOut() == 1 && mType.Out(0).Kind() == reflect.String {
			// We gots ourselves a valid help function
			// Command names are always lower case

			c.helps[strings.ToLower(strings.TrimSuffix(mName, "_help"))] = mVal
		}
	}

	return c
}

// Execute a console implementation
func (c *ConsoleDoer) Run() {
	// Greet the user
	fmt.Print(c.console.Greet(), "\n\n")

	// Prompt the user until they want to be prompted no more
	for {
		fmt.Print(c.console.Prompt()) // Output the console prompt

		// Grab a line from the console
		line, err := c.reader.ReadBytes('\n')

		if len(line) == 0 {
			continue
		}

		if err != nil {
			fmt.Println("Sorry, there was a problem reading that command")
			fmt.Println(err)
		}

		// Sanitize the line
		line = bytes.Trim(line, " \n")

		// Separate the line
		lineSep := bytes.Split(line, []byte{' '})

		// Build arguments slice
		args := make([]string, len(lineSep)-1)
		for i, arg := range lineSep[1:] {
			args[i] = string(arg)
		}

		// Get the line with only the arguments
		var lineWithoutCmd []byte
		if len(line) > len(lineSep[0]) {
			lineWithoutCmd = line[len(lineSep[0])+1:]
		} else {
			lineWithoutCmd = []byte{} // Empty line
		}

		cmd := string(lineSep[0])

		// Finally, execute the line

		if cmd == "help" {
			// It's a help line

			if len(args) == 0 {
				// General help
				c.help("")
			} else if len(args) == 1 {
				// Help on a specific command
				c.help(string(args[0]))
			} else {
				// Too many things...
				c.writer.Write([]byte("I can only help you with a single command, sonny"))
			}
		} else if !c.console.Execute(string(lineSep[0]), args, lineWithoutCmd) {
			break // Break if it's time for the console to exit
		}

		c.history = append(c.history, string(line))
	}
}

func (c *ConsoleDoer) help(command string) {
	if command == "" {
		// User wants some generic help
		c.writer.Write([]byte(c.console.Help()+"\n"))
	} else {
		if helpFunc, ok := c.helps[command]; ok {
			c.writer.Write([]byte(helpFunc.Call(nil)[0].String()+"\n"))
		} else {
			c.writer.Write([]byte("Sorry, that help topic doesn't exist.\n"))
		}
	}
}
