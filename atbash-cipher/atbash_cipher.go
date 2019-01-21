package atbash

import (
	"strings"
	"unicode"
)

// Atbash encrypts the text using an ancient encryption system.
func Atbash(text string) string {
	text = strings.Map(translate, strings.ToLower(text))
	textLength := len(text)
	buf := make([]byte, 0)
	var count int

	for index, r := range []byte(text) {
		buf = append(buf, r)

		count = index + 1
		if count%5 == 0 && count != textLength {
			buf = append(buf, ' ')
		}
	}
	return string(buf)
}

func translate(r rune) rune {
	switch {
	case unicode.IsSpace(r) || r == ',' || r == '.':
		return -1
	case unicode.IsDigit(r):
		return r
	default:
		return 'z' - (r - 'a')
	}
}
