package queue

import (
	"testing"

	q "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Log(`
	Implement a Stack using two queues
	`)

	s := NewStack()

	for i := 3; i <= 8; i++ {
		s.Push(i)
	}

	for i := 8; i >= 3; i-- {
		assert.Equal(t, i, s.Pop())
	}

	assert.Panics(t, func() {
		s.Pop()
	})
}

type Stack struct {
	queue1, queue2 q.Queue[int]
}

func NewStack() *Stack {
	return &Stack{
		queue1: *q.NewQueue(),
		queue2: *q.NewQueue(),
	}
}

func (s *Stack) Push(i int) {
	for !s.queue1.IsEmpty() {
		s.queue2.Enqueue(s.queue1.Dequeue())
	}

	s.queue1.Enqueue(i)

	for !s.queue2.IsEmpty() {
		s.queue1.Enqueue(s.queue2.Dequeue())
	}
}

func (s *Stack) Pop() int {
	if s.queue1.IsEmpty() {
		panic("stack is empty")
	}
	return s.queue1.Dequeue()
}
