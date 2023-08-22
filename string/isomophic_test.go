package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsomorphic(t *testing.T) {
	t.Log(`
	Two strings str1 and str2 are called isomorphic if there is a one-to-one mapping possible for every character of str1 to every character of str2. 
	And all occurrences of every character in ‘str1’ map to the same character in ‘str2’.

	Examples: 

	Input:  str1 = “aab”, str2 = “xxy”
	Output: True
	Explanation: ‘a’ is mapped to ‘x’ and ‘b’ is mapped to ‘y’.

	Input:  str1 = “aab”, str2 = “xyz”
	Output: False
	Explanation: One occurrence of ‘a’ in str1 has ‘x’ in str2 and other occurrence of ‘a’ has ‘y’.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		s1, s2 := "aab", "xxy"
		assert.True(t, RunIsomorphic(s1, s2))
	})

	t.Run("Suite 2", func(t *testing.T) {
		s1, s2 := "aab", "xyz"
		assert.False(t, RunIsomorphic(s1, s2))
	})

	t.Run("Suite 3", func(t *testing.T) {
		s1, s2 := "paper", "title"
		assert.True(t, RunIsomorphic(s1, s2))
	})

}

func RunIsomorphic(s1, s2 string) bool {
	s1Mp := make(map[byte]byte)

	for i := 0; i < len(s1); i++ {
		char1 := s1[i]
		char2 := s2[i]

		if v, ok := s1Mp[char1]; ok {
			if v != char2 {
				return false
			}
			continue
		} else {
			s1Mp[char1] = char2
		}
	}
	return true
}
