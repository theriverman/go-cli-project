package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// application details injected at build time. see Makefile for more details
var (
	AppName          string = ""
	AppBuildDate     string = ""
	AppBuildType     string = ""
	AppSemVersion    string = ""
	AppCopyrightText string = ""
	GitCommit        string = ""
)

// application behaviour
var appVerboseMode bool = false

// demo values for the template
// Refer to the documentation of urfave/cli at https://github.com/urfave/cli
var greetNameValue string
var greetAskMeValue bool

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   fmt.Sprintf("Prints version information about %s", AppName),
	}

	app := &cli.App{
		Name:    AppName,
		Usage:   fmt.Sprintf("The purpose of %s is not explained here yet", AppName),
		Version: AppSemVersion,
		Flags: []cli.Flag{
			// application-level flags
			&cli.BoolFlag{
				Name:        "verbose",
				Destination: &appVerboseMode,
				Usage:       fmt.Sprintf("Runs %s in verbose mode", AppName),
			},
		},
		Before: func(c *cli.Context) error {
			// application-level flags are processed at this point. consider configuring your app-level things in the Before flag, such as, logging
			if appVerboseMode {
				fmt.Println("Verbose mode has been enabled")
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			// this is the default action when executed w/o additional commands
			// remove this action block if you want to show the default help text
			return nil
		},
		Commands: []*cli.Command{
			// application commands can be added here
			// see command `greet` below for a demonstration
			{
				Name:   "greet",
				Usage:  "Greets you appropriately to the configuration",
				Action: actionGreet,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Usage:       "If you define this flag, then I will use this for greeting",
						Value:       "You",
						Destination: &greetNameValue,
						Required:    false,
					},
					&cli.BoolFlag{
						Name:        "ask-me",
						Usage:       "If you set this flag, then I will ask about your wellbeing",
						Value:       false,
						Destination: &greetAskMeValue,
						Required:    false,
					},
				},
			},
			{
				Name:   "version",
				Usage:  fmt.Sprintf("Show the %s version information (detailed)", AppName),
				Action: actionVersion,
			},
		},
		Copyright: AppCopyrightText,
		// see the urfave/cli documentation for all possible options: https://github.com/urfave/cli/blob/master/docs/v2/manual.md
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
