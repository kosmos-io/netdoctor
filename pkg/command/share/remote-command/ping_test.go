package command

import (
	"testing"
)

func TestPingGetTargetStr(t *testing.T) {
	tests := []struct {
		name     string
		ping     *Ping
		expected string
	}{
		{
			name: "Valid IPv4 target IP",
			ping: &Ping{
				TargetIP: "192.168.1.1",
			},
			expected: "192.168.1.1",
		},
		{
			name: "Valid IPv6 target IP",
			ping: &Ping{
				TargetIP: "2001:db8::ff00:42:8329",
			},
			expected: "2001:db8::ff00:42:8329",
		},
		{
			name: "Empty target IP",
			ping: &Ping{
				TargetIP: "",
			},
			expected: "",
		},
		{
			name: "Invalid IP format",
			ping: &Ping{
				TargetIP: "invalid_ip",
			},
			expected: "invalid_ip",
		},
		{
			name: "Localhost IP",
			ping: &Ping{
				TargetIP: "127.0.0.1",
			},
			expected: "127.0.0.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.ping.GetTargetStr()
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestPingGetCommandStr(t *testing.T) {
	tests := []struct {
		name     string
		ping     *Ping
		expected string
	}{
		{
			name: "IPv4 target IP",
			ping: &Ping{
				TargetIP: "192.168.1.1",
			},
			expected: "ping -c 1 192.168.1.1",
		},
		{
			name: "IPv6 target IP",
			ping: &Ping{
				TargetIP: "2001:db8::ff00:42:8329",
			},
			expected: "ping6 -c 1 2001:db8::ff00:42:8329",
		},
		{
			name: "Empty target IP",
			ping: &Ping{
				TargetIP: "",
			},
			expected: "ping -c 1 ", // Invalid command
		},
		{
			name: "Localhost IP",
			ping: &Ping{
				TargetIP: "127.0.0.1",
			},
			expected: "ping -c 1 127.0.0.1",
		},
		{
			name: "Another IPv4 target IP",
			ping: &Ping{
				TargetIP: "10.0.0.1",
			},
			expected: "ping -c 1 10.0.0.1",
		},
		{
			name: "Another IPv6 target IP",
			ping: &Ping{
				TargetIP: "fe80::1ff:fe23:4567:890a",
			},
			expected: "ping6 -c 1 fe80::1ff:fe23:4567:890a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.ping.GetCommandStr()
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestPingParseResult(t *testing.T) {
	tests := []struct {
		name     string
		ping     *Ping
		result   string
		expected *Result
	}{
		{
			name: "Successful ping result",
			ping: &Ping{
				TargetIP: "192.168.1.1",
			},
			result: "PING 192.168.1.1 (192.168.1.1): 1 data byte\n64 bytes from 192.168.1.1: icmp_seq=0 ttl=64 time=0.123 ms\n\n--- 192.168.1.1 ping statistics ---\n1 packets transmitted, 1 packets received, 0% packet loss",
			expected: &Result{
				Status:    CommandSuccessed,
				ResultStr: "ping -c 1 192.168.1.1 PING 192.168.1.1 (192.168.1.1): 1 data byte\n64 bytes from 192.168.1.1: icmp_seq=0 ttl=64 time=0.123 ms\n\n--- 192.168.1.1 ping statistics ---\n1 packets transmitted, 1 packets received, 0% packet loss",
			},
		},
		{
			name: "Failed ping result - no response",
			ping: &Ping{
				TargetIP: "192.168.1.1",
			},
			result: "PING 192.168.1.1 (192.168.1.1): 1 data byte\n\n--- 192.168.1.1 ping statistics ---\n1 packets transmitted, 0 packets received, 100% packet loss",
			expected: &Result{
				Status:    CommandFailed,
				ResultStr: "ping -c 1 192.168.1.1 PING 192.168.1.1 (192.168.1.1): 1 data byte\n\n--- 192.168.1.1 ping statistics ---\n1 packets transmitted, 0 packets received, 100% packet loss",
			},
		},
		{
			name: "Unexpected result format",
			ping: &Ping{
				TargetIP: "192.168.1.1",
			},
			result: "Some random output that does not match expected format",
			expected: &Result{
				Status:    CommandFailed,
				ResultStr: "ping -c 1 192.168.1.1 Some random output that does not match expected format",
			},
		},
		{
			name: "Connection timed out",
			ping: &Ping{
				TargetIP: "192.168.1.1",
			},
			result: "ping: connect: Network is unreachable",
			expected: &Result{
				Status:    CommandFailed,
				ResultStr: "ping -c 1 192.168.1.1 ping: connect: Network is unreachable",
			},
		},
		{
			name: "Server not found",
			ping: &Ping{
				TargetIP: "192.168.1.1",
			},
			result: "ping: 192.168.1.1: Name or service not known",
			expected: &Result{
				Status:    CommandFailed,
				ResultStr: "ping -c 1 192.168.1.1 ping: 192.168.1.1: Name or service not known",
			},
		},
		{
			name: "Valid response but with loss",
			ping: &Ping{
				TargetIP: "192.168.1.1",
			},
			result: "PING 192.168.1.1 (192.168.1.1): 1 data byte\n64 bytes from 192.168.1.1: icmp_seq=0 ttl=64 time=0.123 ms\n\n--- 192.168.1.1 ping statistics ---\n1 packets transmitted, 0 packets received, 100% packet loss",
			expected: &Result{
				Status:    CommandFailed,
				ResultStr: "ping -c 1 192.168.1.1 PING 192.168.1.1 (192.168.1.1): 1 data byte\n64 bytes from 192.168.1.1: icmp_seq=0 ttl=64 time=0.123 ms\n\n--- 192.168.1.1 ping statistics ---\n1 packets transmitted, 0 packets received, 100% packet loss",
			},
		},
		{
			name: "Successful ping with packet loss",
			ping: &Ping{
				TargetIP: "192.168.1.1",
			},
			result: "PING 192.168.1.1 (192.168.1.1): 1 data byte\n64 bytes from 192.168.1.1: icmp_seq=0 ttl=64 time=0.123 ms\n\n--- 192.168.1.1 ping statistics ---\n2 packets transmitted, 1 packets received, 50% packet loss",
			expected: &Result{
				Status:    CommandFailed,
				ResultStr: "ping -c 1 192.168.1.1 PING 192.168.1.1 (192.168.1.1): 1 data byte\n64 bytes from 192.168.1.1: icmp_seq=0 ttl=64 time=0.123 ms\n\n--- 192.168.1.1 ping statistics ---\n2 packets transmitted, 1 packets received, 50% packet loss",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.ping.ParseResult(tt.result)
			if result.Status != tt.expected.Status {
				t.Errorf("%s expected %+v, got %+v", tt.name, tt.expected, result)
			}
		})
	}
}
