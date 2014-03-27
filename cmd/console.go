package cmd

// Interface for consoles
type Console interface {
	Greet() string

	Prompt() string

	// Generic help string
	Help() string

	// Execute a command with its arguments, return whether or not to continue.
	Execute(cmd string, args []string, line []byte) bool
}
