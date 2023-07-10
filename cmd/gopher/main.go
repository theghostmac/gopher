package main

import (
	"fmt"
	"os"

	"github.com/theghostmac/gopher/internal/core/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gopher <command> [arguments]")
		os.Exit(0)
	}

	command := os.Args[1]
	args := os.Args[2:]
	commands.ExecuteCommand(command, args)
}
