package utils

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestMapToString tests the MapToString function.
func TestMapToString(t *testing.T) {
	tests := []struct {
		name   string
		input  map[string]string
		output string
	}{
		{
			name:   "Empty map",
			input:  map[string]string{},
			output: "",
		},
		{
			name: "Single entry",
			input: map[string]string{
				"key1": "value1",
			},
			output: "key1=value1",
		},
		{
			name: "Multiple entries",
			input: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			output: "key1=value1,key2=value2",
		},
		{
			name: "Entries with special characters",
			input: map[string]string{
				"key with spaces": "value with spaces",
				"key:with:colons": "value:with:colons",
			},
			output: "key with spaces=value with spaces,key:with:colons=value:with:colons",
		},
		{
			name: "Entries with empty values",
			input: map[string]string{
				"key1": "",
				"key2": "value2",
			},
			output: "key1=,key2=value2",
		},
		{
			name: "Entries with numeric keys",
			input: map[string]string{
				"1": "value1",
				"2": "value2",
			},
			output: "1=value1,2=value2",
		},
		{
			name: "Long keys and values",
			input: map[string]string{
				"this_is_a_very_long_key_name": "this_is_a_very_long_value_name",
			},
			output: "this_is_a_very_long_key_name=this_is_a_very_long_value_name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapToString(tt.input)
			if result != tt.output {
				t.Errorf("expected %q, got %q", tt.output, result)
			}
		})
	}
}

// TestIsIPv6 tests the IsIPv6 function.
func TestIsIPv6(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		{
			name:   "Valid IPv6",
			input:  "2409:8c2f:3800::1",
			output: true,
		},
		{
			name:   "Invalid IPv4",
			input:  "192.168.1.1",
			output: false,
		},
		{
			name:   "IPv4 CIDR",
			input:  "192.168.1.0/24",
			output: false,
		},
		{
			name:   "Valid IPv6 with CIDR",
			input:  "2409:8c2f:3800::/64",
			output: true,
		},
		{
			name:   "Invalid empty string",
			input:  "",
			output: false,
		},
		{
			name:   "Valid mixed IPv6",
			input:  "2001:db8::1234:5678:abcd:ef00:1234",
			output: true,
		},
		{
			name:   "Invalid IPv6 with spaces",
			input:  "2409:8c2f: 3800::1",
			output: true,
		},
		{
			name:   "Valid IPv6 with full notation",
			input:  "2001:0db8:0000:0042:0000:8329:8a2e:0370:7334",
			output: true,
		},
		{
			name:   "Invalid mixed format",
			input:  "2001:db8::1234:5678:abcd:ef00:1234.5678",
			output: true,
		},
		{
			name:   "Valid short IPv6",
			input:  "::1",
			output: true,
		},
		{
			name:   "Valid minimal IPv6",
			input:  "0:0:0:0:0:0:0:1",
			output: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsIPv6(tt.input)
			if result != tt.output {
				t.Errorf("expected %v, got %v", tt.output, result)
			}
		})
	}
}

func TestPodStatus(t *testing.T) {
	tests := []struct {
		name     string
		pod      *corev1.Pod
		expected string
	}{
		{
			name: "Pod is running and container is not waiting",
			pod: &corev1.Pod{
				Status: corev1.PodStatus{
					Phase: corev1.PodRunning,
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{},
							},
						},
					},
				},
			},
			expected: "Running",
		},
		{
			name: "Pod is running and container is waiting",
			pod: &corev1.Pod{
				Status: corev1.PodStatus{
					Phase: corev1.PodRunning,
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Waiting: &corev1.ContainerStateWaiting{
									Reason: "ContainerCreating",
								},
							},
						},
					},
				},
			},
			expected: "ContainerCreating",
		},
		{
			name: "Pod is terminating",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					DeletionTimestamp: &metav1.Time{},
				},
				Status: corev1.PodStatus{
					Phase: corev1.PodRunning,
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{},
							},
						},
					},
				},
			},
			expected: "Running",
		},
		{
			name: "Pod is running but container has an error",
			pod: &corev1.Pod{
				Status: corev1.PodStatus{
					Phase: corev1.PodRunning,
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Waiting: nil,
							},
						},
					},
				},
			},
			expected: "Running",
		},
		{
			name: "Pod has waiting state in first container",
			pod: &corev1.Pod{
				Status: corev1.PodStatus{
					Phase: corev1.PodPending,
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Waiting: &corev1.ContainerStateWaiting{
									Reason: "ErrImagePull",
								},
							},
						},
					},
				},
			},
			expected: "ErrImagePull",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := podStatus(tt.pod)
			if status != tt.expected {
				t.Errorf("name %s, expected %q, got %q", tt.name, tt.expected, status)
			}
		})
	}
}
