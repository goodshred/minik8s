package v1

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CustomResourceDefinitions returns a CustomResourceDefinitionInformer.
	CustomResourceDefinitions() CustomResourceDefinitionInformer
}
