package apiserver

import (
	genericapiserver "k8s.io/apiserver/pkg/server"
)

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
}

// CompletedConfig same as Config, just to swap private object.
type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

// APIAggregator contains state for a Kubernetes cluster master/api server.
type APIAggregator struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
