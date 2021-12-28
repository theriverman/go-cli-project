package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

/*
This cli.app.go file contains everything related to the main CLI application
All subsequent or lower-level commands are defined in cli.commands.go file.

Refer to the documentation of urfave/cli at https://github.com/urfave/cli
*/

// NewCLIApplication returns a *cli.App pointer which is the primary entrypoint to our CLI application.
// NewCLIApplication is the first call from main().
func NewCLIApplication() (app *cli.App) {
	app = &cli.App{
		Name:    AppName,
		Usage:   "The purpose of My App is not explained here yet",
		Version: AppSemVersion,
		// application-level flags can be define below. these are applicable during the whole runtime
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Destination: &appVerboseMode,
				Usage:       "Runs the application in verbose mode",
			},
		},
		Before: func(c *cli.Context) error {
			// CLI flags are processed at this point. Consider configuring your logging level
			if appVerboseMode {
				fmt.Println("Verbose mode has been enabled")
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			// remove this action block if you want to show the help when the
			// binary gets executed without any arguments
			return nil
		},
		Commands:  getApplicationCommands(),
		Copyright: AppCopyrightText,
	}
	return
}
