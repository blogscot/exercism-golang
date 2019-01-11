// Package diffsquares performs calculations on natural numbers
package diffsquares

// SquareOfSum calculates the square of the sum of the first `num`
// natural numbers.
func SquareOfSum(num int) int {
	var total int
	for n := 1; n <= num; n++ {
		total += n
	}
	return total * total
}

// SumOfSquares calculates the sum of the squares of the first `num`
// natural numbers.
func SumOfSquares(num int) (total int) {
	for n := 1; n <= num; n++ {
		total += n * n
	}
	return
}

// Difference calculates the difference between `SquareOfSum` and `SumOfSquares`
// functions.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
