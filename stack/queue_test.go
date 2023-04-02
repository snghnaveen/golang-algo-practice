package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"

	s "github.com/snghnaveen/golang-algo-practice/internal"
)

func TestImplementQueueUsingStack(t *testing.T) {
	t.Log(`Implement a MyQueue which implements a queue using two stacks`)

	q := NewQueue()

	for i := 3; i <= 8; i++ {
		q.EnQueue(i)
	}

	for i := 3; i <= 8; i++ {
		assert.Equal(t, q.DeQueue(), i)
	}

	assert.Panics(t, func() {
		q.DeQueue()
	})

}

type Queue struct {
	stack1 s.Stack
	stack2 s.Stack
}

func (q *Queue) EnQueue(x int) {
	for !q.stack1.IsEmpty() {
		q.stack2.Push(q.stack1.Pop())
	}

	q.stack1.Push(x)

	for !q.stack2.IsEmpty() {
		q.stack1.Push(q.stack2.Pop())
	}

}

func (q *Queue) DeQueue() int {
	if q.stack1.IsEmpty() {
		panic("queue is empty")
	}
	return q.stack1.Pop()
}

func NewQueue() *Queue {
	return &Queue{
		stack1: *s.NewStack(),
		stack2: *s.NewStack(),
	}
}
