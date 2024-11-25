package share

import (
	"bytes"
	"testing"

	command "github.com/kosmos.io/netdoctor/pkg/command/share/remote-command"
)

func TestPrintResult(t *testing.T) {

	tests := []struct {
		checkData      []*PrintCheckData
		containsString string
	}{
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
			},
			containsString: command.PrintStatus(command.CommandSuccessed),
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: command.PrintStatus(command.CommandFailed),
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: "192.168.8.1",
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: "Node14",
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: command.PrintStatus(command.ExecError),
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: "192.168.3.1",
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "target ip",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: "target ip",
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: "192.168.11.1",
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: "192.168.7.1",
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: "192.168.11.1",
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: "192.168.4.1",
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "NodeTarget",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: "NodeTarget",
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.100.3",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.100.1",
				},
			},
			containsString: "192.168.100.3",
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "255.255.2555.255",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "255.255.2555.0",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: "255.255.2555.255",
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Node14",
					TargetIP:    "192.168.8.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "NodeTarget",
					DstNodeName: "Node16",
					TargetIP:    "255.255.2555.0",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
			},
			containsString: "255.255.2555.0",
		},
		{
			checkData: []*PrintCheckData{
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node1",
					DstNodeName: "Node2",
					TargetIP:    "192.168.1.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node3",
					DstNodeName: "Node4",
					TargetIP:    "192.168.2.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node5",
					DstNodeName: "Node6",
					TargetIP:    "192.168.3.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node7",
					DstNodeName: "Node8",
					TargetIP:    "192.168.4.1",
				},
				{
					Result: command.Result{
						Status:    command.ExecError,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node9",
					DstNodeName: "Node10",
					TargetIP:    "192.168.5.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node15",
					DstNodeName: "Node16",
					TargetIP:    "192.168.9.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node17",
					DstNodeName: "Node18",
					TargetIP:    "192.168.10.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node21",
					DstNodeName: "Node20",
					TargetIP:    "192.168.11.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandSuccessed,
						ResultStr: "Command executed successfully",
					},
					SrcNodeName: "Node22",
					DstNodeName: "Node23",
					TargetIP:    "192.168.12.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node11",
					DstNodeName: "Node12",
					TargetIP:    "192.168.7.1",
				},
				{
					Result: command.Result{
						Status:    command.CommandFailed,
						ResultStr: "Command failed to execute",
					},
					SrcNodeName: "Node13",
					DstNodeName: "Nodetarget",
					TargetIP:    "192.168.8.1",
				},
			},
			containsString: "Nodetarget",
		},
	}

	for _, test := range tests {
		var outputBuffer bytes.Buffer

		PrintResultWithWriter(test.checkData, &outputBuffer)
		output := outputBuffer.String()

		if !contains(output, test.containsString) {
			t.Errorf("Expected output to contain '%s', got: %s", test.containsString, output)
		}
	}

}

func contains(output, substr string) bool {
	return bytes.Contains([]byte(output), []byte(substr))
}
