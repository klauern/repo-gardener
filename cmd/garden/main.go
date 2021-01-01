package main

import (
	"os"

	gardener "github.com/klauern/repo-gardener"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "repo-garden",
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "initialize a new configuration to run tasks across repositories",
				Action: func(c *cli.Context) error {
					return gardener.NewGardenConfig().Template("garden.yaml", afero.NewOsFs())
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("running app")
	}
}
