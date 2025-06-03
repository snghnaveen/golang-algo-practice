package linkedlist

import (
	"sort"
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

// @todo update
func Test012SortLinkedList(t *testing.T) {
	t.Log(`
	You are given a Linked List with nodes that have values 0, 1 or 2. Sort the linked list.
	`)

	t.Run("Suite 1", func(t *testing.T) {

		// Insert values in LL
		ll := i.NewLinkedList()
		seq := []int{0, 1, 2, 2, 1, 2, 0, 2, 1, 0, 0}
		for _, val := range seq {
			ll.Insert(val)
		}

		// Sorting array for assertion
		sort.Ints(seq)

		outLL := RunTest012SortLinkedList(ll)
		node := outLL.Head

		assert.Equal(t, len(seq), outLL.Len)

		p := 0
		for node != nil {
			// Asserting sorted array for each element
			assert.Equal(t, seq[p], node.Info)

			node = node.Next
			p++
		}

	})

}

func RunTest012SortLinkedList(ll *i.LinkedList) i.LinkedList {

	zeroLL := i.NewLinkedList()
	oneLL := i.NewLinkedList()
	twoLL := i.NewLinkedList()

	node := ll.Head
	for node != nil {
		switch node.Info {
		case 0:
			zeroLL.Insert(0)
		case 1:
			oneLL.Insert(1)
		case 2:
			twoLL.Insert(2)
		}

		node = node.Next
	}

	out := mergeFirstLLIntoSecondLL(zeroLL, oneLL)
	out = mergeFirstLLIntoSecondLL(out, twoLL)

	return *out
}

func mergeFirstLLIntoSecondLL(l1 *i.LinkedList, l2 *i.LinkedList) *i.LinkedList {
	node := l2.Head

	for node != nil {
		l1.Insert(node.Info)
		node = node.Next
	}

	return l1
}
