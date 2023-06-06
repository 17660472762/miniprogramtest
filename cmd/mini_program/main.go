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
			appid := ctx.String("appid")
			secret := ctx.String("secret")
			return server.NewBuilder(appid, secret).
				Router().
				Build().
				Run()
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "appid",
				Value:    "",
				Usage:    "",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "secret",
				Value:    "",
				Usage:    "",
				Required: true,
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
