package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubArraySum(t *testing.T) {
	t.Log(`
	Find Subarray with given sum.  (Non-negative Numbers)
	Input: arr[] = {1, 4, 20, 3, 10, 5}, sum = 33
	Output: Sum found between indexes 2 and 4
	Explanation: Sum of elements between indices 2 and 4 is 20 + 3 + 10 = 33
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 4, 20, 3, 10, 5}
		desiredSum := 33
		exp := []int{20, 3, 10}

		assert.Equal(t, exp, RunSubArraySum(desiredSum, in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{1, 4, 0, 0, 3, 10, 5}
		desiredSum := 7
		exp := []int{4, 0, 0, 3}

		assert.Equal(t, exp, RunSubArraySum(desiredSum, in))
	})

}

func RunSubArraySum(desiredSum int, in []int) []int {

	i := 0
	j := 0

	inLen := len(in)

	currentSum := 0
	for i <= inLen-1 && j <= inLen-1 {

		if i == 0 && j == 0 {
			currentSum = in[i]
			j++
		}

		if currentSum == desiredSum {
			break
		} else if currentSum > desiredSum {
			currentSum = currentSum - in[i]
			i++

		} else if currentSum < desiredSum {
			currentSum = currentSum + in[j]
			j++
		}

	}

	return in[i:j]
}
