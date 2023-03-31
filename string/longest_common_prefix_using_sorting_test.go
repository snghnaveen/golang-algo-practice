package string

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefixUsingSorting(t *testing.T) {
	t.Log(`
	Given a set of strings, find the longest common prefix.
	Examples:

	Input: {“geeksforgeeks”, “geeks”, “geek”, “geezer”}
	Output: “gee”

	Input: {“apple”, “ape”, “april”}
	Output: “ap”
	`)

	t.Run("Suite 1", func(t *testing.T) {
		strArr := []string{"geeksforgeeks", "geeks", "geek", "geezer"}

		assert.Equal(t, `gee`,
			RunTestLongestCommonPrefixUsingSorting(strArr))
	})

	t.Run("Suite 2", func(t *testing.T) {
		strArr := []string{"apple", "ape", "april"}

		assert.Equal(t, `ap`,
			RunTestLongestCommonPrefixUsingSorting(strArr))
	})

}

func RunTestLongestCommonPrefixUsingSorting(in []string) string {
	if len(in) == 0 {
		return ""
	}

	if len(in) == 1 {
		return in[0]
	}

	sort.Strings(in)

	end := len(in[0])
	out := ""

	firstStringElement := in[0]
	lastStringElement := in[len(in)-1]

	for x := 0; x < end; x++ {
		if string(firstStringElement[x]) != string(lastStringElement[x]) {
			break
		}
		out = out + string(firstStringElement[x])
	}

	return out
}
