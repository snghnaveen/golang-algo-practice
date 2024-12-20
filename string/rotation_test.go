package string

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckStringAreRotations(t *testing.T) {
	t.Log(`
	Check if given strings are rotations of each other or not
    Input: S1 = ABCD, S2 = CDAB
    Output: Strings are rotations of each other

    Input: S1 = ABCD, S2 = ACBD
    Output: Strings are not rotations of each other`)

	t.Run(`Is Rotation`, func(t *testing.T) {
		str1 := "ABCD"
		str2 := "CDAB"

		assert.True(t, RunCheckStringAreRotations(str1, str2))
	})

	t.Run(`Not Rotation`, func(t *testing.T) {
		str1 := "ABCD"
		str2 := "ACBD"
		assert.False(t, RunCheckStringAreRotations(str1, str2))
	})

}

func RunCheckStringAreRotations(str1, str2 string) bool {
	return strings.Contains(str1+str1, str2)
}
