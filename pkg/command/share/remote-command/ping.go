package command

import (
	"fmt"
	"regexp"
)

var pingReg, _ = regexp.Compile(`PING[\s\S]*1\spackets\stransmitted,\s1\spackets\sreceived,\s0[%]\spacket\sloss[\s\S]*$`)

type Ping struct {
	TargetIP string
}

func (c *Ping) GetTargetStr() string {
	return c.TargetIP
}

func (c *Ping) GetCommandStr() string {
	// execute once
	return fmt.Sprintf("ping -c 1 %s", c.TargetIP)
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
