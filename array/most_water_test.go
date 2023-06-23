package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerWithMostWater(t *testing.T) {
	t.Log(`
	You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
	Find two lines that together with the x-axis form a container, such that the container contains the most water.
	Return the maximum amount of water a container can store.
	Notice that you may not slant the container.

	Link : https://leetcode.com/problems/container-with-most-water/
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		exp := 49
		assert.Equal(t, exp, RunMostWater(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{1, 5, 4, 3}
		exp := 6
		assert.Equal(t, exp, RunMostWater(in))
	})

	t.Run("Suite 3", func(t *testing.T) {
		in := []int{3, 1, 2, 4, 5}
		exp := 12
		assert.Equal(t, exp, RunMostWater(in))
	})
}

func RunMostWater(in []int) int {
	maxArea := 0
	left := 0
	right := len(in) - 1

	for left < right {
		currentArea := (right - left) * min(in[left], in[right])
		maxArea = max(currentArea, maxArea)

		if in[left] > in[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
