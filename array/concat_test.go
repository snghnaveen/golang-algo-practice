package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	t.Log(`

	Given an integer array nums of length n, 
	you want to create an array ans of length 2n where 
	ans[i] == nums[i] and 
	ans[i + n] == nums[i] 
	for 0 <= i < n (0-indexed).
	Specifically, ans is the concatenation of two nums arrays.
	Return the array ans.

	Example 1:
	Input: nums = [1,2,1]
	Output: [1,2,1,1,2,1]
	Explanation: The array ans is formed as follows:
	- ans = [nums[0],nums[1],nums[2],nums[0],nums[1],nums[2]]
	- ans = [1,2,1,1,2,1]
`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 2, 1}
		exp := []int{1, 2, 1, 1, 2, 1}
		assert.Equal(t, exp, RunConcat(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{1, 3, 2, 1}
		exp := []int{1, 3, 2, 1, 1, 3, 2, 1}
		assert.Equal(t, exp, RunConcat(in))
	})
}

func RunConcat(in []int) []int {
	n := len(in)

	out := make([]int, 2*n)
	for i := 0; i < n; i++ {
		out[i] = in[i]
		out[i+n] = in[i]
	}

	return out
}
