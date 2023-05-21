package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSegregate0And1(t *testing.T) {
	t.Log(`
	You are given an array of 0s and 1s in random order. 
	Segregate 0s on left side and 1s on right side of the array 
	[Basically you have to sort the array]. 
	Traverse array only once. 
	Input array   =  [0, 1, 0, 1, 0, 0, 1, 1, 1, 0] 
	Output array =  [0, 0, 0, 0, 0, 1, 1, 1, 1, 1] 
	`)

	in := []int{0, 1, 0, 1, 0, 0, 1, 1, 1, 0}
	exp := []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1}
	assert.Equal(t, exp, RunSegregate0And1(in))

}

func RunSegregate0And1(in []int) []int {
	l := 0
	r := len(in) - 1

	for l < r {
		if in[l] > in[r] {
			in[l], in[r] = in[r], in[l]
			l++
			r--
		} else if in[l] == 0 {
			l++
		} else if in[r] == 1 {
			r--
		}
	}
	return in
}
