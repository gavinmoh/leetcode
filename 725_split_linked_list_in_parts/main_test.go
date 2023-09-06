package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitListToParts(t *testing.T) {

	case1Node1 := &ListNode{Val: 1}
	case1Node2 := &ListNode{Val: 2}
	case1Node3 := &ListNode{Val: 3}

	case1Node1.Next = case1Node2
	case1Node2.Next = case1Node3

	case1Expected := []*ListNode{case1Node1, case1Node2, case1Node3, nil, nil}
	assert.Equal(t, case1Expected, SplitListToParts(case1Node1, 5))
	assert.Nil(t, case1Node1.Next)
	assert.Nil(t, case1Node2.Next)
	assert.Nil(t, case1Node3.Next)

	case2Node1 := &ListNode{Val: 1}
	case2Node2 := &ListNode{Val: 2}
	case2Node3 := &ListNode{Val: 3}
	case2Node4 := &ListNode{Val: 4}
	case2Node5 := &ListNode{Val: 5}
	case2Node6 := &ListNode{Val: 6}
	case2Node7 := &ListNode{Val: 7}
	case2Node8 := &ListNode{Val: 8}
	case2Node9 := &ListNode{Val: 9}
	case2Node10 := &ListNode{Val: 10}

	case2Node1.Next = case2Node2
	case2Node2.Next = case2Node3
	case2Node3.Next = case2Node4
	case2Node4.Next = case2Node5
	case2Node5.Next = case2Node6
	case2Node6.Next = case2Node7
	case2Node7.Next = case2Node8
	case2Node8.Next = case2Node9
	case2Node9.Next = case2Node10

	case2Expected := []*ListNode{case2Node1, case2Node5, case2Node8}
	assert.Equal(t, case2Expected, SplitListToParts(case2Node1, 3))
	assert.Nil(t, case2Node4.Next)
	assert.Nil(t, case2Node7.Next)
	assert.Nil(t, case2Node10.Next)

	case3Node1 := &ListNode{Val: 0}
	case3Node2 := &ListNode{Val: 1}
	case3Node3 := &ListNode{Val: 2}
	case3Node4 := &ListNode{Val: 3}
	case3Node5 := &ListNode{Val: 4}

	case3Node1.Next = case3Node2
	case3Node2.Next = case3Node3
	case3Node3.Next = case3Node4
	case3Node4.Next = case3Node5

	case3Expected := []*ListNode{case3Node1, case3Node3, case3Node5}
	assert.Equal(t, case3Expected, SplitListToParts(case3Node1, 3))
	assert.Nil(t, case3Node2.Next)
	assert.Nil(t, case3Node4.Next)

}
