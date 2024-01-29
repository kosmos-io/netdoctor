package check

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	ctlutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"

	"github.com/kosmos.io/netdoctor/pkg/command"
	"github.com/kosmos.io/netdoctor/pkg/utils"
	"github.com/kosmos.io/netdoctor/pkg/version"
)

var singleCheckExample = templates.Examples(i18n.T(`
        # Check single cluster network connectivity, e.g:
        netctl check --kubeconfig ~/kubeconfig/cluster-kubeconfig
        
        # Check cluster network connectivity, if you need to specify a special image repository, e.g: 
        netctl check -r ghcr.io/kosmos-io
`))

type CommandSingleCheckOptions struct {
	Namespace       string
	ImageRepository string
	Version         string

	Protocol    string
	PodWaitTime int
	Port        string
	HostNetwork bool

	KubeConfig string
	Context    string

	Floater *Floater
}

type PrintSingleCheckData struct {
	command.Result
	SrcNodeName string
	DstNodeName string
	TargetIP    string
}

func NewCmdSingleCheck() *cobra.Command {
	o := &CommandSingleCheckOptions{
		Version: version.GetReleaseVersion().PatchRelease(),
	}
	cmd := &cobra.Command{
		Use:                   "check",
		Short:                 i18n.T("Check single-cluster network connectivity"),
		Long:                  "",
		Example:               singleCheckExample,
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

	flags := cmd.Flags()
	flags.StringVarP(&o.Namespace, "namespace", "n", utils.DefaultNamespace, "Kosmos namespace.")
	flags.StringVarP(&o.ImageRepository, "image-repository", "r", utils.DefaultImageRepository, "Image repository.")
	flags.StringVar(&o.KubeConfig, "kubeconfig", "", "Absolute path to the host kubeconfig file.")
	flags.StringVar(&o.Context, "context", "", "The name of the kubeconfig context.")
	flags.BoolVar(&o.HostNetwork, "host-network", false, "Configure HostNetwork.")
	flags.StringVar(&o.Port, "port", "8889", "Port used by floater.")
	flags.IntVarP(&o.PodWaitTime, "pod-wait-time", "w", 30, "Time for wait pod(floater) launch.")
	flags.StringVar(&o.Protocol, "protocol", string(utils.TCP), "Protocol for the network problem.")

	return cmd
}

func (o *CommandSingleCheckOptions) Complete() error {
	f := NewSingleFloater(o)
	if err := f.CompleteFromKubeConfigPath(o.KubeConfig, o.Context); err != nil {
		return err
	}
	o.Floater = f

	return nil
}

func (o *CommandSingleCheckOptions) Validate() error {
	if len(o.Namespace) == 0 {
		return fmt.Errorf("namespace must be specified")
	}

	return nil
}

func (o *CommandSingleCheckOptions) Run() error {
	var resultData []*PrintCheckData

	if err := o.Floater.CreateFloater(); err != nil {
		return err
	}

	if o.Floater.EnableHostNetwork {
		srcNodeInfos, err := o.Floater.GetNodesInfo()
		if err != nil {
			return fmt.Errorf("get src cluster nodeInfos failed: %s", err)
		}
		resultData = o.RunNative(srcNodeInfos, srcNodeInfos)
	} else {
		srcPodInfos, err := o.Floater.GetPodInfo()
		if err != nil {
			return fmt.Errorf("get src cluster podInfos failed: %s", err)
		}
		resultData = o.RunRange(srcPodInfos, srcPodInfos)
	}

	o.PrintResult(resultData)

	if err := o.Floater.RemoveFloater(); err != nil {
		return err
	}

	return nil
}

func (o *CommandSingleCheckOptions) RunRange(iPodInfos []*FloatInfo, jPodInfos []*FloatInfo) []*PrintCheckData {
	var resultData []*PrintCheckData

	if len(iPodInfos) > 0 && len(jPodInfos) > 0 {
		for _, iPodInfo := range iPodInfos {
			for _, jPodInfo := range jPodInfos {
				for _, ip := range jPodInfo.PodIPs {
					var targetIP string
					var cmdResult *command.Result
					// ToDo RunRange && RunNative func support multiple commands, and the code needs to be optimized
					cmdObj := &command.Ping{
						TargetIP: ip,
					}
					cmdResult = o.Floater.CommandExec(iPodInfo, cmdObj)
					resultData = append(resultData, &PrintCheckData{
						*cmdResult,
						iPodInfo.NodeName, jPodInfo.NodeName, targetIP,
					})
				}
			}
		}
	}

	return resultData
}

func (o *CommandSingleCheckOptions) RunNative(iNodeInfos []*FloatInfo, jNodeInfos []*FloatInfo) []*PrintCheckData {
	var resultData []*PrintCheckData

	if len(iNodeInfos) > 0 && len(jNodeInfos) > 0 {
		for _, iNodeInfo := range iNodeInfos {
			for _, jNodeInfo := range jNodeInfos {
				for _, ip := range jNodeInfo.NodeIPs {
					// ToDo RunRange && RunNative func support multiple commands, and the code needs to be optimized
					cmdObj := &command.Ping{
						TargetIP: ip,
					}
					cmdResult := o.Floater.CommandExec(iNodeInfo, cmdObj)
					resultData = append(resultData, &PrintCheckData{
						*cmdResult,
						iNodeInfo.NodeName, jNodeInfo.NodeName, ip,
					})
				}
			}
		}
	}

	return resultData
}

func (o *CommandSingleCheckOptions) PrintResult(resultData []*PrintCheckData) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"S/N", "SRC_NODE_NAME", "DST_NODE_NAME", "TARGET_IP", "RESULT"})

	tableException := tablewriter.NewWriter(os.Stdout)
	tableException.SetHeader([]string{"S/N", "SRC_NODE_NAME", "DST_NODE_NAME", "TARGET_IP", "RESULT"})

	for index, r := range resultData {
		// klog.Infof(fmt.Sprintf("%s %s %v", r.SrcNodeName, r.DstNodeName, r.IsSucceed))
		row := []string{strconv.Itoa(index + 1), r.SrcNodeName, r.DstNodeName, r.TargetIP, command.PrintStatus(r.Status)}
		if r.Status == command.CommandFailed {
			table.Rich(row, []tablewriter.Colors{
				{},
				{tablewriter.Bold, tablewriter.FgHiRedColor},
				{tablewriter.Bold, tablewriter.FgHiRedColor},
				{tablewriter.Bold, tablewriter.FgHiRedColor},
				{tablewriter.Bold, tablewriter.FgHiRedColor},
			})
		} else if r.Status == command.ExecError {
			tableException.Rich(row, []tablewriter.Colors{
				{},
				{tablewriter.Bold, tablewriter.FgCyanColor},
				{tablewriter.Bold, tablewriter.FgCyanColor},
				{tablewriter.Bold, tablewriter.FgCyanColor},
				{tablewriter.Bold, tablewriter.FgCyanColor},
			})
		} else {
			table.Rich(row, []tablewriter.Colors{
				{},
				{tablewriter.Bold, tablewriter.FgGreenColor},
				{tablewriter.Bold, tablewriter.FgGreenColor},
				{tablewriter.Bold, tablewriter.FgGreenColor},
				{tablewriter.Bold, tablewriter.FgGreenColor},
			})
		}
	}
	fmt.Println("")
	table.Render()
	fmt.Println("")
	tableException.Render()
}
