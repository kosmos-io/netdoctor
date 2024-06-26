package command

import (
	"fmt"

	"github.com/kosmos.io/netdoctor/pkg/utils"
)

type Curl struct {
	TargetIP string
	Port     string
}

func (c *Curl) GetTargetStr() string {
	return fmt.Sprintf("%s:%s", c.TargetIP, c.Port)
}

func (c *Curl) GetCommandStr() string {
	// execute once
	if utils.IsIPv6(c.TargetIP) {
		return fmt.Sprintf("curl -k http://[%s]:%s/", c.TargetIP, c.Port)
	}
	return fmt.Sprintf("curl -k http://%s:%s/", c.TargetIP, c.Port)
}

func (c *Curl) ParseResult(result string) *Result {
	// klog.Infof("curl result parser: %s", result)
	isSucceed := CommandSuccessed
	if result != "OK" {
		isSucceed = CommandFailed
	}
	return &Result{
		Status:    isSucceed,
		ResultStr: fmt.Sprintf("%s %s", c.GetCommandStr(), result),
	}
}
