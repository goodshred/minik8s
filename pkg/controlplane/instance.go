package controlplane

import (
	genericapiserver "k8s.io/apiserver/pkg/server"
	controlplaneapiserver "k8s.io/kubernetes/pkg/controlplane/apiserver"
)

type completedConfig struct {
	ControlPlane controlplaneapiserver.CompletedConfig
}

// CompletedConfig embeds a private pointer that cannot be instantiated outside of this package
type CompletedConfig struct {
	*completedConfig
}

// Instance contains state for a Kubernetes cluster api server instance.
type Instance struct {
	ControlPlane *controlplaneapiserver.Server
}

func (c CompletedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*Instance, error) {

}
