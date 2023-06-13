package concurrency

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestPrintInOrder(t *testing.T) {
	t.Log(`
	https://leetcode.com/problems/print-in-order

	Suppose we have a class:
	public class Foo {
		public void first() { print("first"); }
		public void second() { print("second"); }
		public void third() { print("third"); }
	}

	The same instance of Foo will be passed to three different threads. 
	Thread A will call first(), thread B will call second(), 
	and thread C will call third(). 
	Design a mechanism and modify the program to ensure that second() 
	is executed after first(), and third() is executed after second().
	`)

	defer goleak.VerifyNone(t)

	exp := "123"
	t.Run("Suite 1", func(t *testing.T) {
		in := []int{1, 2, 3}
		assert.Equal(t, exp, RunPrintInOrder(in))
	})
	t.Run("Suite 2", func(t *testing.T) {
		in := []int{3, 2, 1}
		assert.Equal(t, exp, RunPrintInOrder(in))
	})
	t.Run("Suite 3", func(t *testing.T) {
		in := []int{2, 3, 1}
		assert.Equal(t, exp, RunPrintInOrder(in))
	})
}

type PrintInOrder struct {
	c1     chan struct{}
	c2     chan struct{}
	c3     chan struct{}
	Output string
}

func (p *PrintInOrder) First() {
	p.Output = p.Output + "1"
	p.c1 <- struct{}{}
}

func (p *PrintInOrder) Second() {
	<-p.c1
	p.Output = p.Output + "2"
	p.c2 <- struct{}{}
}

func (p *PrintInOrder) Third() {
	<-p.c2
	p.Output = p.Output + "3"
	p.c3 <- struct{}{}
}

func NewPrintInOrder() *PrintInOrder {
	return &PrintInOrder{
		c1:     make(chan struct{}),
		c2:     make(chan struct{}),
		c3:     make(chan struct{}),
		Output: "",
	}
}

func RunPrintInOrder(callOrder []int) string {
	p := NewPrintInOrder()

	for _, call := range callOrder {
		switch call {
		case 1:
			go p.First()
		case 2:
			go p.Second()
		case 3:
			go p.Third()
		default:
			panic("not implemented")
		}
	}

	<-p.c3
	return p.Output
}
