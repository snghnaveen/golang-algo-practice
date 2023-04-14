package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {

	in := []int{10, 20, 30, 40, 50, 60, 70}
	search := 33

	found, _ := RunTestBinarySearch(in, search)
	assert.False(t, found)

	search = 30
	found, idx := RunTestBinarySearch(in, search)
	assert.True(t, found)
	assert.Equal(t, idx, 2)

}

func RunTestBinarySearch(arr []int, search int) (found bool, idx int) {
	left := 0
	right := len(arr) - 1

	for left <= right {

		middle := left + (right-left)/2

		if arr[middle] == search {
			return true, middle
		}

		if arr[middle] > search {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return false, 0
}
