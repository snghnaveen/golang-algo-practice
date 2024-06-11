package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeSort(t *testing.T) {
	t.Log(`
	Given two arrays arr1 and arr2, 
	the elements of arr2 are distinct, and all elements in arr2 are also in arr1.
	Sort the elements of arr1 such that the relative ordering of items in arr1 
	are the same as in arr2. Elements that do not appear in arr2 should be placed 
	at the end of arr1 in ascending order.

	Example 1:
	Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
	Output: [2,2,2,1,4,3,3,9,6,7,19]

	Example 2:
	Input: arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
	Output: [22,28,8,6,17,44]
	`)

	t.Run("Suite 1", func(t *testing.T) {
		arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
		arr2 := []int{2, 1, 4, 3, 9, 6}
		exp := []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}

		assert.Equal(t, exp, RunRelativeSort(arr1, arr2))
	})
	t.Run("Suite 2", func(t *testing.T) {
		arr1 := []int{28, 6, 22, 8, 44, 17}
		arr2 := []int{22, 28, 8, 6}
		exp := []int{22, 28, 8, 6, 17, 44}

		assert.Equal(t, exp, RunRelativeSort(arr1, arr2))
	})
}

func RunRelativeSort(arr1, arr2 []int) []int {
	mpCnt := make(map[int]int)

	for _, val := range arr1 {
		mpCnt[val]++
	}

	var out []int

	for _, val := range arr2 {
		for mpCnt[val] > 0 {
			out = append(out, val)
			mpCnt[val]--
		}
	}

	var end []int
	for k, cnt := range mpCnt {
		for cnt > 0 {
			end = append(end, k)
			cnt--
		}
	}

	sort.Ints(end)
	out = append(out, end...)
	return out
}
