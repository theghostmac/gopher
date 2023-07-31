package commands

import (
	"fmt"
)

func ExecuteCommand(command string, args []string) {
	switch command {
	case "new":
		if len(args) < 1 {
			fmt.Println("Usage: gopher new <project_name> [--web | --app]")
			return
		}
		projectName := args[0]
		flags := args[1:]
		NewProject(projectName, flags)
	case "add":
		AddDependencies(args)
	case "build":
		BuildProject()
	default:
		fmt.Println("Unknown command: ", command)
	}
}
