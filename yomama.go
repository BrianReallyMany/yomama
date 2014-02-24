package main

import (
    "github.com/BrianReallyMany/yomama/cmd"
)



func main() {
    console := MakeMamaConsole()

    cmd.DoConsole(console)
}
