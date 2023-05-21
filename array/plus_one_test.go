package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	t.Log(`
	Given a non-empty array of decimal digits representing a non-negative integer, 
	increment one to the integer and return the resulting array. 
	The digits are stored such that the most significant digit is at the head of the array, 
	and each element in the array contains a single digit.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 2, 3}
		exp := []int{1, 2, 4}
		assert.Equal(t, exp, RunPlusOne(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{1, 2, 9}
		exp := []int{1, 3, 0}
		assert.Equal(t, exp, RunPlusOne(in))
	})
	t.Run("Suite 3", func(t *testing.T) {
		in := []int{9, 9, 9}
		exp := []int{1, 0, 0, 0}
		assert.Equal(t, exp, RunPlusOne(in))
	})
}

func RunPlusOne(in []int) []int {
	for i := len(in) - 1; i >= 0; i-- {
		if in[i] < 9 {
			in[i] = in[i] + 1
			return in
		}
		in[i] = 0
	}

	// If we reach here, it means all digits were 9
	return append([]int{1}, in...)
}
