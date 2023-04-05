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

	assert.Equal(t, 3, RunTestGetMiddleOfLL(ll))

	ll.Insert(6)
	assert.Equal(t, 4, RunTestGetMiddleOfLL(ll))

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
