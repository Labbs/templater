package bootstrap

import (
	"github.com/labbs/templater/config"
	"github.com/rs/zerolog"
)

type Application struct {
	Logger zerolog.Logger
	Config *config.Config
}

func App(c *config.Config) *Application {
	app := &Application{}
	app.Config = c
	app.Logger = InitLogger(*c)

	return app
}
