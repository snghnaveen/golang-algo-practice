package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeInterval(t *testing.T) {
	t.Log(`
	Given an array of intervals where intervals[i] = [starti, endi], 
	merge all overlapping intervals, and return an array of the 
	non-overlapping intervals that cover all the intervals in the input.

	Example 1:
	Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
	Output: [[1,6],[8,10],[15,18]]
	Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

	Example 2:
	Input: intervals = [[1,4],[4,5]]
	Output: [[1,5]]
	Explanation: Intervals [1,4] and [4,5] are considered overlapping.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := [][]int{{1, 3}, {8, 10}, {15, 18}, {2, 6}}
		exp := [][]int{{1, 6}, {8, 10}, {15, 18}}

		assert.Equal(t, exp, RunMergeInterval(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := [][]int{{1, 4}, {4, 5}}
		exp := [][]int{{1, 5}}

		assert.Equal(t, exp, RunMergeInterval(in))
	})

	t.Run("Suite 3", func(t *testing.T) {
		in := [][]int{{6, 8}, {1, 9}, {2, 4}, {4, 7}}
		exp := [][]int{{1, 9}}

		assert.Equal(t, exp, RunMergeInterval(in))
	})
}

func RunMergeInterval(in [][]int) [][]int {
	sort.Slice(in, func(i, j int) bool {
		return in[i][0] < in[j][0]
	})

	out := [][]int{in[0]}
	for i := 1; i < len(in); i++ {
		last := out[len(out)-1][1]
		if in[i][0] > last {
			out = append(out, in[i])
		} else {
			mx := max(out[len(out)-1][1], in[i][1])
			out[len(out)-1][1] = mx
		}
	}
	return out
}
