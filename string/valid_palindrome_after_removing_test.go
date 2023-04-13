package string

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPalindromeAfterRemoving1Char(t *testing.T) {
	t.Log(`
	Find extract characters in a string that can be removed to make it as a palindrome.
	`)

	var tests = []struct {
		in   string
		want string
	}{
		{"cabdac", "b"},
		{"aba", ""},
		{"abd", ""},
		{"aaaab", "b"},
	}

	for _, tt := range tests {
		assert.Equal(t, RunTestIsValidPalindromeAfter1Char(tt.in), tt.want)
	}

}

func RunTestIsValidPalindromeAfter1Char(s string) string {
	strArr := strings.Split(s, "")

	left := 0
	right := len(strArr) - 1

	for left < right {
		if strArr[left] != strArr[right] {

			if isPalindrome(left+1, right, strArr) {
				return strArr[left]
			}

			if isPalindrome(left, right-1, strArr) {
				return strArr[right]
			}
		}
		left++
		right--
	}

	return ""
}

func isPalindrome(left, right int, strArr []string) bool {
	for left < right {
		if strArr[left] != strArr[right] {
			return false
		}

		left++
		right--
	}

	return true
}
