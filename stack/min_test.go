package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinElementInStack(t *testing.T) {
	t.Log(`How would you design a stack which, in addition to push and pop, 
	has a function min which returns the minimum element? 
	Push, pop and min should all operate in 0(1) time.`)

	var minStack MinStack

	minStack.Push(4)
	minStack.Push(3)
	minStack.Push(1)
	minStack.Push(2)

	exp := 1
	assert.Equal(t, exp, minStack.Min())

	minStack.Pop()
	minStack.Pop()

	exp = 3
	assert.Equal(t, exp, minStack.Min())
}

type MinStack struct {
	stack    []int
	minStack []int
}

func (m *MinStack) Push(x int) {
	m.stack = append(m.stack, x)
	if len(m.minStack) == 0 || x <= m.minStack[len(m.minStack)-1] {
		m.minStack = append(m.minStack, x)
	}
}

func (m *MinStack) Pop() {
	if m.stack[len(m.stack)-1] == m.minStack[len(m.minStack)-1] {
		m.minStack = m.minStack[:len(m.minStack)-1]
	}
	m.stack = m.stack[:len(m.stack)-1]
}

func (m *MinStack) Min() int {
	return m.minStack[len(m.minStack)-1]
}
