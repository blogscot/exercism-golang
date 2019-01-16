package cryptosquare

import (
	"bytes"
	"math"
	"strings"
	"unicode"
)

// Encode encodes the given text using the classic square code algorithm.
func Encode(text string) string {
	input := normalise(text)

	cols, rows := getDimensions(len(input))
	var line bytes.Buffer

	for i := 0; i < cols; i++ {
		if i != 0 {
			line.WriteString(" ")
		}
		for j := 0; i+j < cols*rows; j += cols {
			if i+j < len(input) {
				line.WriteString(string(input[i+j]))
			} else {
				line.WriteString(" ")
			}
		}
	}
	return line.String()
}

func normalise(input string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, input)
}

func getDimensions(length int) (int, int) {
	a := int(math.Sqrt(float64(length)))
	switch {
	case a*a == length:
		return a, a
	case (a+1)*a >= length:
		return a + 1, a
	default:
		return a + 1, a + 1
	}
}
