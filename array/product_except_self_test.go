package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	t.Log(`
	Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
	The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
	You must write an algorithm that runs in O(n) time and without using the division operation.

	Example 1:
	Input: nums = [1,2,3,4]
	Output: [24,12,8,6]

	Example 2:
	Input: nums = [-1,1,0,-3,3]
	Output: [0,0,9,0,0]
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 2, 3, 4}
		exp := []int{24, 12, 8, 6}
		assert.Equal(t, RunProductExceptSelf(in), exp)
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{-1, 1, 0, -3, 3}
		exp := []int{0, 0, 9, 0, 0}
		assert.Equal(t, RunProductExceptSelf(in), exp)
	})
}

func RunProductExceptSelf(in []int) []int {
	out := make([]int, len(in))

	prefix := 1
	for i, v := range in {
		out[i] = prefix
		prefix = v * prefix
	}

	postfix := 1
	for i := len(in) - 1; i >= 0; i-- {
		out[i] = out[i] * postfix
		postfix = in[i] * postfix
	}

	return out
}
