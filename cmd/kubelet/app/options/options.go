package options

import (
	kubeletconfig "k8s.io/kubernetes/pkg/kubelet/apis/config"
	"k8s.io/kubernetes/pkg/kubelet/config"
)

type KubeletFlags struct {
	KubeConfig string
	// NodeIP is IP address of the node.
	// If set, kubelet will use this IP address for the node.
	NodeIP string
	// Container-runtime-specific options.
	config.ContainerRuntimeOptions
	// Node Labels are the node labels to add when registering the node in the cluster
	NodeLabels map[string]string
}

// KubeletServer encapsulates all of the parameters necessary for starting up
// a kubelet. These can either be set via command line or directly.
type KubeletServer struct {
	KubeletFlags
	kubeletconfig.KubeletConfiguration
}

// NewKubeletFlags will create a new KubeletFlags with default values
func NewKubeletFlags() *KubeletFlags {
	return &KubeletFlags{
		ContainerRuntimeOptions: *NewContainerRuntimeOptions(),
		NodeLabels:              make(map[string]string),
	}
}

// NewKubeletConfiguration will create a new KubeletConfiguration with default values
func NewKubeletConfiguration() (*kubeletconfig.KubeletConfiguration, error) {
	//scheme, _, err := kubeletscheme.NewSchemeAndCodecs()
	//if err != nil {
	//	return nil, err
	//}
	//versioned := &v1beta1.KubeletConfiguration{}
	//scheme.Default(versioned)
	//if err := scheme.Convert(versioned, config, nil); err != nil {
	//	return nil, err
	//}
	//applyLegacyDefaults(config)
	config := &kubeletconfig.KubeletConfiguration{}
	return config, nil
}
