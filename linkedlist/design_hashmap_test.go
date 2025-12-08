package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type HashMapTestInput struct {
	Action []string
	Value  [][]int
	Output []*int
}

func TestDesignHashMap(t *testing.T) {
	t.Log("Link: https://leetcode.com/problems/design-hashmap/description/")

	t.Log(`
	Design a HashMap without using any built-in hash table libraries.
	Implement the MyHashMap class:
	MyHashMap() initializes the object with an empty map.
	void put(int key, int value) inserts a (key, value) pair into the HashMap. 
	If the key already exists in the map, update the corresponding value.
	int get(int key) returns the value to which the specified key is mapped, 
	or -1 if this map contains no mapping for the key.
	void remove(key) removes the key and its corresponding value if the map contains 
	the mapping for the key.
 
	Example 1:
	Input
	["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
	[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
	Output
	[null, null, null, 1, -1, null, 1, null, -1]

	Explanation
	MyHashMap myHashMap = new MyHashMap();
	myHashMap.put(1, 1); // The map is now [[1,1]]
	myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
	myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
	myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
	myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
	myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
	myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
	myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]
	`)

	// t.Run("", fu)
	intPtr := func(i int) *int { return &i }

	args := HashMapTestInput{
		Action: []string{
			"MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get",
		},
		Value: [][]int{
			{}, {1, 1}, {2, 2}, {1}, {3}, {2, 1}, {2}, {2}, {2},
		},
		Output: []*int{
			nil, nil, nil, intPtr(1), intPtr(-1), nil, intPtr(1), nil, intPtr(-1),
		},
	}

	RunHashMap(t, args)
}

func RunHashMap(t *testing.T, args HashMapTestInput) {
	var hm MyHashMap

	for i, action := range args.Action {
		val := args.Value[i]

		switch action {
		case "MyHashMap":
			hm = HMConstructor()

		case "put":
			key := val[0]
			value := val[1]
			hm.Put(key, value)

		case "get":
			key := val[0]
			expected := args.Output[i]
			assert.Equal(t, *expected, hm.Get(key))

		case "remove":
			key := val[0]
			hm.Remove(key)
		}
	}
}

type HMNode struct {
	Key   int
	Value int
	Next  *HMNode
}

type MyHashMap struct {
	bucket []*HMNode
	size   int
}

func HMConstructor() MyHashMap {
	buckets := make([]*HMNode, 1000)
	for i := range buckets {
		buckets[i] = &HMNode{} // dummy HMNode
	}
	return MyHashMap{bucket: buckets, size: 1000}
}

func (m *MyHashMap) hash(key int) int {
	return key % m.size
}

func (m *MyHashMap) Put(key int, value int) {
	idx := m.hash(key)
	curr := m.bucket[idx]

	for curr.Next != nil {
		if curr.Next.Key == key {
			curr.Next.Value = value
			return
		}
		curr = curr.Next
	}
	curr.Next = &HMNode{Key: key, Value: value}
}

func (m *MyHashMap) Get(key int) int {
	idx := m.hash(key)
	curr := m.bucket[idx]
	for curr.Next != nil {
		if curr.Next.Key == key {
			return curr.Next.Value
		}
		curr = curr.Next
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	idx := m.hash(key)
	dummy := m.bucket[idx]
	curr := dummy.Next
	prev := dummy
	for curr != nil {
		if curr.Key == key {
			prev.Next = curr.Next
			break
		}
		prev = curr
		curr = curr.Next
	}
}
