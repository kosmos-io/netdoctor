package share

import (
	"fmt"
	"sync"

	command "github.com/kosmos.io/netdoctor/pkg/command/share/remote-command"
	"github.com/kosmos.io/netdoctor/pkg/utils"
	"k8s.io/klog/v2"
)

type DoOptions struct {
	Namespace string `json:"namespace,omitempty"`
	Version   string `json:"version,omitempty"`

	Protocol    string `json:"protocol,omitempty"`
	PodWaitTime int    `json:"podWaitTime,omitempty"`
	Port        string `json:"port,omitempty"`

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

func (o *DoOptions) Run() error {
	if err := o.SrcFloater.CreateFloater(); err != nil {
		return err
	}

	if o.DstKubeConfig != "" {
		srcPodInfos, err := o.SrcFloater.GetPodInfo()
		if err != nil {
			return fmt.Errorf("get src cluster podInfos failed: %s", err)
		}

		if err = o.DstFloater.CreateFloater(); err != nil {
			return err
		}
		var dstPodInfos []*FloatInfo
		dstPodInfos, err = o.DstFloater.GetPodInfo()
		if err != nil {
			return fmt.Errorf("get dist cluster podInfos failed: %s", err)
		}

		PrintResult(o.RunRange(srcPodInfos, dstPodInfos))
	} else {
		srcPodInfos, err := o.SrcFloater.GetPodInfo()
		if err != nil {
			return fmt.Errorf("get src cluster podInfos failed: %s", err)
		}
		PrintResult(o.RunRange(srcPodInfos, srcPodInfos))
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

	barctl := utils.NewBar(len(jPodInfos) * len(iPodInfos))

	worker := func(iPodInfo *FloatInfo) {
		for _, jPodInfo := range jPodInfos {
			for _, ip := range jPodInfo.PodIPs {
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
					cmdResult = o.SrcFloater.CommandExec(iPodInfo, cmdObj)
				}
				mutex.Lock()
				resultData = append(resultData, &PrintCheckData{
					*cmdResult,
					iPodInfo.NodeName, jPodInfo.NodeName, targetIP,
				})
				mutex.Unlock()
			}
			barctl.Add(1)
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
