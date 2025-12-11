package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTippingRain(t *testing.T) {
	t.Log(`
	Given n non-negative integers representing an elevation map where the width of each bar is 1, 
	compute how much water it can trap after raining.

	Example 1:
	Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
	Output: 6
	Explanation: The above elevation map (black section) is represented by array 
	[0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

	Example 2:
	Input: height = [4,2,0,3,2,5]
	Output: 9
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		exp := 6
		assert.Equal(t, exp, RunTippingRain(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{4, 2, 0, 3, 2, 5}
		exp := 9
		assert.Equal(t, exp, RunTippingRain(in))
	})

}

// todo revisit
func RunTippingRain(in []int) int {
	l, r := 0, len(in)-1
	lMax, rMax := in[l], in[r]
	res := 0
	for l < r {
		if lMax < rMax {
			l++
			lMax = max(lMax, in[l])
			res = res + lMax - in[l]
		} else {
			r--
			rMax = max(rMax, in[r])
			res = res + rMax - in[r]
		}
	}
	return res
}
