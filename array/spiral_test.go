package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiral(t *testing.T) {
	t.Log(`
	Given an m x n matrix, return all elements of the matrix in spiral order.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		exp := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
		assert.Equal(t, exp, RunSpiralOrder(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}
		exp := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
		assert.Equal(t, exp, RunSpiralOrder(in))
	})

}

func RunSpiralOrder(matrix [][]int) []int {
	var res []int

	top := 0
	right := len(matrix[0]) - 1
	bottom := len(matrix) - 1
	left := 0

	for top <= bottom && left <= right {
		// Traverse Right
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		top++

		// Traverse Bottom
		for j := top; j <= bottom; j++ {
			res = append(res, matrix[j][right])
		}
		right--

		// Traverse Left
		if top <= bottom {
			for j := right; j >= left; j-- {
				res = append(res, matrix[bottom][j])
			}
		}
		bottom--

		// Traverse Up
		if left <= right {
			for j := bottom; j >= top; j-- {
				res = append(res, matrix[j][left])
			}
		}
		left++
	}
	return res
}
