package array

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestNum(t *testing.T) {
	t.Log(`
	Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.
	Since the result may be very large, so you need to return a string instead of an integer.

	Example 1:
	Input: nums = [10,2]
	Output: "210"
	
	Example 2:
	Input: nums = [3,30,34,5,9]
	Output: "9534330"
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{10, 2}
		exp := "210"
		assert.Equal(t, exp, RunLargestNum(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{3, 30, 34, 5, 9}
		exp := "9534330"
		assert.Equal(t, exp, RunLargestNum(in))
	})

	t.Run("Suite 3", func(t *testing.T) {
		in := []int{0, 0, 0, 0}
		exp := "0"
		assert.Equal(t, exp, RunLargestNum(in))
	})
}

func RunLargestNum(in []int) string {
	numsStr := make([]string, len(in))
	for i, n := range in {
		numsStr[i] = strconv.Itoa(n)
	}
	sort.Slice(numsStr, func(i, j int) bool {
		s1 := numsStr[i]
		s2 := numsStr[j]

		if len(s1) == len(s2) {
			return s1 > s2
		}
		return s1+s2 > s2+s1
	})

	res := strings.TrimLeft(strings.Join(numsStr, ""), "0")
	if res == "" {
		res = "0"
	}
	return res
}
