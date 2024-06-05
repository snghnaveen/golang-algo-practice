package string

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonCharacters(t *testing.T) {
	t.Log(`
	Given a string array words, 
	return an array of all characters that show up in all strings within the words 
	(including duplicates). You may return the answer in any order.

	Example 1:
	Input: words = ["bella","label","roller"]
	Output: ["e","l","l"]

	Example 2:
	Input: words = ["cool","lock","cook"]
	Output: ["c","o"]
	`)

	t.Run("Suite 1", func(t *testing.T) {
		words := []string{"bella", "label", "roller"}
		exp := []string{"e", "l", "l"}
		assert.ElementsMatch(t, exp, RunCommonCharacters(words))
	})
	t.Run("Suite 2", func(t *testing.T) {
		words := []string{"cool", "lock", "cook"}
		exp := []string{"c", "o"}
		assert.ElementsMatch(t, exp, RunCommonCharacters(words))
	})
}

func RunCommonCharacters(words []string) []string {
	minCount := make([]int, 26)

	for i := 0; i < 26; i++ {
		minCount[i] = math.MaxInt
	}

	for _, word := range words {
		currChCount := make([]int, 26)

		for _, ch := range word {
			currChCount[ch-'a']++
		}

		for i := 0; i < 26; i++ {
			minCount[i] = min(minCount[i], currChCount[i])
		}
	}

	var res []string
	for i := 0; i < 26; i++ {
		for minCount[i] > 0 {
			res = append(res, string(rune(i+'a')))
			minCount[i]--
		}
	}
	return res
}
