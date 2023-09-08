package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	case1Node1 := &ListNode{Val: 1}
	case1Node2 := &ListNode{Val: 4}
	case1Node3 := &ListNode{Val: 5}
	case1Node1.Next = case1Node2
	case1Node2.Next = case1Node3

	case1Node4 := &ListNode{Val: 1}
	case1Node5 := &ListNode{Val: 3}
	case1Node6 := &ListNode{Val: 4}
	case1Node4.Next = case1Node5
	case1Node5.Next = case1Node6

	case1Node7 := &ListNode{Val: 2}
	case1Node8 := &ListNode{Val: 6}
	case1Node7.Next = case1Node8

	case1 := []*ListNode{case1Node1, case1Node4, case1Node7}
	case1Result := MergeKLists(case1)
	case1Expected := []int{1, 1, 2, 3, 4, 4, 5, 6}
	for _, expected := range case1Expected {
		assert.Equal(t, expected, case1Result.Val)
		case1Result = case1Result.Next
	}

	case2 := []*ListNode{}
	case2Result := MergeKLists(case2)
	assert.Nil(t, case2Result)

	case3 := []*ListNode{nil}
	case3Result := MergeKLists(case3)
	assert.Nil(t, case3Result)
}
