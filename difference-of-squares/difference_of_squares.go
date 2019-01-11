// Package diffsquares performs calculations on natural numbers
package diffsquares

// SquareOfSum calculates the square of the sum of the first `num`
// natural numbers.
func SquareOfSum(num int) int {
	total := (num * (num + 1) / 2)
	return total * total
}

// SumOfSquares calculates the sum of the squares of the first `num`
// natural numbers.
func SumOfSquares(num int) int {
	return num * (num + 1) * (num + num + 1) / 6
}

// Difference calculates the difference between `SquareOfSum` and
// `SumOfSquares` functions.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
