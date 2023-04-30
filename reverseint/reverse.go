// Package reverseint reverses integer.
package reverseint

import (
	"math"
	"strconv"
)

// Reverse gets an integer and reverses it,
// if the reversed integer is outside the signed 32-bit integer range, returns 0.
func Reverse(x int) int {
	isNegative := false

	if x < 0 {
		isNegative = true

		x *= -1
	}

	temp := []byte(strconv.Itoa(x))

	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}

	s := string(temp)

	if isNegative {
		s = "-" + s
	}

	n, _ := strconv.Atoi(s)

	if n < math.MinInt32 || n > math.MaxInt32 {
		return 0
	}

	return n
}
