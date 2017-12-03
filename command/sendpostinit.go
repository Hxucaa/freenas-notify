package command

import (
	"github.com/codegangsta/cli"

	"github.com/hxucaa/freenas-notify/mq"
)

func CmdSendPostInit(c *cli.Context) {
	// Write your code here
	mq.SendPostInit(c.GlobalString("url"))
}