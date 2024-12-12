package v1

import "k8s.io/apimachinery/pkg/runtime/schema"

const GroupName = "apiextensions.k8s.io"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

// GroupVersionResource unambiguously identifies a resource.  It doesn't anonymously include GroupVersion
// to avoid automatic coercion.  It doesn't use a GroupVersion to avoid custom marshalling
type GroupVersionResource struct {
	Group    string
	Version  string
	Resource string
}

// GroupVersion contains the "group" and the "version", which uniquely identifies the API.
type GroupVersion struct {
	Group   string
	Version string
}

// WithResource creates a GroupVersionResource based on the method receiver's GroupVersion and the passed Resource.
func (gv GroupVersion) WithResource(resource string) GroupVersionResource {
	return GroupVersionResource{Group: gv.Group, Version: gv.Version, Resource: resource}
}
