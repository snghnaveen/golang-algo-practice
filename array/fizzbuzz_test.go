package array

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Log(`
	Write a program that returns slice (or print) the numbers from 1 to n
	and for multiples of ‘3’ print “Fizz” instead of the number 
	and for the multiples of ‘5’ print “Buzz”. 
	`)

	assert.Equal(t, RunFizzBuzz(3), []string{"1", "2", "Fizz"})

	assert.Equal(t, RunFizzBuzz(15), []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"})

}

func RunFizzBuzz(n int) []string {
	var out []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			out = append(out, "FizzBuzz")

		} else if i%3 == 0 {
			out = append(out, "Fizz")

		} else if i%5 == 0 {
			out = append(out, "Buzz")

		} else {
			out = append(out, strconv.Itoa(i))
		}
	}
	return out
}
