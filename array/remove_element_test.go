package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	t.Log(`
	Given an array nums and a value val, 
	remove all instances of that value in-place and return the new length. 
	Do not allocate extra space for another array. 
	You must do this by modifying the input array in-place with O(1) extra memory.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in, val := []int{3, 2, 2, 3}, 2
		exp := 2
		assert.Equal(t, exp, RunRemoveElement(in, val))
	})
}

func RunRemoveElement(in []int, val int) int {
	n := len(in)
	i := 0

	for j := 0; j < n; j++ {
		if in[j] != val {
			in[i] = in[j]
			i++
		}
	}
	return i
}
