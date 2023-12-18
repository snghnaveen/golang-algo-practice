package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFistAndLastPostion(t *testing.T) {
	t.Log(`
	Given an array of integers nums sorted in non-decreasing order,
	find the starting and ending position of a given target value.
	If target is not found in the array, return [-1, -1].

	You must write an algorithm with O(log n) runtime complexity.

	Example 1:
	Input: nums = [5,7,7,8,8,10], target = 8
	Output: [3,4]

	Example 2:
	Input: nums = [5,7,7,8,8,10], target = 6
	Output: [-1,-1]

	Example 3:
	Input: nums = [], target = 0
	Output: [-1,-1]
`)
	t.Run("Suite 1", func(t *testing.T) {
		in := []int{5, 7, 7, 8, 8, 10}
		target := 8
		exp := []int{3, 4}
		assert.Equal(t, exp, RunFindFistAndLastPostion(in, target))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{5, 7, 7, 8, 8, 10}
		target := 6
		exp := []int{-1, -1}
		assert.Equal(t, exp, RunFindFistAndLastPostion(in, target))
	})
	t.Run("Suite 3", func(t *testing.T) {
		in := []int{}
		target := 6
		exp := []int{-1, -1}
		assert.Equal(t, exp, RunFindFistAndLastPostion(in, target))
	})
}

func RunFindFistAndLastPostion(nums []int, target int) []int {
	l := binSearch(nums, target, true)
	r := binSearch(nums, target, false)
	return []int{l, r}
}

func binSearch(nums []int, target int, leftBais bool) int {

	l, r := 0, len(nums)-1
	i := -1

	for l <= r {
		m := l + (r-l)/2
		if target > nums[m] {
			l = m + 1
		} else if target < nums[m] {
			r = m - 1
		} else if target == nums[m] {
			i = m
			if leftBais {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return i
}
