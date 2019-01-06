package raindrops

import "strconv"

// Convert returns a string based on whether the given value is
// divisible by 3, 5, or 7.
func Convert(value int) string {
	output := ""

	if value%3 == 0 {
		output += "Pling"
	}
	if value%5 == 0 {
		output += "Plang"
	}
	if value%7 == 0 {
		output += "Plong"
	}

	if output == "" {
		output = strconv.Itoa(value)
	}
	return output
}
