package bootstrap

import "github.com/rs/zerolog"

func InitLogger(logger zerolog.Logger) zerolog.Logger {
	l := logger.With().Str("module", "create").Logger()
	return l
}
