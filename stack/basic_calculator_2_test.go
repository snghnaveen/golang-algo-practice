package stack

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestBasicCalculator2(t *testing.T) {
	t.Log(`Link: https://leetcode.com/problems/basic-calculator-ii/description/`)

	t.Log(`
	Given a string s which represents an expression, 
	evaluate this expression and return its value. 
	The integer division should truncate toward zero.
	You may assume that the given expression is always valid. 
	All intermediate results will be in the range of [-231, 231 - 1].
	Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().

	Example 1:
	Input: s = "3+2*2"
	Output: 7

	Example 2:
	Input: s = " 3/2 "
	Output: 1

	Example 3:
	Input: s = " 3+5 / 2 "
	Output: 5	
	`)

	t.Run("Suite 1", func(t *testing.T) {
		s := "3+2*2"
		exp := 7
		assert.Equal(t, exp, RunBasicCalculator2(s))
	})
	t.Run("Suite 2", func(t *testing.T) {
		s := " 3/2 "
		exp := 1
		assert.Equal(t, exp, RunBasicCalculator2(s))
	})
	t.Run("Suite 3", func(t *testing.T) {
		s := " 3+5 / 2 "
		exp := 5
		assert.Equal(t, exp, RunBasicCalculator2(s))
	})

}

func RunBasicCalculator2(s string) int {
	stack := make([]int, 0)
	res, num := 0, 0
	sign := '+'

	for i, ch := range s {
		// also can be written as ch >= '0' && ch <='9'
		if unicode.IsDigit(ch) {
			num = 10*num + int(ch-'0')
		}

		// i == len(s)-1 to process the last number in the input string
		if !unicode.IsDigit(ch) && !unicode.IsSpace(ch) || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '/':
				stack[len(stack)-1] = stack[len(stack)-1] / num
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * num
			}

			num = 0
			sign = ch

		}

	}

	for _, v := range stack {
		res = res + v
	}
	return res
}
