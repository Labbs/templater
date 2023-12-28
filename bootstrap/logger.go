package bootstrap

import (
	"os"

	"github.com/labbs/templater/config"
	"github.com/rs/zerolog"
)

func InitLogger(c config.Config) zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().
		Timestamp().
		Str("version", c.Version).
		Logger()

	if c.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		logger = logger.With().Caller().Logger()
	}

	if c.PrettyLogs {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return logger
}
