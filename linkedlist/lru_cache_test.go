package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	t.Log(`
	Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
	Implement the LRUCache class:

    LRUCache(int capacity) Initialize the LRU cache with positive size capacity.

    int get(int key) Return the value of the key if the key exists, otherwise return -1.

    void put(int key, int value) Update the value of the key if the key exists. 
	Otherwise, add the key-value pair to the cache. 
	If the number of keys exceeds the capacity from this operation, 
	evict the least recently used key.

	The functions get and put must each run in O(1) average time complexity.
	`)

	intPtr := func(i int) *int {
		return &i
	}

	args := TestInput{
		Action: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
		Value:  [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
		Output: []*int{nil, nil, nil, intPtr(1), nil, intPtr(-1), nil, intPtr(-1), intPtr(3), intPtr(4)},
	}

	RunLRUCache(t, args)
}

type TestInput struct {
	Action []string
	Value  [][]int
	Output []*int
}

func RunLRUCache(t *testing.T, args TestInput) {
	var lru LRUCache
	for i, action := range args.Action {
		value := args.Value[i]
		switch action {
		case "LRUCache":
			lru = Constructor(value[0])
		case "put":
			k := value[0]
			v := value[1]
			lru.Put(k, v)
		case "get":
			k := value[0]
			exp := args.Output[i]
			assert.Equal(t, *exp, lru.Get(k))
		}
	}
}

// actul implementation starts from here
type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	Mp       map[int]*Node
	Left     *Node
	Right    *Node
	Capicity int
}

func Constructor(capacity int) LRUCache {
	left := &Node{}
	right := &Node{}

	left.Next = right
	right.Prev = left
	return LRUCache{
		Mp:       make(map[int]*Node),
		Left:     left,
		Right:    right,
		Capicity: capacity,
	}
}

func (lru *LRUCache) Remove(node *Node) {
	prev, next := node.Prev, node.Next
	prev.Next = next
	next.Prev = prev
}

func (lru *LRUCache) Insert(node *Node) {
	prev, next := lru.Right.Prev, lru.Right
	next.Prev = node
	node.Next = lru.Right

	prev.Next = node
	node.Prev = prev

}

func (lru *LRUCache) Get(key int) int {
	if v, ok := lru.Mp[key]; ok {
		lru.Remove(v)
		lru.Insert(v)
		return v.Val
	}
	return -1

}

func (lru *LRUCache) Put(key int, value int) {
	if v, ok := lru.Mp[key]; ok {
		lru.Remove(v)
	}

	lru.Mp[key] = &Node{Key: key, Val: value}
	lru.Insert(lru.Mp[key])

	if len(lru.Mp) > lru.Capicity {
		next := lru.Left.Next
		lru.Remove(next)
		delete(lru.Mp, next.Key)
	}
}
