package command

import (
	"testing"
)

func TestW(t *testing.T) {
	//创建一个Reveiver
	tTV := TV{}
	//创建一个Command
	tCommand := NewCommand("open", tTV)
	//创建一个调用者
	tInvoke := Invoke{
		Command: tCommand,
	}
	tInvoke.ExecuteCommand()

	tCommand = NewCommand("changechannel", tTV)

	tInvoke = Invoke{
		Command: tCommand,
	}
	tInvoke.ExecuteCommand()

	tCommand = NewCommand("close", tTV)

	tInvoke = Invoke{
		Command: tCommand,
	}
	tInvoke.ExecuteCommand()
}
