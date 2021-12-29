package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// application details injected at build time
var (
	AppName          string = "app" // pretty-formatted
	AppBuildType     string = "unreleased/internal"
	AppBuildDate     string = time.Now().Format("02 Jan 2006 15:04:05") // equals to date '+%c'
	AppSemVersion    string = "no-version"
	AppCopyrightText string = "no copyright"
	GitCommit        string = "commit-id-could-not-be-retrieved"
)

func main() {
	// if you're using the template as intended, this main() function shouldn't be modified at all
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "Prints version information of go-socks5-cli and quit",
	}

	app := NewCLIApplication()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
