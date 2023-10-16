package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPossibleTriangles(t *testing.T) {
	t.Log(`
	Given an unsorted array of positive integers, 
	find the number of triangles that can be formed with three different array elements as 
	three sides of triangles. 
	For a triangle to be possible from 3 values, 
	the sum of any of the two values (or sides) must be greater than the third value (or third side). 

	Input: arr= {4, 6, 3, 7}
	Output: 3
	Explanation: There are three triangles 
	possible {3, 4, 6}, {4, 6, 7} and {3, 6, 7}. 
	Note that {3, 4, 7} is not a possible triangle.  

	Input: arr= {10, 21, 22, 100, 101, 200, 300}.
	Output: 6
	Explanation: There can be 6 possible triangles:
	{10, 21, 22}, {21, 100, 101}, {22, 100, 101}, 
	{10, 100, 101}, {100, 101, 200} and {101, 200, 300}
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{4, 6, 3, 7}
		exp := 3
		assert.Equal(t, exp, RunPossibleTriangles(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{10, 21, 22, 100, 101, 200, 300}
		exp := 6
		assert.Equal(t, exp, RunPossibleTriangles(in))
	})
}

func RunPossibleTriangles(in []int) int {
	sort.Ints(in)

	count := 0
	for i := len(in) - 1; i >= 1; i-- {
		l, r := 0, i-1

		for l < r {
			if in[l]+in[r] > in[i] {
				count = count + r - l
				r--
			} else {
				l++
			}
		}
	}

	return count
}
