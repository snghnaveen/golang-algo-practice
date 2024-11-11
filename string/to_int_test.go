package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	t.Log(`
	Given a string str, the task is to convert the given string into 
	the number without using any inbuilt function.

	Input: str = "985632"
	Output: 985632
	`)

	t.Run("Suite 1", func(t *testing.T) {
		str := "12345"
		exp := 12345
		assert.Equal(t, exp, RunToInt(str))
	})
	t.Run("Suite 2", func(t *testing.T) {
		str := "100000"
		exp := 100000
		assert.Equal(t, exp, RunToInt(str))
	})
}

func RunToInt(s string) int {
	inp := []rune(s)
	var out int
	for i := range inp {
		if inp[i] >= '0' && inp[i] <= '9' {
			out = out*10 + int(inp[i]) - '0'
		}
	}
	return out
}
