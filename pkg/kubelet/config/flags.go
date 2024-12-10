package config

// ContainerRuntimeOptions defines options for the container runtime.
type ContainerRuntimeOptions struct {
	// General Options.

	// RuntimeCgroups that container runtime is expected to be isolated in.
	RuntimeCgroups string
	// PodSandboxImage is the image whose network/ipc namespaces
	// containers in each pod will use.
	PodSandboxImage string
	// Image credential provider plugin options

	// ImageCredentialProviderConfigFile is the path to the credential provider plugin config file.
	// This config file is a specification for what credential providers are enabled and invoked
	// by the kubelet. The plugin config should contain information about what plugin binary
	// to execute and what container images the plugin should be called for.
	// +optional
	ImageCredentialProviderConfigFile string
	// ImageCredentialProviderBinDir is the path to the directory where credential provider plugin
	// binaries exist. The name of each plugin binary is expected to match the name of the plugin
	// specified in imageCredentialProviderConfigFile.
	// +optional
	ImageCredentialProviderBinDir string
}
