package array

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {
	t.Log(`
	Given an integer array nums, find the third maximum distinct number in nums. 
	If the third maximum does not exist, return the maximum number.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{10, 20, 30, 55}
		exp := 20
		assert.Equal(t, exp, RunThirdMax(in))
	})
	t.Run("Suite 1", func(t *testing.T) {
		in := []int{5, 5, 5, 51, 2, 4, 4, 5, 5, 5, 1}
		exp := 4
		assert.Equal(t, exp, RunThirdMax(in))
	})
}

func RunThirdMax(in []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64

	for i := 0; i < len(in); i++ {
		if in[i] == max1 || in[i] == max2 || in[i] == max3 {
			continue
		}

		if in[i] > max1 {
			max1, max2, max3 = in[i], max1, max2
		} else if in[i] > max2 {
			max2, max3 = in[i], max2
		} else if in[i] > max3 {
			max3 = in[i]
		}
	}
	return max3
}
