package ord

import "fmt"

type TV struct{}

func (tv TV) Open() {
	fmt.Println("开机")
}

func (tv TV) Close() {
	fmt.Println("关机")
}

func (tv TV) ChangeChannel() {
	fmt.Println("换台")
}


type Command interface {
	Execute()
}


func NewCommand(t string, tv TV) Command {
	switch t {
	case "open":
		return OpenCommand{
			receiver: tv,
		}
	case "close":
		return CloseCommand{
			receiver: tv,
		}
	case "changechannel":
		return ChangeChannelCommand{
			receiver: tv,
		}
	default:
		return nil
	}
}

type OpenCommand struct {
	receiver TV
}

func (oc OpenCommand) Execute() {
	oc.receiver.Open()
}

type CloseCommand struct {
	receiver TV
}

func (cc CloseCommand) Execute() {
	cc.receiver.Close()
}

type ChangeChannelCommand struct {
	receiver TV
}

func (ccc ChangeChannelCommand) Execute() {
	ccc.receiver.ChangeChannel()
}


type Invoke struct {
	Command
}

func (i Invoke) ExecuteCommand() {
	i.Command.Execute()
}
