package notfoundhandler

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"
)

// New returns an HTTP handler that is meant to be executed at the end of the delegation chain.
// It checks if the request have been made before the server has installed all known HTTP paths.
// In that case it returns a 503 response otherwise it returns a 404.
//
// Note that we don't want to add additional checks to the readyz path as it might prevent fixing bricked clusters.
// This specific handler is meant to "protect" requests that arrive before the paths and handlers are fully initialized.
func New(serializer runtime.NegotiatedSerializer, isMuxAndDiscoveryCompleteFn func(ctx context.Context) bool) *Handler {
	return &Handler{serializer: serializer, isMuxAndDiscoveryCompleteFn: isMuxAndDiscoveryCompleteFn}
}

type Handler struct {
	serializer                  runtime.NegotiatedSerializer
	isMuxAndDiscoveryCompleteFn func(ctx context.Context) bool
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	//panic("implement me")
}
