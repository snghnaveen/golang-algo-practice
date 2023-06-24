package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequentElements(t *testing.T) {
	t.Log(`
	Given an integer array nums and an integer k, return the k most frequent elements. 
	You may return the answer in any order.
	
	Example 1:
	Input: nums = [1,1,1,2,2,3], k = 2
	Output: [1,2]
	
	Example 2:
	Input: nums = [1], k = 1
	Output: [1]
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 1, 1, 2, 2, 3}
		k := 2
		exp := []int{1, 2}

		assert.ElementsMatch(t, exp, RunTopKFreqElements(in, k))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{1}
		k := 1
		exp := []int{1}

		assert.ElementsMatch(t, exp, RunTopKFreqElements(in, k))
	})

}

func RunTopKFreqElements(in []int, k int) []int {
	var out []int

	mp := make(map[int]int)
	for i := 0; i < len(in); i++ {
		if v, ok := mp[in[i]]; ok {
			mp[in[i]] = v + 1
		} else {
			mp[in[i]] = 1
		}
	}

	freq := make([][]int, len(in))

	for k, v := range mp {
		freq[v-1] = append(freq[v-1], k)
	}

	counter := 0
	for i := len(freq) - 1; i >= 0; i-- {
		if counter != k && freq[i] != nil {
			out = append(out, freq[i]...)
			counter++
		}
	}

	return out
}
