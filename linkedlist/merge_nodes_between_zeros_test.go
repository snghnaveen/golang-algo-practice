package linkedlist

import (
	"testing"

	l "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestMergeNodesBetweeenZeros(t *testing.T) {
	t.Log(`https://leetcode.com/problems/merge-nodes-in-between-zeros/description/`)

	t.Log(`
	You are given the head of a linked list, which contains a series of integers separated by 0's. The beginning and end of the linked list will have Node.val == 0.
	For every two consecutive 0's, merge all the nodes lying in between them into a single node whose value is the sum of all the merged nodes. The modified list should not contain any 0's.
	Return the head of the modified linked list.

	Example 1:
	Input: head = [0,3,1,0,4,5,2,0]
	Output: [4,11]
	Explanation: 
	The above figure represents the given linked list. The modified list contains
	- The sum of the nodes marked in green: 3 + 1 = 4.
	- The sum of the nodes marked in red: 4 + 5 + 2 = 11.

	Example 2:
	Input: head = [0,1,0,3,0,2,2,0]
	Output: [1,3,4]
	Explanation: 
	The above figure represents the given linked list. The modified list contains
	- The sum of the nodes marked in green: 1 = 1.
	- The sum of the nodes marked in red: 3 = 3.
	- The sum of the nodes marked in yellow: 2 + 2 = 4.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := l.NewLinkedList()
		for _, v := range []int{0, 3, 1, 0, 4, 5, 2, 0} {
			in.Insert(v)
		}

		exp := l.NewLinkedList()
		for _, v := range []int{4, 11} {
			exp.Insert(v)
		}

		out := l.NewLinkedList()
		out.Head = RunMergeNodesBetweeenZeros(in.Head)

		assert.Equal(t, exp.GetAllValuesAsArr(), out.GetAllValuesAsArr())
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := l.NewLinkedList()
		for _, v := range []int{0, 1, 0, 3, 0, 2, 2, 0} {
			in.Insert(v)
		}

		exp := l.NewLinkedList()
		for _, v := range []int{1, 3, 4} {
			exp.Insert(v)
		}

		out := l.NewLinkedList()
		out.Head = RunMergeNodesBetweeenZeros(in.Head)

		assert.Equal(t, exp.GetAllValuesAsArr(), out.GetAllValuesAsArr())
	})
}

func RunMergeNodesBetweeenZeros(head *l.Node) *l.Node {
	dummy := &l.Node{}
	curr := head
	out := dummy

	sum := 0
	// since first index in 0
	curr = curr.Next
	for curr != nil {
		if curr.Info == 0 {
			out.Next = &l.Node{Info: sum}
			out = out.Next
			sum = 0
		}
		sum = sum + curr.Info
		curr = curr.Next
	}
	return dummy.Next
}
