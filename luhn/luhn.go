package luhn

import (
	"strings"
	"unicode"
)

// Valid determines if a number is a valid Luhn number
func Valid(number string) bool {
	filtered := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, number)

	if len(filtered) < 2 || !allDigits(filtered) {
		return false
	}

	total := 0

	for index, r := range reverse(filtered) {
		digit := int(r - '0')
		if index&1 == 1 {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		total += digit
	}

	return total%10 == 0
}

func allDigits(number string) bool {
	return strings.IndexFunc(number, func(r rune) bool {
		return !unicode.IsNumber(r)
	}) == -1
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
