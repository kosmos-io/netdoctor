package command

import (
	"fmt"
	"regexp"

	"github.com/kosmos.io/netdoctor/pkg/utils"
)

var pingReg, _ = regexp.Compile(`PING[\s\S]*\s0[%]\spacket\sloss[\s\S]*$`)

type Ping struct {
	TargetIP string
}

func (c *Ping) GetTargetStr() string {
	return c.TargetIP
}

func (c *Ping) GetCommandStr() string {
	// execute once
	if utils.IsIPv6(c.TargetIP) {
		return fmt.Sprintf("ping6 -c 1 %s", c.TargetIP)
	} else {
		return fmt.Sprintf("ping -c 1 %s", c.TargetIP)
	}
}

func (c *Ping) ParseResult(result string) *Result {
	// klog.Infof("ping result parser: %s", result)
	isSucceed := CommandSuccessed
	if !pingReg.MatchString(result) {
		isSucceed = CommandFailed
	}
	return &Result{
		Status:    isSucceed,
		ResultStr: fmt.Sprintf("%s %s", c.GetCommandStr(), result),
	}
}
