// Package isogram provides for all your Isogram needs
package isogram

import (
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram
func IsIsogram(text string) bool {
	lettersFound := map[rune]bool{}

	for _, letter := range text {
		if letter == ' ' || letter == '-' {
			continue
		}

		lowercaseLetter := unicode.ToLower(letter)

		if lettersFound[lowercaseLetter] {
			return false
		}
		lettersFound[lowercaseLetter] = true
	}
	return true
}
