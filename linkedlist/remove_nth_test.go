package linkedlist

import (
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	t.Log(`Link: https://leetcode.com/problems/remove-nth-node-from-end-of-list/`)

	t.Log(`
	Given the head of a linked list, 
	remove the nth node from the end of the list and return its head.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		n := 2
		exp := []int{1, 2, 3, 5}
		ll := i.NewLinkedList()
		for _, v := range []int{1, 2, 3, 4, 5} {
			ll.Insert(v)
		}

		assert.Equal(t, exp, RunRemoveNthElements(ll, n).GetAllValuesAsArr())
	})
	t.Run("Suite 2", func(t *testing.T) {
		n := 1
		exp := []int{}
		ll := i.NewLinkedList()
		for _, v := range []int{1} {
			ll.Insert(v)
		}

		assert.Equal(t, exp, RunRemoveNthElements(ll, n).GetAllValuesAsArr())
	})
	t.Run("Suite 3", func(t *testing.T) {
		n := 1
		exp := []int{1}
		ll := i.NewLinkedList()
		for _, v := range []int{1, 2} {
			ll.Insert(v)
		}

		assert.Equal(t, exp, RunRemoveNthElements(ll, n).GetAllValuesAsArr())
	})
}

func RunRemoveNthElements(ll *i.LinkedList, n int) *i.LinkedList {
	dummy := &i.Node{Info: 0, Next: ll.Head}

	left := dummy
	right := ll.Head

	for n > 0 && right != nil {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next
	ll.Head = dummy.Next
	return ll
}
