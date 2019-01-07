// Package scrabble provides scoring for your words
package scrabble

import "unicode"

// Score calculates the scrabble value for the given word
func Score(word string) int {
	total := 0
	for _, letter := range word {
		total += scoreLetter(unicode.ToUpper(letter))
	}
	return total
}

func scoreLetter(letter rune) int {
	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	}
	return 0
}
