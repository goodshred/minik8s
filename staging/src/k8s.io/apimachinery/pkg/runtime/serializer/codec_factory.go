package serializer

import "k8s.io/apimachinery/pkg/runtime"

// CodecFactory provides methods for retrieving codecs and serializers for specific
// versions and content types.
type CodecFactory struct {
	scheme *runtime.Scheme
}
