package wordcount

import (
	"regexp"
	"strings"
)

// Frequency contains the list of words and their number of occurrances.
type Frequency map[string]int

// WordCount calculates the word frequency for the given text.
func WordCount(text string) (freq Frequency) {
	re, _ := regexp.Compile("(\\b[\\w']+\\b)")
	words := re.FindAllString(strings.ToLower(text), -1)

	freq = make(Frequency)

	for _, word := range words {
		if _, ok := freq[word]; !ok {
			freq[word] = 1
		} else {
			freq[word]++
		}
	}
	return
}
