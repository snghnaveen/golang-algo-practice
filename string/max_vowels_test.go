package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxVowels(t *testing.T) {
	t.Log(`
	Given a string s and an integer k, return the maximum number of vowel letters in any substring of s 
	with length k.
	Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

	Example 1:
	Input: s = "abciiidef", k = 3
	Output: 3
	Explanation: The substring "iii" contains 3 vowel letters.

	Example 2:
	Input: s = "aeiou", k = 2
	Output: 2
	Explanation: Any substring of length 2 contains 2 vowels.

	Example 3:
	Input: s = "leetcode", k = 3
	Output: 2
	Explanation: "lee", "eet" and "ode" contain 2 vowels.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := "abciiidef"
		k := 3
		exp := 3
		assert.Equal(t, exp, RunMaxVowels(in, k))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := "aeiou"
		k := 2
		exp := 2
		assert.Equal(t, exp, RunMaxVowels(in, k))
	})

}

func RunMaxVowels(s string, k int) int {
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	res, cnt, l, r := 0, 0, 0, 0

	for r < len(s) {
		if _, ok := vowels[s[r]]; ok {
			cnt++
		}

		if r-l+1 > k {
			if _, ok := vowels[s[l]]; ok {
				cnt--
			}
			l++
		}
		res = max(res, cnt)
		r++
	}

	return res
}
