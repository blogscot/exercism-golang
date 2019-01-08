package accumulate

import "strings"

// Operation transforms a string into another string

// Accumulate performs provided operation on the given collection
func Accumulate(collection []string, operation func(string) string) []string {
	return strings.Map(operation, collection)
}
