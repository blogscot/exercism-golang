package prime

// Nth returns the nth prime number.
func Nth(nth int) (int, bool) {
	if nth == 0 {
		return 0, false
	}
	if nth == 1 {
		return 2, true
	}

	primes := []int{2}
	num := 2

	for count := 1; count < nth; {
		num++
		if isPrime(primes, num) {
			primes = append(primes, num)
			count++
		}
	}
	return num, true
}

func isPrime(currentPrimes []int, num int) bool {
	for _, prime := range currentPrimes {
		if num%prime == 0 {
			return false
		}
	}
	return true
}
