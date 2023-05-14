package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSequence(t *testing.T) {
	t.Log(`
	Given an unsorted array of integers nums, 
	return the length of the longest consecutive elements sequence.
	You must write an algorithm that runs in O(n) time.

	Example 1:
	Input: nums = [100,4,200,1,3,2]
	Output: 4
	Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. 
	Therefore its length is 4.

	Example 2:
	Input: nums = [0,3,7,2,5,8,4,6,0,1]
	Output: 9
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{100, 4, 200, 1, 3, 2}
		exp := 4
		assert.Equal(t, RunLongestSequence(in), exp)
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
		exp := 9
		assert.Equal(t, RunLongestSequence(in), exp)
	})
}

func RunLongestSequence(in []int) int {
	mp := make(map[int]struct{})

	for i := 0; i < len(in); i++ {
		mp[in[i]] = struct{}{}
	}

	var res int

	for i := 0; i < len(in); i++ {
		num := in[i]

		// ignore if current have left sequence
		if _, ok := mp[num-1]; ok {
			continue
		}

		seq := 1
		tmp := num + 1

		for {
			if _, ok := mp[tmp]; ok {
				seq++
				tmp++
			} else {
				break
			}
		}

		if seq > res {
			res = seq
		}

	}

	return res
}
