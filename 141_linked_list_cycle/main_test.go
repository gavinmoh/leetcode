package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	list1 := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}
	case1 := list1
	secondNode := list1.Next
	lastNode := list1.Next.Next.Next
	lastNode.Next = secondNode
	assert.Equal(t, true, HasCycle(case1))

	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	case2 := list2
	lastNode = list2.Next
	lastNode.Next = list2
	assert.Equal(t, true, HasCycle(case2))

	list3 := &ListNode{Val: 1}
	case3 := list3
	assert.Equal(t, false, HasCycle(case3))
}
