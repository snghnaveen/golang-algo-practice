package string

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOneStep(t *testing.T) {
	t.Log(`
	There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character. 
	Given two strings, write a function to check if they are 
	one edit (or zero edits) away. 
	EXAMPLE 
	pale, ple -> true 
	pales, pale -> true 
	pale, bale -> true 
	pale, bae -> false
`)
	t.Run("Suite 1", func(t *testing.T) {
		str1 := "pale"
		str2 := "ple"
		assert.True(t, IsOneEditAway(str1, str2))
	})
	t.Run("Suite 2", func(t *testing.T) {
		str1 := "pales"
		str2 := "pale"
		assert.True(t, IsOneEditAway(str1, str2))
	})
	t.Run("Suite 3", func(t *testing.T) {
		str1 := "pale"
		str2 := "bale"
		assert.True(t, IsOneEditAway(str1, str2))
	})
	t.Run("Suite 4", func(t *testing.T) {
		str1 := "pale"
		str2 := "bae"
		assert.False(t, IsOneEditAway(str1, str2))
	})
	t.Run("Suite 5", func(t *testing.T) {
		str1 := "abc"
		str2 := "abc"
		// false, because they are equal
		assert.False(t, IsOneEditAway(str1, str2))
	})
}

func IsOneEditAway(firstStr, secondStr string) bool {
	first := strings.Split(firstStr, "")
	second := strings.Split(secondStr, "")

	firstLen := len(first)
	secondLen := len(second)

	if firstLen == secondLen {
		return IsOneEditReplace(firstStr, secondStr)
	} else if firstLen-1 == secondLen {
		return IsOneEditInsertOrDelete(firstStr, secondStr)
	} else if firstLen+1 == secondLen {
		return IsOneEditInsertOrDelete(secondStr, firstStr)
	}
	return true
}

func IsOneEditReplace(firstStr, secondStr string) bool {
	first := strings.Split(firstStr, "")
	second := strings.Split(secondStr, "")

	firstLen := len(first)

	isOneEditReplace := false
	for i := 0; i < firstLen; i++ {
		if first[i] != second[i] {
			if isOneEditReplace == true {
				return false
			}
			isOneEditReplace = true
		}
	}

	return isOneEditReplace
}

func IsOneEditInsertOrDelete(firstStr, secondStr string) bool {
	first := strings.Split(firstStr, "")
	second := strings.Split(secondStr, "")

	firstLen := len(first)
	secondLen := len(second)

	i := 0
	j := 0

	for i < firstLen && j < secondLen {
		if first[i] != second[j] {
			if i != j {
				return false
			}
			i++
		} else {
			i++
			j++
		}
	}
	return true

}
