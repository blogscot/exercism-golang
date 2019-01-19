package sieve

// Sieve finds all the primes up to the given number.
func Sieve(num int) []int {
	if num < 2 {
		return []int{}
	}
	grid := make([]bool, num+1)
	primes := []int{2}

	for square := 3; square <= num; square += 2 {
		if !grid[square] {
			primes = append(primes, square)
			for index := square * square; index <= num; index += square {
				grid[index] = true
			}
		}
	}
	return primes
}
