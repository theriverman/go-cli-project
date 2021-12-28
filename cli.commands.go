package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

/*
Refer to the documentation of urfave/cli at https://github.com/urfave/cli
*/

var greetNameValue string
var greetAskMeValue bool

func getApplicationCommands() []*cli.Command {
	return []*cli.Command{
		// your custom application commands can be added here
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
	}
}
