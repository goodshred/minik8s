package options

import (
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilversion "k8s.io/apiserver/pkg/util/version"
	"k8s.io/kubernetes/staging/src/k8s.io/apiserver/pkg/server"
	"time"
)

type ServerRunOptions struct {
	RequestTimeout time.Duration
}

func NewServerRunOptions() *ServerRunOptions {
	return NewServerRunOptionsForComponent(utilversion.DefaultKubeComponent, utilversion.DefaultComponentGlobalsRegistry)

}
func NewServerRunOptionsForComponent(componentName string, componentGlobalsRegistry utilversion.ComponentGlobalsRegistry) *ServerRunOptions {
	defaults := server.NewConfig(serializer.CodecFactory{})
	return &ServerRunOptions{
		RequestTimeout: defaults.RequestTimeout,
	}
}
