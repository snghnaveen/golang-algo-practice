package string

import (
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

	for i := range s {
		// Here, countPalindrome is called with the same index (i, i) to find
		// palindromes of odd (1) length. A palindrome can have a single
		// character in the center (e.g., "racecar" identifies "e" when i is 3).
		// <1>
		// <-3->
		// <--5-->
		out = out + countPalindrome(s, i, i)

		// Here, countPalindrome is called with consecutive indices (i, i+1)
		// to find palindromes of even (2) length. This accounts for cases
		// where the palindrome is centered between two characters (e.g.,
		// "abba" identifies "bb" when i is 1).
		// <1,2>
		// <-2,3->
		// <--3,4-->
		out = out + countPalindrome(s, i, i+1)
	}

	return out
}

func countPalindrome(s string, left, right int) int {
	var count int
	for left >= 0 && right < len(s) && left <= right {
		if s[left] == s[right] {
			count++
			left--
			right++
		} else {
			return count
		}
	}
	return count
}
