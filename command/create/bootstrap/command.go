package bootstrap

import (
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create files from template",
		Flags: Flags(),
		Action: func(c *cli.Context) error {
			app := App(c)

			app.Run()
			return nil
		},
	}
}
