package array

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeader(t *testing.T) {
	t.Log(`
	Write a program to print all the LEADERS in the array. 
	An element is a leader if it is greater than all the elements to its right side. 
	And the rightmost element is always a leader. 

	Input: arr[] = {16, 17, 4, 3, 5, 2}, 
	Output: 17, 5, 2

	Input: arr[] = {1, 2, 3, 4, 5, 2}, 
	Output: 5, 2
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{16, 17, 4, 3, 5, 2}
		exp := []int{17, 5, 2}
		assert.ElementsMatch(t, exp, RunLeader(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 2}
		exp := []int{5, 2}
		assert.ElementsMatch(t, exp, RunLeader(in))
	})
}

func RunLeader(in []int) []int {
	var out []int
	max := math.MinInt
	for i := len(in) - 1; i >= 0; i-- {
		if in[i] > max {
			max = in[i]
			out = append(out, max)
		}
	}
	return out
}
