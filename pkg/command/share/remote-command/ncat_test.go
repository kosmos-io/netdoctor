package command

import (
	"fmt"
	"testing"
)

func TestCmdNcat(t *testing.T) {

	var tests []struct {
		name    string
		ncatCmd Ncat
		want    struct {
			cmd    string
			target string
		}
	}

	// Create 500 random test cases for IPv4
	for i := 0; i < 500; i++ {
		ipv4 := randomIPv4()
		port := randomPort()
		tests = append(tests, struct {
			name    string
			ncatCmd Ncat
			want    struct {
				cmd    string
				target string
			}
		}{
			name: fmt.Sprintf("ipv4_test_%d", i+1),
			ncatCmd: Ncat{
				TargetIP: []string{ipv4},
				Port:     []string{port},
			},
			want: struct {
				cmd    string
				target string
			}{
				cmd:    fmt.Sprintf("nc -w 1 -z -d  -v %s %s 2>&1", ipv4, port),
				target: fmt.Sprintf("IPs: %s; Ports: %s", ipv4, port),
			},
		})
	}

	// Create 500 random test cases for IPv6
	for i := 0; i < 500; i++ {
		ipv6 := randomIPv6()
		port := randomPort()
		tests = append(tests, struct {
			name    string
			ncatCmd Ncat
			want    struct {
				cmd    string
				target string
			}
		}{
			name: fmt.Sprintf("ipv6_test_%d", i+1),
			ncatCmd: Ncat{
				TargetIP: []string{ipv6},
				Port:     []string{port},
			},
			want: struct {
				cmd    string
				target string
			}{
				cmd:    fmt.Sprintf("nc -w 1 -z -d  -v %s %s 2>&1", ipv6, port),
				target: fmt.Sprintf("IPs: %s; Ports: %s", ipv6, port),
			},
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ncatCmd.GetCommandStr() != tt.want.cmd {
				t.Errorf("%s, %s, %s, %s", tt.name, tt.ncatCmd, tt.ncatCmd.GetCommandStr(), tt.want)
			}

			if tt.ncatCmd.GetTargetStr() != tt.want.target {
				t.Errorf("%s, %s, %s, %s", tt.name, tt.ncatCmd, tt.ncatCmd.GetCommandStr(), tt.want)
			}
		})
	}
}
