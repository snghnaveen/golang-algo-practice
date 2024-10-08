package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSwaps(t *testing.T) {
	t.Log(`
	You are given a 0-indexed string s of even length n. 
	The string consists of exactly n / 2 opening brackets '[' and n / 2 closing brackets ']'.
	A string is called balanced if and only if:
	It is the empty string, or
	It can be written as AB, where both A and B are balanced strings, or
	It can be written as [C], where C is a balanced string.
	You may swap the brackets at any two indices any number of times.
	Return the minimum number of swaps to make s balanced.

	Example 1:
	Input: s = "][]["
	Output: 1
	Explanation: You can make the string balanced by swapping index 0 with index 3.
	The resulting string is "[[]]".

	Example 2:
	Input: s = "]]][[["
	Output: 2
	Explanation: You can do the following to make the string balanced:
	- Swap index 0 with index 4. s = "[]][][".
	- Swap index 1 with index 5. s = "[[][]]".
	The resulting string is "[[][]]".

	Example 3:
	Input: s = "[]"
	Output: 0
	Explanation: The string is already balanced.
`)

	t.Run("Suite 1", func(t *testing.T) {
		s := "][]["
		exp := 1
		assert.Equal(t, exp, RunMinSwaps(s))
	})
	t.Run("Suite 2", func(t *testing.T) {
		s := "]]][[["
		exp := 2
		assert.Equal(t, exp, RunMinSwaps(s))
	})
	t.Run("Suite 3", func(t *testing.T) {
		s := "[]"
		exp := 0
		assert.Equal(t, exp, RunMinSwaps(s))
	})
}

func RunMinSwaps(s string) int {
	maxSwap := 0
	closed := 0
	for _, c := range s {
		if c == '[' {
			closed--
		} else {
			closed++
		}

		maxSwap = max(maxSwap, closed)
	}

	// Because every 1 swap will actually get ride of 2 extra closed one.
	return (maxSwap + 1) / 2
}
