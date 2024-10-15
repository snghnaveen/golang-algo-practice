package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagrams(t *testing.T) {
	t.Log(`
	Given two strings s and p, return an array of all the start indices of p's 
	anagrams in s. You may return the answer in any order.
	Example 1:
	Input: s = "cbaebabacd", p = "abc"
	Output: [0,6]
	Explanation:
	The substring with start index = 0 is "cba", which is an anagram of "abc".
	The substring with start index = 6 is "bac", which is an anagram of "abc".

	Example 2:
	Input: s = "abab", p = "ab"
	Output: [0,1,2]
	Explanation:
	The substring with start index = 0 is "ab", which is an anagram of "ab".
	The substring with start index = 1 is "ba", which is an anagram of "ab".
	The substring with start index = 2 is "ab", which is an anagram of "ab".
	`)

	t.Run("Suite 1", func(t *testing.T) {
		s := "cbaebabacd"
		p := "abc"
		exp := []int{0, 6}
		assert.Equal(t, exp, RunFindAnagrams(s, p))
	})
	t.Run("Suite 2", func(t *testing.T) {
		s := "abab"
		p := "ab"
		exp := []int{0, 1, 2}
		assert.Equal(t, exp, RunFindAnagrams(s, p))
	})
}

func RunFindAnagrams(s, p string) []int {
	var res []int

	hasAllChar := func(sCount, pCount map[rune]int) bool {
		for k, v := range pCount {
			if v != sCount[k] {
				return false
			}
		}

		return true
	}

	sCount := make(map[rune]int)
	pCount := make(map[rune]int)

	for _, c := range p {
		pCount[c]++
	}

	l, r := 0, 0
	for _, c := range s {
		sCount[c]++

		if r-l+1 < len(p) {
			r++
			continue
		}

		if hasAllChar(sCount, pCount) {
			res = append(res, l)
		}
		sCount[rune(s[l])]--
		l++
		r++
	}

	return res
}
