package linkedlist

import (
	"sort"
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func Test012SortLinkedList(t *testing.T) {
	t.Log(`Link : https://www.geeksforgeeks.org/problems/given-a-linked-list-of-0s-1s-and-2s-sort-it/1`)

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
	count := make([]int, 3)

	current := ll.Head

	for current != nil {
		count[current.Info]++
		current = current.Next
	}

	idx := 0
	current = ll.Head
	for current != nil {
		if count[idx] == 0 {
			idx++
		} else {
			current.Info = idx
			count[idx]--
			current = current.Next
		}
	}

	return *ll
}
