package main

import (
	"github.com/theghostmac/gopher/internal/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "Gopher",
		Usage: "A tool for managing Go projects",
		Commands: []*cli.Command{
			commands.InitProjectCommand(),
			// TODO: add other commands here.
		},
		// TODO: other parameters.
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
