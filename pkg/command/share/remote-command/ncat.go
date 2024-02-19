package command

import (
	"fmt"
	"strings"

	"github.com/kosmos.io/netdoctor/pkg/utils"
)

type Ncat struct {
	Protocol string
	TargetIP []string
	Port     []string
}

func (c *Ncat) GetCommandStr() string {
	// execute once
	cmdStrList := []string{}
	for _, ip := range c.TargetIP {
		if c.Protocol == string(utils.UDP) {
			// TODO: find a better way
			// netcat send udp packet, if remote server explictly response "reject" icmp packet,
			// netcat will mark failed. If remote server response nothing, netcat mark succeeded.
			// But if remote server connect timetout, server also doesn't response anything,
			// netcat consider it succeeded, it's wrong.
			cmdStrList = append(cmdStrList, fmt.Sprintf("nc -w 1 -z -d -u -v %s %s 2>&1", ip, strings.Join(c.Port, " ")))
		} else {
			cmdStrList = append(cmdStrList, fmt.Sprintf("nc -w 1 -z -d  -v %s %s 2>&1", ip, strings.Join(c.Port, " ")))
		}
	}

	return strings.Join(cmdStrList, " && ")
}

func (c *Ncat) ParseResult(result string) *Result {
	isSucceed := CommandFailed
	index := strings.LastIndex(result, "succeeded")
	if index != -1 {
		isSucceed = CommandSuccessed
	}
	index = strings.LastIndex(result, "Connection refused")
	if index != -1 {
		isSucceed = CommandFailed
	}
	index = strings.LastIndex(result, "timed out")
	if index != -1 {
		isSucceed = CommandFailed
	}
	return &Result{
		Status:    isSucceed,
		ResultStr: fmt.Sprintf("%s %s", c.GetCommandStr(), result),
	}
}
