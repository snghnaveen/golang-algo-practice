package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	i "github.com/snghnaveen/golang-algo-practice/internal"
)

func TestPathSum(t *testing.T) {
	t.Log(`https://leetcode.com/problems/path-sum/description/`)

	t.Log(`
	Given the root of a binary tree and an integer targetSum,
	return true if the tree has a root-to-leaf path such that adding up all the values
	along the path equals targetSum.
	A leaf is a node with no children.

	Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
	Output: true
	`)

	t.Run("Suite 1", func(t *testing.T) {
		b := i.NewBinaryTree()
		b.Root = &i.TNode{Val: 5,
			Left: &i.TNode{Val: 4,
				Left: &i.TNode{
					Val: 11,
					Left: &i.TNode{
						Val: 7,
					},
					Right: &i.TNode{
						Val: 2,
					},
				}},
			Right: &i.TNode{Val: 8,
				Left: &i.TNode{
					Val: 13,
				},
				Right: &i.TNode{
					Val: 4,
					Right: &i.TNode{
						Val: 1,
					},
				}},
		}
		targetSum := 22
		assert.True(t, RunPathSum(b.Root, targetSum))
	})

	t.Run("Suite 2", func(t *testing.T) {
		b := i.NewBinaryTree()
		b.Root = &i.TNode{
			Val: 1,
			Left: &i.TNode{
				Val: 2,
			},
			Right: &i.TNode{
				Val: 3,
			},
		}
		targetSum := 5
		assert.False(t, RunPathSum(b.Root, targetSum))
	})

}

func RunPathSum(p *i.TNode, targetSum int) bool {
	return dfs(0, targetSum, p)
}

func dfs(currentSum, targetSum int, n *i.TNode) bool {
	if n == nil {
		return false
	}

	currentSum = currentSum + n.Val
	if n.Left == nil && n.Right == nil {
		return currentSum == targetSum
	}

	return dfs(currentSum, targetSum, n.Left) || dfs(currentSum, targetSum, n.Right)
}
