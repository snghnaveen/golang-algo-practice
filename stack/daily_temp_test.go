package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	t.Log(`
	Given an array of integers temperatures represents the daily temperatures, 
	return an array answer such that answer[i] is the number of days you have to 
	wait after the ith day to get a warmer temperature. 
	If there is no future day for which this is possible, keep answer[i] == 0 instead.

	Example 1:
	Input: temperatures = [73,74,75,71,69,72,76,73]
	Output: [1,1,4,2,1,1,0,0]

	Example 2:
	Input: temperatures = [30,40,50,60]
	Output: [1,1,1,0]

	Example 3:
	Input: temperatures = [30,60,90]
	Output: [1,1,0]
	`)

	t.Run("Suite 1", func(t *testing.T) {
		temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
		exp := []int{1, 1, 4, 2, 1, 1, 0, 0}
		assert.Equal(t, exp, RunDailyTemperatures(temperatures))
	})
	t.Run("Suite 2", func(t *testing.T) {
		temperatures := []int{30, 40, 50, 60}
		exp := []int{1, 1, 1, 0}
		assert.Equal(t, exp, RunDailyTemperatures(temperatures))
	})
	t.Run("Suite 3", func(t *testing.T) {
		temperatures := []int{30, 60, 90}
		exp := []int{1, 1, 0}
		assert.Equal(t, exp, RunDailyTemperatures(temperatures))
	})
}

func RunDailyTemperatures(temps []int) []int {
	res := make([]int, len(temps))
	stack := make([]int, 0)

	for i, temp := range temps {
		for len(stack) > 0 && temp > temps[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return res
}
