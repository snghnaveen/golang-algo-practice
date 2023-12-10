package tree

import (
	i "github.com/snghnaveen/golang-algo-practice/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSame(t *testing.T) {
	t.Log(`
	Given the roots of two binary trees p and q, write a function to check if they are the same or not.
	Two binary trees are considered the same if they are structurally identical,
	and the nodes have the same value.

	Example 1:
	Input: p = [1,2,3], q = [1,2,3]
	Output: true

	Example 2:
	Input: p = [1,2], q = [1,null,2]
	Output: false

	Example 3:
	Input: p = [1,2,1], q = [1,1,2]
	Output: false
	`)

	t.Run("Suite 1", func(t *testing.T) {
		b1 := i.NewBinaryTree()
		b1.Root = &i.TNode{Val: 1,
			Left:  &i.TNode{Val: 2},
			Right: &i.TNode{Val: 3},
		}

		b2 := i.NewBinaryTree()
		b2.Root = &i.TNode{Val: 1,
			Left:  &i.TNode{Val: 2},
			Right: &i.TNode{Val: 3},
		}

		assert.True(t, RunIsSame(b1.Root, b2.Root))
	})
}

func RunIsSame(p *i.TNode, q *i.TNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	return RunIsSame(p.Left, q.Left) && RunIsSame(p.Right, q.Right)
}
