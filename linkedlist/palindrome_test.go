package linkedlist

import (
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	t.Log(`
	Implement a function to check if a linked list is a palindrome.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		ll := i.NewLinkedList()
		for _, v := range []int{1, 2, 3, 4, 3, 2, 1} {
			ll.Insert(v)
		}

		assert.True(t, RunIsPalindrome(ll))
		assert.True(t, RunIsPalindromeRev(ll))
	})

	t.Run("Suite 2", func(t *testing.T) {
		ll := i.NewLinkedList()
		for _, v := range []int{1, 2, 3, 4, 3, 7, 1} {
			ll.Insert(v)
		}

		assert.False(t, RunIsPalindrome(ll))
		assert.False(t, RunIsPalindromeRev(ll))
	})

	t.Run("Suite 3", func(t *testing.T) {
		ll := i.NewLinkedList()
		for _, v := range []int{1, 2, 3, 3, 2, 1} {
			ll.Insert(v)
		}

		assert.True(t, RunIsPalindrome(ll))
		assert.True(t, RunIsPalindromeRev(ll))
	})
}

func RunIsPalindromeRev(l *i.LinkedList) bool {
	slow := l.Head
	fast := l.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse 2nd half
	cur := slow
	var prev *i.Node
	for cur != nil {
		nextNode := cur.Next

		cur.Next = prev
		prev = cur

		cur = nextNode
	}
	slow = prev

	// compare
	listA := l.Head
	listB := slow
	for listA != nil && listB != nil {
		if listA.Info != listB.Info {
			return false
		}
		listA = listA.Next
		listB = listB.Next
	}

	return true
}

// using stack
func RunIsPalindrome(l *i.LinkedList) bool {
	slow := l.Head
	fast := l.Head

	stk := i.NewStack()
	for fast != nil && fast.Next != nil {
		stk.Push(slow.Info)
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Has odd number of elements, so skip the middle element
	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		top := stk.Pop()
		if top != slow.Info {
			return false
		}
		slow = slow.Next
	}

	return true
}
