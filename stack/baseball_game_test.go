package stack

import (
	"strconv"
	"testing"

	s "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestBaseballGame(t *testing.T) {
	t.Log(`
	You are keeping the scores for a baseball game with strange rules. 
	At the beginning of the game, you start with an empty record.
	You are given a list of strings operations, 
	where operations[i] is the ith operation you must apply to the record and is one of the following:

	An integer x.
	Record a new score of x.
	'+'.
	Record a new score that is the sum of the previous two scores.
	'D'.
	Record a new score that is the double of the previous score.
	'C'.
	Invalidate the previous score, removing it from the record.
	Return the sum of all the scores on the record after applying all the operations.

	The test cases are generated such that the answer and all intermediate calculations 
	fit in a 32-bit integer and that all operations are valid.

	Example 1:
	Input: ops = ["5","2","C","D","+"]
	Output: 30
	Explanation:
	"5" - Add 5 to the record, record is now [5].
	"2" - Add 2 to the record, record is now [5, 2].
	"C" - Invalidate and remove the previous score, record is now [5].
	"D" - Add 2 * 5 = 10 to the record, record is now [5, 10].
	"+" - Add 5 + 10 = 15 to the record, record is now [5, 10, 15].
	The total sum is 5 + 10 + 15 = 30.

	Example 2:
	Input: ops = ["5","-2","4","C","D","9","+","+"]
	Output: 27
	Explanation:
	"5" - Add 5 to the record, record is now [5].
	"-2" - Add -2 to the record, record is now [5, -2].
	"4" - Add 4 to the record, record is now [5, -2, 4].
	"C" - Invalidate and remove the previous score, record is now [5, -2].
	"D" - Add 2 * -2 = -4 to the record, record is now [5, -2, -4].
	"9" - Add 9 to the record, record is now [5, -2, -4, 9].
	"+" - Add -4 + 9 = 5 to the record, record is now [5, -2, -4, 9, 5].
	"+" - Add 9 + 5 = 14 to the record, record is now [5, -2, -4, 9, 5, 14].
	The total sum is 5 + -2 + -4 + 9 + 5 + 14 = 27.

	Example 3:
	Input: ops = ["1","C"]
	Output: 0
	Explanation:
	"1" - Add 1 to the record, record is now [1].
	"C" - Invalidate and remove the previous score, record is now [].
	Since the record is empty, the total sum is 0.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []string{"5", "2", "C", "D", "+"}
		exp := 30
		assert.Equal(t, exp, RunBaseballGame(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
		exp := 27
		assert.Equal(t, exp, RunBaseballGame(in))
	})
	t.Run("Suite 3", func(t *testing.T) {
		in := []string{"1", "C"}
		exp := 0
		assert.Equal(t, exp, RunBaseballGame(in))
	})

}

func RunBaseballGame(operations []string) int {
	stk := s.NewStringStack()
	for _, o := range operations {
		switch o {
		case "D":
			num := con(stk.Peek())
			stk.Push(strconv.Itoa(num * 2))
		case "C":
			stk.Pop()
		case "+":
			a := con(stk.Pop())
			b := con(stk.Pop())
			stk.Push(strconv.Itoa(b))
			stk.Push(strconv.Itoa(a))

			stk.Push(strconv.Itoa(a + b))
		default:
			stk.Push(o)
		}
	}

	res := 0
	for !stk.IsEmpty() {
		res = res + con(stk.Pop())
	}
	return res
}

func con(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic("unable to parse")
	}
	return v

}
