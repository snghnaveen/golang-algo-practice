package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3Sum(t *testing.T) {
	t.Log(`
	Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that 
	i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
	Notice that the solution set must not contain duplicate triplets.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		assert.Equal(t, Run3Sum(
			[]int{-1, 0, 1, 2, -1, -4}),
			[][]int{
				{
					-1, -1, 2,
				},
				{
					-1, 0, 1,
				},
			},
		)
	})

	t.Run("Suite 2", func(t *testing.T) {
		assert.Equal(t, Run3Sum(
			[]int{0, -1, 2, -3, 1}),
			[][]int{
				{
					-3, 1, 2,
				},
				{
					-1, 0, 1,
				},
			},
		)
	})
}

func Run3Sum(in []int) [][]int {
	sort.Ints(in)
	n := len(in)

	var out [][]int

	for num1Idx := 0; num1Idx < n-2; num1Idx++ {

		// Check for duplicate value 2nd index onwards
		if num1Idx > 0 && in[num1Idx] == in[num1Idx-1] {
			continue
		}

		num2Idx := num1Idx + 1
		num3Idx := n - 1

		for num2Idx < num3Idx {
			sum := in[num1Idx] + in[num2Idx] + in[num3Idx]

			if sum == 0 {
				out = append(out, []int{in[num1Idx], in[num2Idx], in[num3Idx]})

				num3Idx--

				// remove duplicate from right
				for num2Idx < num3Idx && in[num3Idx] == in[num3Idx+1] {
					num3Idx--
				}

			} else if sum > 0 {
				num3Idx--
			} else {
				num2Idx++
			}
		}
	}

	return out
}
