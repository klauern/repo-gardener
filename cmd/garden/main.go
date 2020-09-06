package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
