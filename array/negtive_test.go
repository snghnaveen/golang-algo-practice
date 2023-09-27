package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignof(t *testing.T) {
	t.Log(`
	You are given an integer array nums. Let product be the product of all values in the array nums.
	Return signFunc(product).

	Example 1:
	Input: nums = [-1,-2,-3,-4,3,2,1]
	Output: 1
	Explanation: The product of all values in the array is 144, and signFunc(144) = 1
	
	Example 2:
	Input: nums = [1,5,0,2,-3]
	Output: 0
	Explanation: The product of all values in the array is 0, and signFunc(0) = 0
	
	Example 3:
	Input: nums = [-1,1,-1,1,-1]
	Output: -1
	Explanation: The product of all values in the array is -1, and signFunc(-1) = -1
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{-1, -2, -3, -4, 3, 2, 1}
		exp := 1
		assert.Equal(t, exp, RunSignFunc(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{-1, 1, -1, 1, -1}
		exp := -1
		assert.Equal(t, exp, RunSignFunc(in))
	})
}

func RunSignFunc(in []int) int {
	negCnt := 0
	for i := 0; i < len(in); i++ {
		if in[i] == 0 {
			return 0
		}

		if in[i] < 0 {
			negCnt++
		}
	}
	if negCnt%2 != 0 {
		return -1
	}
	return 1
}
