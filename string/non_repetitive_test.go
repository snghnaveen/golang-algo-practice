package string

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstNonRepetitive(t *testing.T) {
	t.Log(`
	For a given string, print first non-repetitive letter from string
	Input: “geeksforgeeks”
	Output: f
	Explanation: As ‘f’ is first character in the string which does not repeat.
	`)

	assert.Equal(t, RunFirstNonRepetitive("geeksforgeeks"), "f")
	assert.Equal(t, RunFirstNonRepetitive("algorithm"), "a")
	assert.Equal(t, RunFirstNonRepetitive("bbaa"), "")

}

func RunFirstNonRepetitive(s string) string {
	mp := make(map[string]int)
	s = strings.ToLower(s)
	sArr := strings.Split(s, "")

	for i := 0; i < len(sArr); i++ {
		if v, ok := mp[sArr[i]]; ok {
			mp[sArr[i]] = v + 1
		} else {
			mp[sArr[i]] = 1
		}
	}

	for i := 0; i < len(sArr); i++ {
		if v, ok := mp[sArr[i]]; ok && v == 1 {
			return sArr[i]
		}
	}
	return ""
}
