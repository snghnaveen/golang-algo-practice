package linkedlist

import (
	"testing"

	l "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
	t.Log(`
	Given the head of a linked list and an integer val, 
	remove all the nodes of the linked list that has Node.val == val, 
	and return the new head.

	Example 1:
	Input: head = [1,2,6,3,4,5,6], val = 6
	Output: [1,2,3,4,5]

	Example 2:
	Input: head = [], val = 1
	Output: []

	Example 3:
	Input: head = [7,7,7,7], val = 7
	Output: []
	`)

	t.Run("Suite 1", func(t *testing.T) {
		ll := l.NewLinkedList()
		for _, v := range []int{1, 2, 6, 3, 4, 5, 6} {
			ll.Insert(v)
		}

		val := 6
		RunRemoveElements(ll, val)

		exp := []int{1, 2, 3, 4, 5}
		assert.Equal(t, exp, ll.GetAllValuesAsArr())
	})

	t.Run("Suite 2", func(t *testing.T) {
		ll := l.NewLinkedList()

		val := 6
		RunRemoveElements(ll, val)

		exp := []int{}
		assert.Equal(t, exp, ll.GetAllValuesAsArr())
	})

	t.Run("Suite 2", func(t *testing.T) {
		ll := l.NewLinkedList()

		val := 6
		RunRemoveElements(ll, val)

		exp := []int{}
		assert.Equal(t, exp, ll.GetAllValuesAsArr())
	})

	t.Run("Suite 3", func(t *testing.T) {
		ll := l.NewLinkedList()
		for _, v := range []int{7, 7, 7, 7} {
			ll.Insert(v)
		}

		val := 7
		RunRemoveElements(ll, val)

		exp := []int{}
		assert.Equal(t, exp, ll.GetAllValuesAsArr())
	})
}

func RunRemoveElements(ll *l.LinkedList, val int) {
	dummy := &l.Node{Next: ll.Head}
	prev := dummy
	curr := ll.Head

	for curr != nil {
		if curr.Info == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}

	ll.Head = dummy.Next
}
