package queue

import (
	"testing"

	q "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Log(`
	Give an algorithm for reversing a queue Q
	`)

	t.Run("Suite 1", func(t *testing.T) {

		q := q.NewQueue()

		inArr := []int{1, 3, 5, 6, 7, 9, 11, 15}
		for i := 0; i < len(inArr); i++ {
			q.Enqueue(inArr[i])
		}

		RunReverse(q)

		for i := len(inArr) - 1; i >= 0; i-- {
			assert.Equal(t, inArr[i], q.Dequeue())
		}

		assert.True(t, q.IsEmpty())
	})

}

func RunReverse(queue *q.Queue[int]) {
	stk := q.NewStack()
	for !queue.IsEmpty() {
		stk.Push(queue.Dequeue())
	}

	for !stk.IsEmpty() {
		queue.Enqueue(stk.Pop())
	}
}
