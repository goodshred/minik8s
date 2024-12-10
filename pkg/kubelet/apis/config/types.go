package config

// KubeletConfiguration contains the configuration for the Kubelet
type KubeletConfiguration struct {
	// enableServer enables Kubelet's secured server.
	// Note: Kubelet's insecure port is controlled by the readOnlyPort option.
	EnableServer bool
}
