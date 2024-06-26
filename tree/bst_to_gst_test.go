package tree

import (
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestBstToGst(t *testing.T) {
	t.Log(`
	Given the root of a Binary Search Tree (BST), 
	convert it to a Greater Tree such that every key of the original BST is changed 
	to the original key plus the sum of all keys greater than the original key in BST.
	As a reminder, a binary search tree is a tree that satisfies these constraints:

	The left subtree of a node contains only nodes with keys less than the node's key.
	The right subtree of a node contains only nodes with keys greater than the node's key.
	Both the left and right subtrees must also be binary search trees.

	Example 1:
	Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
	Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]

	Example 2:
	Input: root = [0,null,1]
	Output: [1,null,1]
`)

	t.Run("Suite 1", func(t *testing.T) {
		t1 := i.NewBinaryTree()
		for _, i := range []int{4, 1, 6, 0, 2, 5, 7, 3, 8} {
			t1.Insert(i)
		}

		out := RunBstToGst(t1.Root)
		bt := i.NewBinaryTree()
		bt.Root = out

		exp := []int{30, 36, 21, 36, 35, 26, 15, 33, 8}
		assert.Equal(t, exp, bt.TraverseLevelOrder())
	})
}

func RunBstToGst(root *i.TNode) *i.TNode {
	sum := 0
	bstToGst(root, &sum)
	return root
}

func bstToGst(root *i.TNode, sum *int) *i.TNode {
	if root != nil {
		bstToGst(root.Right, sum)
		root.Val = root.Val + *sum
		*sum = root.Val
		bstToGst(root.Left, sum)
	}
	return root
}
