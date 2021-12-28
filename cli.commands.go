package main

import (
	"fmt"
	"runtime"

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
			Name:  "greet",
			Usage: "Greets you appropriately to the configuration",
			Action: func(c *cli.Context) error {
				if len(greetNameValue) > 0 {
					greetNameValue = ", " + greetNameValue
				}
				fmt.Printf("Hello%s\n", greetNameValue)
				if greetAskMeValue {
					fmt.Println("How are you?")
				}
				return nil
			},
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
			Name:  "version",
			Usage: fmt.Sprintf("Show the %s version information (detailed)", AppName),
			Action: func(c *cli.Context) error {
				fmt.Println(AppName + ":")
				fmt.Printf("  Version: %s\n", AppSemVersion)
				fmt.Printf("  Go version: %s\n", runtime.Version())
				fmt.Printf("  Git commit: %s\n", GitCommit)
				fmt.Printf("  Built: %s\n", AppBuildDate)
				fmt.Printf("  OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
				fmt.Printf("  Build type: %s\n", AppBuildType)
				return nil
			},
		},
	}
}
