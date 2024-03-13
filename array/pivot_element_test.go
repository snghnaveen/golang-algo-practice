package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotInteger(t *testing.T) {
	t.Log(`
	Given a positive integer n, find the pivot integer x such that:
	The sum of all elements between 1 and x inclusively equals the sum of 
	all elements between x and n inclusively.
	Return the pivot integer x. 
	If no such integer exists, return -1. 
	It is guaranteed that there will be at most one pivot index for the given input.
	
	Example 1:
	Input: n = 8
	Output: 6
	Explanation: 6 is the pivot integer since: 1 + 2 + 3 + 4 + 5 + 6 = 6 + 7 + 8 = 21.

	Example 2:
	Input: n = 1
	Output: 1
	Explanation: 1 is the pivot integer since: 1 = 1.

	Example 3:
	// Input: n = 4
	Output: -1
	Explanation: It can be proved that no such integer exist.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		n := 8
		exp := 6
		assert.Equal(t, exp, RunPivotInteger(n))
	})
	t.Run("Suite 2", func(t *testing.T) {
		n := 1
		exp := 1
		assert.Equal(t, exp, RunPivotInteger(n))
	})
	t.Run("Suite 3", func(t *testing.T) {
		n := 4
		exp := -1
		assert.Equal(t, exp, RunPivotInteger(n))
	})
}

// @TODO can be solved using binaray search
func RunPivotInteger(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total = total + i
	}

	left := 0
	for i := 1; i <= n; i++ {
		if left == total-i-left {
			return i
		}
		left = left + i
	}
	return -1
}
