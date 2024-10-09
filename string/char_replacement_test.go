package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	t.Log(`
	You are given a string s and an integer k. 
	You can choose any character of the string and change 
	it to any other uppercase English character. 
	You can perform this operation at most k times.
	Return the length of the longest substring containing the 
	same letter you can get after performing the above operations.

	Example 1:
	Input: s = "ABAB", k = 2
	Output: 4
	Explanation: Replace the two 'A's with two 'B's or vice versa.

	Example 2:
	Input: s = "AABABBA", k = 1
	Output: 4
	Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
	The substring "BBBB" has the longest repeating letters, which is 4.
	There may exists other ways to achieve this answer too.
`)

	t.Run("Suite 1", func(t *testing.T) {
		s := "ABAB"
		k := 2
		exp := 4
		assert.Equal(t, exp, RunCharacterReplacement(s, k))
	})
	t.Run("Suite 2", func(t *testing.T) {
		s := "AABABBA"
		k := 1
		exp := 4
		assert.Equal(t, exp, RunCharacterReplacement(s, k))
	})
}

func RunCharacterReplacement(s string, k int) int {
	var res int

	maxFreq := func(mpCnt map[byte]int) int {
		maxFreq := 0
		for _, v := range mpCnt {
			maxFreq = max(maxFreq, v)
		}
		return maxFreq
	}

	l, r := 0, 0
	mpCnt := make(map[byte]int)
	for r < len(s) {
		mpCnt[s[r]]++

		if (r-l+1)-maxFreq(mpCnt) <= k {
			res = max(res, r-l+1)
		} else {
			mpCnt[s[l]]--
			l++
		}
		r++
	}
	return res
}
