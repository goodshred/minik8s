package kuberuntime

import internalapi "cri-api/pkg/apis"

// KubeGenericRuntime is a interface contains interfaces for container runtime and command.
type KubeGenericRuntime interface {
	kubecontainer.Runtime
	kubecontainer.StreamingRuntime
	kubecontainer.CommandRunner
}

type kubeGenericRuntimeManager struct {
	// gRPC service clients
	runtimeService internalapi.RuntimeService
}
// NewKubeGenericRuntimeManager creates a new kubeGenericRuntimeManager
func NewKubeGenericRuntimeManager(runtimeService internalapi.RuntimeService)(KubeGenericRuntime, error) {

	kubeRuntimeManager := &kubeGenericRuntimeManager{

	}
}
	return kubeRuntimeManager, nil
