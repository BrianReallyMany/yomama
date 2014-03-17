package main

import (
	"github.com/BrianReallyMany/yomama/cmd"
	"github.com/BrianReallyMany/yomama/ui"
	"os"
)

func main() {
	console := ui.MakeMamaConsole()

	cmd.DoConsole(console, os.Stdin)
}
