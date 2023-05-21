package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesSorted(t *testing.T) {
	t.Log(`
	Given a sorted array nums, 
	remove the duplicates in-place such that each element appears only once 
	and returns the new length.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 1, 2}
		exp := 2
		assert.Equal(t, exp, RunRemoveDuplicatesSorted(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		exp := 5
		assert.Equal(t, exp, RunRemoveDuplicatesSorted(in))
	})
}

func RunRemoveDuplicatesSorted(in []int) int {
	n := len(in)
	i := 0

	for j := 1; j < n; j++ {
		if in[i] != in[j] {
			i++
			in[i] = in[j]
		}
	}
	return i + 1
}
