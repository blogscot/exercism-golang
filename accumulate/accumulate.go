package accumulate

// Operation transforms a string into another string
type Operation = func(string) string

// Accumulate performs provided operation on the given collection
func Accumulate(collection []string, operation Operation) (output []string) {
	for _, elem := range collection {
		output = append(output, operation(elem))
	}
	return
}
