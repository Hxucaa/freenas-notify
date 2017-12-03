package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Hxucaa"
	app.Email = "lancezhu77@gmail.com"
	app.Usage = "Send FreeNAS status message to RabbitMQ"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
