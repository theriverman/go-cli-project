package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// application details injected at build time
var AppName string = "my-app" // pretty-formatted
var AppBuildType string = "unreleased/internal"
var AppBuildDate string = time.Now().Format("02 Jan 2006 15:04:05") // equals to date '+%c'
var AppSemVersion, GitCommit string

// application behaviour
var appVerboseMode bool = false

func main() {
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
