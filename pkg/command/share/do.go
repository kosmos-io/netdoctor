package share

import (
	"fmt"
	"sync"

	command "github.com/kosmos.io/netdoctor/pkg/command/share/remote-command"
	"github.com/kosmos.io/netdoctor/pkg/utils"

	progressbar "github.com/schollz/progressbar/v3"
	"k8s.io/klog/v2"
)

type ModeType string

const (
	Node      ModeType = "node"
	Pod       ModeType = "pod"
	NodeToPod ModeType = "node-to-pod"
	PodToNode ModeType = "pod-to-node"
)

type DoOptions struct {
	Namespace string   `json:"namespace,omitempty"`
	Version   string   `json:"version,omitempty"`
	Mode      ModeType `json:"mode,omitempty"`

	Protocol                 string   `json:"protocol,omitempty"`
	PodWaitTime              int      `json:"podWaitTime,omitempty"`
	Port                     string   `json:"port,omitempty"`
	CustomizedTargetPortList []string `json:"customizedTargetPortList,omitempty"`
	CustomizedTargetIPList   []string `json:"customizedTargetIPList,omitempty"`
	TargetDNSServer          string   `json:"targetDNSServer,omitempty"`
	TargetHostToLookup       string   `json:"targetHostToLookup,omitempty"`

	MaxNum     int  `json:"maxNum,omitempty"`
	CmdTimeout int  `json:"cmdTimeout,omitempty"`
	AutoClean  bool `json:"autoClean,omitempty"`

	SrcKubeConfig      string `json:"srcKubeConfig,omitempty"`
	SrcImageRepository string `json:"srcImageRepository,omitempty"`
	DstKubeConfig      string `json:"dstKubeConfig,omitempty"`
	DstImageRepository string `json:"dstImageRepository,omitempty"`

	SrcFloater *Floater `json:"-"`
	DstFloater *Floater `json:"-"`

	ResumeRecord []*PrintCheckData `json:"-"`
}

func (o *DoOptions) GetEnableHostNetwork(isSrc bool) bool {
	if isSrc {
		switch o.Mode {
		case Node, NodeToPod:
			return true
		case Pod, PodToNode:
			return false
		}
	} else {
		switch o.Mode {
		case Node, PodToNode:
			return true
		case NodeToPod, Pod:
			return false
		}
	}
	return false
}

func (o *DoOptions) GetInfo(floater *Floater, isSrc bool) ([]*FloatInfo, error) {
	if isSrc {
		switch o.Mode {
		case Node, NodeToPod:
			return floater.GetNodesInfo()
		case Pod, PodToNode:
			return floater.GetPodInfo()
		}
	} else {
		switch o.Mode {
		case Node, PodToNode:
			return floater.GetNodesInfo()
		case NodeToPod, Pod:
			return floater.GetPodInfo()
		}
	}
	return floater.GetPodInfo()
}

func (o *DoOptions) Run() error {
	if err := o.SrcFloater.CreateFloater(); err != nil {
		return err
	}

	if o.DstKubeConfig != "" {
		srcInfos, err := o.GetInfo(o.SrcFloater, true)
		if err != nil {
			return fmt.Errorf("get src cluster podInfos failed: %s", err)
		}

		if err = o.DstFloater.CreateFloater(); err != nil {
			return err
		}
		var dstInfos []*FloatInfo
		dstInfos, err = o.GetInfo(o.DstFloater, false)
		if err != nil {
			return fmt.Errorf("get dist cluster podInfos failed: %s", err)
		}

		PrintResult(o.RunRange(srcInfos, dstInfos))
	} else {
		srcInfos, err := o.GetInfo(o.SrcFloater, true)
		if err != nil {
			return fmt.Errorf("get src cluster podInfos failed: %s", err)
		}

		var dstInfos []*FloatInfo
		dstInfos, err = o.GetInfo(o.SrcFloater, false)
		if err != nil {
			return fmt.Errorf("get dist cluster podInfos failed: %s", err)
		}

		PrintResult(o.RunRange(srcInfos, dstInfos))
	}

	if o.AutoClean {
		if err := o.Clean(); err != nil {
			return err
		}
	}

	// save options for resume
	o.SaveOpts()

	return nil
}

