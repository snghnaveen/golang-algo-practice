package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeSumQueryImmutable(t *testing.T) {
	t.Log(`https://leetcode.com/problems/range-sum-query-immutable/description/`)

	t.Log(`
	Given an integer array nums, handle multiple queries of the following type:
	Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.

	Implement the NumArray class:
	NumArray(int[] nums) Initializes the object with the integer array nums.
	int sumRange(int left, int right) Returns the sum of the elements of nums 
	between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).


	Example 1:

	Input
	["NumArray", "sumRange", "sumRange", "sumRange"]
	[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
	Output
	[null, 1, -1, -3]

	Explanation
	NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
	numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
	numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
	numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3
	`)

	intPtr := func(i int) *int {
		return &i
	}

	args := TestInputRangeSum{
		Action: []string{"NumArray", "sumRange", "sumRange", "sumRange"},
		Value: [][]int{
			{-2, 0, 3, -5, 2, -1}, // nums for constructor
			{0, 2},
			{2, 5},
			{0, 5},
		},
		Output: []*int{
			nil,
			intPtr(1),
			intPtr(-1),
			intPtr(-3),
		},
	}

	RunNumArray(t, args)

}

type TestInputRangeSum struct {
	Action []string
	Value  [][]int
	Output []*int
}

func RunNumArray(t *testing.T, args TestInputRangeSum) {
	var numArr NumArray

	for i, action := range args.Action {
		value := args.Value[i]

		switch action {
		case "NumArray":
			// value is a single element: the nums array
			numArr = ConstructorRangeSum(value)

		case "sumRange":
			left := value[0]
			right := value[1]
			exp := args.Output[i]
			assert.Equal(t, *exp, numArr.SumRange(left, right))
		}
	}
}

type NumArray struct {
	Elements []int
}

func ConstructorRangeSum(nums []int) NumArray {
	n := NumArray{}
	n.Elements = nums

	for i := 1; i < len(nums); i++ {
		n.Elements[i] = n.Elements[i-1] + n.Elements[i]
	}
	return n
}

func (n *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return n.Elements[right]
	}
	return n.Elements[right] - n.Elements[left-1]
}
