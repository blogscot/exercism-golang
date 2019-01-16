// Package grains defines operations on an imaginary chessboard, on which
// the first square contains one grain of wheat and each subsequent square
// holds double that amount.
package grains

import (
	"fmt"
)

const maxSquares int = 64

// Square calculates the number of grains on a given chessboard square.
func Square(num int) (uint64, error) {
	if num <= 0 || num > maxSquares {
		return 0, fmt.Errorf("invalid square: %d", num)
	}
	return uint64(1) << uint64(num-1), nil
}

// Total calculates the total number of grains on a chessboard.
func Total() (total uint64) {
	for square := 1; square <= maxSquares; square++ {
		if count, err := Square(square); err == nil {
			total += count
		}
	}
	return
}
