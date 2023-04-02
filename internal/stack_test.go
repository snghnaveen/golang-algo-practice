package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	t.Run("Should Panic on empty pop", func(t *testing.T) {
		s := NewStack()

		assert.Panics(t, func() {
			s.Pop()
		})
	})
	t.Run("push and pop", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		s := NewStack()

		for i := 0; i < len(in); i++ {
			s.Push(in[i])
		}

		for i := len(in) - 1; i >= 0; i-- {
			assert.Equal(t, in[i], s.Pop())
		}
	})

	t.Run("Is empty", func(t *testing.T) {
		s := NewStack()
		assert.True(t, s.IsEmpty())

		s.Push(1)
		assert.False(t, s.IsEmpty())

		_ = s.Pop()
		assert.True(t, s.IsEmpty())
	})
}
