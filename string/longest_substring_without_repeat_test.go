package string

import (
	"strings"
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
}

func RunLongestSubstringWithoutRepeatingChars(s string) int {
	var max int
	sArr := strings.Split(s, "")
	mp := make(map[string]struct{})

	l := 0
	r := 0
	for i := 0; i < len(sArr); i++ {
		if _, ok := mp[sArr[i]]; ok {
			delete(mp, sArr[l])
			l++
		} else {
			mp[sArr[i]] = struct{}{}
			tmpMax := r - l + 1
			if tmpMax > max {
				max = tmpMax
			}
			r++
		}
	}

	return max
}
