package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefixUsingSorting(t *testing.T) {
	t.Log(`
	Given an array, check whether the array is in sorted order with recursion.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		arr := []int{1, 4, 11, 14, 17}

		assert.Equal(t, true,
			RunIsSortedWithRecursion(arr))
	})

	t.Run("Suite 2", func(t *testing.T) {
		arr := []int{1, 3, 4, 2, 5}

		assert.Equal(t, false,
			RunIsSortedWithRecursion(arr))
	})

}

func RunIsSortedWithRecursion(in []int) bool {
	return checkSorted(in, len(in))
}

func checkSorted(in []int, n int) bool {

	// Array has one or no element or the
	// rest are already checked and approved.
	if n == 0 || n == 1 {
		return true
	}

	if in[n-1] < in[n-2] {
		return false
	}

	return checkSorted(in, n-1)
}
