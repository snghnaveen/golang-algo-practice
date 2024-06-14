package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinIncrementForUnique(t *testing.T) {
	t.Log(`
	You are given an integer array nums. In one move, 
	you can pick an index i where 0 <= i < nums.length and increment nums[i] by 1.
	Return the minimum number of moves to make every value in nums unique.
	The test cases are generated so that the answer fits in a 32-bit integer.

	Example 1:
	Input: nums = [1,2,2]
	Output: 1
	Explanation: After 1 move, the array could be [1, 2, 3].

	Example 2:
	Input: nums = [3,2,1,2,1,7]
	Output: 6
	Explanation: After 6 moves, the array could be [3, 4, 1, 2, 5, 7].
	It can be shown with 5 or less moves that it is impossible for the array to have all unique values.
	`)
	t.Run("Suite 1", func(t *testing.T) {
		nums := []int{1, 2, 2}
		exp := 1
		assert.Equal(t, exp, RunMinIncrementForUnique(nums))
	})
	t.Run("Suite 2", func(t *testing.T) {
		nums := []int{3, 2, 1, 2, 1, 7}
		exp := 6
		assert.Equal(t, exp, RunMinIncrementForUnique(nums))
	})
}

func RunMinIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	if len(nums) < 2 {
		return 0
	}
	var out int

	i, j := 0, 1

	for j < len(nums) {
		if nums[i] > nums[j] {
			x := nums[i] - nums[j]
			out = out + x
			nums[j] = nums[j] + x
		}

		if nums[i] == nums[j] {
			nums[j]++
			out++
		}

		i++
		j++
	}
	return out
}
