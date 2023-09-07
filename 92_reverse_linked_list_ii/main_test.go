package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBetween(t *testing.T) {
	case1Node1 := &ListNode{Val: 1}
	case1Node2 := &ListNode{Val: 2}
	case1Node3 := &ListNode{Val: 3}
	case1Node4 := &ListNode{Val: 4}
	case1Node5 := &ListNode{Val: 5}

	case1Node1.Next = case1Node2
	case1Node2.Next = case1Node3
	case1Node3.Next = case1Node4
	case1Node4.Next = case1Node5

	ReverseBetween(case1Node1, 2, 4)

	// 1 -> 4 -> 3 -> 2 -> 5
	assert.Equal(t, case1Node1.Next, case1Node4)
	assert.Equal(t, case1Node4.Next, case1Node3)
	assert.Equal(t, case1Node3.Next, case1Node2)
	assert.Equal(t, case1Node2.Next, case1Node5)

	case2Node1 := &ListNode{Val: 3}
	case2Node2 := &ListNode{Val: 5}

	case2Node1.Next = case2Node2

	case2Expected := ReverseBetween(case2Node1, 1, 2)
	assert.Equal(t, case2Expected, case2Node2)

}
