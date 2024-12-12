package apiserver

import (
	externalinformers "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
}
type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

type CustomResourceDefinitions struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
	// provided for easier embedding
	Informers externalinformers.SharedInformerFactory
}

func (c completedConfig) New(delegationTarget genericapiserver.DelegationTarget) (*CustomResourceDefinitions, error) {
	return nil, nil
}
