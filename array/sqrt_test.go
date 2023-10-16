package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	t.Log(`
	Given a non-negative integer x, return the square root of x rounded down to the nearest integer. 
	The returned integer should be non-negative as well.
	You must not use any built-in exponent function or operator.
	For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
 
	Example 1:
	Input: x = 4
	Output: 2
	Explanation: The square root of 4 is 2, so we return 2.

	Example 2:
	Input: x = 8
	Output: 2
	Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 
	2 is returned.

	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := 4
		exp := 2
		assert.Equal(t, exp, RunSqrt(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := 8
		exp := 2
		assert.Equal(t, exp, RunSqrt(in))
	})

}

func RunSqrt(in int) int {
	left := 0
	right := in

	nearestSquareRoot := 0
	for left <= right {
		middle := left + (right-left)/2

		calucatedSquare := middle * middle
		if calucatedSquare == in {
			return middle
		}

		if calucatedSquare > in {
			right = middle - 1
		} else if calucatedSquare < in {
			nearestSquareRoot = middle
			left = middle + 1
		}
	}
	return nearestSquareRoot
}
