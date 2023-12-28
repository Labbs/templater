package bootstrap

import (
	"github.com/labbs/templater/command/create/config"
	"github.com/labbs/templater/command/create/run"
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
)

func App(c *cli.Context) *run.Run {
	app := &run.Run{}
	app.Logger = InitLogger(c.App.Metadata["logger"].(zerolog.Logger))
	app.Config = config.AppConfig
	return app
}
