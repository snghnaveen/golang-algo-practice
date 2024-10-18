package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	t.Log(`
	You are given an m x n integer matrix matrix with the following two properties:
	Each row is sorted in non-decreasing order.
	The first integer of each row is greater than the last integer of the previous row.
	Given an integer target, return true if target is in matrix or false otherwise.
	You must write a solution in O(log(m * n)) time complexity.

	Example 1:
	Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
	Output: true

	Example 2:
	Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
	Output: false
	`)

	t.Run("Suite 1", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target := 3
		assert.True(t, RunSearchMatrix(matrix, target))
	})
	t.Run("Suite 2", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target := 13
		assert.False(t, RunSearchMatrix(matrix, target))
	})
}

func RunSearchMatrix(matrix [][]int, target int) bool {
	top := 0
	bottom := len(matrix) - 1

	// find the desired row
	row := -1
	for top <= bottom {
		midRow := top + (bottom-top)/2
		l := 0
		r := len(matrix[midRow]) - 1
		if matrix[midRow][l] <= target && target <= matrix[midRow][r] {
			row = midRow
			break
		} else if target < matrix[midRow][l] {
			bottom = midRow - 1
		} else {
			top = midRow + 1
		}
	}

	if row == -1 {
		return false
	}

	l := 0
	r := len(matrix[row]) - 1

	for l <= r {
		mid := l + (r-l)/2
		if target == matrix[row][mid] {
			return true
		} else if target < matrix[row][mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
