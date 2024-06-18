package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitAssignment(t *testing.T) {
	t.Log(`
	You have n jobs and m workers. You are given three arrays: difficulty, profit, and worker where:

	difficulty[i] and profit[i] are the difficulty and the profit of the ith job, and
	worker[j] is the ability of jth worker (i.e., the jth worker can only complete a job with
	difficulty at most worker[j]).
	Every worker can be assigned at most one job, but one job can be completed multiple times.

	For example, if three workers attempt the same job that pays $1, 
	then the total profit will be $3. If a worker cannot complete any job, their profit is $0.
	Return the maximum profit we can achieve after assigning the workers to the jobs.

	Example 1:
	Input: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
	Output: 100
	Explanation: Workers are assigned jobs of difficulty [4,4,6,6] and they get a profit of 
	[20,20,30,30] separately.

	Example 2:
	Input: difficulty = [85,47,57], profit = [24,66,99], worker = [40,25,25]
	Output: 0
	`)

	t.Run("Suite 1", func(t *testing.T) {
		difficulty := []int{2, 4, 6, 8, 10}
		profit := []int{10, 20, 30, 40, 50}
		worker := []int{4, 5, 6, 7}
		exp := 100

		assert.Equal(t, exp, RunMaxProfitAssignment(difficulty, profit, worker))
	})
	t.Run("Suite 2", func(t *testing.T) {
		difficulty := []int{85, 47, 57}
		profit := []int{24, 66, 99}
		worker := []int{40, 25, 25}
		exp := 0

		assert.Equal(t, exp, RunMaxProfitAssignment(difficulty, profit, worker))
	})
}

func RunMaxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	type job struct {
		diff, prof int
	}
	jobs := make([]job, len(difficulty))

	for i := range difficulty {
		jobs[i] = job{
			diff: difficulty[i],
			prof: profit[i],
		}
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].diff < jobs[j].diff
	})

	sort.Ints(worker)

	var res, maxProfit, j int

	for _, w := range worker {
		for j < len(jobs) && jobs[j].diff <= w {
			maxProfit = max(maxProfit, jobs[j].prof)
			j++
		}
		res = res + maxProfit
	}
	return res
}
