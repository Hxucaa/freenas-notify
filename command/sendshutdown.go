package command

import (
	"github.com/codegangsta/cli"

	"github.com/hxucaa/freenas-notify/mq"
)

func CmdSendShutdown(c *cli.Context) {
	// Write your code here
	mq.SendShutdown(c.GlobalString("url"))
}