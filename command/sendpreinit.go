package command

import (
	"github.com/codegangsta/cli"

	"github.com/hxucaa/freenas-notify/mq"
)

func CmdSendPreInit(c *cli.Context) {
	// Write your code here
	mq.SendPreInit(c.GlobalString("url"))
}