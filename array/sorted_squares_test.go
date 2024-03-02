package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	t.Log(`
	Given an integer array nums sorted in non-decreasing order, 
	return an array of the squares of each number sorted in non-decreasing order.

	Example 1:
	Input: nums = [-4,-1,0,3,10]
	Output: [0,1,9,16,100]
	Explanation: After squaring, the array becomes [16,1,0,9,100].
	After sorting, it becomes [0,1,9,16,100].

	Example 2:
	Input: nums = [-7,-3,2,3,11]
	Output: [4,9,9,49,121]
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{-4, -1, 0, 3, 10}
		exp := []int{0, 1, 9, 16, 100}

		assert.Equal(t, exp, RunSortedSquares(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{-7, -3, 2, 3, 11}
		exp := []int{4, 9, 9, 49, 121}

		assert.Equal(t, exp, RunSortedSquares(in))
	})

}

func RunSortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	l := 0
	r := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		sqrNum := 0
		if absInt(nums[l]) > absInt(nums[r]) {
			sqrNum = nums[l]
			l++
		} else {
			sqrNum = nums[r]
			r--
		}

		square := sqrNum * sqrNum
		res[i] = square

	}
	return res
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
