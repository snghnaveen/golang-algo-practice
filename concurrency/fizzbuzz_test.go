package concurrency

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestFizzBuzz(t *testing.T) {
	t.Log(`
	Write a program that returns slice (or print) the numbers from 1 to n
	and for multiples of ‘3’ print “Fizz” instead of the number 
	and for the multiples of ‘5’ print “Buzz” using go routine. 
	`)

	defer goleak.VerifyNone(t)

	t.Run("Suite 1", func(t *testing.T) {
		in := 3
		exp := []string{"1", "2", "Fizz"}
		assert.Equal(t, exp, RunFizzBuzz(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := 15
		exp := []string{
			"1", "2", "Fizz",
			"4", "Buzz", "Fizz",
			"7", "8", "Fizz",
			"Buzz", "11", "Fizz",
			"13", "14", "FizzBuzz",
		}

		assert.Equal(t, exp, RunFizzBuzz(in))
	})
}

type FizzBuzz struct {
	chFizz, chBuzz, chFizzBuzz, chNum, chNotify chan int
	chDone                                      chan struct{}
	res                                         []string
}

func NewFizzBuzz() *FizzBuzz {
	return &FizzBuzz{
		chFizz:     make(chan int),
		chBuzz:     make(chan int),
		chFizzBuzz: make(chan int),
		chNum:      make(chan int),
		chNotify:   make(chan int),
		chDone:     make(chan struct{}),
		res:        make([]string, 0),
	}
}

func (fb *FizzBuzz) fizz() {
	for i := range fb.chFizz {
		if i%3 == 0 && i%5 != 0 {
			fb.res = append(fb.res, "Fizz")
			fb.chDone <- struct{}{}
		}
	}
}

func (fb *FizzBuzz) buzz() {
	for i := range fb.chBuzz {
		if i%3 != 0 && i%5 == 0 {
			fb.res = append(fb.res, "Buzz")
			fb.chDone <- struct{}{}
		}
	}
}

func (fb *FizzBuzz) fizzbuz() {
	for i := range fb.chFizzBuzz {
		if i%3 == 0 && i%5 == 0 {
			fb.res = append(fb.res, "FizzBuzz")
			fb.chDone <- struct{}{}
		}
	}
}

func (fb *FizzBuzz) num() {
	for i := range fb.chNum {
		if i%3 != 0 && i%5 != 0 {
			fb.res = append(fb.res, strconv.Itoa(i))
			fb.chDone <- struct{}{}
		}
	}
}

func (fb *FizzBuzz) closeAllChans() {
	close(fb.chNotify)
	close(fb.chFizz)
	close(fb.chBuzz)
	close(fb.chFizzBuzz)
	close(fb.chNum)
}

func RunFizzBuzz(n int) []string {
	fb := NewFizzBuzz()

	go fb.fizz()
	go fb.buzz()
	go fb.fizzbuz()
	go fb.num()

	for i := 1; i <= n; i++ {
		fb.chFizz <- i
		fb.chBuzz <- i
		fb.chFizzBuzz <- i
		fb.chNum <- i

		<-fb.chDone
	}

	fb.closeAllChans()

	return fb.res
}
