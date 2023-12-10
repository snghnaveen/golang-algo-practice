package tree

import (
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	t.Log(`
	You are given two binary trees root1 and root2.
	Imagine that when you put one of them to cover the other,
	some nodes of the two trees are overlapped while the others are not.
	You need to merge the two trees into a new binary tree.
	The merge rule is that if two nodes overlap,
	then sum node values up as the new value of the merged node.
	Otherwise, the NOT null node will be used as the node of the new tree.
	Return the merged tree.
	Note: The merging process must start from the root nodes of both trees.


	Example 1:
	Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
	Output: [3,4,5,5,4,null,7]

	Example 2:
	Input: root1 = [1], root2 = [1,2]
	Output: [2,2]
`)

	t.Run("Suite 1", func(t *testing.T) {
		t1 := i.NewBinaryTree()
		t1.Root = &i.TNode{Val: 1,
			Left:  &i.TNode{Val: 3, Left: &i.TNode{Val: 5}},
			Right: &i.TNode{Val: 2},
		}

		t2 := i.NewBinaryTree()
		t2.Root = &i.TNode{Val: 2,
			Left:  &i.TNode{Val: 1, Right: &i.TNode{Val: 4}},
			Right: &i.TNode{Val: 3, Right: &i.TNode{Val: 7}},
		}

		out := RunMerge(t1.Root, t2.Root)
		bt := i.NewBinaryTree()
		bt.Root = out

		exp := []int{3, 4, 5, 5, 4, 7}
		assert.Equal(t, exp, bt.TraverseLevelOrder())
	})
}

func RunMerge(t1 *i.TNode, t2 *i.TNode) *i.TNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	v1, v2 := 0, 0

	var t1Left *i.TNode
	var t1Right *i.TNode
	if t1 != nil {
		v1 = t1.Val
		t1Left = t1.Left
		t1Right = t1.Right
	}

	var t2Left *i.TNode
	var t2Right *i.TNode
	if t2 != nil {
		v2 = t2.Val
		t2Right = t2.Right
		t2Left = t2.Left
	}

	root := &i.TNode{Val: v1 + v2}
	root.Left = RunMerge(t1Left, t2Left)
	root.Right = RunMerge(t1Right, t2Right)
	return root
}
