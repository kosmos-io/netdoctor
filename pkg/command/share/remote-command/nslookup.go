package command

import (
	"fmt"
	"strings"
)

type Nslookup struct {
	TargetHost string
	DNSServer  string
}

func (c *Nslookup) GetTargetStr() string {
	var targethost string
	var targetdns string
	if len(c.TargetHost) == 0 {
		targethost = fmt.Sprintf("dns:%s", "kubernetes.default.svc.cluster.local")
	} else {
		targethost = c.TargetHost
	}
	if len(c.DNSServer) == 0 {
		targetdns = "coredns"
	} else {
		targetdns = c.DNSServer
	}
	return fmt.Sprintf("host: %s; dns: %s", targethost, targetdns)
}

func (c *Nslookup) GetCommandStr() string {
	// execute once
	if c.TargetHost == "" {
		c.TargetHost = "kubernetes.default.svc.cluster.local"
	}
	return fmt.Sprintf("nslookup %s %s", c.TargetHost, c.DNSServer)
}

func (c *Nslookup) ParseResult(result string) *Result {
	// klog.Infof("curl result parser: %s", result)
	isSucceed := CommandSuccessed
	index := strings.LastIndex(result, "server can't find")
	if index != -1 {
		isSucceed = CommandFailed
	}
	index = strings.LastIndex(result, "connection timed out")
	if index != -1 {
		isSucceed = CommandFailed
	}
	return &Result{
		Status:    isSucceed,
		ResultStr: fmt.Sprintf("%s %s", c.GetCommandStr(), result),
	}
}
