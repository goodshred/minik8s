package apiserver

import genericapiserver "k8s.io/apiserver/pkg/server"

// Server is a struct that contains a generic control plane apiserver instance
// that can be run to start serving the APIs.
type Server struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}
