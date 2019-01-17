package wordcount

import (
	"strings"
	"unicode"
)

// Frequency contains the list of words and their number of occurrances.
type Frequency map[string]int

// WordCount calculates the word frequency for the given text.
func WordCount(text string) (freq Frequency) {
	freq = make(Frequency)

	words := strings.FieldsFunc(strings.ToLower(text), func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})

	for _, word := range words {
		freq[strings.Trim(word, "'")]++
	}
	return
}
