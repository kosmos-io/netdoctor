package cmd

import (
	"fmt"
	"strconv"

	"github.com/kosmos.io/netdoctor/pkg/command/share"
	"github.com/kosmos.io/netdoctor/pkg/utils"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	ctlutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var resumeExample = templates.Examples(i18n.T(`
        # re-Check cluster network connectivity, e.g:
        netdoctor resume 
`))

type ResumeOptions struct {
	DoOption *share.DoOptions
}

func NewResumeCmd() *cobra.Command {
	o := &ResumeOptions{}
	cmd := &cobra.Command{
		Use:                   "resume",
		Short:                 i18n.T("Resume check network connectivity"),
		Long:                  "",
		Example:               resumeExample,
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

func (o *ResumeOptions) LoadConfig() {
	fromConfig := &share.DoOptions{}
	if err := utils.ReadOpt(fromConfig); err == nil {
		once.Do(func() {
			klog.Infof("use config from file!!!!!!")
		})
		o.DoOption = fromConfig
	}
}

func (o *ResumeOptions) Complete() error {
	o.LoadConfig()
	if o.DoOption == nil {
		return fmt.Errorf("config.json load error")
	}

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

	var resumeData []*share.PrintCheckData

	err := utils.ReadResume(&resumeData)
	if err != nil {
		klog.Error("read resumeData error")
	}

	o.DoOption.ResumeRecord = resumeData

	return nil
}

func (o *ResumeOptions) Validate() error {
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
	}

	return nil
}

func (o *ResumeOptions) Run() error {
	return o.DoOption.Run()
}
