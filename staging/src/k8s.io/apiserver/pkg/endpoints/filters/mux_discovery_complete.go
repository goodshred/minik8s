package filters

import "context"

type muxAndDiscoveryIncompleteKeyType int

const (
	// muxAndDiscoveryIncompleteKey is a key under which a protection signal for all requests made before the server have installed all known HTTP paths is stored in the request's context
	muxAndDiscoveryIncompleteKey muxAndDiscoveryIncompleteKeyType = iota
)

// NoMuxAndDiscoveryIncompleteKey checks if the context contains muxAndDiscoveryIncompleteKey.
// The presence of the key indicates the request has been made when the HTTP paths weren't installed.
func NoMuxAndDiscoveryIncompleteKey(ctx context.Context) bool {
	muxAndDiscoveryCompleteProtectionKeyValue, _ := ctx.Value(muxAndDiscoveryIncompleteKey).(string)
	return len(muxAndDiscoveryCompleteProtectionKeyValue) == 0
}
