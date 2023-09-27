package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNondecreasing(t *testing.T) {
	t.Log(`Given an array nums with n integers, 
	your task is to check if it could become non-decreasing by modifying at most one element.
	We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) 
	such that (0 <= i <= n - 2).

	Example 1:
	Input: nums = [4,2,3]
	Output: true
	Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
	
	Example 2:
	Input: nums = [4,2,1]
	Output: false
	Explanation: You cannot get a non-decreasing array by modifying at most one element.
	`)
	t.Run("Suite 1", func(t *testing.T) {
		in := []int{4, 2, 3}
		assert.True(t, RunNondecreasing(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{4, 2, 1}
		assert.False(t, RunNondecreasing(in))
	})

}

func RunNondecreasing(in []int) bool {
	modified := false
	for i := 0; i < len(in)-1; i++ {
		if in[i] <= in[i+1] {
			continue
		}

		if modified {
			return false
		}

		//         4
		// [3, 4 , 2͇ , 3]
		//     i  i+1
		if i == 0 || in[i+1] >= in[i-1] {
			in[i] = in[i+1]
		} else {

			in[i+1] = in[i]
		}
		modified = true
	}

	return true
}
