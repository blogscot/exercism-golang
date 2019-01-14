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

	for index, r := range filtered {
		digit := int(r - '0')
		if (len(filtered)-1-index)&1 == 1 {
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
