package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNiceSubarrays(t *testing.T) {
	t.Log(`
	Given an array of integers nums and an integer k. 
	A continuous subarray is called nice if there are k odd numbers on it.
	Return the number of nice sub-arrays.

	Example 1:
	Input: nums = [1,1,2,1,1], k = 3
	Output: 2
	Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and [1,2,1,1].

	Example 2:
	Input: nums = [2,4,6], k = 1
	Output: 0
	Explanation: There are no odd numbers in the array.

	Example 3:
	Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
	Output: 16
	`)

	t.Run("Suite 1", func(t *testing.T) {
		nums := []int{1, 1, 2, 1, 1}
		k := 3
		exp := 2
		assert.Equal(t, exp, RunNiceSubarray(nums, k))
	})
	t.Run("Suite 2", func(t *testing.T) {
		nums := []int{2, 4, 6}
		k := 1
		exp := 0
		assert.Equal(t, exp, RunNiceSubarray(nums, k))
	})
	t.Run("Suite 3", func(t *testing.T) {
		nums := []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
		k := 2
		exp := 16
		assert.Equal(t, exp, RunNiceSubarray(nums, k))
	})
}

func RunNiceSubarray(nums []int, k int) int {
	for i, num := range nums {
		nums[i] = num % 2
	}

	mpPrefixSum := make(map[int]int)
	mpPrefixSum[0] = 1 // if currSum - k == 0 then we need to include it.
	res, currSum := 0, 0
	for _, num := range nums {
		currSum = currSum + num
		diff := currSum - k
		if v, ok := mpPrefixSum[diff]; ok {
			res = res + v
		}
		mpPrefixSum[currSum]++
	}
	return res
}
