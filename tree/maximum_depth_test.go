package tree

import (
	"container/list"
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"

	"github.com/stretchr/testify/assert"
)

func TestMaximumDepth(t *testing.T) {
	t.Log("Link: https://leetcode.com/problems/maximum-depth-of-binary-tree/")

	t.Log(`
	Given the root of a binary tree, return its maximum depth.
	A binary tree's maximum depth is the number of nodes along the
	longest path from the root node down to the farthest leaf node.

	Example 1:
	root = [3,9,20,null,null,15,7]
	Output: 3

	Example 2:
	Input: root = [1,null,2]
	Output: 2
	`)

	t.Run("Suite 1", func(t *testing.T) {
		bt := i.NewBinaryTree()

		// [3,9,20,null,null,15,7
		bt.Root = &i.TNode{Val: 3}
		bt.Root.Left = &i.TNode{Val: 9}
		bt.Root.Right = &i.TNode{Val: 20}
		bt.Root.Right.Left = &i.TNode{Val: 15}
		bt.Root.Right.Right = &i.TNode{Val: 7}

		exp := 3
		assert.Equal(t, exp, RunMaximumDepth(bt.Root))
		assert.Equal(t, exp, RunMaximumDepthRec(bt.Root))
	})
}

func RunMaximumDepth(root *i.TNode) int {
	var height int
	if root != nil {

		queue := list.New()
		queue.PushBack(root)

		for queue.Len() > 0 {
			levelSize := queue.Len()
			for j := 0; j < levelSize; j++ {

				element := queue.Front()
				queue.Remove(element)

				current := element.Value.(*i.TNode)

				if current.Left != nil {
					queue.PushBack(current.Left)
				}

				if current.Right != nil {
					queue.PushBack(current.Right)
				}

			}

			height++
		}
	}

	return height
}

func RunMaximumDepthRec(root *i.TNode) int {
	if root == nil {
		return 0
	}
	leftHeight := RunMaximumDepth(root.Left)
	rightHeight := RunMaximumDepth(root.Right)

	return 1 + max(leftHeight, rightHeight)
}
