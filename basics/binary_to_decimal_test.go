package basics

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDecimal(t *testing.T) {
	t.Log(`
	Program for Binary to Decimal Binary Conversion
	`)

	t.Run("Suite 1", func(t *testing.T) {
		in := 10001
		exp := 17
		// 1 * 2^0 = 1
		// 0 * 2^1 = 0
		// 0 * 2^2 = 0
		// 0 * 2^3 = 0
		// 1 * 2^4 = 16
		assert.Equal(t, exp, RunToDecimal(in))
	})

	t.Run("Suite 2", func(t *testing.T) {
		in := 10111
		exp := 23
		assert.Equal(t, exp, RunToDecimal(in))
	})
	t.Run("Suite 3", func(t *testing.T) {
		in := 10111
		exp := 23
		assert.Equal(t, exp, RunToDecimal2(in))
	})
}

func RunToDecimal(n int) int {
	str := strconv.Itoa(n)
	var inp []int
	for _, v := range str {
		inp = append(inp, int(v)-'0')
	}

	var res int
	var p int
	for i := len(inp) - 1; i >= 0; i-- {
		out := inp[i] * int(math.Pow(2, float64(p)))
		p++
		res = out + res
	}
	return res
}

func RunToDecimal2(n int) int {
	var res int
	temp := n
	base := 1
	for temp > 0 {
		last := temp % 10
		temp = temp / 10
		res = res + last*base
		base = base * 2

	}
	return res
}
