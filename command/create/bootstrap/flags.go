package bootstrap

import (
	"github.com/labbs/templater/command/create/config"
	"github.com/urfave/cli/v2"
)

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "values",
			Aliases:     []string{"v"},
			Usage:       "values file",
			Destination: &config.AppConfig.ValuesFile,
		},
		&cli.StringFlag{
			Name:        "template",
			Aliases:     []string{"t"},
			Usage:       "template file",
			Destination: &config.AppConfig.TemplateFiles,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "output file",
			Destination: &config.AppConfig.OutputFiles,
		},
	}
}
