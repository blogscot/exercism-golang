package accumulate

// Accumulate performs provided operation on the given collection
func Accumulate(collection []string, operation func(string) string) []string {
	output := make([]string, 0)
	for _, elem := range collection {
		output = append(output, operation(elem))
	}
	return output
}
