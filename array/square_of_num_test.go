package array

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgeSquareSum(t *testing.T) {
	t.Log(`
	Given a non-negative integer c, 
	decide whether there're two integers a and b such that a2 + b2 = c.

	Example 1:
	Input: c = 5
	Output: true
	Explanation: 1 * 1 + 2 * 2 = 5

	Example 2:
	Input: c = 3
	Output: false
	`)
	t.Run("Suite 1", func(t *testing.T) {
		c := 5
		assert.True(t, RunJudgeSquareSum(c))
	})
	t.Run("Suite 2", func(t *testing.T) {
		c := 3
		assert.False(t, RunJudgeSquareSum(c))
	})

}

func RunJudgeSquareSum(c int) bool {
	left := 0
	right := int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left--
		}
	}
	return false
}
