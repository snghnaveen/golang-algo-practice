package linkedlist

import (
	"testing"

	l "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestMergeSorted(t *testing.T) {
	t.Log(`
	You are given the heads of two sorted linked lists list1 and list2.
	Merge the two lists into one sorted list. 
	The list should be made by splicing together the nodes of the first two lists.
	Return the head of the merged linked list.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		ll1 := l.NewLinkedList()
		for _, v := range []int{1, 2, 4} {
			ll1.Insert(v)
		}
		ll2 := l.NewLinkedList()
		for _, v := range []int{1, 3, 4} {
			ll2.Insert(v)
		}

		exp := []int{1, 1, 2, 3, 4, 4}

		assert.Equal(t, exp, RunMergeSorted(ll1, ll2).GetAllValuesAsArr())
	})

	t.Run("Suite 2", func(t *testing.T) {
		ll1 := l.NewLinkedList()
		ll2 := l.NewLinkedList()

		exp := []int{}
		assert.Equal(t, exp, RunMergeSorted(ll1, ll2).GetAllValuesAsArr())
	})
	t.Run("Suite 3", func(t *testing.T) {
		ll1 := l.NewLinkedList()
		ll2 := l.NewLinkedList()
		ll2.Insert(0)

		exp := []int{0}
		assert.Equal(t, exp, RunMergeSorted(ll1, ll2).GetAllValuesAsArr())
	})

}

func RunMergeSorted(ll1, ll2 *l.LinkedList) *l.LinkedList {
	out := l.NewLinkedList()

	l1Node, l2Node := ll1.Head, ll2.Head
	for l1Node != nil && l2Node != nil {
		if l1Node.Info > l2Node.Info {
			out.Insert(l2Node.Info)
			l2Node = l2Node.Next
		} else if l1Node.Info < l2Node.Info {
			out.Insert(l1Node.Info)
			l1Node = l1Node.Next
		} else {
			out.Insert(l1Node.Info)
			out.Insert(l2Node.Info)
			l1Node = l1Node.Next
			l2Node = l2Node.Next
		}
	}

	for l1Node != nil {
		out.Insert(l1Node.Info)
		l1Node = l1Node.Next
	}

	for l2Node != nil {
		out.Insert(l2Node.Info)
		l2Node = l2Node.Next
	}

	return out
}
