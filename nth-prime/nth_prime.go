package prime

import "math"

// Nth returns the nth prime number.
func Nth(nth int) (int, bool) {
	if nth == 0 {
		return 0, false
	}
	num := 1

	for count := 0; count < nth; {
		num++
		if isPrime(num) {
			count++
		}
	}
	return num, true
}

func isPrime(num int) bool {
	if num%2 == 0 {
		return num == 2
	}
	limit := int(math.Sqrt(float64(num))) + 1
	for n := 3; n < limit; n += 2 {
		if num%n == 0 {
			return false
		}
	}
	return true
}
