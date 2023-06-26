package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort012(t *testing.T) {
	t.Log(`
	Given an array A[] consisting of only 0s, 1s, and 2s. 
	The task is to write a function that sorts the given array. 
	The functions should put all 0s first, then all 1s and all 2s in last.

	Input: {0, 1, 2, 0, 1, 2}
	Output: {0, 0, 1, 1, 2, 2}

	Input: {0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	Output: {0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2}
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{0, 1, 2, 0, 1, 2}
		exp := []int{0, 0, 1, 1, 2, 2}
		assert.Equal(t, exp, RunSort012(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
		exp := []int{0, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 2, 2}
		assert.Equal(t, exp, RunSort012(in))
	})
}

func RunSort012(in []int) []int {
	left := 0
	mid := 0
	right := len(in) - 1

	for mid <= right {
		switch in[mid] {
		case 0:
			in[left], in[mid] = in[mid], in[left]
			left++
			mid++
		case 1:
			mid++
		case 2:
			in[right], in[mid] = in[mid], in[right]
			right--
		}
	}
	return in
}
