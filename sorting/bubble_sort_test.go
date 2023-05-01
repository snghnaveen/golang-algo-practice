package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	t.Log(`Bubble Sort`)

	in := []int{7, 6, 4, 3, 1}
	out := []int{1, 3, 4, 6, 7}

	RunBubbleSort(in)

	assert.Equal(t, in, out)
}

func RunBubbleSort(in []int) {
	n := len(in)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
			}
		}
	}
}
