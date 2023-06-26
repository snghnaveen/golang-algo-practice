package linkedlist

import (
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Log(`
	You are given two non-empty linked lists representing two non-negative integers.
	The digits are stored in reverse order, 
	and each of their nodes contains a single digit. 
	Add the two numbers and return the sum as a linked list.

	You may assume the two numbers do not contain any leading zero, 
	except the number 0 itself.
	
	Example 1:
	Input: l1 = [2,4,3], l2 = [5,6,4]
	Output: [7,0,8]
	Explanation: 342 + 465 = 807

	Example 2:
	Input: l1 = [0], l2 = [0]
	Output: [0]

	Example 3:
	Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	Output: [8,9,9,9,0,0,0,1]
		`)

	t.Run("Suite 1", func(t *testing.T) {

		// Insert values in L1
		l1 := i.NewLinkedList()
		for _, val := range []int{9, 9, 9, 9, 9, 9, 9} {
			l1.Insert(val)
		}

		// Insert values in L2
		l2 := i.NewLinkedList()
		for _, val := range []int{9, 9, 9, 9} {
			l2.Insert(val)
		}

		outLL := RunAddTwoNumbers(l1, l2)
		assert.Equal(t, []int{8, 9, 9, 9, 0, 0, 0, 1}, outLL.GetAllValuesAsArr())
	})

	t.Run("Suite 2", func(t *testing.T) {

		// Insert values in L1
		l1 := i.NewLinkedList()
		for _, val := range []int{2, 4, 3} {
			l1.Insert(val)
		}

		// Insert values in L2
		l2 := i.NewLinkedList()
		for _, val := range []int{5, 6, 4} {
			l2.Insert(val)
		}

		outLL := RunAddTwoNumbers(l1, l2)
		assert.Equal(t, []int{7, 0, 8}, outLL.GetAllValuesAsArr())
	})

}

func RunAddTwoNumbers(l1, l2 *i.LinkedList) *i.LinkedList {
	out := i.NewLinkedList()

	node1 := l1.Head
	node2 := l2.Head

	var carry int
	for node1 != nil || node2 != nil || carry != 0 {
		var val1, val2 int
		if node1 != nil {
			val1 = node1.Info
			node1 = node1.Next
		}

		if node2 != nil {
			val2 = node2.Info
			node2 = node2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10
		sum = sum % 10
		out.Insert(sum)
	}

	return out
}
