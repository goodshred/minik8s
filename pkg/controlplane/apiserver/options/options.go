package options

import (
	genericoptions "k8s.io/apiserver/pkg/server/options"
)

type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions
}

// completedServerRunOptions is a private wrapper that enforces a call of Complete() before Run can be invoked.
type completedOptions struct {
	Options
}

type CompletedOptions struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedOptions
}

// NewOptions creates a new ServerRunOptions object with default parameters
func NewOptions() *Options {
	s := Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
	}
	return &s
}
