package internal

import (
	"container/list"
)

type TNode struct {
	Val   int
	Left  *TNode
	Right *TNode
}

type BinaryTree struct {
	Root *TNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Insert(Val int) {
	if t.Root == nil {
		t.Root = &TNode{Val: Val}
		return
	}
	t.insertTNode(t.Root, Val)
}

func (t *BinaryTree) insertTNode(node *TNode, Val int) {
	if Val < node.Val {
		if node.Left == nil {
			node.Left = &TNode{Val: Val}
		} else {
			t.insertTNode(node.Left, Val)
		}
	} else if Val > node.Val {
		if node.Right == nil {
			node.Right = &TNode{Val: Val}
		} else {
			t.insertTNode(node.Right, Val)
		}
	}
}

func (t *BinaryTree) Search(Val int) bool {
	return t.searchTNode(t.Root, Val)
}

func (t *BinaryTree) searchTNode(node *TNode, Val int) bool {
	if node == nil {
		return false
	}

	if Val < node.Val {
		return t.searchTNode(node.Left, Val)
	} else if Val > node.Val {
		return t.searchTNode(node.Right, Val)
	}

	return true
}

func (t *BinaryTree) TraverseInOrder() []int {
	return t.traverseInOrder(t.Root)
}

func (t *BinaryTree) traverseInOrder(node *TNode) []int {
	res := []int{}
	if node == nil {
		return res
	}

	resLeft := t.traverseInOrder(node.Left)
	res = append(res, resLeft...)

	res = append(res, node.Val)

	resRight := t.traverseInOrder(node.Right)
	res = append(res, resRight...)

	return res
}

func (t *BinaryTree) TraversePreOrder() []int {
	return t.traversePreOrder(t.Root)
}

func (t *BinaryTree) traversePreOrder(node *TNode) []int {
	res := []int{}
	if node == nil {
		return res
	}

	res = append(res, node.Val)

	resLeft := t.traversePreOrder(node.Left)
	res = append(res, resLeft...)

	resRight := t.traversePreOrder(node.Right)
	res = append(res, resRight...)

	return res
}

func (t *BinaryTree) TraversePostOrder() []int {
	return t.traversePostOrder(t.Root)
}

func (t *BinaryTree) traversePostOrder(node *TNode) []int {
	res := []int{}
	if node == nil {
		return res
	}

	resLeft := t.traversePostOrder(node.Left)
	res = append(res, resLeft...)

	resRight := t.traversePostOrder(node.Right)
	res = append(res, resRight...)

	res = append(res, node.Val)

	return res
}

func (t *BinaryTree) TraverseLevelOrder() []int {
	return t.traverseLevelOrder(t.Root)
}

func (t *BinaryTree) traverseLevelOrder(node *TNode) []int {
	res := []int{}
	if node == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(node)

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)

		current := element.Value.(*TNode)

		res = append(res, current.Val)

		if current.Left != nil {
			queue.PushBack(current.Left)
		}

		if current.Right != nil {
			queue.PushBack(current.Right)
		}
	}

	return res
}
