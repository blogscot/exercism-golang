// Package isogram provides for all your Isogram needs
package isogram

import (
	"regexp"
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram
func IsIsogram(text string) bool {
	re, _ := regexp.Compile("([- ])")
	letters := re.ReplaceAllString(text, "")

	lettersSet := map[rune]rune{}

	for _, letter := range letters {
		lettersSet[unicode.ToLower(letter)] = letter
	}

	return len(letters) == len(lettersSet)
}
