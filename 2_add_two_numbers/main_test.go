package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	list := addTwoNumbers(l1, l2)
	assert.Equal(t, 7, list.Val)
	assert.Equal(t, 0, list.Next.Val)
	assert.Equal(t, 8, list.Next.Next.Val)
}
