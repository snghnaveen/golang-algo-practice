package concurrency

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestPrintFooBarAlternately(t *testing.T) {
	t.Log(`
	https://leetcode.com/problems/print-foobar-alternately/

	Suppose you are given the following code:

	class FooBar {
		public void foo() {
			for (int i = 0; i < n; i++) {
			print("foo");
			}
		}

		public void bar() {
			for (int i = 0; i < n; i++) {
			print("bar");
			}
		}
	}
	The same instance of FooBar will be passed to two different threads:

	thread A will call foo(), while
	thread B will call bar().
	Modify the given program to output "foobar" n times.
	`)

	defer goleak.VerifyNone(t)

	// for generating expected test output
	tFooBar := func(in int) string {
		out := ""
		for i := 1; in != 0 && i <= in; i++ {
			out = out + "foobar"
		}
		return out
	}

	t.Run("Suite 1", func(t *testing.T) {
		in := 2
		assert.Equal(t, tFooBar(in), RunPrintFooBarAlternately(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := 10
		assert.Equal(t, tFooBar(in), RunPrintFooBarAlternately(in))
	})
	t.Run("Suite 3", func(t *testing.T) {
		in := 0
		assert.Equal(t, tFooBar(in), RunPrintFooBarAlternately(in))
	})

}

type FooBar struct {
	n            int
	fooMu, barMu *sync.Mutex
	output       string
	wg           *sync.WaitGroup
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:      n,
		fooMu:  &sync.Mutex{},
		barMu:  &sync.Mutex{},
		output: "",
		wg:     &sync.WaitGroup{},
	}
}

func (fb *FooBar) Foo() {
	for i := 0; i < fb.n; i++ {
		fb.fooMu.Lock()
		fb.output = fb.output + "foo"
		fb.barMu.Unlock()
	}
}

func (fb *FooBar) Bar() {
	for i := 0; i < fb.n; i++ {
		fb.barMu.Lock()
		fb.output = fb.output + "bar"
		fb.fooMu.Unlock()
	}
	fb.wg.Done()
}

func RunPrintFooBarAlternately(in int) string {
	fb := NewFooBar(in)

	fb.barMu.Lock()
	fb.wg.Add(1)
	go fb.Foo()
	go fb.Bar()
	fb.wg.Wait()

	return fb.output
}
