package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"

	s "github.com/snghnaveen/golang-algo-practice/internal"
)

func TestAsteroidCollision(t *testing.T) {
	t.Log(`
	We are given an array asteroids of integers representing asteroids in a row.
	For each asteroid, the absolute value represents its size, 
	and the sign represents its direction 
	(positive meaning right, negative meaning left). 
	Each asteroid moves at the same speed.

	Find out the state of the asteroids after all collisions. 
	If two asteroids meet, the smaller one will explode. 
	If both are the same size, both will explode. 
	Two asteroids moving in the same direction will never meet.

	

	Example 1:
	Input: asteroids = [5,10,-5]
	Output: [5,10]
	Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.

	Example 2:
	Input: asteroids = [8,-8]
	Output: []
	Explanation: The 8 and -8 collide exploding each other.

	Example 3:
	Input: asteroids = [10,2,-5]
	Output: [10]
	Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		asteroids := []int{5, 10, -5}
		exp := []int{5, 10}
		assert.Equal(t, exp, RunAsteroidCollision(asteroids))
	})
	t.Run("Suite 2", func(t *testing.T) {
		asteroids := []int{8, -8}
		exp := []int{}
		assert.Equal(t, exp, RunAsteroidCollision(asteroids))
	})
	t.Run("Suite 3", func(t *testing.T) {
		asteroids := []int{10, 2, -5}
		exp := []int{10}
		assert.Equal(t, exp, RunAsteroidCollision(asteroids))
	})
}

func RunAsteroidCollision(asteroids []int) []int {
	stk := s.NewStack()

	for _, asteroid := range asteroids {
		for !stk.IsEmpty() && asteroid < 0 && stk.Peek() > 0 {
			diff := asteroid + stk.Peek()

			if diff < 0 {
				stk.Pop()
			} else if diff > 0 {
				asteroid = 0
			} else {
				asteroid = 0
				stk.Pop()
			}
		}
		if asteroid != 0 {
			stk.Push(asteroid)
		}
	}

	res := []int{}
	for !stk.IsEmpty() {
		res = append([]int{stk.Pop()}, res...)
	}

	return res
}
