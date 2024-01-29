package command

import (
	"fmt"

	"github.com/kosmos.io/netdoctor/pkg/utils"
)

const (
	ExecError = iota
	CommandSuccessed
	CommandFailed
)

type Result struct {
	Status    int
	ResultStr string
}

type Command interface {
	GetCommandStr() string
	ParseResult(string) *Result
}

func ParseError(err error) *Result {
	return &Result{
		Status:    ExecError,
		ResultStr: fmt.Sprintf("exec error: %s", err),
	}
}

func PrintStatus(status int) string {
	if status == ExecError {
		return "EXCEPTION"
	}
	if status == CommandSuccessed {
		return "SUCCESSED"
	}
	if status == CommandFailed {
		return "FAILED"
	}
	return "UNEXCEPTIONED"
}

func NewCmd(protocol string, args ...string) Command {
	if protocol == string(utils.TCP) {
		return &Curl{
			TargetIP: args[0],
			Port:     args[1],
		}
	} else {
		return &Ping{
			TargetIP: args[0],
		}
	}
}
