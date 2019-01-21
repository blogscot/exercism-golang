package atbash

import (
	"bytes"
	"strings"
	"unicode"
)

// Atbash encrypts the text using an ancient encryption system.
func Atbash(text string) string {
	text = strings.Map(translate, strings.ToLower(text))
	textLength := len(text)
	var buf bytes.Buffer
	var count int

	for index, r := range text {
		buf.WriteString(string(r))

		count = index + 1
		if count%5 == 0 && count != textLength {
			buf.WriteString(" ")
		}
	}
	return buf.String()
}

func translate(r rune) rune {
	switch {
	case unicode.IsSpace(r) || r == ',' || r == '.':
		return -1
	case unicode.IsDigit(r):
		return r
	default:
		return rune(byte('z') - (byte(r) - byte('a')))
	}
}
