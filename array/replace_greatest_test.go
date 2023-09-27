package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceElementsGreatestElementOnRightSide(t *testing.T) {
	t.Log(`
	Given an array arr, replace every element in that array with the greatest element among the elements to its right, 
	and replace the last element with -1.
	After doing so, return the array.

	Example 1:
	Input: arr = [17,18,5,4,6,1]
	Output: [18,6,6,6,1,-1]
	Explanation: 
	- index 0 --> the greatest element to the right of index 0 is index 1 (18).
	- index 1 --> the greatest element to the right of index 1 is index 4 (6).
	- index 2 --> the greatest element to the right of index 2 is index 4 (6).
	- index 3 --> the greatest element to the right of index 3 is index 4 (6).
	- index 4 --> the greatest element to the right of index 4 is index 5 (1).
	- index 5 --> there are no elements to the right of index 5, so we put -1.

	Example 2:
	Input: arr = [400]
	Output: [-1]
	Explanation: There are no elements to the right of index 0.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{17, 18, 5, 4, 6, 1}
		exp := []int{18, 6, 6, 6, 1, -1}

		assert.Equal(t, exp, RunReplaceElementsGreatestElementOnRightSide(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{400}
		exp := []int{-1}

		assert.Equal(t, exp, RunReplaceElementsGreatestElementOnRightSide(in))
	})

}

func RunReplaceElementsGreatestElementOnRightSide(arr []int) []int {
	out := make([]int, len(arr))
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		out[i] = max
		if arr[i] > max {
			max = arr[i]
		}

	}
	return out
}
