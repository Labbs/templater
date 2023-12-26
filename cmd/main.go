// Path: cmd/main.go
package main

import (
	"log"
	"os"

	"github.com/labbs/templater/bootstrap"
	"github.com/labbs/templater/config"
	"github.com/labbs/templater/render"
	"github.com/urfave/cli/v2"
)

// Version is the current version of the application.
// This value is overwritten by the build process.
var version = "development"

func main() {
	flags := bootstrap.Flags()
	appConfig := &config.AppConfig
	appConfig.Version = version

	app := cli.NewApp()
	app.Name = "Templater"
	app.Version = version

	app.Commands = []*cli.Command{
		{
			Name:  "render",
			Usage: "render a template",
			Flags: flags,
			Action: func(c *cli.Context) error {
				appBootstrap := bootstrap.App(appConfig)

				return render.Render(appBootstrap)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
