package main

import (
	"os"

	"github.com/17660472762/miniprogramtest/pkg/server"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "mini_program",
		Usage: "",
		Action: func(ctx *cli.Context) error {
			return server.NewBuilder().
				Router().
				Build().
				Run()
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
