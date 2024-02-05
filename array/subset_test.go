package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubset(t *testing.T) {
	t.Log(`
	Given an integer array nums of unique elements, return all possible subsets.
	The solution set must not contain duplicate subsets.
	
	Example 1:
	Input: nums = [1,2,3]
	Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

	Example 2:
	Input: nums = [0]
	Output: [[],[0]]
	`)
	t.Run("Suite 1", func(t *testing.T) {
		nums := []int{1, 2, 3}
		exp := [][]int{
			{},
			{1},
			{2},
			{3},
			{1, 2},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		}
		assert.ElementsMatch(t, exp, RunSubset(nums))
	})

	t.Run("Suite 2", func(t *testing.T) {
		nums := []int{3, 2, 4, 1}

		exp := [][]int{
			{}, {3}, {2}, {2, 3}, {4}, {3, 4}, {2, 4}, {2, 3, 4}, {1}, {1, 3}, {1, 2}, {1, 2, 3}, {1, 4}, {1, 3, 4}, {1, 2, 4}, {1, 2, 3, 4},
		}
		assert.ElementsMatch(t, exp, RunSubset(nums))
	})
}

func RunSubset(nums []int) [][]int {
	var res [][]int
	var subset []int
	subsetDFS(nums, &res, subset, 0)

	for _, x := range res {
		sort.Ints(x)
	}
	return res
}

func subsetDFS(nums []int, res *[][]int, subset []int, i int) {
	if i >= len(nums) {
		temp := make([]int, len(subset))
		copy(temp, subset)
		*res = append(*res, temp)
		return
	}

	subset = append(subset, nums[i])
	subsetDFS(nums, res, subset, i+1)

	subset = subset[:len(subset)-1]
	subsetDFS(nums, res, subset, i+1)
}
