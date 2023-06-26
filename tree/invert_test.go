package tree

import (
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"

	"github.com/stretchr/testify/assert"
)

func TestInvert(t *testing.T) {
	t.Log(`
	Given the root of a binary tree, invert the tree, and return its root.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		tree := i.NewBinaryTree()

		for _, i := range []int{4, 2, 7, 1, 3, 6, 9} {
			tree.Insert(i)
		}

		initial := []int{4, 2, 7, 1, 3, 6, 9}
		assert.Equal(t, initial, tree.TraverseLevelOrder())

		exp := []int{4, 7, 2, 9, 6, 3, 1}
		out := RunInvert(tree.Root)

		bt := i.NewBinaryTree()
		bt.Root = out

		assert.Equal(t, exp, bt.TraverseLevelOrder())
	})

	t.Run("Suite 2", func(t *testing.T) {
		tree := i.NewBinaryTree()

		for _, i := range []int{2, 1, 3} {
			tree.Insert(i)
		}

		initial := []int{2, 1, 3}
		assert.Equal(t, initial, tree.TraverseLevelOrder())

		exp := []int{2, 3, 1}
		out := RunInvert(tree.Root)

		bt := i.NewBinaryTree()
		bt.Root = out

		assert.Equal(t, exp, bt.TraverseLevelOrder())
	})
}

func RunInvert(tree *i.TNode) *i.TNode {
	if tree != nil {
		tree.Left, tree.Right = RunInvert(tree.Right), RunInvert(tree.Left)
	}
	return tree
}
