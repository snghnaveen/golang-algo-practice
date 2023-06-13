package basics

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBinary(t *testing.T) {
	t.Log(`
	Program for Decimal to Binary Conversion
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := 17
		exp := 10001

		// 17 / 2 = 8 |  17 %2 [1]
		// 8 / 2 = 4  |  8 %2 [0]
		// 4 / 2 = 2  |  4 %2 [0]
		// 2 / 2 = 1  |  2 %2 [0]
		// 1 / 2 = 0 |  1 %2 [1]
		assert.Equal(t, exp, RunToBinary(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := 23
		exp := 10111

		// 23 / 2 = 11 [1]
		// 11 / 2 = 5 [1]
		// 5 / 2 = 2 [1]
		// 2 / 2 = 1 [0]
		// 1 / 2 = 0 [1]
		assert.Equal(t, exp, RunToBinary(in))
	})
}

func RunToBinary(n int) int {
	var out []int
	for n > 0 {
		out = append(out, n%2)
		n = n / 2
	}

	// convert []int to int
	var buf bytes.Buffer
	for i := len(out) - 1; i >= 0; i-- {
		buf.WriteString(fmt.Sprintf("%d", out[i]))
	}

	binaryRes, err := strconv.Atoi(buf.String())
	if err != nil {
		panic(err)
	}

	return binaryRes
}
