package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateArray(t *testing.T) {
	t.Log(`
	Given an array of integers arr[] of size N and an integer, 
	the task is to rotate the array elements to the left by d positions.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7}
		rotateBy := 2
		expected := []int{3, 4, 5, 6, 7, 1, 2}

		assert.Equal(t, RunTestRotateArray(in, rotateBy), expected)
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{3, 4, 5, 6, 7, 1, 2}
		rotateBy := 2
		expected := []int{5, 6, 7, 1, 2, 3, 4}

		assert.Equal(t, RunTestRotateArray(in, rotateBy), expected)
	})
}

func RunTestRotateArray(arr []int, x int) []int {
	n := len(arr)

	// The modulo operation is used to ensure that the rotation amount X is within the range of the array size n.
	// Consider an example where the length of the array is 5 and we want to rotate it by 7 positions. If we simply use X as the rotation amount, we would end up rotating the array by 7 positions, which is equivalent to rotating it by 2 positions (i.e., 7 % 5 = 2).
	// Without the modulo operation, we would end up rotating the array by more than one full rotation, which is not what we want. By taking the modulo of X and n, we ensure that the actual rotation amount is always within the range of the array size.
	// We calculate the modulo of X and n using the expression X = X % n. This ensures that the rotation amount is always less than the size of the array.
	x = x % n // in case X is greater than n

	arr = append(arr[x:], arr[:x]...)
	return arr
}
