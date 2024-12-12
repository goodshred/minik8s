package server

import "net/http"

type emptyDelegate struct {
	// handler is called at the end of the delegation chain
	// when a request has been made against an unregistered HTTP path the individual servers will simply pass it through until it reaches the handler.
	handler http.Handler
}

// DelegationTarget is an interface which allows for composition of API servers with top level handling that works
// as expected.
type DelegationTarget interface {
}

// GenericAPIServer contains state for a Kubernetes cluster api server.
type GenericAPIServer struct {
}

// NewEmptyDelegateWithCustomHandler allows for registering a custom handler usually for special handling of 404 requests
func NewEmptyDelegateWithCustomHandler(handler http.Handler) DelegationTarget {
	return emptyDelegate{handler}
}
