package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromeSubstring(t *testing.T) {
	t.Log(`
	Given a string s, return the longest palindromic substring in s.
	Example 1:
	Input: s = "babad"
	Output: "bab"
	Explanation: "aba" is also a valid answer.
	
	Example 2:
	Input: s = "cbbd"
	Output: "bb"
	`)
	t.Run("Suite 1", func(t *testing.T) {
		s := "babad"
		exp := []string{"bab", "aba"}
		assert.Contains(t, exp, RunLongestPalindromeSubstring(s))
	})
	t.Run("Suite 2", func(t *testing.T) {
		s := "cbbd"
		exp := []string{"bb"}
		assert.Contains(t, exp, RunLongestPalindromeSubstring(s))
	})
}

func RunLongestPalindromeSubstring(s string) string {
	var res string

	for i := range s {
		pOdd := getPlaindrome(s, i, i)
		if len(pOdd) > len(res) {
			res = pOdd
		}
		pEven := getPlaindrome(s, i, i+1)
		if len(pEven) > len(res) {
			res = pEven
		}
	}

	return res
}

func getPlaindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && l <= r && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
