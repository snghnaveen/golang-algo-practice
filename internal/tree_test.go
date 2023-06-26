package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	bt := NewBinaryTree()

	bt.Root = &TNode{Val: 1}
	bt.Root.Left = &TNode{Val: 2}
	bt.Root.Right = &TNode{Val: 3}
	bt.Root.Left.Left = &TNode{Val: 4}
	bt.Root.Left.Right = &TNode{Val: 5}

	t.Run("InOrderTraversal", func(t *testing.T) {
		assert.Equal(t, []int{4, 2, 5, 1, 3}, bt.TraverseInOrder())
	})

	t.Run("PreOrderTraversal", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 4, 5, 3}, bt.TraversePreOrder())
	})

	t.Run("PostOrderTraversal", func(t *testing.T) {
		assert.Equal(t, []int{4, 5, 2, 3, 1}, bt.TraversePostOrder())
	})

	t.Run("LevelOrderTraversal", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, bt.TraverseLevelOrder())
	})
}
