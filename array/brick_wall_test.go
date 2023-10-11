package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrickWall(t *testing.T) {
	t.Log(`
	There is a rectangular brick wall in front of you with n rows of bricks. 
	The ith row has some number of bricks each of the same height (i.e., one unit) but they can be of 
	different widths. The total width of each row is the same.
	Draw a vertical line from the top to the bottom and cross the least bricks. 
	If your line goes through the edge of a brick, then the brick is not considered as crossed. 
	You cannot draw a line just along one of the two vertical edges of the wall, in which case the 
	line will obviously cross no bricks.
	Given the 2D array wall that contains the information about the wall, 
	return the minimum number of crossed bricks after drawing such a vertical line.
	
	Example 1:
	Input: wall = [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]
	Output: 2

	Example 2:
	Input: wall = [[1],[1],[1]]
	Output: 3
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := [][]int{
			{1, 2, 2, 1},
			{3, 1, 2},
			{1, 3, 2},
			{2, 4},
			{3, 1, 2},
			{1, 3, 1, 1},
		}

		exp := 2
		assert.Equal(t, exp, RunBrickWall(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := [][]int{
			{1},
			{1},
			{1},
		}

		exp := 3
		assert.Equal(t, exp, RunBrickWall(in))
	})

}

func RunBrickWall(in [][]int) int {
	mpCntGap := make(map[int]int)

	for r := 0; r < len(in); r++ {
		total := 0
		for i := 0; i < len(in[r])-1; i++ {
			val := in[r][i]
			total = total + val

			if v, ok := mpCntGap[total]; ok {
				mpCntGap[total] = v + 1
			} else {
				mpCntGap[total] = 1
			}
		}
	}

	maxGap := 0
	for _, gap := range mpCntGap {
		if gap > maxGap {
			maxGap = gap
		}
	}
	return len(in) - maxGap
}
