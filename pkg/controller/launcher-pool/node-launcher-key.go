package launcherpool

import (
	"fmt"
)

// NodeLauncherKey defines the unique identifier for a (Node, LauncherConfig) pair
type NodeLauncherKey struct {
	LauncherConfigName string
	NodeName           string
}

func (k NodeLauncherKey) String() string {
	return fmt.Sprintf("%s/%s", k.LauncherConfigName, k.NodeName)
}

// MapToLoggable converts a map of NodeLauncherKey to int32 values into a string representation.
// This function formats the map as a string with the format "{namespace/name/node:count, ...}"
// for debugging and logging purposes.
func MapToLoggable[Key interface {
	comparable
	fmt.Stringer
}, Val any](m map[Key]Val) map[string]any {
	result := make(map[string]any, len(m))
	for k, v := range m {
		result[k.String()] = v
	}
	return result
}
