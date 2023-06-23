package string

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromicSubstring(t *testing.T) {
	t.Log(`
	Given a string s, return the number of palindromic substrings in it.
	A string is a palindrome when it reads the same backward as forward.
	A substring is a contiguous sequence of characters within the string.

	Example 1:

	Input: s = "abc"
	Output: 3
	Explanation: Three palindromic strings: "a", "b", "c".
	Example 2:

	Input: s = "aaa"
	Output: 6
	Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := "aaa"
		exp := 6

		assert.Equal(t, exp, RunPalindromicSubstring(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := "abc"
		exp := 3

		assert.Equal(t, exp, RunPalindromicSubstring(in))
	})

	t.Run("Suite 3", func(t *testing.T) {
		in := "aaab"
		exp := 7

		assert.Equal(t, exp, RunPalindromicSubstring(in))
	})
}

func RunPalindromicSubstring(s string) int {
	var out int

	arr := strings.Split(s, "")

	for i := 0; i <= len(arr); i++ {

		// odd
		// <1>
		// <-3->
		// <--5-->
		out = out + countPalindrome(arr, i, i)

		// even
		// <2>
		// <-4->
		// <--6-->
		out = out + countPalindrome(arr, i, i+1)
	}

	return out
}

func countPalindrome(arr []string, left, right int) int {
	var count int
	for left >= 0 && right < len(arr) && left <= right {
		if arr[left] == arr[right] {
			count++
			left--
			right++
		} else {
			return count
		}
	}
	return count
}
