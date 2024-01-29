package utils

const (
	DefaultNamespace       = "kosmos-system"
	DefaultImageRepository = "ghcr.io/kosmos-io"
)

type Protocol string

const (
	TCP  Protocol = "tcp"
	UDP  Protocol = "udp"
	IPv4 Protocol = "ipv4"
)
