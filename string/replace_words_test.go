package string

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceWords(t *testing.T) {
	t.Log(`
	In English, we have a concept called root, 
	which can be followed by some other word to form another longer word - 
	let's call this word derivative. 
	For example, when the root "help" is followed by the word "ful", 
	we can form a derivative "helpful".
	Given a dictionary consisting of many roots and a sentence consisting of words separated 
	by spaces, replace all the derivatives in the sentence with the root forming it. 
	If a derivative can be replaced by more than one root, replace it with the root that has 
	the shortest length.
	Return the sentence after the replacement.
	
	Example 1:
	Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
	Output: "the cat was rat by the bat"

	Example 2:
	Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
	Output: "a a b c"
	`)

	t.Run("Suite 1", func(t *testing.T) {
		dictonary := []string{"cat", "bat", "rat"}
		sentence := "the cattle was rattled by the battery"
		exp := "the cat was rat by the bat"
		assert.Equal(t, exp, RunReplaceWords(dictonary, sentence))
	})

	t.Run("Suite 2", func(t *testing.T) {
		dictonary := []string{"a", "b", "c"}
		sentence := "aadsfasf absbs bbab cadsfafs"
		exp := "a a b c"
		assert.Equal(t, exp, RunReplaceWords(dictonary, sentence))
	})
}

func RunReplaceWords(dictonary []string, sentence string) string {
	words := strings.Fields(sentence)

	dictMp := make(map[string]struct{})
	for _, v := range dictonary {
		dictMp[v] = struct{}{}
	}

	for idx, word := range words {
		for i := 1; i < len(word); i++ {
			prefix := word[:i]
			if _, ok := dictMp[prefix]; ok {
				words[idx] = prefix
				break
			}
		}
	}
	return strings.Join(words, " ")
}
