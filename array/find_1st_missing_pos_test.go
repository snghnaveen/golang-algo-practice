package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstMissingPostiveNum(t *testing.T) {
	t.Log(`Link: https://leetcode.com/problems/first-missing-positive/description/`)

	t.Log(`
	Given an unsorted array arr[] with both positive and negative elements, 
	the task is to find the smallest positive number missing from the array.
	Note: You can modify the original array.

	Examples:

	Input:  arr[] = {2, 3, 7, 6, 8, -1, -10, 15}
	Output: 1

	Input:  arr[] = { 2, 3, -7, 6, 8, 1, -10, 15 }
	Output: 4

	Input: arr[] = {1, 1, 0, -1, -2}
	Output: 2
	`)
	t.Run("Suite 1", func(t *testing.T) {
		in := []int{2, 3, -7, 6, 8, 1, -10, 15}
		exp := 4
		assert.Equal(t, exp, RunFindFirstMissingPostiveNum(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{2, 3, 1, -3, 8, 2}
		exp := 4
		assert.Equal(t, exp, RunFindFirstMissingPostiveNum(in))
	})

	t.Run("Suite 3", func(t *testing.T) {
		in := []int{2, 3, 7, 6, 8, -1, -10, 15}
		exp := 1
		assert.Equal(t, exp, RunFindFirstMissingPostiveNum(in))
	})

	t.Run("Suite 4", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		exp := 6
		assert.Equal(t, exp, RunFindFirstMissingPostiveNum(in))
	})

}

func RunFindFirstMissingPostiveNum(in []int) int {
	inLen := len(in)
	for i := 0; i < inLen; i++ {

		if in[i] > 0 && in[i] < inLen {

			val := in[i]
			position := val - 1
			// swap the element to respective location
			// for example :
			// 4 should sit in chair/index 3
			// 1 should sit in chair/index 0
			if in[position] != val {
				in[position], in[i] = in[i], in[position]

				// check if after swap the element is in correct position or not
				i--
			}
		}
	}

	for i := 0; i < inLen; i++ {
		if in[i] != i+1 {
			return i + 1
		}
	}

	return inLen + 1
}
