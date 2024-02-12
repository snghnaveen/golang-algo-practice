package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsContinuousSubArr(t *testing.T) {
	t.Log(`
	Given an integer array nums and an integer k, 
	return true if nums has a good subarray or false otherwise.
	A good subarray is a subarray where:
	its length is at least two, and
	the sum of the elements of the subarray is a multiple of k.

	Note that:
	A subarray is a contiguous part of the array.
	An integer x is a multiple of k if there exists an integer n such that x = n * k.
	0 is always a multiple of k.
	

	Example 1:
	Input: nums = [23,2,4,6,7], k = 6
	Output: true
	Explanation: [2, 4] is a continuous subarray of size 2 whose elements sum up to 6.

	Example 2:
	Input: nums = [23,2,6,4,7], k = 6
	Output: true
	Explanation: [23, 2, 6, 4, 7] is an continuous subarray of size 5 whose elements sum up to 42.
	42 is a multiple of 6 because 42 = 7 * 6 and 7 is an integer.

	Example 3:
	Input: nums = [23,2,6,4,7], k = 13
	Output: false
	`)
	t.Run("Suite 1", func(t *testing.T) {
		in := []int{23, 2, 4, 6, 7}
		k := 6
		assert.True(t, RunIsContinuousSubArr(in, k))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{23, 2, 6, 4, 7}
		k := 6
		assert.True(t, RunIsContinuousSubArr(in, k))
	})
	t.Run("Suite 3", func(t *testing.T) {
		in := []int{23, 2, 6, 4, 7}
		k := 13
		assert.False(t, RunIsContinuousSubArr(in, k))
	})

}

func RunIsContinuousSubArr(nums []int, k int) bool {
	reminderIdxMp := map[int]int{
		0: -1,
	}

	currentSum := 0

	for i, num := range nums {
		currentSum = currentSum + num
		currentRem := currentSum % k
		if remIdx, ok := reminderIdxMp[currentRem]; ok {
			if i-remIdx > 1 {
				fmt.Println(reminderIdxMp)
				return true
			}
		} else {
			reminderIdxMp[currentRem] = i
		}
	}
	return false
}
