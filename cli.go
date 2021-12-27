package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func NewCLIApplication() (app *cli.App) {
	app = &cli.App{
		Name:    AppName,
		Usage:   "The purpose of My App is not explained here yet",
		Version: AppSemVersion,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Destination: &appVerboseMode,
				Usage:       "Runs the application in verbose mode",
			},
		},
		Before: func(c *cli.Context) error {
			// CLI flags are processed at this point. Consider configuring your logging level
			fmt.Println("Verbose mode enabled:", appVerboseMode)
			return nil
		},
		Action: func(c *cli.Context) error {
			// remove this action block if you want to show the help when the
			// binary gets executed without any arguments
			return nil
		},
		Commands: getApplicationCommands(),
	}
	return
}
