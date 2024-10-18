package array

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertDeleteGetRandom_O_n(t *testing.T) {
	t.Log(`
	Implement the RandomizedSet class:
	RandomizedSet() Initializes the RandomizedSet object.

	bool insert(int val) Inserts an item val into the set if not present. 
	Returns true if the item was not present, false otherwise.
	
	bool remove(int val) Removes an item val from the set if present. 
	Returns true if the item was present, false otherwise.
	
	int getRandom() Returns a random element from the current set of elements
	(it's guaranteed that at least one element exists when this method is called). 
	Each element must have the same probability of being returned.
	
	You must implement the functions of the class such that each function works 
	in average O(1) time complexity.
	`)

	args := TestInput{
		Action: []string{"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"},
		Value:  []int{0, 1, 2, 2, 0, 1, 2, 0},
		Output: []interface{}{nil, true, false, true, 2, true, false, 2},
	}

	RunInsertDeleteGetRandom_O_n(t, args)

}

type TestInput struct {
	Action []string
	Value  []int
	Output []interface{}
}

func RunInsertDeleteGetRandom_O_n(t *testing.T, args TestInput) {
	var rSet RandomizedSet
	for i, action := range args.Action {
		value := args.Value[i]
		output := args.Output[i]
		switch action {
		case "RandomizedSet":
			rSet = Constructor()
		case "insert":
			assert.Equal(t, output, rSet.Insert(value))
		case "remove":
			assert.Equal(t, output, rSet.Remove(value))
		case "getRandom":
			rSet.GetRandom()
		}
	}

}

// actual implementation starts from here
type RandomizedSet struct {
	list []int
	mp   map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		list: make([]int, 0),
		mp:   make(map[int]int),
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.mp[val]; ok {
		return false
	}

	r.list = append(r.list, val)
	idx := len(r.list) - 1
	r.mp[val] = idx
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	if idx, ok := r.mp[val]; ok {
		lastElement := r.list[len(r.list)-1]
		r.list[idx] = lastElement
		r.mp[lastElement] = idx
		r.list = r.list[:len(r.list)-1]
		delete(r.mp, val)
		return true
	}
	return false
}

func (r *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(r.list))
	return r.list[index]
}
