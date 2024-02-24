package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	t.Log(`
	Given n pairs of parentheses, 
	write a function to generate all combinations of well-formed parentheses.

	Example 1:
	Input: n = 3
	Output: ["((()))","(()())","(())()","()(())","()()()"]

	Example 2:
	Input: n = 1
	Output: ["()"]
	`)

	t.Run("Suite 1", func(t *testing.T) {
		n := 3
		exp := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
		assert.ElementsMatch(t, exp, RunGenerateParenthesis(n))
	})
	t.Run("Suite 2", func(t *testing.T) {
		n := 1
		exp := []string{"()"}
		assert.ElementsMatch(t, exp, RunGenerateParenthesis(n))
	})
}

func RunGenerateParenthesis(n int) []string {
	ans := make([]string, 0)
	current := make([]byte, n*2)
	openN, closeN := 0, 0
	generateParenthesisBackTrack(n, openN, closeN, current, &ans)
	return ans
}

func generateParenthesisBackTrack(n, openN, closeN int, current []byte, ans *[]string) {
	if openN == closeN && openN == n {
		*ans = append(*ans, string(current))
	}

	if openN < n {
		current[openN+closeN] = '('
		generateParenthesisBackTrack(n, openN+1, closeN, current, ans)
	}

	if closeN < openN {
		current[openN+closeN] = ')'
		generateParenthesisBackTrack(n, openN, closeN+1, current, ans)
	}
}
