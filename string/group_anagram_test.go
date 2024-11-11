package string

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagram(t *testing.T) {

	t.Log(`
	Given an array of strings strs, group the anagrams together. You can return the answer in any order.
	An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
	typically using all the original letters exactly once.
	
	Example 1:
	Input: strs = ["eat","tea","tan","ate","nat","bat"]
	Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
	Example 2:
	
	Input: strs = [""]
	Output: [[""]]
	Example 3:
	
	Input: strs = ["a"]
	Output: [["a"]]
	`)

	t.Run("Suite 1", func(t *testing.T) {
		assert.ElementsMatch(t, RunGroupAnagram([]string{"eat", "tea", "tan", "ate", "nat", "bat"}),
			[][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			})
	})

	t.Run("Suite 2", func(t *testing.T) {
		assert.ElementsMatch(t, RunGroupAnagram([]string{"a"}),
			[][]string{
				{"a"},
			})
	})
	t.Run("Suite 2", func(t *testing.T) {
		assert.ElementsMatch(t, RunGroupAnagram([]string{""}),
			[][]string{
				{""},
			})
	})

}

func RunGroupAnagram(in []string) [][]string {
	var out [][]string

	mp := make(map[string][]string)

	n := len(in)
	for i := 0; i < n; i++ {
		// sort string
		sArr := strings.Split(in[i], "")
		sort.Strings(sArr)
		s := strings.Join(sArr, "")

		mp[s] = append(mp[s], in[i])

	}

	for _, val := range mp {
		out = append(out, val)
	}

	return out
}
