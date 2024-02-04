package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddString(t *testing.T) {
	t.Log(`Add 2 strings`)
	t.Run("Suite 1", func(t *testing.T) {
		num1 := "13"
		num2 := "24"
		exp := 37
		assert.Equal(t, exp, RunAddString(num1, num2))
	})
	t.Run("Suite 2", func(t *testing.T) {
		num1 := "99"
		num2 := "10"
		exp := 109
		assert.Equal(t, exp, RunAddString(num1, num2))
	})

}

func RunAddString(num1, num2 string) int {

	n1Len, n2Len := len(num1), len(num2)

	for n1Len > n2Len {
		num2 = "0" + num2
		n2Len++
	}

	for n2Len > n1Len {
		num1 = "0" + num1
		n1Len++
	}

	carry := 0
	result := 0
	decimalPlace := 1

	for i := n1Len - 1; i >= 0; i-- {
		digit1 := int(num1[i] - '0')
		digit2 := int(num2[i] - '0')
		sumDigit := digit1 + digit2 + carry

		carry = sumDigit / 10

		result = result + (sumDigit%10)*decimalPlace
		decimalPlace = decimalPlace * 10
	}

	if carry > 0 {
		result = result + carry*decimalPlace
	}

	return result
}
