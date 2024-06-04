package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	t.Log(`
	Given a string s which consists of lowercase or uppercase letters, 
	return the length of the longest 
	palindrome that can be built with those letters.

	Letters are case sensitive, for example, "Aa" is not considered a palindrome.

	Example 1:
	Input: s = "abccccdd"
	Output: 7
	Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.

	Example 2:
	Input: s = "a"
	Output: 1
	Explanation: The longest palindrome that can be built is "a", whose length is 1.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		s := "abccccdd"
		exp := 7
		assert.Equal(t, exp, RunLongestPlaindrome(s))
	})

	t.Run("Suite 2", func(t *testing.T) {
		s := "a"
		exp := 1
		assert.Equal(t, exp, RunLongestPlaindrome(s))
	})
	t.Run("Suite 3", func(t *testing.T) {
		s := "aaabaaa"
		exp := 7
		assert.Equal(t, exp, RunLongestPlaindrome(s))
	})
}

func RunLongestPlaindrome(s string) int {
	count := 0
	mpCnt := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		mpCnt[s[i]]++
	}

	for _, v := range mpCnt {
		count = count + ((v / 2) * 2)
	}

	if count < len(s) {
		count++
	}

	return count
}
