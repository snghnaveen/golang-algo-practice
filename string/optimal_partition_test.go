package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptimalPartition(t *testing.T) {
	t.Log(`
	Given a string s, partition the string into one or more substrings such that the characters 
	in each substring are unique. 
	That is, no letter appears in a single substring more than once.
	Return the minimum number of substrings in such a partition.
	Note that each character should belong to exactly one substring in a partition.

	Example 1:
	Input: s = "abacaba"
	Output: 4
	Explanation:
	Two possible partitions are ("a","ba","cab","a") and ("ab","a","ca","ba").
	It can be shown that 4 is the minimum number of substrings needed.

	Example 2:
	Input: s = "ssssss"
	Output: 6
	Explanation:
	The only valid partition is ("s","s","s","s","s","s").
	`)

	t.Run("Suite 1", func(t *testing.T) {
		str := "abacaba"
		exp := 4
		assert.Equal(t, exp, RunOptimalPartition(str))
	})
	t.Run("Suite 2", func(t *testing.T) {
		str := "ssssss"
		exp := 6
		assert.Equal(t, exp, RunOptimalPartition(str))
	})
}

func RunOptimalPartition(s string) int {
	visited := make(map[string]struct{})
	res := 1
	for _, c := range s {
		if _, ok := visited[string(c)]; ok {
			res++
			clear(visited)
		}
		visited[string(c)] = struct{}{}
	}
	return res
}
