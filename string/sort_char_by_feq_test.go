package string

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortCharactersByFrequency(t *testing.T) {
	t.Log(`
	Given a string s, sort it in decreasing order based on the frequency of the characters. 
	The frequency of a character is the number of times it appears in the string.
	Return the sorted string. If there are multiple answers, return any of them.

	Example 1:
	Input: s = "tree"
	Output: "eert"
	Explanation: 'e' appears twice while 'r' and 't' both appear once.
	So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := "tree"
		exp := []string{"eert", "eetr"}
		assert.Contains(t, exp, RunSortCharactersByFrequency(in))
	})
}

func RunSortCharactersByFrequency(s string) string {
	// The choice of 255 as the length of the freq slice covers
	// all possible ASCII characters (including the extended set).
	freq := make([]int, 255)

	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		if freq[b[i]] == freq[b[j]] {
			// Comparing ASCII values ensures a consistent and deterministic order
			// for characters with the same frequency.
			// For example, if both 'a' and 'b' appear twice in the input string,
			// we want to ensure that 'b' comes before 'a' in the sorted result.
			return b[i] > b[j]
		}
		return freq[b[i]] > freq[b[j]]
	})
	return string(b)
}
