package stack

import (
	"strconv"
	"testing"

	s "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestReverseNotation(t *testing.T) {
	t.Log(`
	You are given an array of strings tokens that represents 
	an arithmetic expression in a Reverse Polish Notation.
	Evaluate the expression. 
	Return an integer that represents the value of the expression.

	Note that:

	The valid operators are '+', '-', '*', and '/'.
	Each operand may be an integer or another expression.
	The division between two integers always truncates toward zero.
	There will not be any division by zero.
	The input represents a valid arithmetic expression in a reverse polish notation.
	The answer and all the intermediate calculations can be represented in a 32-bit integer.
	

	Example 1:
	Input: tokens = ["2","1","+","3","*"]
	Output: 9
	Explanation: ((2 + 1) * 3) = 9

	Example 2:
	Input: tokens = ["4","13","5","/","+"]
	Output: 6
	Explanation: (4 + (13 / 5)) = 6

	Example 3:
	Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
	Output: 22
	Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
	= ((10 * (6 / (12 * -11))) + 17) + 5
	= ((10 * (6 / -132)) + 17) + 5
	= ((10 * 0) + 17) + 5
	= (0 + 17) + 5
	= 17 + 5
	= 22
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []string{"2", "1", "+", "3", "*"}
		exp := 9
		assert.Equal(t, exp, RunReverseNotation(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []string{"4", "13", "5", "/", "+"}
		exp := 6
		assert.Equal(t, exp, RunReverseNotation(in))
	})
	t.Run("Suite 3", func(t *testing.T) {
		in := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
		exp := 22
		assert.Equal(t, exp, RunReverseNotation(in))
	})
}

func RunReverseNotation(tokens []string) int {
	stk := s.NewStringStack()

	for _, token := range tokens {

		switch token {
		case "*":
			b := toInt(stk.Pop())
			a := toInt(stk.Pop())
			token = strconv.Itoa(a * b)
		case "+":
			b := toInt(stk.Pop())
			a := toInt(stk.Pop())
			token = strconv.Itoa(a + b)
		case "-":
			b := toInt(stk.Pop())
			a := toInt(stk.Pop())
			token = strconv.Itoa(a - b)
		case "/":
			b := toInt(stk.Pop())
			a := toInt(stk.Pop())
			token = strconv.Itoa(a / b)
		}
		stk.Push(token)
	}

	return toInt(stk.Pop())
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("unable to convert to string to int")
	}

	return i
}
