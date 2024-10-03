package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEraseOverlapIntervals(t *testing.T) {
	t.Log(`
	Given an array of intervals intervals where intervals[i] = [starti, endi], 
	return the minimum number of intervals you need to remove to make 
	the rest of the intervals non-overlapping.
	Note that intervals which only touch at a point are non-overlapping. 
	For example, [1, 2] and [2, 3] are non-overlapping.

	Example 1:
	Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
	Output: 1
	Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.

	Example 2:
	Input: intervals = [[1,2],[1,2],[1,2]]
	Output: 2
	Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.

	Example 3:
	Input: intervals = [[1,2],[2,3]]
	Output: 0
	Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
		exp := 1
		assert.Equal(t, exp, RunEraseOverlapIntervals(intervals))
	})
	t.Run("Suite 2", func(t *testing.T) {
		intervals := [][]int{{1, 2}, {1, 2}, {1, 2}}
		exp := 2
		assert.Equal(t, exp, RunEraseOverlapIntervals(intervals))
	})
	t.Run("Suite 3", func(t *testing.T) {
		intervals := [][]int{{1, 2}, {2, 3}}
		exp := 0
		assert.Equal(t, exp, RunEraseOverlapIntervals(intervals))
	})
}

func RunEraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prev := intervals[0][1]
	nonOverlappCount := 1
	for i := 1; i < len(intervals); i++ {
		if prev <= intervals[i][0] {
			nonOverlappCount++
			prev = intervals[i][1]
		}
	}

	return len(intervals) - nonOverlappCount
}
