package string

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidAnagram(t *testing.T) {
	t.Log(`
	An anagram of a string is another string that contains the same characters, 
	only the order of characters can be different.
	 For example, “abcd” and “dabc” are an anagram of each other.
	`)
	t.Run("Suite 1", func(t *testing.T) {
		assert.True(t, RunIsValidAnagram("abcd", "dabc"))
	})
	t.Run("Suite 2", func(t *testing.T) {
		assert.False(t, RunIsValidAnagram("abxd", "dabc"))
	})

}

func RunIsValidAnagram(str1 string, str2 string) bool {
	str1Arr := strings.Split(str1, "")
	str2Arr := strings.Split(str2, "")

	str1Len := len(str1Arr)
	str2Len := len(str2Arr)

	if str1Len != str2Len {
		return false
	}

	mp1 := make(map[string]int)

	for i := 0; i < str1Len; i++ {
		if v, ok := mp1[str1Arr[i]]; ok {
			mp1[str1Arr[i]] = v + 1
		}
		mp1[str1Arr[i]] = 1
	}

	for i := 0; i < str2Len; i++ {
		if v, ok := mp1[str2Arr[i]]; ok {
			if v == 1 {
				delete(mp1, str2Arr[i])
			} else {
				mp1[str2Arr[i]] = v - 1
			}
		} else {
			// char does not exists
			return false
		}
	}
	return len(mp1) == 0
}
