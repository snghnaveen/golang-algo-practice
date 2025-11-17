package concurrency

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestOddEven(t *testing.T) {
	t.Log(`
	Even–Odd Number Printer Using Goroutines and Channels
	Description:
	Write a Go program that uses two goroutines to print numbers from 1 to N in 
	perfect alternating order:

	1 2 3 4 5 6 ...
	One goroutine should print odd numbers.
	Another goroutine should print even numbers.
	The two goroutines must coordinate using unbuffered channels to ensure 
	that the numbers are printed strictly in order.

	Requirements:
	No busy waiting — you cannot use time.Sleep.
	Use two unbuffered channels, e.g., oddChan and evenChan, for synchronization.
	Ensure no race conditions or missed numbers.
	The program should gracefully terminate after printing all numbers.
	Input:
	N (integer): The maximum number to print.

	Expected Output Example (for N=10):
	1 2 3 4 5 6 7 8 9 10
	`)

	defer goleak.VerifyNone(t)

	exp := func(n int) []int {
		var out []int
		for i := 1; i <= n; i++ {
			out = append(out, i)
		}
		return out
	}
	t.Run("Suite 1", func(t *testing.T) {
		n := 10
		assert.Equal(t, exp(n), RunOddEven(n))
	})
	t.Run("Suite 2", func(t *testing.T) {
		n := 535
		assert.Equal(t, exp(n), RunOddEven(n))
	})
	t.Run("Suite 3", func(t *testing.T) {
		n := 9000
		assert.Equal(t, exp(n), RunOddEven(n))
	})
}

func RunOddEven(n int) []int {
	var out []int
	var mu sync.Mutex

	oddChan := make(chan struct{})
	evenChan := make(chan struct{})

	done := make(chan struct{}, 2)

	go func() {
		for i := 1; i <= n; i = i + 2 {
			<-oddChan
			{
				mu.Lock()
				out = append(out, i)
				mu.Unlock()
			}
			if i+1 <= n {
				evenChan <- struct{}{}
			}
		}
		done <- struct{}{}
	}()

	go func() {
		for i := 2; i <= n; i = i + 2 {
			<-evenChan
			{
				mu.Lock()
				out = append(out, i)
				mu.Unlock()
			}
			if i+1 <= n {
				oddChan <- struct{}{}
			}
		}
		done <- struct{}{}
	}()

	oddChan <- struct{}{}

	<-done
	<-done
	return out
}
