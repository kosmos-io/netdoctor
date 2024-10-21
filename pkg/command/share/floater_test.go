package share

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func compareStringArray(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func TestPodIPToArray(t *testing.T) {
	tests := []struct {
		name        string
		pods        []corev1.PodIP
		expectedCMs []string
		pass        bool
	}{
		{
			name: "Test Case 1",
			pods: []corev1.PodIP{
				{
					IP: "10.0.0.1",
				},
				{
					IP: "10.0.0.2",
				},
			},
			expectedCMs: []string{
				"10.0.0.1",
				"10.0.0.2",
			},
			pass: true,
		},
		{
			name: "Test Case 2",
			pods: []corev1.PodIP{
				{
					IP: "10.0.1.1",
				},
				{
					IP: "10.0.1.2",
				},
			},
			expectedCMs: []string{
				"10.0.1.1",
				"10.0.1.2",
			},
			pass: true,
		},
		{
			name: "Test Case 3",
			pods: []corev1.PodIP{
				{
					IP: "192.168.1.1",
				},
				{
					IP: "192.168.1.2",
				},
			},
			expectedCMs: []string{
				"192.168.1.1",
				"192.168.1.2",
			},
			pass: true,
		},
		{
			name: "Test Case 4",
			pods: []corev1.PodIP{
				{
					IP: "172.16.0.1",
				},
				{
					IP: "172.16.0.2",
				},
			},
			expectedCMs: []string{
				"172.16.0.1",
				"172.16.0.2",
			},
			pass: true,
		},
		{
			name: "Test Case 5",
			pods: []corev1.PodIP{
				{
					IP: "192.0.2.1",
				},
				{
					IP: "192.0.2.2",
				},
			},
			expectedCMs: []string{
				"192.0.2.1",
				"192.0.2.2",
			},
			pass: true,
		},
		{
			name: "Test Case 6",
			pods: []corev1.PodIP{
				{
					IP: "203.0.113.1",
				},
				{
					IP: "203.0.113.2",
				},
			},
			expectedCMs: []string{
				"203.0.113.1",
				"203.0.113.2",
			},
			pass: true,
		},
		{
			name: "Test Case 7",
			pods: []corev1.PodIP{
				{
					IP: "198.51.100.1",
				},
				{
					IP: "198.51.100.2",
				},
			},
			expectedCMs: []string{
				"198.51.100.1",
				"198.51.100.2",
			},
			pass: true,
		},
		{
			name: "Test Case 8",
			pods: []corev1.PodIP{
				{
					IP: "192.168.100.1",
				},
				{
					IP: "192.168.100.2",
				},
			},
			expectedCMs: []string{
				"192.168.100.1",
				"192.168.100.2",
			},
			pass: true,
		},
		{
			name: "Test Case 9",
			pods: []corev1.PodIP{
				{
					IP: "10.1.0.1",
				},
				{
					IP: "10.1.0.2",
				},
			},
			expectedCMs: []string{
				"10.1.0.1",
				"10.1.0.2",
			},
			pass: true,
		},
		{
			name: "Test Case 10",
			pods: []corev1.PodIP{
				{
					IP: "172.31.0.1",
				},
				{
					IP: "172.31.0.2",
				},
			},
			expectedCMs: []string{
				"172.31.0.1",
				"172.31.0.2",
			},
			pass: true,
		},
	}

	for _, test := range tests {
		if got := podIPToArray(test.pods); compareStringArray(got, test.expectedCMs) != test.pass {
			t.Errorf("PodIPToArray() = %v, want %v", got, test.expectedCMs)
		}
	}
}

func TestNodeIPToArray(t *testing.T) {
	tests := []struct {
		name        string
		node        corev1.Node
		expectedCMs []string
	}{
		{
			name: "Single Internal IP",
			node: corev1.Node{
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    corev1.NodeInternalIP,
							Address: "192.168.1.1",
						},
					},
				},
			},
			expectedCMs: []string{
				"192.168.1.1",
			},
		},
		{
			name: "Single External IP",
			node: corev1.Node{
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    corev1.NodeExternalIP,
							Address: "203.0.113.1",
						},
					},
				},
			},
			expectedCMs: []string{},
		},
		{
			name: "Multiple Addresses - Mixed Types",
			node: corev1.Node{
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    corev1.NodeInternalIP,
							Address: "192.168.1.1",
						},
						{
							Type:    corev1.NodeExternalIP,
							Address: "203.0.113.1",
						},
					},
				},
			},
			expectedCMs: []string{
				"192.168.1.1",
			},
		},
		{
			name: "No Addresses",
			node: corev1.Node{
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{},
				},
			},
			expectedCMs: []string{},
		},
		{
			name: "Internal and External IPs with Duplicates",
			node: corev1.Node{
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    corev1.NodeInternalIP,
							Address: "192.168.1.1",
						},
						{
							Type:    corev1.NodeInternalIP,
							Address: "192.168.1.1",
						},
						{
							Type:    corev1.NodeExternalIP,
							Address: "203.0.113.1",
						},
					},
				},
			},
			expectedCMs: []string{
				"192.168.1.1",
				"192.168.1.1",
			},
		},
		{
			name: "Multiple Internal IPs",
			node: corev1.Node{
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    corev1.NodeInternalIP,
							Address: "192.168.1.1",
						},
						{
							Type:    corev1.NodeInternalIP,
							Address: "192.168.1.2",
						},
					},
				},
			},
			expectedCMs: []string{
				"192.168.1.1",
				"192.168.1.2",
			},
		},
		{
			name: "Multiple External IPs",
			node: corev1.Node{
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    corev1.NodeExternalIP,
							Address: "203.0.113.1",
						},
						{
							Type:    corev1.NodeExternalIP,
							Address: "203.0.113.2",
						},
					},
				},
			},
			expectedCMs: []string{},
		},
		{
			name: "Internal and External IPs",
			node: corev1.Node{
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    corev1.NodeInternalIP,
							Address: "192.168.1.1",
						},
						{
							Type:    corev1.NodeExternalIP,
							Address: "203.0.113.1",
						},
						{
							Type:    corev1.NodeExternalIP,
							Address: "203.0.113.2",
						},
					},
				},
			},
			expectedCMs: []string{
				"192.168.1.1",
			},
		},
		{
			name: "Node with Hostname",
			node: corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node-1",
				},
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    corev1.NodeHostName,
							Address: "node-1.example.com",
						},
					},
				},
			},
			expectedCMs: []string{},
		},
		{
			name: "Multiple Addresses Including Hostname",
			node: corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node-1",
				},
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    corev1.NodeInternalIP,
							Address: "192.168.1.1",
						},
						{
							Type:    corev1.NodeHostName,
							Address: "node-1.example.com",
						},
					},
				},
			},
			expectedCMs: []string{
				"192.168.1.1",
			},
		},
		{
			name: "IPv6 Address",
			node: corev1.Node{
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    corev1.NodeInternalIP,
							Address: "2001:db8::1",
						},
					},
				},
			},
			expectedCMs: []string{
				"2001:db8::1",
			},
		},
		{
			name: "IPv4 and IPv6 Addresses",
			node: corev1.Node{
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    corev1.NodeInternalIP,
							Address: "192.168.1.1",
						},
						{
							Type:    corev1.NodeInternalIP,
							Address: "2001:db8::1",
						},
					},
				},
			},
			expectedCMs: []string{
				"192.168.1.1",
				"2001:db8::1",
			},
		},
		{
			name: "Node with Empty Address",
			node: corev1.Node{
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    corev1.NodeInternalIP,
							Address: "",
						},
					},
				},
			},
			expectedCMs: []string{
				"",
			},
		},
	}

	for _, test := range tests {
		if got := nodeIPToArray(test.node); !compareStringArray(got, test.expectedCMs) {
			t.Errorf("NodeIPToArray() = %v, want %v, name %v", got, test.expectedCMs, test.name)
		}
	}
}
