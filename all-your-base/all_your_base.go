package allyourbase

import (
	"errors"
	"math"
)

// ConvertToBase converts a sequence of digits from the input base to the output base.
func ConvertToBase(inBase int, digits []int, outBase int) ([]int, error) {
	if inBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}
	if outBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}
	if len(digits) == 0 || allZeros(digits) {
		return []int{0}, nil
	}
	if !digitsValid(inBase, digits) {
		return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
	}

	total := 0
	var pow float64

	for index := len(digits) - 1; index >= 0; index-- {
		total += digits[index] * int(math.Pow(float64(inBase), pow))
		pow++
	}

	var outDigits []int
	for total > 0 {
		outDigits = append([]int{total % outBase}, outDigits...)
		total = total / outBase
	}

	return outDigits, nil
}

func digitsValid(base int, digits []int) bool {
	for _, num := range digits {
		if num < 0 || num >= base {
			return false
		}
	}
	return true
}

func allZeros(numbers []int) bool {
	for _, num := range numbers {
		if num != 0 {
			return false
		}
	}
	return true
}
