package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var App = cli.App{
	Name:        "yulctl",
	HelpName:    "yulmails control",
	Usage:       "manage yulmails from CLI",
	Description: "yulctl is the command line tool to start yulmails services (entrypoint, workers, etc.)",
	Version:     "0.1.0",
	Commands: []cli.Command{
		cli.Command{
			Name:        "entrypoint",
			Aliases:     []string{"e"},
			Usage:       "start the entrypoint",
			Description: "the entrypoint is a SMTP server in order to receive emails",
			Action: func(c *cli.Context) error {
				fmt.Println("entrypoint lol")
				return nil
			},
		},
	},
}
