package server

import (
	"k8s.io/apimachinery/pkg/runtime/serializer"
	serverstore "k8s.io/apiserver/pkg/server/storage"
	"net/http"
	"time"
)

// Config is a structure used to configure a GenericAPIServer.
// Its members are sorted roughly in order of importance for composers.
type Config struct {
	Serializer serializer.CodecFactory
	// BuildHandlerChainFunc allows you to build custom handler chains by decorating the apiHandler.
	BuildHandlerChainFunc func(apiHandler http.Handler, c *Config) (secure http.Handler)
	// If specified, all requests except those which match the LongRunningFunc predicate will timeout
	// after this duration.
	RequestTimeout time.Duration
	// MergedResourceConfig indicates which groupVersion enabled and its resources enabled/disabled.
	// This is composed of genericapiserver defaultAPIResourceConfig and those parsed from flags.
	// If not specify any in flags, then genericapiserver will only enable defaultAPIResourceConfig.
	MergedResourceConfig *serverstore.ResourceConfig
}

type completedConfig struct {
	*Config
}

type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

// NewConfig returns a Config struct with the default values
func NewConfig(codecs serializer.CodecFactory) *Config {
	return &Config{
		Serializer:            codecs,
		BuildHandlerChainFunc: DefaultBuildHandlerChain,
		RequestTimeout:        time.Duration(60) * time.Second,
	}
}

func DefaultBuildHandlerChain(apiHandler http.Handler, c *Config) http.Handler {
	handler := apiHandler
	return handler
}
