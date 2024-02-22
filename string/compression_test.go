package string

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressionMechanism(t *testing.T) {
	t.Log(`
	Implement a method to perform basic string compression using the counts 
	of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the 
	"compressed" string would not become smaller than the original string, your method should return 
	the original string. You can assume the string has only uppercase and lowercase letters (a - z). 
	`)
	t.Run("Suite 1", func(t *testing.T) {
		in := "aabcccccaaax"
		exp := "a2b1c5a3x1"
		assert.Equal(t, exp, RunGetCompression(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := "abcccccaaax"
		exp := "a1b1c5a3x1"
		assert.Equal(t, exp, RunGetCompression(in))
	})
	t.Run("Suite 3", func(t *testing.T) {
		in := "abcccccaaaxm"
		exp := "a1b1c5a3x1m1"
		assert.Equal(t, exp, RunGetCompression(in))
	})
}

func RunGetCompression(in string) string {
	compressed := ""
	count := 1

	for i := 0; i < len(in); i++ {
		if i+1 < len(in) && in[i] == in[i+1] {
			count++
		} else {
			compressed = compressed + string(in[i]) + strconv.Itoa(count)
			count = 1
		}
	}
	return compressed
}
