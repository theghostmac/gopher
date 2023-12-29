package commands

import (
	"fmt"
	"github.com/theghostmac/gopher/internal/core"
	"github.com/urfave/cli/v2"
	"os"
)

func TrackCommand() *cli.Command {
	return &cli.Command{
		Name:  "track",
		Usage: "Track TODOs in the codebase to show progress",
		Action: func(context *cli.Context) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get current working directory: %w", err)
			}

			todos, err := core.TrackTODOs(cwd)
			if err != nil {
				return err
			}

			fmt.Println("Found", len(todos), "unattended TODOs: ")
			for _, todo := range todos {
				fmt.Printf("- [%s:%d] %s\n", todo.FilePath, todo.LineNumber, todo.Content)
			}
			return nil
		},
	}
}
