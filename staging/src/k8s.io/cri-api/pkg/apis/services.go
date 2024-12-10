package apis

import (
	"context"
	runtimeapi "cri-api/pkg/apis/runtime/v1"
)

// RuntimeService interface should be implemented by a container runtime.
// The methods should be thread-safe.
type RuntimeService interface {
	ContainerManager
}

// ContainerManager contains methods to manipulate containers managed by a
// container runtime. The methods are thread-safe.
type ContainerManager interface {
	// CreateContainer creates a new container in specified PodSandbox.
	CreateContainer(ctx context.Context, podSandboxID string, config *runtimeapi.ContainerConfig, sandboxConfig *runtimeapi.PodSandboxConfig) (string, error)
	// StartContainer starts the container.
	StartContainer(ctx context.Context, containerID string) error
}
