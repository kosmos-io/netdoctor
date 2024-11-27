package cmd

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	ctlutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"

	"github.com/kosmos.io/netdoctor/pkg/command/share"
	"github.com/kosmos.io/netdoctor/pkg/utils"
)

var (
	once sync.Once
)

var checkExample = templates.Examples(i18n.T(`
        # Check cluster network connectivity, e.g:
        netctl check 
`))

type CheckOptions struct {
	DoOption *share.DoOptions
}

func NewCheckCmd() *cobra.Command {
	o := &CheckOptions{}
	cmd := &cobra.Command{
		Use:                   "check",
		Short:                 i18n.T("Check Kubernetes Cluster Network Connectivity."),
		Long:                  "",
		Example:               checkExample,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctlutil.CheckErr(o.Complete())
			ctlutil.CheckErr(o.Validate())
			ctlutil.CheckErr(o.Run())
			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	return cmd
}

func (o *CheckOptions) LoadConfig() {
	fromConfig := &share.DoOptions{}
	if err := utils.ReadOpt(fromConfig); err == nil {
		if len(fromConfig.Mode) == 0 {
			fromConfig.Mode = share.Pod
		}
		once.Do(func() {
			klog.Infof("use config from file!!!!!!")
		})
		o.DoOption = fromConfig
	}
}

func (o *CheckOptions) Complete() error {
	o.LoadConfig()

	if o.DoOption == nil {
		return fmt.Errorf("config.json load error")
	}

	srcFloater := &share.Floater{
		Namespace:         o.DoOption.Namespace,
		Name:              share.DefaultFloaterName,
		ImageRepository:   o.DoOption.SrcImageRepository,
		Version:           o.DoOption.Version,
		PodWaitTime:       o.DoOption.PodWaitTime,
		Port:              o.DoOption.Port,
		EnableHostNetwork: o.DoOption.GetEnableHostNetwork(true),
		EnableAnalysis:    false,
	}
	if err := srcFloater.CompleteFromKubeConfigPath(o.DoOption.SrcKubeConfig, ""); err != nil {
		return err
	}
	o.DoOption.SrcFloater = srcFloater

	if o.DoOption.DstKubeConfig == "" {
		o.DoOption.DstFloater = srcFloater
	} else {
		dstFloater := &share.Floater{
			Namespace:         o.DoOption.Namespace,
			Name:              share.DefaultFloaterName,
			ImageRepository:   o.DoOption.DstImageRepository,
			Version:           o.DoOption.Version,
			PodWaitTime:       o.DoOption.PodWaitTime,
			Port:              o.DoOption.Port,
			EnableHostNetwork: o.DoOption.GetEnableHostNetwork(false),
			EnableAnalysis:    false,
		}
		if err := dstFloater.CompleteFromKubeConfigPath(o.DoOption.DstKubeConfig, ""); err != nil {
			return err
		}
		o.DoOption.DstFloater = dstFloater
	}

	return nil
}

func (o *CheckOptions) Validate() error {
	if len(o.DoOption.Namespace) == 0 {
		return fmt.Errorf("namespace must be specified")
	}
	if len(o.DoOption.CustomizedTargetPortList) != 0 {
		for _, port := range o.DoOption.CustomizedTargetPortList {
			portInt, err := strconv.Atoi(port)
			if err != nil {
				return fmt.Errorf("invalid port: %s", port)
			} else if portInt <= 0 || portInt > 65535 {
				return fmt.Errorf("invalid port: %d", portInt)
			}
		}
		if len(o.DoOption.CustomizedTargetIPList) == 0 {
			return fmt.Errorf("if CustomizedTargetPortList is not null, CustomizedTargetIPList should be assigned")
		}
	}

	return nil
}

func (o *CheckOptions) Run() error {
	return o.DoOption.Run()
}
