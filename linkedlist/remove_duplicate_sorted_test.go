package linkedlist

import (
	"testing"

	l "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestDuplicateElementsSorted(t *testing.T) {
	t.Log(`
	Given the head of a sorted linked list, delete all duplicates 
	such that each element appears only once. 
	Return the linked list sorted as well.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		ll := l.NewLinkedList()
		for _, v := range []int{1, 1, 1, 1, 2} {
			ll.Insert(v)
		}

		DuplicateElementsSorted(ll)

		exp := []int{1, 2}
		assert.Equal(t, exp, ll.GetAllValuesAsArr())
	})

	t.Run("Suite 2", func(t *testing.T) {
		ll := l.NewLinkedList()
		for _, v := range []int{1, 1, 2, 3, 3} {
			ll.Insert(v)
		}

		DuplicateElementsSorted(ll)

		exp := []int{1, 2, 3}
		assert.Equal(t, exp, ll.GetAllValuesAsArr())
	})

	t.Run("Suite 3", func(t *testing.T) {
		ll := l.NewLinkedList()

		DuplicateElementsSorted(ll)

		exp := []int{}
		assert.Equal(t, exp, ll.GetAllValuesAsArr())
	})

	t.Run("Suite 4", func(t *testing.T) {
		ll := l.NewLinkedList()
		for _, v := range []int{7, 7, 7, 7} {
			ll.Insert(v)
		}

		DuplicateElementsSorted(ll)

		exp := []int{7}
		assert.Equal(t, exp, ll.GetAllValuesAsArr())
	})
}

func DuplicateElementsSorted(ll *l.LinkedList) {
	node := ll.Head

	for node != nil {
		for node.Next != nil && node.Next.Info == node.Info {
			node.Next = node.Next.Next
		}
		node = node.Next
	}
}
