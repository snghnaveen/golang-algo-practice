package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstIdxFistOccurance(t *testing.T) {
	t.Log(`
	Given two strings needle and haystack, return the index of the first occurrence of needle in haystack,
	or -1 if needle is not part of haystack.

	Example 1:
	Input: haystack = "sadbutsad", needle = "sad"
	Output: 0
	Explanation: "sad" occurs at index 0 and 6.
	The first occurrence is at index 0, so we return 0.

	Example 2:
	Input: haystack = "leetcode", needle = "leeto"
	Output: -1
	Explanation: "leeto" did not occur in "leetcode", so we return -1.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		haystack := "sadbutsad"
		needle := "sad"
		exp := 0
		assert.Equal(t, exp, RunFirstIdxFistOccurance(haystack, needle))
	})

	t.Run("Suite 2", func(t *testing.T) {
		haystack := "leetcode"
		needle := "leeto"
		exp := -1
		assert.Equal(t, exp, RunFirstIdxFistOccurance(haystack, needle))
	})
	t.Run("Suite 3", func(t *testing.T) {
		haystack := "a"
		needle := "a"
		exp := 0
		assert.Equal(t, exp, RunFirstIdxFistOccurance(haystack, needle))
	})
	t.Run("Suite 3", func(t *testing.T) {

	})
}

func RunFirstIdxFistOccurance(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
