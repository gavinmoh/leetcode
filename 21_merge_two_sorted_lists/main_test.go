package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func MapToArray(list *ListNode) []int {
	if list == nil {
		return []int{}
	}

	var result []int
	for list != nil {
		result = append(result, list.Val)
		list = list.Next
	}
	return result
}

func TestMergeTwoLists(t *testing.T) {
	// 	Input: list1 = [1,2,4], list2 = [1,3,4]
	// Output: [1,1,2,3,4,4]

	case1List1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	case1List2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	case1Expected := []int{1, 1, 2, 3, 4, 4}
	assert.Equal(t, case1Expected, MapToArray(MergeTwoLists(case1List1, case1List2)))

	// 	Input: list1 = [], list2 = []
	// Output: []

	var case2List1 *ListNode
	var case2List2 *ListNode
	case2Expected := []int{}
	assert.Equal(t, case2Expected, MapToArray(MergeTwoLists(case2List1, case2List2)))

	// Input: list1 = [], list2 = [0]
	// Output: [0]

	var case3List1 *ListNode
	case3List2 := &ListNode{Val: 0}
	case3Expected := []int{0}
	assert.Equal(t, case3Expected, MapToArray(MergeTwoLists(case3List1, case3List2)))
}
