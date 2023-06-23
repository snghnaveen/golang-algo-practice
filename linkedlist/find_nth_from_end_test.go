package linkedlist

import (
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestFindNthFromEnd(t *testing.T) {
	t.Log(`
		Find nth node from the end of a Linked List
		`)

	ll := i.NewLinkedList()
	sample := []int{1, 2, 3, 4}
	for _, val := range sample {
		ll.Insert(val)
	}

	in := ll

	t.Run("Suite 1", func(t *testing.T) {
		nth := 1
		exp := 4
		assert.Equal(t, exp, RunTestGetNthFromEnd(in, nth))
	})

	t.Run("Suite 2", func(t *testing.T) {
		nth := 3
		exp := 2
		assert.Equal(t, exp, RunTestGetNthFromEnd(in, nth))
	})

	t.Run("Suite 3", func(t *testing.T) {
		nth := 4
		exp := 1
		assert.Equal(t, exp, RunTestGetNthFromEnd(in, nth))
	})

	assert.Equal(t, RunTestGetNthFromEnd(ll, 3), 2)
	assert.Equal(t, RunTestGetNthFromEnd(ll, 4), 1)
}

func RunTestGetNthFromEnd(ll *i.LinkedList, n int) int {
	llLen := ll.GetLen()
	required := llLen - n + 1
	node := ll.Head

	current := 0
	for node != nil {
		current++

		if current == required {
			return node.Info
		}

		node = node.Next
	}
	return 0

}
