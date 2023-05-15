package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestTimeBuyAndSellStock(t *testing.T) {
	t.Log(`
	You are given an array prices where prices[i] is the price of a given stock on the ith day.
	You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
	Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

	Example 1:
	Input: prices = [7,1,5,3,6,4]
	Output: 5
	Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
	Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
	
	Example 2:
	Input: prices = [7,6,4,3,1]
	Output: 0
	Explanation: In this case, no transactions are done and the max profit = 0.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := []int{7, 1, 5, 3, 6, 4}
		exp := 5
		assert.Equal(t, RunBestTimeBuyAndSellStock(in), exp)
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := []int{7, 6, 4, 3, 1}
		exp := 0
		assert.Equal(t, RunBestTimeBuyAndSellStock(in), exp)
	})

	t.Run("Suite 3", func(t *testing.T) {
		in := []int{1, 6, 4, 3, 10}
		exp := 9
		assert.Equal(t, RunBestTimeBuyAndSellStock(in), exp)
	})
}

func RunBestTimeBuyAndSellStock(in []int) int {
	var maxProfit int

	l := 0
	r := l + 1

	for r < len(in) {
		if in[l] > in[r] {
			l++
			r++
		} else {
			currentProfit := in[r] - in[l]
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
			r++
		}
	}

	return maxProfit
}
