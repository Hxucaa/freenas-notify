package main

import (
	"fmt"
	"os"

	"github.com/hxucaa/freenas-notify/command"
	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name: "url",
		Value: "",
		Usage: "The URL of RabbitMQ server.",

	},
}

var Commands = []cli.Command{
	{
		Name:   "send",
		Usage:  "Send a message to message queue",
		Subcommands: []cli.Command{
			{
				Name: "pre-init",
				Usage: "Pre initilization status",
				Action: command.CmdSendPreInit,
				After: func(c *cli.Context) error {
					fmt.Println("Successfully sent pre init message.")
					return nil
				},
			},
			{
				Name: "post-init",
				Usage: "Post initilization status",
				Action: command.CmdSendPostInit,
				After: func(c *cli.Context) error {
					fmt.Println("Successfully sent post init message.")
					return nil
				},
			},
			{
				Name: "shutdown",
				Usage: "Shutdown status",
				Action: command.CmdSendShutdown,
				After: func(c *cli.Context) error {
					fmt.Println("Successfully sent shutdown message.")
					return nil
				},
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}