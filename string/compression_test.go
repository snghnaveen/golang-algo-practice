package string

import (
	"strconv"
	"strings"
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
	strArr := strings.Split(in, "")
	strArrLen := len(strArr)

	if strArrLen == 0 {
		return ""
	}

	out := strArr[0]
	count := 1

	for i := 1; i < strArrLen; i++ {

		current := strArr[i]
		previous := strArr[i-1]

		if current == previous {
			count++
		} else {
			out += strconv.Itoa(count)
			count = 1
			out += current
		}
	}

	out += strconv.Itoa(count)
	return out
}
