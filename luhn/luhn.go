package luhn

import (
	"unicode"
)

// Valid determines if a number is a valid Luhn number
func Valid(number string) bool {
	total := 0
	count := 0

	for index := len(number) - 1; index >= 0; index-- {
		r := rune(number[index])
		if r == ' ' {
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}
		digit := int(r - '0')
		if count&1 == 1 {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		total += digit
		count++
	}

	if count < 2 {
		return false
	}

	return total%10 == 0
}
