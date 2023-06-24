package stack

import (
	"sort"
	"testing"

	s "github.com/snghnaveen/golang-algo-practice/internal"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {

	t.Log(`
	Write a program to sort a stack such that the smallest items are on the top. 
	You can use an additional temporary stack, 
	but you may not copy the elements into any other data structure (such as an array). 
	The stack supports the following operations: 
		push, pop, peek, and is Empty	
	`)

	t.Run("Suite 1", func(t *testing.T) {
		stk := s.NewStack()
		inArr := []int{19, 16, 14, 12, 11, 23, 47}
		for _, v := range inArr {
			stk.Push(v)
		}

		outStk := RunSort(stk)

		sort.Ints(inArr)
		for i := 0; i < len(inArr); i++ {
			assert.Equal(t, inArr[i], outStk.Pop())
		}
	})
	t.Run("Suite 2", func(t *testing.T) {
		stk := s.NewStack()
		inArr := []int{199, 288, 1, 87, 55, 23, 54}
		for _, v := range inArr {
			stk.Push(v)
		}

		outStk := RunSort(stk)

		sort.Ints(inArr)
		for i := 0; i < len(inArr); i++ {
			assert.Equal(t, inArr[i], outStk.Pop())
		}
	})
}

func RunSort(stk *s.Stack[int]) *s.Stack[int] {
	tempStk := s.NewStack()

	for !stk.IsEmpty() {
		tmp := stk.Pop()

		for !tempStk.IsEmpty() && tempStk.Peek() < tmp {
			stk.Push(tempStk.Pop())
		}

		tempStk.Push(tmp)
	}

	return tempStk
}
