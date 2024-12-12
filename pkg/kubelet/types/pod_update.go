package types

// PodOperation defines what changes will be made on a pod configuration.
type PodOperation int

// These constants identify the PodOperations that can be made on a pod configuration.
const (
	// SET is the current pod configuration.
	SET PodOperation = iota
	// ADD signifies pods that are new to this source.
	ADD
	// DELETE signifies pods that are gracefully deleted from this source.
	DELETE
	// REMOVE signifies pods that have been removed from this source.
	REMOVE
	// UPDATE signifies pods have been updated in this source.
	UPDATE
	// RECONCILE signifies pods that have unexpected status in this source,
	// kubelet should reconcile status with this source.
	RECONCILE
)

// PodUpdate defines an operation sent on the channel. You can add or remove single services by
// sending an array of size one and Op == ADD|REMOVE (with REMOVE, only the ID is required).
// For setting the state of the system to a given state for this source configuration, set
// Pods as desired and Op to SET, which will reset the system state to that specified in this
// operation for this source channel. To remove all pods, set Pods to empty object and Op to SET.
//
// Additionally, Pods should never be nil - it should always point to an empty slice. While
// functionally similar, this helps our unit tests properly check that the correct PodUpdates
// are generated.
type PodUpdate struct {
	Pods   []*v1.Pod
	Op     PodOperation
	Source string
}
