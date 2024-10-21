package command

import (
	"testing"
)

func TestNslookupGetTargetStr(t *testing.T) {
	tests := []struct {
		name     string
		nslookup *Nslookup
		expected string
	}{
		{
			name: "Default target and DNS server",
			nslookup: &Nslookup{
				TargetHost: "",
				DNSServer:  "",
			},
			expected: "host: dns:kubernetes.default.svc.cluster.local; dns: coredns",
		},
		{
			name: "Custom target host",
			nslookup: &Nslookup{
				TargetHost: "example.com",
				DNSServer:  "",
			},
			expected: "host: example.com; dns: coredns",
		},
		{
			name: "Custom DNS server",
			nslookup: &Nslookup{
				TargetHost: "",
				DNSServer:  "8.8.8.8",
			},
			expected: "host: dns:kubernetes.default.svc.cluster.local; dns: 8.8.8.8",
		},
		{
			name: "Custom target host and DNS server",
			nslookup: &Nslookup{
				TargetHost: "example.com",
				DNSServer:  "8.8.8.8",
			},
			expected: "host: example.com; dns: 8.8.8.8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.nslookup.GetTargetStr()
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestNslookupGetCommandStr(t *testing.T) {
	tests := []struct {
		name     string
		nslookup *Nslookup
		expected string
	}{
		{
			name: "Default target host and DNS server",
			nslookup: &Nslookup{
				TargetHost: "",
				DNSServer:  "",
			},
			expected: "nslookup kubernetes.default.svc.cluster.local ",
		},
		{
			name: "Custom target host",
			nslookup: &Nslookup{
				TargetHost: "example.com",
				DNSServer:  "",
			},
			expected: "nslookup example.com ",
		},
		{
			name: "Custom DNS server",
			nslookup: &Nslookup{
				TargetHost: "",
				DNSServer:  "8.8.8.8",
			},
			expected: "nslookup kubernetes.default.svc.cluster.local 8.8.8.8",
		},
		{
			name: "Custom target host and DNS server",
			nslookup: &Nslookup{
				TargetHost: "example.com",
				DNSServer:  "8.8.8.8",
			},
			expected: "nslookup example.com 8.8.8.8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.nslookup.GetCommandStr()
			if result != tt.expected {
				t.Errorf("%s expected %q, got %q", tt.name, tt.expected, result)
			}
		})
	}
}

func TestNslookupParseResult(t *testing.T) {
	tests := []struct {
		name     string
		nslookup *Nslookup
		result   string
		expected *Result
	}{
		{
			name: "Successful command",
			nslookup: &Nslookup{
				TargetHost: "example.com",
				DNSServer:  "coredns",
			},
			result: "Server:     coredns\nAddress:    10.96.0.10\n\nName:      example.com\nAddress 1: 93.184.216.34",
			expected: &Result{
				Status:    CommandSuccessed,
				ResultStr: "nslookup example.com coredns Server:     coredns\nAddress:    10.96.0.10\n\nName:      example.com\nAddress 1: 93.184.216.34",
			},
		},
		{
			name: "Command failed due to server not found",
			nslookup: &Nslookup{
				TargetHost: "example.com",
				DNSServer:  "coredns",
			},
			result: ";; connection timed out; no servers could be reached",
			expected: &Result{
				Status:    CommandFailed,
				ResultStr: "nslookup example.com coredns ;; connection timed out; no servers could be reached",
			},
		},
		{
			name: "Command failed due to server can't find",
			nslookup: &Nslookup{
				TargetHost: "nonexistent.example.com",
				DNSServer:  "coredns",
			},
			result: "Server:     coredns\nAddress:    10.96.0.10\n\n** server can't find nonexistent.example.com: NXDOMAIN",
			expected: &Result{
				Status:    CommandFailed,
				ResultStr: "nslookup nonexistent.example.com coredns Server:     coredns\nAddress:    10.96.0.10\n\n** server can't find nonexistent.example.com: NXDOMAIN",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.nslookup.ParseResult(tt.result)
			if result.Status != tt.expected.Status || result.ResultStr != tt.expected.ResultStr {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
