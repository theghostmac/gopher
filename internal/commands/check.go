package commands

import (
	"fmt"
	"github.com/theghostmac/gopher/internal/core"
	"github.com/urfave/cli/v2"
	"os"
)

func CheckProjectCommand() *cli.Command {
	return &cli.Command{
		Name:  "check",
		Usage: "Check the project for compilation errors",
		Action: func(context *cli.Context) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get current working directory: %w", err)
			}

			err = core.CheckProject(cwd)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
