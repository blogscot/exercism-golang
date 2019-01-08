package etl

import "strings"

// Transform transform legacy data into new data
func Transform(legacy map[int][]string) map[string]int {
	out := make(map[string]int)

	for k, list := range legacy {
		for _, elem := range list {
			out[strings.ToLower(elem)] = k
		}
	}
	return out
}
