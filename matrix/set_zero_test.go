package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeroes(t *testing.T) {
	t.Log(`
	Given an m x n integer matrix matrix, if an element is 0, 
	set its entire row and column to 0's.
	You must do it in place.

	Example 1:
	Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
	Output: [[1,0,1],[0,0,0],[1,0,1]]

	Example 2:
	Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
	Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
	`)
	t.Run("Suite 1", func(t *testing.T) {
		matrix := [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}
		exp := [][]int{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		}

		RunSetZeroes(matrix)

		assert.Equal(t, exp, matrix)
	})
	t.Run("Suite 2", func(t *testing.T) {
		matrix := [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}
		exp := [][]int{
			{0, 0, 0, 0},
			{0, 4, 5, 0},
			{0, 3, 1, 0},
		}

		RunSetZeroes(matrix)

		assert.Equal(t, exp, matrix)

	})
}

func RunSetZeroes(matrix [][]int) {
	zeroRow := false

	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				if r > 0 {
					matrix[r][0] = 0
				} else {
					zeroRow = true
				}
			}
		}
	}

	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[0]); c++ {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for r := 0; r < len(matrix); r++ {
			matrix[r][0] = 0
		}
	}

	if zeroRow {
		for c := 0; c < len(matrix[0]); c++ {
			matrix[0][c] = 0
		}
	}
}
