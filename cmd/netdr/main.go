package main

import (
	"k8s.io/component-base/cli"
	"k8s.io/kubectl/pkg/cmd/util"

	"github.com/kosmos.io/netdoctor/pkg/netdr"
)

func main() {
	cmd := netdr.NewNetDoctorCtlCommand()
	if err := cli.RunNoErrOutput(cmd); err != nil {
		util.CheckErr(err)
	}
}
