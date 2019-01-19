package sieve

import "math"

// Sieve finds all the primes up to the given number.
func Sieve(num int) (primes []int) {
	for n := 2; n <= num; n++ {
		if isPrime(n) {
			primes = append(primes, n)
		}
	}
	return
}

func isPrime(num int) bool {
	if num%2 == 0 {
		return num == 2
	}
	limit := int(math.Sqrt(float64(num)))
	for n := 3; n <= limit; n += 2 {
		if num%n == 0 {
			return false
		}
	}
	return true
}
