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

	digits := toIntArray(reverse(filtered))

	for index := 0; index < len(digits); index++ {
		if index&1 == 1 {
			doubled := digits[index] * 2
			if doubled > 9 {
				doubled -= 9
			}
			digits[index] = doubled
		}
	}

	return sum(digits)%10 == 0
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

func sum(nums []int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}

func toIntArray(number string) (digits []int) {
	for _, r := range number {
		digits = append(digits, int(r-'0'))
	}
	return
}
