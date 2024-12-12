package storage

import "k8s.io/apimachinery/pkg/runtime/schema"

type ResourceConfig struct {
	GroupVersionConfigs map[schema.GroupVersion]bool
	ResourceConfigs     map[schema.GroupVersionResource]bool
}

func (o *ResourceConfig) ResourceEnabled(resource schema.GroupVersionResource) bool {
	return true
}
