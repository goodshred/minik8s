package apiserver

import genericapiserver "k8s.io/apiserver/pkg/server"

type completedConfig struct {
	Generic genericapiserver.CompletedConfig
}

// CompletedConfig embeds a private pointer that cannot be instantiated outside of this package
type CompletedConfig struct {
	*completedConfig
}
