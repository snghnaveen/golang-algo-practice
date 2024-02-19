package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	t.Log(`
	Given an input string s, reverse the order of the words.	
	A word is defined as a sequence of non-space characters. 
	The words in s will be separated by at least one space.
	Return a string of the words in reverse order concatenated by a single space.
	Note that s may contain leading or trailing spaces or multiple spaces between two words. 
	The returned string should only have a single space separating the words. '
	Do not include any extra spaces.

	Example 1:
	Input: s = "the sky is blue"
	Output: "blue is sky the"

	Example 2:
	Input: s = "  hello world  "
	Output: "world hello"
	Explanation: Your reversed string should not contain leading or trailing spaces.

	Example 3:
	Input: s = "a good   example"
	Output: "example good a"
	Explanation: You need to reduce multiple spaces between two words to a single space 
	in the reversed string.
`)

	t.Run("Suite 1", func(t *testing.T) {
		in := "the sky is blue"
		exp := "blue is sky the"
		assert.Equal(t, exp, RunReverseWords(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := "  hello world  "
		exp := "world hello"
		assert.Equal(t, exp, RunReverseWords(in))
	})
	t.Run("Suite 3", func(t *testing.T) {
		in := "a good   example"
		exp := "example good a"
		assert.Equal(t, exp, RunReverseWords(in))
	})
}

func RunReverseWords(s string) string {
	rev := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}

		j := i

		for j < len(s) && s[j] != ' ' {
			j++
		}

		if rev == "" {
			rev = s[i:j]
		} else {
			rev = s[i:j] + " " + rev
		}
		i = j
	}
	return rev
}
