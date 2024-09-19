package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstringWithoutRepeatingChars(t *testing.T) {
	t.Log(`
	Given a string s, find the length of the longest substring without repeating characters.

	Example 1:
	Input: s = "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3.

	Example 2:
	Input: s = "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.

	Example 3:
	Input: s = "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.
	Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := "abcabcbb"
		exp := 3
		assert.Equal(t, exp, RunLongestSubstringWithoutRepeatingChars(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := "bbbbb"
		exp := 1
		assert.Equal(t, exp, RunLongestSubstringWithoutRepeatingChars(in))
	})
	t.Run("Suite 3", func(t *testing.T) {
		in := "pwwkew"
		exp := 3
		assert.Equal(t, exp, RunLongestSubstringWithoutRepeatingChars(in))
	})
	t.Run("Suite 4", func(t *testing.T) {
		in := "aab"
		exp := 2
		assert.Equal(t, exp, RunLongestSubstringWithoutRepeatingChars(in))
	})
}

func RunLongestSubstringWithoutRepeatingChars(s string) int {
	var maxLen int
	mp := make(map[rune]struct{})

	l := 0
	r := 0
	for _, v := range s {
		for l < r {
			if _, ok := mp[v]; ok {
				delete(mp, rune(s[l]))
				l++
			} else {
				break
			}
		}
		mp[v] = struct{}{}
		maxLen = max(maxLen, r-l+1)
		r++
	}

	return maxLen
}
