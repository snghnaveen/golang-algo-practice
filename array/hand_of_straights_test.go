package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNStraightHand(t *testing.T) {
	t.Log(`
	Alice has some number of cards and she wants to rearrange the cards into groups 
	so that each group is of size groupSize, and consists of groupSize consecutive cards.
	Given an integer array hand where hand[i] is the value written on the ith card and an 
	integer groupSize, return true if she can rearrange the cards, or false otherwise.
	
	Example 1:
	Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
	Output: true
	Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]

	Example 2:
	Input: hand = [1,2,3,4,5], groupSize = 4
	Output: false
	Explanation: Alice's hand can not be rearranged into groups of 4.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
		groupSize := 3
		assert.True(t, RunIsNStraightHand(hand, groupSize))
	})
	t.Run("Suite 2", func(t *testing.T) {
		hand := []int{1, 2, 3, 4, 5}
		groupSize := 4
		assert.False(t, RunIsNStraightHand(hand, groupSize))
	})
}

func RunIsNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	mpCount := make(map[int]int)

	for _, card := range hand {
		mpCount[card]++
	}

	for _, card := range hand {
		if mpCount[card] == 0 {
			continue
		}

		for i := 0; i < groupSize; i++ {
			if mpCount[card+i] == 0 {
				return false
			}
			mpCount[card+i]--
		}
	}

	return true
}
