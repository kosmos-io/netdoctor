package command

import "fmt"

type Wrapper struct {
	Cmd Command
}

func (w Wrapper) GetCommandStr() string {
	return fmt.Sprintf(`nsenter --target "1" --mount --uts --ipc --net --pid -- %s`, w.Cmd.GetCommandStr())
}

func (w Wrapper) ParseResult(str string) *Result {
	return w.Cmd.ParseResult(str)
}

func (w Wrapper) GetTargetStr() string {
	return w.Cmd.GetTargetStr()
}
