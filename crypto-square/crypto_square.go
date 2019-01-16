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

	numCols, numRows := getDimensions(len(input))
	chunks := chunk(input, numCols)

	var line bytes.Buffer
	var out []string

	for colIndex := 0; colIndex < numCols; colIndex++ {
		for rowIndex := 0; rowIndex < numRows; rowIndex++ {
			line.WriteString(string(chunks[rowIndex][colIndex]))
		}
		out = append(out, line.String())
		line.Reset()
	}

	return strings.Join(out, " ")
}

func chunk(input string, size int) (out [][]rune) {

	for len(input) > 0 {
		if len(input) > size {
			out = append(out, []rune(input[:size]))
			input = input[size:]
		} else {
			for len(input) < size {
				input += " "
			}
			out = append(out, []rune(input))
			input = ""
		}
	}
	return
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
	if a*a == length {
		return a, a
	}
	if (a+1)*a >= length {
		return a + 1, a
	}
	return a + 1, a + 1
}
