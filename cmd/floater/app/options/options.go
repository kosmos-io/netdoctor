package options

import (
	"github.com/spf13/pflag"
)

type Options struct {
	KubeConfig string
}

// NewOptions builds a default agent options.
func NewOptions() *Options {
	return &Options{}
}

// AddFlags adds flags of agent to the specified FlagSet
func (o *Options) AddFlags(fs *pflag.FlagSet) {
}
