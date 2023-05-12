package string

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArePermutation(t *testing.T) {
	t.Log(`
	Write a function to check whether two given strings are Permutation of each other or not.
	A Permutation of a string is another string that contains same characters,
	 only the order of characters can be different. 
	For example, “abcd” and “dabc” are Permutation of each other.
	`)

	t.Run("When string are in Permutation", func(t *testing.T) {

		str1 := "abcd"
		str2 := "dabc"
		assert.Equal(t, true, RunArePermutation(str1, str2))
	})

	t.Run("When string are not in Permutation", func(t *testing.T) {
		str1 := "abcd"
		str2 := "qwer"
		assert.Equal(t, false, RunArePermutation(str1, str2))
	})

}

func RunArePermutation(str1, str2 string) bool {

	str1Arr := strings.Split(str1, "")
	str2Arr := strings.Split(str2, "")

	sort.Strings(str1Arr)
	sort.Strings(str2Arr)

	n1 := len(str1Arr)
	n2 := len(str2Arr)

	if n2 != n1 {
		return false
	}

	for i := 0; i < n1; i++ {
		if str1Arr[i] != str2Arr[i] {
			return false
		}
	}

	return true
}
