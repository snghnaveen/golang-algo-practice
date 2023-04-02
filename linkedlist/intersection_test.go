package linkedlist

import (
	"testing"

	i "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	t.Log(`
		Given two (singly) linked lists, determine if the two lists intersect. 
		Return the intersecting node. 
		Note that the intersection is defined based on reference, not value.
		That is, if the kth node of the first linked list is the exact same node (by reference) as the jth node of the second linked list, then they are intersecting.
		`)
	head1 := i.Node{Info: 10}
	head2 := i.Node{Info: 3}

	n1 := i.Node{Info: 6}
	head2.Next = &n1

	n2 := i.Node{Info: 9}
	head2.Next.Next = &n2

	n3 := i.Node{Info: 15}
	head1.Next = &n3
	head2.Next.Next.Next = &n3

	n4 := i.Node{Info: 30}
	head1.Next.Next = &n4

	head1.Next.Next.Next = nil

	var l1 i.LinkedList
	l1.Head = &head1

	var l2 i.LinkedList
	l2.Head = &head2

	// l1 = [10 15 30]
	//l2 [3 6 9 15 30]
	assert.True(t, RunTestIntersection(l1, l2))
}

func RunTestIntersection(l1, l2 i.LinkedList) bool {
	node1 := l1.Head
	node2 := l2.Head

	count1, count2 := 0, 0

	for node1 != nil {
		count1++
		node1 = node1.Next
	}

	// reset pointer
	node1 = l1.Head

	for node2 != nil {
		count2++
		node2 = node2.Next
	}

	// reset pointer
	node2 = l2.Head

	diff := abs(count1 - count2)

	if count1 > count2 {
		for i := 0; i < diff; i++ {
			node1 = node1.Next
		}
	} else {
		for i := 0; i < diff; i++ {
			node2 = node2.Next
		}
	}

	addMp := make(map[*i.Node]struct{})

	for node1 != nil && node2 != nil {
		addMp[node2] = struct{}{}
		if _, ok := addMp[node1]; ok {
			return true
		}

		node1 = node1.Next
		node2 = node2.Next
	}

	return false
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
