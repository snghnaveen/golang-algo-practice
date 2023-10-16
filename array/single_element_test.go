package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNonDuplicate(t *testing.T) {
	t.Log(`
	You are given a sorted array consisting of only integers where every element appears exactly twice, 
	except for one element which appears exactly once.
	Return the single element that appears only once.
	Your solution must run in O(log n) time and O(1) space.
	
	Example 1:
	Input: nums = [1,1,2,3,3,4,4,8,8]
	Output: 2

	Example 2:
	Input: nums = [3,3,7,7,10,11,11]
	Output: 10
		`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
		exp := 2
		assert.Equal(t, exp, RunSingleNonDuplicate(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{3, 3, 7, 7, 10, 11, 11}
		exp := 10
		assert.Equal(t, exp, RunSingleNonDuplicate(in))
	})
}

func RunSingleNonDuplicate(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := l + ((r - l) / 2)

		current := nums[m]

		if (m-1 < 0 || nums[m-1] != current) &&
			(m+1 == len(nums) || nums[m+1] != current) {
			return current
		}

		// [1,1,2,3,3,4,4,8,8]
		leftSize := 0
		if m-1 >= 0 && nums[m-1] == current {
			leftSize = m - 1
		} else {
			leftSize = m
		}

		if leftSize%2 == 0 {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return 0
}
