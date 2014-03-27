package main

import (
	"github.com/BrianReallyMany/yomama/cmd"
	"github.com/BrianReallyMany/yomama/ui"
	"os"
)

func main() {
	console := ui.MakeMamaConsole()

	doer := cmd.NewConsoleDoer(console, os.Stdin, os.Stdout)

	doer.Run()
}
