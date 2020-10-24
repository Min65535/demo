package main

import (
	"github.com/magiclvzs/antnet"
)

func main() {
	antnet.StartServer("tcp://:6666", antnet.MsgTypeCmd, &antnet.EchoMsgHandler{}, nil)
	antnet.WaitForSystemExit()
}
