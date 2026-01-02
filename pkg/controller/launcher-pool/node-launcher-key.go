package launcherpool

import (
	"fmt"
)

// NodeLauncherKey defines the unique identifier for a (Node, LauncherConfig) pair
type NodeLauncherKey struct {
	LauncherConfigName      string
	LauncherConfigNamespace string
	NodeName                string
}

func (k NodeLauncherKey) String() string {
	return fmt.Sprintf("%s/%s/%s", k.LauncherConfigNamespace, k.LauncherConfigName, k.NodeName)
}

// MapToLoggable converts a map of NodeLauncherKey to int32 values into a string representation.
// This function formats the map as a string with the format "{namespace/name/node:count, ...}"
// for debugging and logging purposes.
func MapToLoggable(m map[NodeLauncherKey]int32) map[string]int32 {
	result := make(map[string]int32, len(m))
	for k, v := range m {
		result[k.String()] = v
	}
	return result
}
