package options

import (
	controlplane "k8s.io/kubernetes/pkg/controlplane/apiserver/options"
)

// completedOptions is a private wrapper that enforces a call of Complete() before Run can be invoked.
type completedOptions struct {
	controlplane.CompletedOptions
}

type CompletedOptions struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedOptions
}

// Complete set default ServerRunOptions.
// Should be called after kube-apiserver flags parsed.
func (opts *ServerRunOptions) Complete() (CompletedOptions, error) {
	return CompletedOptions{}, nil
}
