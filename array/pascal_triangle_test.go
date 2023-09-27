package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalTriangle(t *testing.T) {
	t.Log(`
	Given an integer numRows, return the first numRows of Pascal's triangle.
	In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
	
	Example 1:
	Input: numRows = 5
	Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
	
	Example 2:
	Input: numRows = 1
	Output: [[1]]
	`)

	t.Run("Suite 1", func(t *testing.T) {
		row := 5
		exp := [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}
		assert.Equal(t, exp, RunPascalTriangle(row))
	})
}

func RunPascalTriangle(row int) [][]int {
	res := [][]int{{1}}

	for i := 2; i <= row; i++ {

		temp := res[len(res)-1]
		temp = append(temp, 0)
		temp = append([]int{0}, temp...)

		row := []int{}
		for j := 0; j <= len(res[len(res)-1]); j++ {
			row = append(row, temp[j]+temp[j+1])
		}

		res = append(res, row)
	}

	return res
}
