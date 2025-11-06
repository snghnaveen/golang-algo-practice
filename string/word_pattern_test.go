package string

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordPattern(t *testing.T) {
	t.Log(`Link: https://leetcode.com/problems/word-pattern/`)

	t.Log(`
	Given a pattern and a string s, find if s follows the same pattern.
	Here follow means a full match, such that there is a bijection between 
	a letter in pattern and a non-empty word in s. 
	Specifically:
	Each letter in pattern maps to exactly one unique word in s.
	Each unique word in s maps to exactly one letter in pattern.
	No two letters map to the same word, and no two words map to the same letter.
	
	Example 1:
	Input: pattern = "abba", s = "dog cat cat dog"
	Output: true
	Explanation:
	The bijection can be established as:
	'a' maps to "dog".
	'b' maps to "cat".
	
	Example 2:
	Input: pattern = "abba", s = "dog cat cat fish"
	Output: false

	Example 3:
	Input: pattern = "aaaa", s = "dog cat cat dog"
	Output: false
	`)

	t.Run("Suite 1", func(t *testing.T) {
		pattern := "abba"
		s := "dog cat cat dog"
		assert.True(t, RunWordPattern(pattern, s))
	})

	t.Run("Suite 2", func(t *testing.T) {
		pattern := "abba"
		s := "dog cat cat fish"
		assert.False(t, RunWordPattern(pattern, s))
	})
	t.Run("Suite 3", func(t *testing.T) {
		pattern := "aaaa"
		s := "dog cat cat dog"
		assert.False(t, RunWordPattern(pattern, s))
	})
}

func RunWordPattern(pattern string, s string) bool {
	pMap := make(map[rune]string)
	sMap := make(map[string]rune)

	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	for i, ch := range pattern {
		word := words[i]

		mappedWord, pExists := pMap[ch]
		mappedChar, sExists := sMap[word]

		if pExists && mappedWord != word {
			return false
		}
		if sExists && mappedChar != ch {
			return false
		}

		pMap[ch] = word
		sMap[word] = ch
	}

	return true
}
