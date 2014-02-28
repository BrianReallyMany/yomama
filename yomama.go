package main

import (
    "github.com/BrianReallyMany/yomama/cmd"
    "os"
)



func main() {
    console := MakeMamaConsole()

    cmd.DoConsole(console, os.Stdin)
}
