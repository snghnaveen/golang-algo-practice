package internal

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
	Len  int
}
type Node struct {
	Info int
	Next *Node
}

func (l *LinkedList) Insert(d int) {
	newNode := &Node{Info: d, Next: nil}

	if l.Head == nil {
		l.Head = newNode
	} else {
		node := l.Head
		for node.Next != nil {
			node = node.Next
		}
		node.Next = newNode
	}
	l.Len++
}

func (l *LinkedList) InsertAtHead(d int) {
	newNode := &Node{Info: d, Next: nil}

	if l.Head == nil {
		l.Head = newNode
	} else {
		head := l.Head
		newNode.Next = head
		l.Head = newNode
	}
	l.Len++
}

func (l *LinkedList) DeleteAtIndex(idx int) {
	if l.Len == 0 {
		return
	}

	p := l.Head
	q := l.Head.Next
	for index := 0; index < idx-1; index++ {
		p = p.Next
		q = q.Next
	}

	p.Next = q.Next
	l.Len--
}

func (l *LinkedList) DeleteFirstMatchingElement(find int) bool {
	if l.Len == 0 {
		return false
	}

	p := l.Head
	q := l.Head.Next

	if p.Info == find {
		l.DeleteAtHead()
		return true
	}

	for q.Info != find && q.Next != nil {
		p = p.Next
		q = q.Next
	}

	if q.Info == find {
		p.Next = q.Next
		l.Len--
		return true
	}

	return false
}

func (l *LinkedList) DeleteAtHead() {
	node := l.Head
	if node != nil {
		l.Head = node.Next
		l.Len--
	}
}

func (l *LinkedList) DeleteAtTail() {
	node := l.Head

	// list is empty
	if node == nil {
		return
	}

	// only 1 element in linked list
	if node.Next == nil {
		l.Head = nil
		l.Len--
		return
	}

	secondLastNode := l.Head

	for secondLastNode.Next.Next != nil {
		secondLastNode = secondLastNode.Next
	}
	secondLastNode.Next = nil
	l.Len--
}

func Show(l *LinkedList) {
	node := l.Head
	for node != nil {
		fmt.Printf("-> %v ", node.Info)
		node = node.Next
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Len: 0}
}

// GetAllValuesAsArr return slice of values stored in slice
func GetAllValuesAsArr(l *LinkedList) []int {
	var out []int
	node := l.Head
	for node != nil {
		out = append(out, node.Info)
		node = node.Next
	}
	return out
}

func (l *LinkedList) GetAllValuesAsArr() []int {
	var out []int
	var i int
	l.Len = i
	node := l.Head
	for node != nil {
		out = append(out, node.Info)

		i++
		l.Len = i

		node = node.Next
	}
	return out
}

func (l *LinkedList) Reverse() {
	var prev Node
	current := l.Head
	var next Node

	for current != nil {
		next = *current.Next

		current.Next = &prev
		prev = *current

		current = &next
	}
	l.Head = &prev
}

func (l *LinkedList) GetLen() int {
	return l.Len
}
