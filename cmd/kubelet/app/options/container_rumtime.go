package options

import "k8s.io/kubernetes/pkg/kubelet/config"

const (
	// When these values are updated, also update test/utils/image/manifest.go
	defaultPodSandboxImageName    = "registry.k8s.io/pause"
	defaultPodSandboxImageVersion = "3.10"
)

var (
	defaultPodSandboxImage = defaultPodSandboxImageName +
		":" + defaultPodSandboxImageVersion
)

// NewContainerRuntimeOptions will create a new ContainerRuntimeOptions with
// default values.
func NewContainerRuntimeOptions() *config.ContainerRuntimeOptions {
	return &config.ContainerRuntimeOptions{
		PodSandboxImage: defaultPodSandboxImage,
	}
}
