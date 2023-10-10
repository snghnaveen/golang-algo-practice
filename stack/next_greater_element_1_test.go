package stack

import (
	"testing"

	s "github.com/snghnaveen/golang-algo-practice/internal"
	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElement1(t *testing.T) {
	t.Log(`
	The next greater element of some element x in an array is the first greater element that is to the right of x in the same array.
	You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.
	For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and determine the next greater element of nums2[j] in nums2. If there is no next greater element, then the answer for this query is -1.
	Return an array ans of length nums1.length such that ans[i] is the next greater element as described above.

	Example 1:
	Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
	Output: [-1,3,-1]
	Explanation: The next greater element for each value of nums1 is as follows:
	- 4 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
	- 1 is underlined in nums2 = [1,3,4,2]. The next greater element is 3.
	- 2 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.

	Example 2:
	Input: nums1 = [2,4], nums2 = [1,2,3,4]
	Output: [3,-1]
	Explanation: The next greater element for each value of nums1 is as follows:
	- 2 is underlined in nums2 = [1,2,3,4]. The next greater element is 3.
	- 4 is underlined in nums2 = [1,2,3,4]. There is no next greater element, so the answer is -1.
 

	Constraints:
	1 <= nums1.length <= nums2.length <= 1000
	0 <= nums1[i], nums2[i] <= 104
	All integers in nums1 and nums2 are unique.
	All the integers of nums1 also appear in nums2.
	`)

	t.Run("Suite 1", func(t *testing.T) {
		nums1 := []int{4, 1, 2}
		nums2 := []int{1, 3, 4, 2}

		exp := []int{-1, 3, -1}

		assert.Equal(t, exp, RunNextGreaterElement1(nums1, nums2))
	})

	t.Run("Suite 2", func(t *testing.T) {
		nums1 := []int{4, 1, 2}
		nums2 := []int{2, 1, 3, 4}

		exp := []int{-1, 3, 3}

		assert.Equal(t, exp, RunNextGreaterElement1(nums1, nums2))
	})

}

func RunNextGreaterElement1(nums1, nums2 []int) []int {
	mpValIdx := make(map[int]int)

	out := make([]int, len(nums1))
	for i, v := range nums1 {
		mpValIdx[v] = i

		out[i] = -1 // default value
	}

	stk := s.NewStack()

	for _, v := range nums2 {
		for !stk.IsEmpty() && stk.Peek() < v {
			stkVal := stk.Pop()
			stkValIdx := mpValIdx[stkVal]
			out[stkValIdx] = v
		}

		if _, ok := mpValIdx[v]; ok {
			stk.Push(v)
		}
	}

	return out
}
