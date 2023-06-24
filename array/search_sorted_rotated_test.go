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
		in := []int{4, 5, 6, 7, 0, 1, 2}
		target := 0
		exp := 4
		assert.Equal(t, exp, RunSearchSortedRotated(target, in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{4, 5, 6, 7, 0, 1, 2}
		target := 3
		exp := -1
		assert.Equal(t, exp, RunSearchSortedRotated(target, in))
	})
}

func RunSearchSortedRotated(target int, in []int) int {
	left, right := 0, len(in)-1

	for left <= right {
		mid := left + (right-left)/2

		if in[mid] == target {
			return mid
		}

		if in[left] < in[mid] {
			if in[left] <= target && target < in[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if in[mid] < target && target <= in[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
