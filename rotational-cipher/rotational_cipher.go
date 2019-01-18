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
	const lowercaseA = byte('a')
	const uppercaseA = byte('A')

	switch {
	case letter >= lowercaseA:
		return (((letter - lowercaseA) + byte(amount)) % 26) + lowercaseA
	default:
		return (((letter - uppercaseA) + byte(amount)) % 26) + uppercaseA
	}
}
