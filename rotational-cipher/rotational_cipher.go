package rotationalcipher

import (
	"bytes"
	"unicode"
)

// RotationalCipher shifts each letter in the text by given amount.
// Note that punctuation is unaffected.
func RotationalCipher(text string, shift int) string {
	var out bytes.Buffer

	for _, r := range text {
		if unicode.IsLetter(r) {
			out.WriteString(string(translate(byte(r), shift)))
		} else {
			out.WriteString(string(r))
		}

	}
	return out.String()
}

func translate(letter byte, amount int) byte {
	var offset = byte('A')

	if letter >= byte('a') {
		offset = byte('a')
	}
	return (((letter - offset) + byte(amount)) % 26) + offset
}
