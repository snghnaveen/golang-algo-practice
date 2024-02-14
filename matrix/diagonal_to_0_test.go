package matrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalTo0(t *testing.T) {
	t.Log(`
	Given an N*N matrix. 
	The task is to convert the elements of diagonal of a matrix to 0.
	
	Examples: 
	Input: mat[][] = 
	{{ 2, 1, 7 },
	{ 3, 7, 2 },
	{ 5, 4, 9 }}
	Output: 
	{ {0, 1, 0},
	{3, 0, 2},
	{0, 4, 0}}

	Input:  mat[][] = 
	{{1, 3, 5, 6, 7},
	{3, 5, 3, 2, 1},
	{1, 2, 3, 4, 5},
	{7, 9, 2, 1, 6},
	{9, 1, 5, 3, 2}}
	Output: 
	{{0, 3, 5, 6, 0},
	{3, 0, 3, 0, 1},
	{1, 2, 0, 4, 5},
	{7, 0, 2, 0, 6},
	{0, 1, 5, 3, 0}}
 `)
	t.Run("Suite 1", func(t *testing.T) {
		in := [][]int{
			{2, 1, 7},
			{3, 7, 2},
			{5, 4, 9},
		}

		exp := [][]int{
			{0, 1, 0},
			{3, 0, 2},
			{0, 4, 0},
		}
		assert.Equal(t, exp, RunDiagonalTo0(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := [][]int{
			{1, 3, 5, 6, 7},
			{3, 5, 3, 2, 1},
			{1, 2, 3, 4, 5},
			{7, 9, 2, 1, 6},
			{9, 1, 5, 3, 2},
		}

		exp := [][]int{
			{0, 3, 5, 6, 0},
			{3, 0, 3, 0, 1},
			{1, 2, 0, 4, 5},
			{7, 0, 2, 0, 6},
			{0, 1, 5, 3, 0},
		}
		assert.Equal(t, exp, RunDiagonalTo0(in))
	})
}

func RunDiagonalTo0(arr [][]int) [][]int {
	row := len(arr)
	col := len(arr)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			if i+j+1 == col {
				fmt.Println(i, j, 1, col)

			}

			if i == j || i+j+1 == col {
				arr[i][j] = 0
			}
		}
	}
	return arr
}
