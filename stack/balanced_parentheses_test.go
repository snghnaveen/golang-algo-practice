package stack

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	s "github.com/snghnaveen/golang-algo-practice/internal"
)

func TestBalancedParentheses(t *testing.T) {

	t.Log(`
		Given an expression string exp, write a program to examine whether the pairs and the orders of 
		“{“, “}”, “(“, “)”, “[“, “]” are correct in the given expression.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		exp := "[()]{}{[()()]()}"
		assert.True(t, RunTestBalancedParentheses(exp))
	})
	t.Run("Suite 2", func(t *testing.T) {
		exp := "[(])"
		assert.False(t, RunTestBalancedParentheses(exp))
	})
	t.Run("Suite 3", func(t *testing.T) {
		exp := "][()]"
		assert.False(t, RunTestBalancedParentheses(exp))
	})
}

func RunTestBalancedParentheses(str string) bool {
	stk := s.NewStringStack()

	for _, v := range strings.Split(str, "") {

		if v == "(" || v == "[" || v == "{" {
			stk.Push(v)
			continue
		}

		// If current character is not opening
		// bracket, then it must be closing. So stack
		// cannot be empty at this point.

		if stk.IsEmpty() {
			return false
		}

		switch v {
		case "]":
			if stk.Pop() != "[" {
				return false
			}

		case "}":
			if stk.Pop() != "{" {
				return false
			}
		case ")":
			if stk.Pop() != "(" {
				return false
			}
		}
	}

	// Check Empty Stack
	return stk.IsEmpty()
}
