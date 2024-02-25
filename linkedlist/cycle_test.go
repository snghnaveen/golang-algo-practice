package linkedlist

import (
	"testing"

	l "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	t.Log(`
	Given head, the head of a linked list, 
	determine if the linked list has a cycle in it.
	There is a cycle in a linked list if there is some node in 
	the list that can be reached again by continuously following the next pointer. 
	Internally, pos is used to denote the index of the node that tail's 
	next pointer is connected to. Note that pos is not passed as a parameter.
	Return true if there is a cycle in the linked list. Otherwise, return false.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		ll := l.NewLinkedList()
		for _, v := range []int{1, 2, 3, 4, 5} {
			ll.Insert(v)
		}
		assert.False(t, RunHasCycle(ll.Head))
	})
	t.Run("Suite 2", func(t *testing.T) {
		node1 := l.Node{
			Info: 12,
		}

		ll := l.Node{
			Info: 1,
			Next: &l.Node{
				Info: 2,
				Next: &node1,
			},
		}

		node1.Next = &ll

		assert.True(t, RunHasCycle(&node1))
	})
}

func RunHasCycle(head *l.Node) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
