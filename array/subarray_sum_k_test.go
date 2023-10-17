package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubArraySumEqualsK(t *testing.T) {
	t.Log(`
	Given an array of integers nums and an integer k, 
	return the total number of subarrays whose sum equals to k.
	A subarray is a contiguous non-empty sequence of elements within an array.
	
	Example 1:
	Input: nums = [1,1,1], k = 2
	Output: 2

	Example 2:
	Input: nums = [1,2,3], k = 3
	Output: 2
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 1, 1}
		k := 2
		exp := 2
		assert.Equal(t, exp, RunSubArraySumEqualsK(in, k))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{1, -1, 1, 1, 1, 1}
		k := 3
		exp := 4
		assert.Equal(t, exp, RunSubArraySumEqualsK(in, k))
	})

}

func RunSubArraySumEqualsK(nums []int, k int) int {
	res := 0
	currSum := 0
	mpPrefixCount := make(map[int]int)
	mpPrefixCount[0] = 1

	for _, num := range nums {
		currSum = currSum + num
		diff := currSum - k

		if v, ok := mpPrefixCount[diff]; ok {
			res = res + v
		}

		m := mpPrefixCount[currSum]
		mpPrefixCount[currSum] = m + 1
	}

	return res
}
