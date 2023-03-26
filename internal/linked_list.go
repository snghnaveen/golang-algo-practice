package internal

import "fmt"

type Node struct {
	Info int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Len  int
}

func (l *LinkedList) Insert(d int) {
	newNode := &Node{Info: d, Next: nil}

	if l.Head == nil {
		l.Len++
		l.Head = newNode
	} else {
		node := l.Head
		for node.Next != nil {
			node = node.Next
		}
		l.Len++
		node.Next = newNode
	}
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
