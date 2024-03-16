package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxLength(t *testing.T) {
	t.Log(`
	Given a binary array nums, 
	return the maximum length of a contiguous subarray with an equal number of 0 and 1.

	Example 1:
	Input: nums = [0,1]
	Output: 2
	Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.

	Example 2:
	Input: nums = [0,1,0]
	Output: 2
	Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{0, 1}
		exp := 2
		assert.Equal(t, exp, RunFindMaxLength(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{0, 1, 0}
		exp := 2
		assert.Equal(t, exp, RunFindMaxLength(in))
	})
}

func RunFindMaxLength(nums []int) int {
	res := 0
	sum := 0
	mpSumIdx := map[int]int{
		0: -1,
	}

	for currIdx, num := range nums {
		if num == 1 {
			sum++
		} else {
			sum--
		}

		if idx, ok := mpSumIdx[sum]; ok {
			res = max(res, currIdx-idx)
		} else {
			mpSumIdx[sum] = currIdx
		}
	}
	return res
}
