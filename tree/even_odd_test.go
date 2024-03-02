package tree

import (
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestEvenOdd(t *testing.T) {
	t.Log(`
	A binary tree is named Even-Odd if it meets the following conditions:

	The root of the binary tree is at level index 0, 
	its children are at level index 1, 
	their children are at level index 2, etc.

	For every even-indexed level, 
	all nodes at the level have odd integer values in strictly increasing order 
	(from left to right).

	For every odd-indexed level, 
	all nodes at the level have even integer values in strictly decreasing order 
	(from left to right).

	Given the root of a binary tree, 
	return true if the binary tree is Even-Odd, otherwise return false.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		t1 := i.NewBinaryTree()
		t1.Root = &i.TNode{
			Val: 1,
			Left: &i.TNode{
				Val: 10,
				Left: &i.TNode{
					Val:   3,
					Left:  &i.TNode{Val: 12},
					Right: &i.TNode{Val: 8},
				},
			},
			Right: &i.TNode{
				Val: 4,
				Left: &i.TNode{
					Val:  7,
					Left: &i.TNode{Val: 6},
				},
				Right: &i.TNode{
					Val:   9,
					Right: &i.TNode{Val: 2},
				},
			},
		}
		assert.True(t, RunEvenOdd(t1.Root))
	})

	t.Run("Suite 2", func(t *testing.T) {
		t1 := i.NewBinaryTree()
		t1.Root = &i.TNode{
			Val: 5,
			Left: &i.TNode{
				Val:   9,
				Left:  &i.TNode{Val: 3},
				Right: &i.TNode{Val: 5},
			},
			Right: &i.TNode{
				Val:  2,
				Left: &i.TNode{Val: 7},
			},
		}
		assert.False(t, RunEvenOdd(t1.Root))
	})
}

func RunEvenOdd(root *i.TNode) bool {
	queue := make([]*i.TNode, 1)
	height := 0
	queue[0] = root

	for len(queue) > 0 {
		size := len(queue)

		var prev *i.TNode
		for i := 0; i < size; i++ {
			current := queue[i]

			if height%2 == 0 {
				if current.Val%2 == 0 {
					return false
				}
				if prev != nil && current.Val <= prev.Val {
					return false
				}
			} else {
				if current.Val%2 != 0 {
					return false
				}
				if prev != nil && current.Val >= prev.Val {
					return false
				}
			}
			prev = current
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		queue = queue[size:]
		height++
	}

	return true
}
