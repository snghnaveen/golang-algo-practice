package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectionOfTwoArrays(t *testing.T) {
	t.Log(`
		Find intersection of two unsorted arrays
	`)
	arrA := []int{7, 1, 5, 2, 3, 6}
	arrB := []int{3, 8, 6, 20, 7}

	assert.ElementsMatch(t, RunTestIntersection(arrA, arrB), []int{3, 6, 7})

}

func RunTestIntersection(arrA, arrB []int) []int {
	var out []int

	sort.Ints(arrA)
	sort.Ints(arrB)

	lenA, lenB := len(arrA), len(arrB)
	i, j := 0, 0

	for i < lenA && j < lenB {
		if arrA[i] == arrB[j] {
			out = append(out, arrA[i])
			i++
			j++
		} else if arrA[i] < arrB[j] {
			i++
		} else {
			j++
		}
	}

	return out
}
