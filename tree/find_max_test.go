package tree

import (
	"math"
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"

	"github.com/stretchr/testify/assert"
)

func TestFindMaximum(t *testing.T) {
	t.Log(`
	Give an algorithm for finding maximum element in binary tree.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		bt := i.NewBinaryTree()

		for _, v := range []int{5, 1, 3, 9, 10, 2, 4, 7, 6} {
			bt.Insert(v)
		}

		exp := 10
		assert.Equal(t, exp, RunFindMaximum(bt.Root))
	})

}

func RunFindMaximum(root *i.TNode) int {
	max := math.MinInt
	if root != nil {
		leftMax := RunFindMaximum(root.Left)
		rightMax := RunFindMaximum(root.Right)

		if leftMax > rightMax {
			max = leftMax
		} else {
			max = rightMax
		}

		if root.Val > max {
			max = root.Val
		}
	}
	return max
}
