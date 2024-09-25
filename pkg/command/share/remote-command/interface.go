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
	GetTargetStr() string
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
		return "SUCCEEDED"
	}
	if status == CommandFailed {
		return "FAILED"
	}
	return "UNEXCEPTIONED"
}

func NewCmd(protocol string, args ...any) Command {
	switch args[1].(type) {
	case []string:
		return &Ncat{
			Protocol: protocol,
			TargetIP: args[0].([]string),
			Port:     args[1].([]string),
		}
	default:
		if protocol == string(utils.TCP) {
			return &Curl{
				TargetIP: args[0].(string),
				Port:     args[1].(string),
			}
		} else if protocol == string(utils.DNS) {
			return &Nslookup{
				TargetHost: args[0].(string),
				DNSServer:  args[1].(string),
			}
		} else {
			return &Ping{
				TargetIP: args[0].(string),
			}
		}
	}
}
