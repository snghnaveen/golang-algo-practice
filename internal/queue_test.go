package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("Should Panic on empty dequeue", func(t *testing.T) {
		q := NewQueue()

		assert.Panics(t, func() {
			q.Dequeue()
		})
	})
	t.Run("Enqueue and Dequeue", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		q := NewQueue()

		for i := 0; i < len(in); i++ {
			q.Enqueue(in[i])
		}

		for i := 0; i < len(in); i++ {
			assert.Equal(t, in[i], q.Dequeue())
		}
	})

	t.Run("Is empty", func(t *testing.T) {
		q := NewQueue()
		assert.True(t, q.IsEmpty())

		q.Enqueue(1)
		assert.False(t, q.IsEmpty())

		_ = q.Dequeue()
		assert.True(t, q.IsEmpty())
	})
}
