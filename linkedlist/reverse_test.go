package linkedlist

import (
	"slices"
	"testing"

	l "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Log(`
	Given the head of a singly linked list, reverse the list, and return the reversed list.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		ll := l.NewLinkedList()

		arr := []int{1, 2, 6, 3, 4, 5, 6}
		for _, v := range arr {
			ll.Insert(v)
		}

		RunReverse(ll)

		slices.Reverse(arr) // reverse arr for test
		assert.Equal(t, arr, ll.GetAllValuesAsArr())
	})
	t.Run("Suite 2", func(t *testing.T) {
		ll := l.NewLinkedList()

		arr := []int{10, 11, 15, 2}
		for _, v := range arr {
			ll.Insert(v)
		}

		RunReverse(ll)

		slices.Reverse(arr) // reverse arr for test
		assert.Equal(t, arr, ll.GetAllValuesAsArr())
	})

}

func RunReverse(ll *l.LinkedList) {
	var prev *l.Node
	curr := ll.Head

	for curr != nil {
		nxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}

	ll.Head = prev
}
