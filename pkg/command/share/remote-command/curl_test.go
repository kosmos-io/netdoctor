package command

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"testing"
)

// Helper function to generate random IPv4 address
func randomIPv4() string {
	return fmt.Sprintf("%d.%d.%d.%d", randInt(1, 255), randInt(0, 255), randInt(0, 255), randInt(0, 255))
}

// Helper function to generate random IPv6 address
func randomIPv6() string {
	ip := net.ParseIP(fmt.Sprintf("2409:8c2f:3800::%x", randInt(1, 10000)))
	return ip.String()
}

// Helper function to generate random port
func randomPort() string {
	return strconv.Itoa(randInt(1000, 9999))
}

// Generate random number between min and max
func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func TestCmdCurl(t *testing.T) {

	var tests []struct {
		name    string
		curlCmd Curl
		want    struct {
			curl   string
			target string
		}
	}

	// Create 500 random test cases for IPv4
	for i := 0; i < 500; i++ {
		ipv4 := randomIPv4()
		port := randomPort()
		tests = append(tests, struct {
			name    string
			curlCmd Curl
			want    struct {
				curl   string
				target string
			}
		}{
			name: fmt.Sprintf("ipv4_test_%d", i+1),
			curlCmd: Curl{
				TargetIP: ipv4,
				Port:     port,
			},
			want: struct {
				curl   string
				target string
			}{
				curl:   fmt.Sprintf("curl -k http://%s:%s/", ipv4, port),
				target: fmt.Sprintf("%s:%s", ipv4, port),
			},
		})
	}

	// Create 500 random test cases for IPv6
	for i := 0; i < 500; i++ {
		ipv6 := randomIPv6()
		port := randomPort()
		tests = append(tests, struct {
			name    string
			curlCmd Curl
			want    struct {
				curl   string
				target string
			}
		}{
			name: fmt.Sprintf("ipv6_test_%d", i+1),
			curlCmd: Curl{
				TargetIP: ipv6,
				Port:     port,
			},
			want: struct {
				curl   string
				target string
			}{
				curl:   fmt.Sprintf("curl -k http://[%s]:%s/", ipv6, port),
				target: fmt.Sprintf("%s:%s", ipv6, port),
			},
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.curlCmd.GetCommandStr() != tt.want.curl {
				t.Errorf("%s, %s, %s, %s", tt.name, tt.curlCmd, tt.curlCmd.GetCommandStr(), tt.want)
			}

			if tt.curlCmd.GetTargetStr() != tt.want.target {
				t.Errorf("%s, %s, %s, %s", tt.name, tt.curlCmd, tt.curlCmd.GetCommandStr(), tt.want)
			}
		})
	}
}
