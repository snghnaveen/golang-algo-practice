package linkedlist

import (
	"testing"

	l "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestReorder(t *testing.T) {
	t.Log(`
	You are given the head of a singly linked-list. The list can be represented as:

	L0 → L1 → … → Ln - 1 → Ln
	Reorder the list to be on the following form:

	L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
	You may not modify the values in the list's nodes. Only nodes themselves may be changed.

	Input: head = [1,2,3,4]
	Output: [1,4,2,3]

	Input: head = [1,2,3,4,5]
	Output: [1,5,2,4,3]
	`)
	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 2, 3, 4}
		exp := []int{1, 4, 2, 3}

		ll := l.NewLinkedList()
		for _, v := range in {
			ll.Insert(v)
		}
		RunReorderList(ll)
		assert.Equal(t, exp, ll.GetAllValuesAsArr())
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		exp := []int{1, 5, 2, 4, 3}

		ll := l.NewLinkedList()
		for _, v := range in {
			ll.Insert(v)
		}
		RunReorderList(ll)
		assert.Equal(t, exp, ll.GetAllValuesAsArr())
	})
}

func RunReorderList(ll *l.LinkedList) {
	head := ll.Head
	if head == nil || head.Next == nil {
		return
	}

	// Find the middle of the list
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half of the list
	var prev *l.Node
	cur := slow.Next
	slow.Next = nil // Split the list into two halves
	for cur != nil {
		nextNode := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextNode
	}

	// Merge the two halves
	first, second := head, prev
	for first != nil && second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}
