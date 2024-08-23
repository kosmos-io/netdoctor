package cmd

import (
	"fmt"

	"github.com/kosmos.io/netdoctor/pkg/command/share"
	"github.com/kosmos.io/netdoctor/pkg/utils"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	ctlutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var cleanExample = templates.Examples(i18n.T(`
        # re-Check cluster network connectivity, e.g:
        netdoctor clean 
`))

type CleanOptions struct {
	DoOption *share.DoOptions
}

func NewCleanCmd() *cobra.Command {
	o := &CleanOptions{}
	cmd := &cobra.Command{
		Use:                   "clean",
		Short:                 i18n.T("Clean check network connectivity"),
		Long:                  "",
		Example:               cleanExample,
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

func (o *CleanOptions) LoadConfig() {
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

func (o *CleanOptions) Complete() error {
	o.LoadConfig()

	srcfloater := &share.Floater{
		Namespace:         o.DoOption.Namespace,
		Name:              share.DefaultFloaterName,
		ImageRepository:   o.DoOption.SrcImageRepository,
		Version:           o.DoOption.Version,
		PodWaitTime:       o.DoOption.PodWaitTime,
		Port:              o.DoOption.Port,
		EnableHostNetwork: o.DoOption.GetEnableHostNetwork(true),
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
			EnableHostNetwork: o.DoOption.GetEnableHostNetwork(false),
			EnableAnalysis:    false,
		}
		if err := dstfloater.CompleteFromKubeConfigPath(o.DoOption.DstKubeConfig, ""); err != nil {
			return err
		}
		o.DoOption.DstFloater = dstfloater
	}

	return nil
}

func (o *CleanOptions) Validate() error {
	if len(o.DoOption.Namespace) == 0 {
		return fmt.Errorf("namespace must be specified")
	}

	return nil
}

func (o *CleanOptions) Run() error {
	return o.DoOption.Clean()
}
