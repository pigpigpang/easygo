package main

import (
	"context"
	"easygo/internal"
	"easygo/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
)

// VERSION 版本号
var VERSION = "0.0.2"

func main() {
	logger.SetVersion(VERSION)
	ctx := logger.NewTagContext(context.Background(), "__main__")

	app := cli.NewApp()
	app.Name = "easy go"
	app.Version = VERSION
	app.Usage = "Easy go"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "mode",
			Aliases: []string{"m"},
			Usage:   "run mode",
		},
	}

	app.Action = func(c *cli.Context) error {
		return internal.Run(
			ctx,
			internal.SetMode(c.String("mode")),
			internal.SetVersion(VERSION),
		)
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}
