package commands

import (
	"fmt"
	"github.com/theghostmac/gopher/internal/core"
	"github.com/urfave/cli/v2"
)

func InitProjectCommand() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "Initialize a new Go project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "type",
				Aliases: []string{"t"},
				Value:   "general",
				Usage:   "Specify the project type (web, cli, general)",
			},
		},
		Action: func(c *cli.Context) error {
			projectName := c.Args().First()
			projectType := core.ProjectType(c.String("type"))

			err := core.CreateNewProject(projectName, projectType)
			if err != nil {
				fmt.Println("Error creating project: ", err)
				return err
			}

			fmt.Println("Project created successfully")
			return nil
		},
	}
}
