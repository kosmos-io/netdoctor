package cmd

import (
	"fmt"

	"sync"

	"github.com/kosmos.io/netdoctor/pkg/command/share"
	"github.com/kosmos.io/netdoctor/pkg/utils"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	ctlutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	once sync.Once
)

var checkExample = templates.Examples(i18n.T(`
        # Check cluster network connectivity, e.g:
        netdoctor check 
`))

type CheckOptions struct {
	DoOption *share.DoOptions
}

func NewCheckCmd() *cobra.Command {
	o := &CheckOptions{}
	cmd := &cobra.Command{
		Use:                   "check",
		Short:                 i18n.T("Check network connectivity"),
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
		once.Do(func() {
			klog.Infof("use config from file!!!!!!")
		})
		o.DoOption = fromConfig
	}
}

func (o *CheckOptions) Complete() error {
	o.LoadConfig()

	srcfloater := &share.Floater{
		Namespace:         o.DoOption.Namespace,
		Name:              share.DefaultFloaterName,
		ImageRepository:   o.DoOption.SrcImageRepository,
		Version:           o.DoOption.Version,
		PodWaitTime:       o.DoOption.PodWaitTime,
		Port:              o.DoOption.Port,
		EnableHostNetwork: false,
		EnableAnalysis:    false,
	}
	if err := srcfloater.CompleteFromKubeConfigPath(o.DoOption.SrcKubeConfig, ""); err != nil {
		return err
	}
	o.DoOption.SrcFloater = srcfloater

	if o.DoOption.DstKubeConfig == "" {
		o.DoOption.DstFloater = srcfloater
	} else {
		dstfloater := &share.Floater{
			Namespace:         o.DoOption.Namespace,
			Name:              share.DefaultFloaterName,
			ImageRepository:   o.DoOption.DstImageRepository,
			Version:           o.DoOption.Version,
			PodWaitTime:       o.DoOption.PodWaitTime,
			Port:              o.DoOption.Port,
			EnableHostNetwork: false,
			EnableAnalysis:    false,
		}
		if err := dstfloater.CompleteFromKubeConfigPath(o.DoOption.DstKubeConfig, ""); err != nil {
			return err
		}
		o.DoOption.DstFloater = dstfloater
	}

	return nil
}

func (o *CheckOptions) Validate() error {
	if len(o.DoOption.Namespace) == 0 {
		return fmt.Errorf("namespace must be specified")
	}

	return nil
}

func (o *CheckOptions) Run() error {
	return o.DoOption.Run()
}
