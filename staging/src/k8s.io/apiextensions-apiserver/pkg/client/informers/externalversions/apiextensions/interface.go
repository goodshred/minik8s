package apiextensions

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions/apiextensions/v1"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1 provides access to shared informers for resources in V1.
	V1() v1.Interface
}
