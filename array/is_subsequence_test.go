package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfSubsequence(t *testing.T) {
	t.Log(`
	Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
	A subsequence of a string is a new string that is formed from the original 
	string by deleting some (can be none) of the characters without disturbing the 
	relative positions of the remaining characters. 
	(i.e., "ace" is a subsequence of "abcde" while "aec" is not).

	Example 1:
	Input: s = "abc", t = "ahbgdc"
	Output: true

	Example 2:
	Input: s = "axc", t = "ahbgdc"
	Output: false

	What is subsequence and subString?
    Subsequence: a sequence that appears in the same relative order, but not necessarily contiguous.
    SubString: a contiguous sequence of symbols that appears in the same relative order as the original string.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		str := "ahbgdc"
		subStr := "abc"
		assert.True(t, RunIsSubsequence(str, subStr))
	})

	t.Run("Suite 2", func(t *testing.T) {
		str := "axc"
		subStr := "ahbgdc"
		assert.False(t, RunIsSubsequence(str, subStr))
	})

}

func RunIsSubsequence(str, subStr string) bool {
	strLen := len(str)
	subStrLen := len(subStr)

	if subStrLen > strLen {
		return false
	}

	p1, p2 := 0, 0
	for i := 0; i < strLen; i++ {
		if str[p1] == subStr[p2] {
			p1++
			p2++
		} else {
			p1++
		}
	}

	return p2 == subStrLen
}
