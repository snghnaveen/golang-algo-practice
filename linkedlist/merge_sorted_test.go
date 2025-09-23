package linkedlist

import (
	"testing"

	l "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestMergeSorted(t *testing.T) {
	t.Log(`Link: https://leetcode.com/problems/merge-two-sorted-lists/`)

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

		merged := l.NewLinkedList()
		merged.Head = RunMergeSorted(ll1, ll2)

		assert.Equal(t, exp, merged.GetAllValuesAsArr())
	})

	t.Run("Suite 2", func(t *testing.T) {
		ll1 := l.NewLinkedList()
		ll2 := l.NewLinkedList()

		exp := []int{}

		merged := l.NewLinkedList()
		merged.Head = RunMergeSorted(ll1, ll2)

		assert.Equal(t, exp, merged.GetAllValuesAsArr())
	})
	t.Run("Suite 3", func(t *testing.T) {
		ll1 := l.NewLinkedList()
		ll2 := l.NewLinkedList()
		ll2.Insert(0)

		exp := []int{0}

		merged := l.NewLinkedList()
		merged.Head = RunMergeSorted(ll1, ll2)

		assert.Equal(t, exp, merged.GetAllValuesAsArr())
	})

}

func RunMergeSorted(ll1, ll2 *l.LinkedList) *l.Node {
	dummy := &l.Node{}
	tail := dummy

	node1 := ll1.Head
	node2 := ll2.Head

	for node1 != nil && node2 != nil {
		if node1.Info < node2.Info {
			tail.Next = node1
			node1 = node1.Next
		} else {
			tail.Next = node2
			node2 = node2.Next
		}
		tail = tail.Next
	}

	if node1 != nil {
		tail.Next = node1
	} else if node2 != nil {
		tail.Next = node2
	}

	return dummy.Next
}
