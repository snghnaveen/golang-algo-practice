package linkedlist

import (
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestFindMiddleOfLL(t *testing.T) {
	t.Log(`
		Find the middle of a given linked list
		`)

	ll := i.NewLinkedList()
	for i := 1; i <= 5; i++ {
		ll.Insert(i)
	}

	t.Run("Suite 1", func(t *testing.T) {
		exp := 3
		assert.Equal(t, exp, RunTestGetMiddleOfLL(ll))
	})

	t.Run("Suite 2", func(t *testing.T) {
		ll.Insert(6)

		exp := 4
		assert.Equal(t, exp, RunTestGetMiddleOfLL(ll))
	})
}

func RunTestGetMiddleOfLL(ll *i.LinkedList) int {
	slowPtr := ll.Head
	fastPtr := ll.Head

	for fastPtr != nil && fastPtr.Next != nil {

		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
	}

	return slowPtr.Info
}
