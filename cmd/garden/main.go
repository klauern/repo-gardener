package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "repo-garden",
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "initialize a new configuration to run tasks across repositories",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("running app")
	}
}
