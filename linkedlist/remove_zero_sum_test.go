package linkedlist

import (
	"testing"

	l "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestRemoveZeroSumSublists(t *testing.T) {
	t.Log(`
	Given the head of a linked list, 
	we repeatedly delete consecutive sequences of nodes that sum to 0 until there 
	are no such sequences.

	After doing so, return the head of the final linked list.  
	You may return any such answer.
	
	Example 1:
	Input: head = [1,2,-3,3,1]
	Output: [3,1]
	Note: The answer [1,2,1] would also be accepted.

	Example 2:
	Input: head = [1,2,3,-3,4]
	Output: [1,2,4]

	Example 3:
	Input: head = [1,2,3,-3,-2]
	Output: [1]
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 2, -3, 3, 1}
		exp := [][]int{
			{3, 1},
			{1, 2, 1},
		}

		ll := l.NewLinkedList()
		for _, v := range in {
			ll.Insert(v)
		}
		assert.Contains(t, exp, RunRemoveZeroSumSublists(ll).GetAllValuesAsArr())
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{1, 2, 3, -3, 4}
		exp := [][]int{
			{1, 2, 4},
		}

		ll := l.NewLinkedList()
		for _, v := range in {
			ll.Insert(v)
		}
		assert.Contains(t, exp, RunRemoveZeroSumSublists(ll).GetAllValuesAsArr())
	})
}

// todo revisit
func RunRemoveZeroSumSublists(ll *l.LinkedList) *l.LinkedList {
	head := ll.Head
	dummy := &l.Node{Info: 0, Next: head}
	prefixSum := make(map[int]*l.Node)
	prefixSum[0] = dummy
	sum := 0

	for node := head; node != nil; node = node.Next {
		sum += node.Info
		if prev, found := prefixSum[sum]; found {
			// Remove nodes with zero sum
			sumToRemove := sum
			for n := prev.Next; n != node; n = n.Next {
				sumToRemove += n.Info
				delete(prefixSum, sumToRemove)
			}
			prev.Next = node.Next
		} else {
			prefixSum[sum] = node
		}
	}
	ll.Head = dummy.Next
	return ll
}
