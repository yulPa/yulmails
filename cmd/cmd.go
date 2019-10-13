package cmd

import (
	"github.com/urfave/cli"

	"gitlab.com/tortuemat/yulmails/services/entrypoint"
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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "smtp-config",
					Value: "/etc/yulmails/smtp.json",
					Usage: "absolute path to the SMTP config file",
				},
			},
			Action: func(c *cli.Context) error {
				return entrypoint.StartSMTP(c.String("smtp-config"))
			},
		},
	},
}
