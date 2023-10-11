package stack

import (
	"testing"

	s "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElement(t *testing.T) {
	t.Log(`
	Given an array arr[ ] of size N having elements, the task is to find the next greater element for each element of the array in order of their appearance in the array.
	Next greater element of an element in the array is the nearest element on the right which is greater than the current element.
	If there does not exist next greater of current element, then next greater element for current element is -1. For example, next greater of the last element is always -1.
	
	Example 1:
	Input: 
	N = 4, arr[] = [1 3 2 4]
	Output:
	3 4 4 -1
	Explanation:
	In the array, the next larger element 
	to 1 is 3 , 3 is 4 , 2 is 4 and for 4 ? 
	since it doesn't exist, it is -1.

	Example 2:
	Input: 
	N = 5, arr[] [6 8 0 1 3]
	Output:
	8 -1 1 3 -1

	Explanation:
	In the array, the next larger element to 
	6 is 8, for 8 there is no larger elements 
	hence it is -1, for 0 it is 1 , for 1 it 
	is 3 and then for 3 there is no larger 
	element on right and hence -1.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 3, 2, 4}
		exp := []int{3, 4, 4, -1}

		assert.Equal(t, exp, RunNextGreaterElement(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{6, 8, 0, 1, 3}
		exp := []int{8, -1, 1, 3, -1}

		assert.Equal(t, exp, RunNextGreaterElement(in))
	})
}

func RunNextGreaterElement(in []int) []int {
	out := make([]int, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = -1 // default value
	}

	stk := s.NewStack()

	for i, v := range in {

		for !stk.IsEmpty() && in[stk.Peek()] < v {
			idx := stk.Pop()
			out[idx] = v
		}
		stk.Push(i)

	}

	return out
}
