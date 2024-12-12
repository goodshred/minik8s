package version

import "sync"

// ComponentGlobals stores the global variables for a component for easy access.
type ComponentGlobals struct {
}
type componentGlobalsRegistry struct {
	componentGlobals map[string]*ComponentGlobals
	mutex            sync.RWMutex
	// list of component name to emulation version set from the flag.
	emulationVersionConfig []string
	// map of component name to the list of feature gates set from the flag.
	featureGatesConfig map[string][]string
	// set stores if the Set() function for the registry is already called.
	set bool
}

// DefaultComponentGlobalsRegistry is the global var to store the effective versions and feature gates for all components for easy access.
// Example usage:
// // register the component effective version and feature gate first
// _, _ = utilversion.DefaultComponentGlobalsRegistry.ComponentGlobalsOrRegister(utilversion.DefaultKubeComponent, utilversion.DefaultKubeEffectiveVersion(), utilfeature.DefaultMutableFeatureGate)
// wardleEffectiveVersion := utilversion.NewEffectiveVersion("1.2")
// wardleFeatureGate := featuregate.NewFeatureGate()
// utilruntime.Must(utilversion.DefaultComponentGlobalsRegistry.Register(apiserver.WardleComponentName, wardleEffectiveVersion, wardleFeatureGate, false))
//
//	cmd := &cobra.Command{
//	 ...
//		// call DefaultComponentGlobalsRegistry.Set() in PersistentPreRunE
//		PersistentPreRunE: func(*cobra.Command, []string) error {
//			if err := utilversion.DefaultComponentGlobalsRegistry.Set(); err != nil {
//				return err
//			}
//	 ...
//		},
//		RunE: func(c *cobra.Command, args []string) error {
//			// call utilversion.DefaultComponentGlobalsRegistry.Validate() somewhere
//		},
//	}
//
// flags := cmd.Flags()
// // add flags
// utilversion.DefaultComponentGlobalsRegistry.AddFlags(flags)
var DefaultComponentGlobalsRegistry ComponentGlobalsRegistry = NewComponentGlobalsRegistry()

const (
	DefaultKubeComponent = "kube"
)

type ComponentGlobalsRegistry interface {
	// Reset removes all stored ComponentGlobals, configurations, and version mappings.
	Reset()
}

func NewComponentGlobalsRegistry() *componentGlobalsRegistry {
	return &componentGlobalsRegistry{
		componentGlobals:       make(map[string]*ComponentGlobals),
		emulationVersionConfig: nil,
		featureGatesConfig:     nil,
	}
}

func (r *componentGlobalsRegistry) Reset() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.componentGlobals = make(map[string]*ComponentGlobals)
	r.emulationVersionConfig = nil
	r.featureGatesConfig = nil
	r.set = false
}
