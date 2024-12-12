package app

import (
	apiextensionsapiserver "k8s.io/apiextensions-apiserver/pkg/apiserver"
	aggregatorapiserver "k8s.io/kube-aggregator/pkg/apiserver"
	"k8s.io/kubernetes/cmd/kube-apiserver/app/options"
	"k8s.io/kubernetes/pkg/controlplane"
)

type completedConfig struct {
	Options       options.CompletedOptions
	Aggregator    aggregatorapiserver.CompletedConfig
	KubeAPIs      controlplane.CompletedConfig
	ApiExtensions apiextensionsapiserver.CompletedConfig
}
type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}
