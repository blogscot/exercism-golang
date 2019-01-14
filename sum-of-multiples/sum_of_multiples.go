package summultiples

// SumMultiples calculates the sum of all the unique multiples
// of the divisors below the limit.
func SumMultiples(limit int, divisors ...int) int {
	multiples := make(map[int]struct{})

	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for index := divisor; index < limit; index += divisor {
			multiples[index] = struct{}{}
		}
	}

	keys := make([]int, 0, len(multiples))

	for key := range multiples {
		keys = append(keys, key)
	}
	return sum(keys)
}

func sum(nums []int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}
