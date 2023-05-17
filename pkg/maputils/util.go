package maputils

import "strings"

func MapKeysToString[T any](m map[string]T) string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return "[" + strings.Join(keys, ",") + "]"
}
