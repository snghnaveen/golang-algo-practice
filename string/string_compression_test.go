package string

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	t.Log(`https://leetcode.com/problems/string-compression/description/`)

	t.Log(`
	Given an array of characters chars, compress it using the following algorithm:
	Begin with an empty string s. For each group of consecutive repeating characters in chars:
	If the group's length is 1, append the character to s.
	Otherwise, append the character followed by the group's length.
	The compressed string s should not be returned separately, but instead, be stored in the input character array chars. 
	Note that group lengths that are 10 or longer will be split into multiple characters in chars.
	After you are done modifying the input array, return the new length of the array.
	You must write an algorithm that uses only constant extra space.
	Note: The characters in the array beyond the returned length do not matter and should be ignored.
	
	Example 1:
	Input: chars = ["a","a","b","b","c","c","c"]
	Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
	Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".
	
	Example 2:
	Input: chars = ["a"]
	Output: Return 1, and the first character of the input array should be: ["a"]
	Explanation: The only group is "a", which remains uncompressed since it's a single character.
	
	Example 3:
	Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
	Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
	Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".
	`)

	t.Run("Suite 1", func(t *testing.T) {
		exp := 6
		expVal := "a2b2c3"
		in := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
		assert.Equal(t, exp, RunCompress(in))
		assert.Equal(t, string(in[:exp]), expVal)
	})
	t.Run("Suite 2", func(t *testing.T) {
		exp := 1
		expVal := "a"
		in := []byte{'a'}
		assert.Equal(t, exp, RunCompress(in))
		assert.Equal(t, string(in[:exp]), expVal)
	})

	t.Run("Suite 3", func(t *testing.T) {
		exp := 4
		expVal := "ab12"
		in := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
		assert.Equal(t, exp, RunCompress(in))
		assert.Equal(t, string(in[:exp]), expVal)
	})
}

func RunCompress(chars []byte) int {
	idx := 0
	i := 0
	n := len(chars)
	for i < n {
		curr := chars[i]
		count := 0
		for i < n && chars[i] == curr {
			i++
			count++
		}

		if count == 1 {
			chars[idx] = curr
			idx++
		} else {
			chars[idx] = curr
			idx++

			for _, digit := range strconv.Itoa(count) {
				chars[idx] = byte(digit)
				idx++
			}
		}
	}
	return idx
}
