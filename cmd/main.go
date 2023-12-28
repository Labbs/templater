// Path: cmd/main.go
package main

import (
	"log"
	"os"

	"github.com/labbs/templater/bootstrap"
	"github.com/labbs/templater/config"

	createBootstrap "github.com/labbs/templater/command/create/bootstrap"

	// "github.com/labbs/templater/render"
	"github.com/urfave/cli/v2"
)

// Version is the current version of the application.
// This value is overwritten by the build process.
var version = "development"

func main() {
	app := cli.NewApp()

	app.Name = "Templater"
	app.Usage = "A template engine"

	app.Version = version
	appConfig := &config.AppConfig
	appConfig.Version = version

	app.Flags = bootstrap.Flags()
	app.Before = func(c *cli.Context) error {
		c.App.Metadata["logger"] = bootstrap.InitLogger(*appConfig)
		return nil
	}

	app.Commands = []*cli.Command{}

	app.Commands = append(app.Commands, createBootstrap.Command())

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
