package concurrency

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestSumOfSquare(t *testing.T) {
	t.Log(`
	Write a Go program that concurrently calculates the sum of squares of numbers 
	from 1 to N, where N is a given positive integer. 
	The program should use goroutines to split the calculation across multiple 
	tasks and then combine the results to get the final sum.
	`)

	defer goleak.VerifyNone(t)
	t.Run("Suite 1", func(t *testing.T) {
		in := 2
		exp := 5

		assert.Equal(t, exp, RunSumOfSquare(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := 8
		exp := 204

		assert.Equal(t, exp, RunSumOfSquare(in))
	})

	t.Run("Suite 3", func(t *testing.T) {
		in := 100
		exp := 338350

		assert.Equal(t, exp, RunSumOfSquare(in))
	})
}

func RunSumOfSquare(N int) int {

	var wg sync.WaitGroup
	resultCh := make(chan int)

	numGoroutines := 3
	chunkSize := (N + numGoroutines - 1) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			var sum int
			for j := start; j <= end && j <= N; j++ {
				sum = sum + j*j
			}
			resultCh <- sum
		}(i*chunkSize+1, (i+1)*chunkSize)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var totalSum int
	for res := range resultCh {
		totalSum = totalSum + res
	}
	return totalSum
}
