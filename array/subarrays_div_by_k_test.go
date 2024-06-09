package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraysDivByK(t *testing.T) {
	t.Log(`
	Given an integer array nums and an integer k, 
	return the number of non-empty subarrays that have a sum divisible by k.
	A subarray is a contiguous part of an array.

	Example 1:
	Input: nums = [4,5,0,-2,-3,1], k = 5
	Output: 7
	Explanation: There are 7 subarrays with a sum divisible by k = 5:
	[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]

	Example 2:
	Input: nums = [5], k = 9
	Output: 0
	`)

	t.Run("Suite 1", func(t *testing.T) {
		nums := []int{4, 5, 0, -2, -3, 1}
		k := 5
		exp := 7
		assert.Equal(t, exp, RunSubArraysDivByK(nums, k))
	})
	t.Run("Suite 2", func(t *testing.T) {
		nums := []int{5}
		k := 9
		exp := 0
		assert.Equal(t, exp, RunSubArraysDivByK(nums, k))
	})
}

func RunSubArraysDivByK(nums []int, k int) int {
	mpRemCnt := map[int]int{
		0: 1,
	}

	count := 0
	prefixSum := 0
	for _, num := range nums {
		prefixSum = prefixSum + num
		currRem := prefixSum % k
		if currRem < 0 {
			currRem = currRem + k
		}

		if v, ok := mpRemCnt[currRem]; ok {
			count = count + v
		}

		mpRemCnt[currRem]++
	}

	return count
}
