package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSubArrayLen(t *testing.T) {
	t.Log(`
	Given an array of positive integers nums and a positive integer target, 
	return the minimal length of a subarray whose sum is greater than or equal to target. 
	If there is no such subarray, return 0 instead.

	Example 1:
	Input: target = 7, nums = [2,3,1,2,4,3]
	Output: 2
	Explanation: The subarray [4,3] has the minimal length under the problem constraint.
	
	Example 2:
	Input: target = 4, nums = [1,4,4]
	Output: 1

	Example 3:
	Input: target = 11, nums = [1,1,1,1,1,1,1,1]
	Output: 0
`)
	t.Run("Suite 1", func(t *testing.T) {
		target := 7
		nums := []int{2, 3, 1, 2, 4, 3}
		exp := 2
		assert.Equal(t, exp, RunMinSubArrayLen(target, nums))
	})
	t.Run("Suite 2", func(t *testing.T) {
		target := 4
		nums := []int{1, 4, 4}
		exp := 1
		assert.Equal(t, exp, RunMinSubArrayLen(target, nums))
	})
	t.Run("Suite 3", func(t *testing.T) {
		target := 11
		nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
		exp := 0
		assert.Equal(t, exp, RunMinSubArrayLen(target, nums))
	})
}

func RunMinSubArrayLen(target int, nums []int) int {
	res := len(nums) + 1

	l, sum := 0, 0
	for r := 0; r < len(nums); r++ {
		sum = sum + nums[r]

		for sum >= target {
			res = min(res, r-l+1)
			sum = sum - nums[l]
			l++
		}
	}

	if res == len(nums)+1 {
		return 0
	}

	return res
}
