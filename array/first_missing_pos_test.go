package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPostive(t *testing.T) {
	t.Log(`
	Given an unsorted integer array nums. 
	Return the smallest positive integer that is not present in nums.
	You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.

	Example 1:
	Input: nums = [1,2,0]
	Output: 3
	Explanation: The numbers in the range [1,2] are all in the array.

	Example 2:
	Input: nums = [3,4,-1,1]
	Output: 2
	Explanation: 1 is in the array but 2 is missing.

	Example 3:
	Input: nums = [7,8,9,11,12]
	Output: 1
	Explanation: The smallest positive integer 1 is missing.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		nums := []int{1, 2, 0}
		exp := 3
		assert.Equal(t, exp, RunFirstMissingPositive(nums))
	})

	t.Run("Suite 2", func(t *testing.T) {
		nums := []int{3, 4, -1, 1}
		exp := 2
		assert.Equal(t, exp, RunFirstMissingPositive(nums))
	})
	t.Run("Suite 3", func(t *testing.T) {
		nums := []int{7, 8, 9, 11, 12}
		exp := 1
		assert.Equal(t, exp, RunFirstMissingPositive(nums))
	})
}

func RunFirstMissingPositive(nums []int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] < len(nums) {
			val := nums[i]
			pos := val - 1
			if nums[pos] != val {
				nums[pos], nums[i] = nums[i], nums[pos]
				i--
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
