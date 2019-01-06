package hamming

import "errors"

// Distance calculates the hamming distance between to strings
func Distance(string1, string2 string) (int, error) {

	if len(string1) != len(string2) {
		return -1, errors.New("unequal string lengths")
	}

	distance := 0

	for i := range string1 {
		if string1[i] != string2[i] {
			distance++
		}
	}
	return distance, nil
}
