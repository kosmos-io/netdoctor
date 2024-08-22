package utils

const (
	DefaultNamespace       = "kosmos-system"
	DefaultImageRepository = "ghcr.io/kosmos-io"
	DefaultKubeConfigPath  = "~/.kube/config"
)

type Protocol string

const (
	TCP  Protocol = "tcp"
	UDP  Protocol = "udp"
	ICMP Protocol = "icmp"
	DNS  Protocol = "dns"
)
