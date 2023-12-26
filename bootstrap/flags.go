package bootstrap

import (
	"github.com/labbs/templater/config"
	"github.com/urfave/cli/v2"
)

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Aliases:     []string{"d"},
			Usage:       "debug mode",
			Destination: &config.AppConfig.Debug,
		},
		&cli.BoolFlag{
			Name:        "pretty-logs",
			Aliases:     []string{"p"},
			Usage:       "pretty logs",
			Value:       true,
			Destination: &config.AppConfig.PrettyLogs,
		},
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
