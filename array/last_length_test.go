package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	t.Log(`
	Given a string s consisting of words and spaces, return the length of the last word in the string.
	A word is a maximal substring consisting of non-space characters only.


	Example 1:
	Input: s = "Hello World"
	Output: 5
	Explanation: The last word is "World" with length 5.

	Example 2:
	Input: s = "   fly me   to   the moon  "
	Output: 4
	Explanation: The last word is "moon" with length 4.

	Example 3:
	Input: s = "luffy is still joyboy"
	Output: 6
	Explanation: The last word is "joyboy" with length 6.

	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := "Hello World"
		exp := 5
		assert.Equal(t, exp, RunLengthOfLastWord(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := "   fly me   to   the moon  "
		exp := 4
		assert.Equal(t, exp, RunLengthOfLastWord(in))
	})

	t.Run("Suite 3", func(t *testing.T) {
		in := "luffy is still joyboy"
		exp := 6
		assert.Equal(t, exp, RunLengthOfLastWord(in))
	})
}

func RunLengthOfLastWord(s string) int {
	i, l := len(s)-1, 0
	for s[i] == ' ' {
		i--
	}

	for i > 0 && s[i] != ' ' {
		l++
		i--
	}

	return l
}
