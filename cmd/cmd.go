package cmd

import (
	"github.com/urfave/cli"

	"github.com/yulpa/yulmails/services/entrypoint"
	"github.com/yulpa/yulmails/services/worker"
	"github.com/yulpa/yulmails/services/sender"
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
			Usage:       "start an entrypoint",
			Description: "the entrypoint is a SMTP server in order to receive emails",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "smtp-config",
					Value: "/etc/yulmails/smtp.json",
					Usage: "absolute path to the SMTP config file",
				},
			},
			Action: func(c *cli.Context) error {
				return entrypoint.StartSMTP(c.String("smtp-config"))
			},
		},
		cli.Command{
			Name:        "worker",
			Aliases:     []string{"w"},
			Usage:       "start a worker",
			Description: "a worker is a dedicated resource in order to fetch emails from the queue and compute them",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "worker-config",
					Value: "/etc/yulmails/worker.json",
					Usage: "absolute path to the worker config file",
				},
			},
			Action: func(c *cli.Context) error {
				return worker.StartWorker(c.String("worker-config"))
			},
		},
		cli.Command{
			Name:        "sender",
			Aliases:     []string{"s"},
			Usage:       "start a sender",
			Description: "a sender is a dedicated resource in order to fetch emails from the queue and send them",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "sender-config",
					Value: "/etc/yulmails/sender.json",
					Usage: "absolute path to the sender config file",
				},
			},
			Action: func(c *cli.Context) error {
				return sender.StartSender(c.String("sender-config"))
			},
		},
	},
}
