package armstrong

import "math"

// IsNumber determines if a number is an armstrong number
func IsNumber(num int) bool {
	total := 0.0
	digits := getDigits(num)
	lenDigits := float64(len(digits))

	for _, digit := range digits {
		total += math.Pow(float64(digit), lenDigits)
	}
	return int(total) == num
}

func getDigits(num int) (digits []int) {
	var digit int

	for num > 0 {
		digit = num % 10
		num = num / 10
		digits = append(digits, digit)
	}
	return
}
