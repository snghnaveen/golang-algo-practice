package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {

	t.Run("Insert", func(t *testing.T) {
		ll := NewLinkedList()

		expected := []int{1, 2, 3, 4, 5, 6}
		for _, v := range expected {
			ll.Insert(v)
		}

		assert.Equal(t, len(expected), ll.Len)
		assert.ElementsMatch(t, expected, GetAllValuesAsArr(ll))
	})
	t.Run("InsertAtHead", func(t *testing.T) {
		ll := NewLinkedList()

		ll.Insert(1)
		ll.Insert(2)
		ll.Insert(3)

		ll.InsertAtHead(0)

		expected := []int{0, 1, 2, 3}
		assert.Equal(t, len(expected), ll.Len)
		assert.ElementsMatch(t, expected, GetAllValuesAsArr(ll))
	})
	t.Run("DeleteAtHead", func(t *testing.T) {
		t.Run("When elements are present", func(t *testing.T) {
			ll := NewLinkedList()

			ll.Insert(1)
			ll.Insert(2)
			ll.Insert(3)
			ll.DeleteAtHead()

			expected := []int{2, 3}
			assert.Equal(t, len(expected), ll.Len)
			assert.ElementsMatch(t, expected, GetAllValuesAsArr(ll))
		})

		t.Run("When elements are not present", func(t *testing.T) {
			ll := NewLinkedList()
			ll.DeleteAtHead()

			expected := []int{}

			assert.Equal(t, len(expected), ll.Len)
			assert.ElementsMatch(t, expected, GetAllValuesAsArr(ll))
		})
		t.Run("When only 1 element is present", func(t *testing.T) {
			ll := NewLinkedList()
			ll.Insert(1)
			ll.DeleteAtHead()
			ll.DeleteAtHead()

			expected := []int{}

			assert.Equal(t, len(expected), ll.Len)
			assert.ElementsMatch(t, expected, GetAllValuesAsArr(ll))
		})

	})
	t.Run("DeleteAtTail", func(t *testing.T) {
		t.Run("When elements are present", func(t *testing.T) {
			ll := NewLinkedList()

			ll.Insert(1)
			ll.Insert(2)
			ll.Insert(3)
			ll.DeleteAtTail()

			in := []int{1, 2}
			assert.Equal(t, len(in), ll.Len)
			assert.ElementsMatch(t, in, GetAllValuesAsArr(ll))
		})
		t.Run("When elements are not present", func(t *testing.T) {
			ll := NewLinkedList()
			ll.DeleteAtTail()

			in := []int{}

			assert.Equal(t, len(in), ll.Len)
			assert.ElementsMatch(t, in, GetAllValuesAsArr(ll))
		})
		t.Run("When only 1 element is present", func(t *testing.T) {
			ll := NewLinkedList()
			ll.Insert(1)
			ll.DeleteAtTail()
			ll.DeleteAtTail()

			expected := []int{}

			assert.Equal(t, len(expected), ll.Len)
			assert.ElementsMatch(t, expected, GetAllValuesAsArr(ll))
		})
	})

	t.Run("DeleteFirstMatchingElement", func(t *testing.T) {
		t.Run("When elements are present and value is found", func(t *testing.T) {
			ll := NewLinkedList()

			ll.Insert(1)
			ll.Insert(2)
			ll.Insert(3)
			isDeleted := ll.DeleteFirstMatchingElement(1)

			expected := []int{2, 3}
			assert.Equal(t, len(expected), ll.Len)
			assert.ElementsMatch(t, expected, GetAllValuesAsArr(ll))
			assert.True(t, isDeleted)

			isDeleted = ll.DeleteFirstMatchingElement(3)

			expected = []int{2}
			assert.Equal(t, len(expected), ll.Len)
			assert.ElementsMatch(t, expected, GetAllValuesAsArr(ll))
			assert.True(t, isDeleted)

		})
		t.Run("When elements are present and value is not found", func(t *testing.T) {
			ll := NewLinkedList()

			ll.Insert(2)
			ll.Insert(3)
			ll.Insert(5)

			isDeleted := ll.DeleteFirstMatchingElement(10)

			expected := []int{2, 3, 5}
			assert.Equal(t, len(expected), ll.Len)
			assert.ElementsMatch(t, expected, GetAllValuesAsArr(ll))
			assert.False(t, isDeleted)
		})
		t.Run("When elements are present and duplicate values are found", func(t *testing.T) {
			ll := NewLinkedList()

			ll.Insert(3)
			ll.Insert(2)
			ll.Insert(5)
			ll.Insert(3)
			ll.Insert(5)

			isDeleted := ll.DeleteFirstMatchingElement(5)

			expected := []int{3, 2, 3, 5}
			assert.Equal(t, len(expected), ll.Len)
			assert.ElementsMatch(t, expected, GetAllValuesAsArr(ll))
			assert.True(t, isDeleted)
		})

		t.Run("When elements are not present", func(t *testing.T) {
			ll := NewLinkedList()
			isDeleted := ll.DeleteFirstMatchingElement(1)

			expected := []int{}

			assert.Equal(t, len(expected), ll.Len)
			assert.ElementsMatch(t, expected, GetAllValuesAsArr(ll))
			assert.False(t, isDeleted)
		})
		t.Run("When only 1 element is present and matched", func(t *testing.T) {
			ll := NewLinkedList()
			ll.Insert(1)
			isDeleted := ll.DeleteFirstMatchingElement(1)

			expected := []int{}

			assert.Equal(t, len(expected), ll.Len)
			assert.ElementsMatch(t, expected, GetAllValuesAsArr(ll))
			assert.True(t, isDeleted)
		})
	})

}
