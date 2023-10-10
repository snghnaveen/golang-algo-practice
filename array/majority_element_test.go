package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajortiyElement(t *testing.T) {
	t.Log(`
	Given an array nums of size n, return the majority element.
	The majority element is the element that appears more than ⌊n / 2⌋ times. 
	You may assume that the majority element always exists in the array.
	
	Example 1:
	Input: nums = [3,2,3]
	Output: 3

	Example 2:
	Input: nums = [2,2,1,1,1,2,2]
	Output: 2
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{3, 2, 3}
		exp := 3
		assert.Equal(t, exp, RunMajortiyElement(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{2, 2, 1, 1, 1, 2, 2}
		exp := 2
		assert.Equal(t, exp, RunMajortiyElement(in))
	})
}

func RunMajortiyElement(in []int) int {
	res, count := 0, 0
	for _, v := range in {
		if count == 0 {
			res = v
		}

		if res == v {
			count++
		} else {
			count--
		}
	}

	return res
}
