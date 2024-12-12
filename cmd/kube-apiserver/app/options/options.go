package options

import (
	controlplaneapiserver "k8s.io/kubernetes/pkg/controlplane/apiserver/options"
)

type ServerRunOptions struct {
	*controlplaneapiserver.Options // embedded to avoid noise in existing consumers
}

// NewServerRunOptions creates and returns ServerRunOptions according to the given featureGate and effectiveVersion of the server binary to run.
func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{
		Options: controlplaneapiserver.NewOptions(),
	}
	return &s
}
