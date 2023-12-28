package config

import "github.com/urfave/cli/v2"

var AppConfig Config = Config{}

type Config struct {
	ValuesFile    cli.StringSlice
	TemplateFiles string
	OutputFiles   string
}
