package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchSortedRotated(t *testing.T) {
	t.Log(`
	There is an integer array nums sorted in ascending order (with distinct values).
	Given the array nums after the possible rotation and an integer target, 
	return the index of target if it is in nums, or -1 if it is not in nums.

	You must write an algorithm with O(log n) runtime complexity.

	Example 1:

	Input: nums = [4,5,6,7,0,1,2], target = 0
	Output: 4

	Example 2:
	Input: nums = [4,5,6,7,0,1,2], target = 3
	Output: -1

	Example 3:
	Input: nums = [1], target = 0
	Output: -1
	`)

	t.Run("Suite 1", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 0
		exp := 4
		assert.Equal(t, exp, RunSearchSortedRotated(nums, target))
	})

	t.Run("Suite 2", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 3
		exp := -1
		assert.Equal(t, exp, RunSearchSortedRotated(nums, target))
	})
}

func RunSearchSortedRotated(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
