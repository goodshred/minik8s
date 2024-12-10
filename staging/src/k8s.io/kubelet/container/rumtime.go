package container

import (
	"context"
	runtimeapi "cri-api/pkg/apis/runtime/v1"
)

// Runtime interface defines the interfaces that should be implemented
// by a container runtime.
// Thread safety is required from implementations of this interface.
type Runtime interface {
	// GetPods returns a list of containers grouped by pods. The boolean parameter
	// specifies whether the runtime returns all containers including those already
	// exited and dead containers (used for garbage collection).
	GetPods(ctx context.Context, all bool) ([]*Pod, error)
	// ImageService provides methods to image-related methods.
	ImageService
}

// ImageService interfaces allows to work with image service.
type ImageService interface {
	// PullImage pulls an image from the network to local storage using the supplied
	// secrets if necessary. It returns a reference (digest or ID) to the pulled image.
	PullImage(ctx context.Context, image ImageSpec, pullSecrets []v1.Secret, podSandboxConfig *runtimeapi.PodSandboxConfig) (string, error)
}
