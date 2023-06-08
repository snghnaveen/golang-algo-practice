package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveNegPosToEnd(t *testing.T) {
	t.Log(`
	An array contains both positive and negative numbers in random order. 
	Rearrange the array elements so that all negative numbers appear before all positive numbers.
	
	Examples : 
	Input: -12, 11, -13, -5, 6, -7, 5, -3, -6
	Output: -12 -13 -5 -7 -3 -6 11 6 5
	`)
	t.Run("Suite 1", func(t *testing.T) {
		in := []int{-12, 11, -13, -5, 6, -7, 5, -3, -6}
		exp := []int{-12, -6, -13, -5, -3, -7, 5, 6, 11}
		assert.Equal(t, exp, RunMoveNegPosToEnd(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{1, -2, -5, 3, -4, 7, -6, 8}
		exp := []int{-6, -2, -5, -4, 3, 7, 1, 8}
		assert.Equal(t, exp, RunMoveNegPosToEnd(in))
	})
}

func RunMoveNegPosToEnd(in []int) []int {
	l := 0
	r := len(in) - 1

	for l <= r {
		if in[l] < 0 {
			l++
		} else if in[r] > 0 {
			r--
		} else if in[l] > 0 && in[r] < 0 {
			in[l], in[r] = in[r], in[l]
			l++
			r--
		}
	}
	return in
}
