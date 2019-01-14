package summultiples

// SumMultiples calculates the sum of all the unique multiples
// of the divisors below the limit.
func SumMultiples(limit int, divisors ...int) (total int) {
	multiples := make(map[int]struct{})

	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for index := divisor; index < limit; index += divisor {
			if _, ok := multiples[index]; !ok {
				multiples[index] = struct{}{}
				total += index
			}
		}
	}
	return
}
