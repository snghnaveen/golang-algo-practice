package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertInterval(t *testing.T) {
	t.Log(`
	Given a set of non-overlapping intervals and a new interval, 
	insert the interval at correct position. 
	If the insertion results in overlapping intervals, 
	then merge the overlapping intervals. 
	Assume that the set of non-overlapping intervals is sorted on the basis of start time, 
	to find correct position of insertion.

	Example 1:
	Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
	Output: [[1,5],[6,9]]

	Example 2:
	Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
	Output: [[1,2],[3,10],[12,16]]
	Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := [][]int{{1, 3}, {6, 9}}
		newInterval := []int{2, 5}
		exp := [][]int{{1, 5}, {6, 9}}

		assert.Equal(t, exp, RunInsetInterval(in, newInterval))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
		newInterval := []int{4, 8}
		exp := [][]int{{1, 2}, {3, 10}, {12, 16}}

		assert.Equal(t, exp, RunInsetInterval(in, newInterval))
	})
	t.Run("Suite 3", func(t *testing.T) {
		in := [][]int{{5, 9}, {12, 15}}
		newInterval := []int{1, 3}
		exp := [][]int{{1, 3}, {5, 9}, {12, 15}}

		assert.Equal(t, exp, RunInsetInterval(in, newInterval))
	})
}

func RunInsetInterval(in [][]int, newInterval []int) [][]int {
	var out [][]int

	for i := 0; i < len(in); i++ {
		if newInterval[1] < in[i][0] {
			out = append(out, newInterval)
			out = append(out, in[i:]...)
			return out
		} else if newInterval[0] > in[i][1] {
			out = append(out, in[i])
		} else {
			newInterval = []int{min(newInterval[0], in[i][0]), max(newInterval[1], in[i][1])}
		}
	}

	out = append(out, newInterval)

	return out
}