func (o *DoOptions) Clean() error {
	if err := o.SrcFloater.RemoveFloater(); err != nil {
		return err
	}

	if o.DstKubeConfig != "" {
		if err := o.DstFloater.RemoveFloater(); err != nil {
			return err
		}
	}
	return nil
}

func (o *DoOptions) SaveOpts() {
	if err := utils.WriteOpt(o); err != nil {
		klog.Fatal(err)
	} else {
		klog.Info("write opts success")
	}
}

func (o *DoOptions) SkipPod(podInfo *FloatInfo) bool {
	// is check:  no skip
	if len(o.ResumeRecord) == 0 {
		return false
	}
	// is resume: filt
	for _, r := range o.ResumeRecord {
		if r.SrcNodeName == podInfo.NodeName {
			return false
		}
	}
	return true
}

func (o *DoOptions) Skip(podInfo *FloatInfo, targetIP string) bool {
	// is check:  no skip
	if len(o.ResumeRecord) == 0 {
		return false
	}
	// is resume: filt
	for _, r := range o.ResumeRecord {
		if r.SrcNodeName == podInfo.NodeName && r.TargetIP == targetIP {
			return false
		}
	}
	return true
}

func (o *DoOptions) RunRange(iPodInfos []*FloatInfo, jPodInfos []*FloatInfo) []*PrintCheckData {
	var resultData []*PrintCheckData
	mutex := sync.Mutex{}

	var barctl *progressbar.ProgressBar

	if len(o.CustomizedTargetIPList) != 0 && len(o.CustomizedTargetPortList) != 0 ||
		o.Protocol == string(utils.DNS) {
		barctl = utils.NewBar(len(iPodInfos))
	} else {
		barctl = utils.NewBar(len(jPodInfos) * len(iPodInfos))
	}

	worker := func(iPodInfo *FloatInfo) {
		var cmdObj command.Command
		if len(o.CustomizedTargetIPList) != 0 && len(o.CustomizedTargetPortList) != 0 {
			cmdObj = command.NewCmd(o.Protocol, o.CustomizedTargetIPList, o.CustomizedTargetPortList)
		} else if o.Protocol == string(utils.DNS) {
			cmdObj = command.NewCmd(o.Protocol, o.TargetHostToLookup, o.TargetDNSServer)
		} else {
			for _, jPodInfo := range jPodInfos {
				targetIPs := jPodInfo.PodIPs
				needWrapper := false
				if len(jPodInfo.NodeIPs) != 0 {
					targetIPs = jPodInfo.NodeIPs
					needWrapper = true
				}
				for _, ip := range targetIPs {
					var targetIP string
					var err error
					var cmdResult *command.Result
					targetIP = ip
					if err != nil {
						cmdResult = command.ParseError(err)
					} else {
						// isSkip
						if o.Skip(iPodInfo, targetIP) {
							continue
						}
						// ToDo RunRange && RunNative func support multiple commands, and the code needs to be optimized
						cmdObj := command.NewCmd(o.Protocol, targetIP, o.Port)
						if needWrapper {
							cmdObj = command.Wrapper{
								Cmd: cmdObj,
							}
						}
						cmdResult = o.SrcFloater.CommandExec(iPodInfo, cmdObj)
					}
					mutex.Lock()
					resultData = append(resultData, &PrintCheckData{
						*cmdResult,
						iPodInfo.NodeName, jPodInfo.NodeName, targetIP,
					})
					mutex.Unlock()
				}
				err := barctl.Add(1)
				if err != nil {
					klog.Error("processs bar event add error")
				}
			}
			return
		}
		if o.SkipPod(iPodInfo) {
			return
		}
		cmdResult := o.SrcFloater.CommandExec(iPodInfo, cmdObj)
		mutex.Lock()
		resultData = append(resultData, &PrintCheckData{
			*cmdResult,
			iPodInfo.NodeName, iPodInfo.NodeName, cmdObj.GetTargetStr(),
		})
		mutex.Unlock()
		err := barctl.Add(1)
		if err != nil {
			klog.Error("processs bar event add error")
		}
	}

	var wg sync.WaitGroup
	ch := make(chan struct{}, o.MaxNum)

	if len(iPodInfos) > 0 && len(jPodInfos) > 0 {
		for _, iPodInfo := range iPodInfos {
			podInfo := iPodInfo
			ch <- struct{}{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				worker(podInfo)
				<-ch
			}()
		}
	}

	wg.Wait()

	return resultData
}
