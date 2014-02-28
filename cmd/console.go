package cmd

import (
    "bufio"
    "bytes"
    "fmt"
    "io"
)

// Interface for consoles
type Console interface {
    Greet() string
    Prompt() string

    // Execute a command with its arguments, return whether or not to continue.
    Execute(cmd string, args []string, line []byte) bool
}

// Execute a console implementation
// TODO: Possibly supply buffer to read from?
func DoConsole(c Console, ioReader io.Reader) {
    // Greet the user
    fmt.Print(c.Greet(), "\n\n")

    // Create a buffer for line reading
    reader := bufio.NewReader(ioReader)

    // Prompt the user until they want to be prompted no more
    for {
        fmt.Print(c.Prompt()) // Output the console prompt

        // Grab a line from the console
        line, err := reader.ReadBytes('\n')

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

        // Finally, execute the line
        if !c.Execute(string(lineSep[0]), args, lineWithoutCmd) {
            break // Break if it's time for the console to exit
        } 
    }
}
