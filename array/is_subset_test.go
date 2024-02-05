package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubset(t *testing.T) {
	t.Log(`
	Given two arrays: arr1[0..m-1] and arr2[0..n-1]. 
	Find whether arr2[] is a subset of arr1[] or not. 
	Both arrays are not in sorted order. 
	It may be assumed that elements in both arrays are distinct.
	Examples: 

	Input: arr1[] = {11, 1, 13, 21, 3, 7}, arr2[] = {11, 3, 7, 1} 
	Output: arr2[] is a subset of arr1[]

	Input: arr1[] = {1, 2, 3, 4, 5, 6}, arr2[] = {1, 2, 4} 
	Output: arr2[] is a subset of arr1[]


	Input: arr1[] = {10, 5, 2, 23, 19}, arr2[] = {19, 5, 3} 
	Output: arr2[] is not a subset of arr1[] 
	`)

	t.Run("Suite 1", func(t *testing.T) {
		assert.True(t, RunIsSubset(
			[]int{11, 1, 13, 21, 3, 7},
			[]int{11, 3, 7, 1},
		))
	})

	t.Run("Suite 2", func(t *testing.T) {
		assert.True(t, RunIsSubset(
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 2, 4},
		))
	})

	t.Run("Suite 2", func(t *testing.T) {
		assert.False(t, RunIsSubset(
			[]int{10, 5, 2, 23, 19},
			[]int{19, 5, 3},
		))
	})
}

func RunIsSubset(arr1, arr2 []int) bool {
	sort.Ints(arr1)
	sort.Ints(arr2)

	arr1Len := len(arr1)
	arr2Len := len(arr2)

	var i, j int

	for i < arr1Len && j < arr2Len {
		if arr1[i] == arr2[j] {
			i++
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else if arr1[i] > arr2[j] {
			return false
		}
	}

	return true
}
