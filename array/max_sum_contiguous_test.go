package array

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSumContiguousArr(t *testing.T) {
	t.Log(`
	Given an integer array nums, 
	find the subarray with the largest sum, and return its sum.

	Example 1:
	Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
	Output: 6
	Explanation: The subarray [4,-1,2,1] has the largest sum 6.

	Example 2:
	Input: nums = [1]
	Output: 1
	Explanation: The subarray [1] has the largest sum 1.

	Example 3:
	Input: nums = [5,4,-1,7,8]
	Output: 23
	Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
		
	# Kadane's Algorithm
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		exp := 6

		assert.Equal(t, exp, RunMaxSumContiguousArr(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{5, 4, -1, 7, 8}
		exp := 23

		assert.Equal(t, exp, RunMaxSumContiguousArr(in))
	})
}

func RunMaxSumContiguousArr(in []int) int {
	maxSum := math.MinInt

	sum := 0
	for i := 0; i < len(in); i++ {
		sum = sum + in[i]
		if sum > maxSum {
			maxSum = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}
