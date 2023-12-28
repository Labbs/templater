package run

import (
	"github.com/labbs/templater/command/create/config"
	"github.com/rs/zerolog"
)

type Run struct {
	Logger zerolog.Logger
	Config config.Config
}
